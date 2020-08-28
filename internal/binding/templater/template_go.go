package templater

import (
	"bytes"
	"fmt"
	"go/format"
	"path/filepath"
	"strings"

	"github.com/therecipe/qt/internal/binding/converter"
	"github.com/therecipe/qt/internal/binding/parser"
	"github.com/therecipe/qt/internal/cmd"
	"github.com/therecipe/qt/internal/utils"
)

func GoTemplate(module string, stub bool, mode int, pkg, target, tags string) []byte {

	if !utils.QT_GEN_GO_WRAPPER() && module == "AndroidExtras" {
		utils.Save(utils.GoQtPkgPath(strings.ToLower(module), "utils-androidextras_android.go"), utils.Load(filepath.Join(strings.TrimSpace(utils.GoListOptional("{{.Dir}}", "github.com/therecipe/qt/internal", "-find", "get files dir")), "/binding/files/utils-androidextras_android.go")))
	} else if module == "Qml" {
		cont := utils.Load(filepath.Join(strings.TrimSpace(utils.GoListOptional("{{.Dir}}", "github.com/therecipe/qt/internal", "-find", "get files dir")), "/binding/files/utils-qml.go"))
		if utils.QT_API_NUM(utils.QT_VERSION()) < 5060 {
			cont = strings.Replace(cont, "Property2(", "Property3(", -1)
		}
		if utils.QT_API_NUM(utils.QT_VERSION()) < 5050 {
			cont = "package qml"
		}
		utils.Save(utils.GoQtPkgPath(strings.ToLower(module), "utils-qml.go"), cont)

		//TODO: make layer generic instead
		cont = strings.Replace(cont, "package qml", "package interop", -1)
		cont = strings.Replace(cont, "qml.", "interop.", -1)
		cont = strings.Replace(cont, "QJSValue", "PseudoQJSValue", -1)
		cont = strings.Replace(cont, "QJSEngine", "PseudoQJSEngine", -1)
		if utils.ExistsDir(utils.GoQtPkgPath(strings.ToLower("interop"))) {
			utils.Save(utils.GoQtPkgPath(strings.ToLower("interop"), "interop.go"), cont)
		}
	}

	utils.Log.WithField("module", module).Debug("generating go")
	parser.State.Target = target

	bb := new(bytes.Buffer)
	defer bb.Reset()

	if mode != MOC {
		module = "Qt" + module
	}

	if !utils.QT_GEN_GO_WRAPPER() && !(UseStub(stub, module, mode) || UseJs()) {
		var m string
		if module != "QtCore" {
			m = "core."
		}
		fmt.Fprintf(bb, "func cGoFreePacked(ptr unsafe.Pointer) { %vNewQByteArrayFromPointer(ptr).DestroyQByteArray() }\n", m)
		fmt.Fprintf(bb, "func cGoUnpackString(s C.struct_%v_PackedString) string { defer cGoFreePacked(s.ptr)\n if int(s.len) == -1 {\n return C.GoString(s.data)\n }\n return C.GoStringN(s.data, C.int(s.len)) }\n", strings.Title(module))
		fmt.Fprintf(bb, "func cGoUnpackBytes(s C.struct_%v_PackedString) []byte { defer cGoFreePacked(s.ptr)\n if int(s.len) == -1 {\n gs := C.GoString(s.data)\n return []byte(gs)\n }\n return C.GoBytes(unsafe.Pointer(s.data), C.int(s.len)) }\n", strings.Title(module))
	}

	if UseJs() {
		fmt.Fprint(bb, "func jsGoUnpackString(s string) string { dec, _ := hex.DecodeString(s)\n return string(dec)\n }\n") //TODO: calling it cGoUnpackString won't work, bug in go wasm ?
		fmt.Fprint(bb, "func jsGoUnpackBytes(s string) []byte { dec, _ := hex.DecodeString(s)\n return dec\n }\n")
	}

	if !utils.QT_GEN_GO_WRAPPER() {
		fmt.Fprint(bb, "func unpackStringList(s string) []string {\nif len(s) == 0 {\nreturn make([]string, 0)\n}\nreturn strings.Split(s, \"¡¦!\")\n}\n")
	}

	if module == "QtAndroidExtras" && utils.QT_VERSION_NUM() >= 5060 {
		fmt.Fprint(bb, "func QAndroidJniEnvironment_ExceptionCatch() error {\n")
		if UseStub(stub, module, mode) || UseJs() {
			fmt.Fprint(bb, "return nil\n")
		} else {
			fmt.Fprint(bb, "var err error\n")
			fmt.Fprint(bb, "if QAndroidJniEnvironment_ExceptionCheck() {\n tmpExcPtr := QAndroidJniEnvironment_ExceptionOccurred()\nQAndroidJniEnvironment_ExceptionClear()\n")
			fmt.Fprint(bb, "tmpExc := NewQAndroidJniObject6(tmpExcPtr)\n")
			fmt.Fprint(bb, "err = errors.New(tmpExc.CallMethodString2(\"toString\", \"()Ljava/lang/String;\"))\n")
			fmt.Fprint(bb, "tmpExc.DestroyQAndroidJniObject()\n")
			fmt.Fprint(bb, "}\nreturn err\n")
		}
		fmt.Fprint(bb, "}\n\n")

		if UseStub(stub, module, mode) || UseJs() {
			fmt.Fprint(bb, "func (e *QAndroidJniEnvironment) ExceptionCatch() error { return nil }\n")
		} else {
			fmt.Fprint(bb, "func (e *QAndroidJniEnvironment) ExceptionCatch() error { return QAndroidJniEnvironment_ExceptionCatch() }\n")
		}
	}

	for _, class := range parser.SortedClassesForModule(module, true) {

		class.Stub = UseStub(stub, module, mode)

		if mode != MINIMAL || (mode == MINIMAL && class.Export) {

			if mode != MOC {
				fmt.Fprintf(bb, "type %v struct {\n%v\n}\n\n",

					class.Name,

					func() string {
						if class.Bases == "" {
							if utils.QT_GEN_GO_WRAPPER() {
								return "internal.Internal"
							}
							return "ptr unsafe.Pointer"
						}

						var bb = new(bytes.Buffer)
						defer bb.Reset()

						for _, parentClassName := range class.GetBases() {
							var parentClass, ok = parser.State.ClassMap[parentClassName]
							if !ok {
								continue
							}
							if parentClass.Module == class.Module {
								fmt.Fprintf(bb, "%v\n", parentClassName)
							} else {
								fmt.Fprintf(bb, "%v.%v\n", goModule(parentClass.Module), parentClassName)
							}
						}

						return bb.String()
					}(),
				)
			}

			fmt.Fprintf(bb, "type %v_ITF interface {\n%v%v\n}\n\n",

				class.Name,

				func() string {
					var bb = new(bytes.Buffer)
					defer bb.Reset()

					for _, parentClassName := range class.GetBases() {
						var parentClass, ok = parser.State.ClassMap[parentClassName]
						if !ok {
							continue
						}
						if parentClass.Module == class.Module {
							fmt.Fprintf(bb, "%v_ITF\n", parentClassName)
						} else {
							fmt.Fprintf(bb, "%v.%v_ITF\n", goModule(parentClass.Module), parentClassName)
						}
					}

					return bb.String()
				}(),

				fmt.Sprintf("%[1]v_PTR() *%[1]v\n", class.Name),
			)

			fmt.Fprintf(bb, "func (ptr *%[1]v) %[1]v_PTR() *%[1]v {\nreturn ptr\n}\n\n", class.Name)

			if class.Bases == "" {
				if !utils.QT_GEN_GO_WRAPPER() {
					fmt.Fprintf(bb, "func (ptr *%v) Pointer() unsafe.Pointer {\nif ptr != nil {\nreturn ptr.ptr\n}\nreturn nil\n}\n\n", class.Name)
					fmt.Fprintf(bb, "func (ptr *%v) SetPointer(p unsafe.Pointer) {\nif ptr != nil {\nptr.ptr = p\n}\n}\n\n", class.Name)
				} else {
					fmt.Fprintf(bb, "func (ptr *%v) Pointer() unsafe.Pointer {\nif ptr != nil {\nreturn unsafe.Pointer(ptr.Internal.Pointer())\n}\nreturn nil\n}\n\n", class.Name)
					fmt.Fprintf(bb, "func (ptr *%v) SetPointer(p unsafe.Pointer) {\nif ptr != nil {\nptr.Internal.SetPointer(uintptr(p))\n}\n}\n\n", class.Name)
				}
			} else {
				fmt.Fprintf(bb, "func (ptr *%v) Pointer() unsafe.Pointer {\nif ptr != nil {\nreturn ptr.%v_PTR().Pointer()\n}\nreturn nil\n}\n\n", class.Name, class.GetBases()[0])

				fmt.Fprintf(bb, "func (ptr *%v) SetPointer(p unsafe.Pointer) {\nif ptr != nil{\n%v}\n}\n",

					class.Name,

					func() string {
						var bb = new(bytes.Buffer)
						defer bb.Reset()

						for _, parentClassName := range class.GetBases() {
							fmt.Fprintf(bb, "ptr.%v_PTR().SetPointer(p)\n", parentClassName)
						}

						return bb.String()
					}(),
				)
			}

			fmt.Fprintf(bb, `
func PointerFrom%[1]v(ptr %[2]v_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.%[2]v_PTR().Pointer()
	}
	return nil
}
`, strings.Title(class.Name), class.Name)

			if class.Module == parser.MOC {
				fmt.Fprintf(bb, `
func New%[1]vFromPointer(ptr unsafe.Pointer) (n *%[2]v) {
	if gPtr, ok := qt.Receive(ptr); !ok {
		n = new(%[2]v)
		n.SetPointer(ptr)
	} else {
		switch deduced := gPtr.(type) {
			case *%[2]v:
				n = deduced

			case *%[3]v:
				n = &%[2]v{%[4]v: *deduced }

			default:
				n = new(%[2]v)
				n.SetPointer(ptr)
		}
	}%[5]v
	return
}
`,
					strings.Title(class.Name),
					class.Name,
					func() string {
						bc := class.GetBases()[0]
						if class.Module == parser.State.ClassMap[bc].Module {
							return bc
						}
						return fmt.Sprintf("%v.%v", goModule(parser.State.ClassMap[bc].Module), bc)
					}(),
					class.GetBases()[0],
					func() string {
						if utils.QT_GEN_GO_WRAPPER() {
							bc := class.GetBases()[0]
							if class.Module != parser.State.ClassMap[bc].Module {
								bc = fmt.Sprintf("%v.%v", goModule(parser.State.ClassMap[bc].Module), bc)
							}
							return fmt.Sprintf("\nn.InitFromInternal(uintptr(ptr), \"%v\")", bc)
						}
						return ""
					}())
			} else {
				if utils.QT_GEN_GO_WRAPPER() {
					if len(class.GetBases()) >= 1 {
						fmt.Fprintf(bb, `
func (n *%v) InitFromInternal(ptr uintptr, name string) {
	%v
}
`, class.Name, func() string {
							var bb = new(bytes.Buffer)
							defer bb.Reset()
							for _, parentClassName := range class.GetBases() {
								fmt.Fprintf(bb, "n.%v_PTR().InitFromInternal(uintptr(ptr), name)\n", parentClassName)
							}
							return bb.String()
						}())

						fmt.Fprintf(bb, `
func (n *%v) ClassNameInternalF()string {
	return n.%v_PTR().ClassNameInternalF()
}
`, class.Name, class.GetBases()[0])
					} else {
						fmt.Fprintf(bb, `
func (n *%v) ClassNameInternalF()string {
	return n.Internal.ClassNameInternalF()
}
`, class.Name)
					}

					fmt.Fprintf(bb, `
func New%[1]vFromPointer(ptr unsafe.Pointer) (n *%[2]v) {
	n = new(%[2]v)
	n.InitFromInternal(uintptr(ptr), "%[3]v.%[2]v")
	return
}
`, strings.Title(class.Name), class.Name, goModule(class.Module))

				} else {
					fmt.Fprintf(bb, `
func New%[1]vFromPointer(ptr unsafe.Pointer) (n *%[2]v) {
	n = new(%[2]v)
	n.SetPointer(ptr)
	return
}
`, strings.Title(class.Name), class.Name)
				}
			}

			if !class.HasDestructor() {
				if UseStub(stub, module, mode) {
					fmt.Fprintf(bb, "\nfunc (ptr *%v) Destroy%v() {\n}\n\n", class.Name, strings.Title(class.Name))
				} else if utils.QT_GEN_GO_WRAPPER() {
					//TODO:
					fmt.Fprintf(bb, "\nfunc (ptr *%v) Destroy%v() {\n}\n\n", class.Name, strings.Title(class.Name))
				} else if !class.IsSubClassOfQObject() {
					var ds string
					if class.HasCallbackFunctions() {
						ds = "\nqt.DisconnectAllSignals(ptr.Pointer(), \"\")"
					}
					var free string
					if !UseJs() {
						free = "C.free(ptr.Pointer())" //TODO: free c ptr in js/wasm ?
					}
					fmt.Fprintf(bb, `func (ptr *%[1]v) Destroy%[1]v() {
						if ptr != nil {
							qt.SetFinalizer(ptr, nil)
							%v
							%v
							ptr.SetPointer(nil)
						}
					}
					`, class.Name, ds, free)
				}
			}

			if mode == MOC {

				if len(class.Constructors) > 0 {
					if strings.ToLower(class.Constructors[0])[0] == class.Constructors[0][0] {
						fmt.Fprintf(bb, "func (this *%v) %v() { this.%v() }\n", class.Name, strings.Title(class.Constructors[0]), class.Constructors[0])
					}
				}

				if UseJs() {
					if parser.UseWasm() {
						fmt.Fprintf(bb, "//export callback%[1]v_Constructor\nfunc callback%[1]v_Constructor(_ js.Value, args []js.Value) interface{} {", class.Name)
						fmt.Fprint(bb, "\nptr := uintptr(args[0].Int())\n")
					} else {
						fmt.Fprintf(bb, "//export callback%[1]v_Constructor\nfunc callback%[1]v_Constructor(ptr uintptr) {", class.Name)
					}
					fmt.Fprintf(bb, "this := New%vFromPointer(unsafe.Pointer(ptr))\nqt.Register(unsafe.Pointer(ptr), this)\n", strings.Title(class.Name))
				} else {
					if utils.QT_GEN_GO_WRAPPER() {
						fmt.Fprintf(bb, "func callback%[1]v_Constructor(ptr unsafe.Pointer) *%[1]v {", class.Name)
					} else {
						fmt.Fprintf(bb, "//export callback%[1]v_Constructor\nfunc callback%[1]v_Constructor(ptr unsafe.Pointer) {", class.Name)
					}
					fmt.Fprintf(bb, "this := New%vFromPointer(ptr)\nqt.Register(ptr, this)\n", strings.Title(class.Name))
				}

				var lastModule string
				for _, bcn := range class.GetAllBases() {
					if bc := parser.State.ClassMap[bcn]; bc.Module != class.Module {
						if len(bc.Constructors) > 0 && lastModule != bc.Module {
							fmt.Fprintf(bb, "this.%v.%v()\n", strings.Title(bc.Name), strings.Title(bc.Constructors[0]))
						}
						lastModule = bc.Module
					}
				}

				for _, bcn := range append(class.GetAllBases(), class.Name) {
					if bc, ok := parser.State.ClassMap[bcn]; ok {
						for _, f := range bc.Functions {
							if f.Connect == 0 || !f.IsMocFunction {
								continue
							}

							if class.Name != bcn {
								if parser.UseJs() {
									fmt.Fprintf(bb, "qt.DisconnectSignal(unsafe.Pointer(ptr), \"%v\")\n", f.Name)
								} else {
									fmt.Fprintf(bb, "qt.DisconnectSignal(ptr, \"%v\")\n", f.Name)
								}
							}
						}
					}
				}

				connect := func(class *parser.Class, local bool) {
					for _, bcn := range append(class.GetAllBases(), class.Name) {
						if bc, ok := parser.State.ClassMap[bcn]; ok {
							for _, f := range bc.Functions {
								if f.Connect == 0 || !f.IsMocFunction {
									continue
								}
								if (local && f.Target != "") || (!local && f.Target == "") {
									continue
								}

								name := f.Name
								if f.Inbound {
									name = strings.Title(name)
								}

								if f.Connect == 1 {
									if f.Target == "" {
										fmt.Fprintf(bb, "this.Connect%v(this.%v)\n", strings.Title(name), name)
									} else {
										tUpper := strings.Split(f.Target, ".")
										tUpper[len(tUpper)-1] = strings.Title(tUpper[len(tUpper)-1])

										if strings.Count(f.Target, ".") >= 2 || (len(strings.Split(f.Target, ".")) == 2 && strings.Split(f.Target, ".")[0] != "this" && strings.Split(f.Target, ".")[1][:1] == strings.ToLower(strings.Split(f.Target, ".")[1][:1])) {
											fmt.Fprintf(bb, "this.Connect%v(%v)\n", strings.Title(name), strings.Join(tUpper, "."))
										} else {
											fmt.Fprintf(bb, "this.Connect%v(%v.%v)\n", strings.Title(name), f.Target, strings.Title(name))
										}
									}
								} else {
									if f.Target != "" {
										tCon := strings.Split(f.Target, ".")
										tCon[len(tCon)-1] = "Connect" + strings.Title(tCon[len(tCon)-1])

										if strings.Count(f.Target, ".") >= 2 || (len(strings.Split(f.Target, ".")) == 2 && strings.Split(f.Target, ".")[0] != "this" && strings.Split(f.Target, ".")[1][:1] == strings.ToLower(strings.Split(f.Target, ".")[1][:1])) {
											fmt.Fprintf(bb, "%v(this.%v)\n", strings.Join(tCon, "."), name)
										} else {
											fmt.Fprintf(bb, "%v.Connect%v(this.%v)\n", f.Target, strings.Title(name), name)
										}
									}
								}
							}
						}
					}

					for _, bcn := range append(class.GetAllBases(), class.Name) {
						if bc, ok := parser.State.ClassMap[bcn]; ok {
							for _, p := range bc.Properties {
								if p.Connect == 0 {
									continue
								}
								if (local && p.Target != "") || (!local && p.Target == "") {
									continue
								}

								name := p.Name
								if p.Inbound {
									name = strings.Title(name)
								}
								name = strings.TrimSuffix(name, "z__")

								if p.Connect == 1 {
									if p.Target == "" {
										if p.ConnectGet || !(p.ConnectSet || p.ConnectChanged) {
											fmt.Fprintf(bb, "this.Connect%v(this.%v)\n",
												func() string {
													if p.Output == "bool" && !strings.HasPrefix(name, "is") {
														return "Is" + strings.Title(name)
													}
													return strings.Title(name)
												}(),
												func() string {
													if p.Output == "bool" && !strings.HasPrefix(name, "is") {
														return "is" + strings.Title(name)
													}
													return name
												}())
										}
										if p.ConnectSet || !(p.ConnectGet || p.ConnectChanged) {
											fmt.Fprintf(bb, "this.ConnectSet%v(this.set%v)\n", strings.Title(name), strings.Title(name))
										}
										if p.ConnectChanged || !(p.ConnectGet || p.ConnectSet) {
											fmt.Fprintf(bb, "this.Connect%vChanged(this.%vChanged)\n", strings.Title(name), name)
										}
									} else {
										t := p.Target
										if strings.Count(t, ".") < 2 {
											if !(len(strings.Split(p.Target, ".")) == 2 && strings.Split(p.Target, ".")[0] != "this" && strings.Split(p.Target, ".")[1][:1] == strings.ToLower(strings.Split(p.Target, ".")[1][:1])) {
												t = p.Target + "." + name
											}
										}

										tSet := strings.Split(t, ".")
										tSet[len(tSet)-1] = "ConnectSet" + strings.Title(tSet[len(tSet)-1])

										tChanged := strings.Split(t, ".")
										tChanged[len(tChanged)-1] = strings.Title(tChanged[len(tChanged)-1]) + "Changed"

										tUpper := strings.Split(t, ".")
										tUpper[len(tUpper)-1] = strings.Title(tUpper[len(tUpper)-1])

										tIs := strings.Split(t, ".")
										if p.Output == "bool" && !strings.HasPrefix(tIs[len(tIs)-1], "is") {
											tIs[len(tIs)-1] = "ConnectIs" + strings.Title(tIs[len(tIs)-1])
										} else {
											tIs[len(tIs)-1] = "Connect" + strings.Title(tIs[len(tIs)-1])
										}

										if p.ConnectGet || !(p.ConnectSet || p.ConnectChanged) {
											fmt.Fprintf(bb, "%v(this.%v)\n",
												strings.Join(tIs, "."),
												func() string {
													if p.Output == "bool" && !strings.HasPrefix(name, "is") {
														return "Is" + strings.Title(name)
													}
													return strings.Title(name)
												}())
										}
										if p.ConnectSet || !(p.ConnectGet || p.ConnectChanged) {
											fmt.Fprintf(bb, "%v(this.Set%v)\n", strings.Join(tSet, "."), strings.Title(name))
										}
										if p.ConnectChanged || !(p.ConnectGet || p.ConnectSet) {
											fmt.Fprintf(bb, "this.Connect%vChanged(%v)\n", strings.Title(name), strings.Join(tChanged, "."))
										}
									}
								} else {
									if p.Target != "" {
										t := p.Target
										if strings.Count(t, ".") < 2 {
											if !(len(strings.Split(p.Target, ".")) == 2 && strings.Split(p.Target, ".")[0] != "this" && strings.Split(p.Target, ".")[1][:1] == strings.ToLower(strings.Split(p.Target, ".")[1][:1])) {
												t = p.Target + "." + name
											}
										}

										tSet := strings.Split(t, ".")
										tSet[len(tSet)-1] = "Set" + strings.Title(tSet[len(tSet)-1])

										tChanged := strings.Split(t, ".")
										tChanged[len(tChanged)-1] = "Connect" + strings.Title(tChanged[len(tChanged)-1]) + "Changed"

										tUpper := strings.Split(t, ".")
										tUpper[len(tUpper)-1] = strings.Title(tUpper[len(tUpper)-1])

										tIs := strings.Split(t, ".")
										if p.Output == "bool" && !strings.HasPrefix(tIs[len(tIs)-1], "is") {
											tIs[len(tIs)-1] = "Is" + strings.Title(tIs[len(tIs)-1])
										} else {
											tIs[len(tIs)-1] = strings.Title(tIs[len(tIs)-1])
										}

										if p.ConnectGet || !(p.ConnectSet || p.ConnectChanged) {
											fmt.Fprintf(bb, "this.Connect%v(%v)\n",
												func() string {
													if p.Output == "bool" && !strings.HasPrefix(name, "is") {
														return "Is" + strings.Title(name)
													}
													return strings.Title(name)
												}(), strings.Join(tIs, "."))
										}
										if p.ConnectSet || !(p.ConnectGet || p.ConnectChanged) {
											fmt.Fprintf(bb, "this.ConnectSet%v(%v)\n", strings.Title(name), strings.Join(tSet, "."))
										}
										if p.ConnectChanged || !(p.ConnectGet || p.ConnectSet) {
											fmt.Fprintf(bb, "%v(this.%vChanged)\n", strings.Join(tChanged, "."), strings.Title(name))
										}
									}
								}
							}
						}
					}
				}

				connect(class, true)

				if len(class.Constructors) > 0 {
					fmt.Fprintf(bb, "this.%v()\n", class.Constructors[0])
				}

				connect(class, false)

				if utils.QT_GEN_GO_WRAPPER() {
					bb.WriteString("\nreturn this\n")
				}

				if UseJs() {
					if parser.UseWasm() {
						bb.WriteString("\nreturn nil\n")
					}
				}

				fmt.Fprint(bb, "}\n\n")
			}

			if class.Name == "QVariant" {
				fmt.Fprint(bb, "type qVariant_ITF interface { ToVariant() *QVariant }\n")
				fmt.Fprint(bb, "func  NewQVariant1(i interface{}) *QVariant {\n")
				fmt.Fprint(bb, "switch d:= i.(type) {\n")

				has := make(map[string]struct{})
				fmt.Fprint(bb, "case *QVariant:\nreturn d\n")
				has["QVariant"] = struct{}{}

				for _, f := range class.Functions {
					if f.Meta == parser.CONSTRUCTOR && len(f.Parameters) == 1 && f.IsSupported() {
						v := f.Parameters[0].Value
						gt := converter.GoType(f, v, f.Parameters[0].PureGoType)
						if _, ok := has[gt]; ok {
							continue
						}
						if gt == "string" && !strings.Contains(f.Parameters[0].Value, "const char") {
							continue
						}
						if gt == "map[string]*QVariant" && !strings.Contains(f.Parameters[0].Value, "const QMap<") {
							continue
						}
						has[gt] = struct{}{}

						if c, ok := parser.IsClass(parser.CleanValue(v)); ok && parser.State.ClassMap[c].IsSupported() {
							fmt.Fprintf(bb, "case *%v:\n", gt)
						} else {
							fmt.Fprintf(bb, "case %v:\n", gt)
						}
						fmt.Fprintf(bb, "return NewQVariant%v(d)\n", f.OverloadNumber)
					}
				}
				fmt.Fprint(bb, "case qVariant_ITF:\nreturn d.ToVariant()\n")
				fmt.Fprint(bb, `default:
s := reflect.ValueOf(i)
if s.Kind() == reflect.Ptr {
	s = s.Elem()
}
switch s.Kind() {
	case reflect.Struct:
		tmp := make(map[string]*QVariant, s.NumField())
		for id := 0; id < s.NumField(); id++ {
			field := s.Field(id)
			if !field.CanInterface() {
				continue
			}
			if tag, ok := s.Type().Field(id).Tag.Lookup("json"); ok {
				switch {
				case (strings.HasSuffix(tag, ",omitempty") && isZero(field)) || tag == "-":
				case strings.Count(tag, ",") == 0:
					tmp[tag] = NewQVariant1(field.Interface())
				default:
					tags := strings.Split(tag, ",")
					name := s.Type().Field(id).Name
					if len(tags[0]) != 0 {
						name = tags[0]
					}
					v := NewQVariant1(field.Interface())
					if tags[1] == "string" {
						v = NewQVariant1(v.ToString()) //TODO: pure go type conversion
					}
					tmp[name] = v
				}
			} else {
				tmp[s.Type().Field(id).Name] = NewQVariant1(field.Interface())
			}
		}
		return NewQVariant1(tmp)

	case reflect.Slice, reflect.Array:
		tmp := make([]*QVariant, s.Len())
		for id := 0; id < s.Len(); id++ {
			if field := s.Index(id); field.CanInterface() {
				tmp[id] = NewQVariant1(field.Interface())
			}
		}
		return NewQVariant1(tmp)

	case reflect.Map:
		tmp := make(map[string]*QVariant, s.Len())
		for _, id := range s.MapKeys() {
			if field := s.MapIndex(id); field.CanInterface() {
				if key := NewQVariant1(id.Interface()); key.CanConvert(int(QMetaType__QString)) {
					tmp[key.ToString()] = NewQVariant1(field.Interface())
				}
			}
		}
		return NewQVariant1(tmp)

	case reflect.UnsafePointer:
		return NewQVariant1(uint64(uintptr(s.Interface().(unsafe.Pointer))))

	case reflect.Uintptr:
		return NewQVariant1(uint64(s.Interface().(uintptr)))
}

if i != nil && s.Type().ConvertibleTo(reflect.TypeOf(int8(0))) {
	if s.Kind() == reflect.Int64 {
		return NewQVariant1(s.Convert(reflect.TypeOf(int64(0))).Interface())
	}
	return NewQVariant1(s.Interface())
}

return NewQVariant()
}
}

func isZero(v reflect.Value) bool {
	switch v.Kind() {
	case reflect.Bool:
		return !v.Bool()
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return v.Int() == 0
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr:
		return v.Uint() == 0
	case reflect.Float32, reflect.Float64:
		return math.Float64bits(v.Float()) == 0
	case reflect.Complex64, reflect.Complex128:
		c := v.Complex()
		return math.Float64bits(real(c)) == 0 && math.Float64bits(imag(c)) == 0
	case reflect.Array:
		for i := 0; i < v.Len(); i++ {
			if !isZero(v.Index(i)) {
				return false
			}
		}
		return true
	case reflect.Chan, reflect.Func, reflect.Interface, reflect.Map, reflect.Ptr, reflect.Slice, reflect.UnsafePointer:
		return v.IsNil()
	case reflect.String:
		return v.Len() == 0
	case reflect.Struct:
		for i := 0; i < v.NumField(); i++ {
			if !isZero(v.Field(i)) {
				return false
			}
		}
		return true
	default:
		// This should never happens, but will act as a safeguard for
		// later, as a default value doesn't makes sense here.
		panic(&reflect.ValueError{"reflect.Value.IsZero", v.Kind()})
	}
}
`)

				//

				fmt.Fprint(bb, "func (ptr *QVariant) ToInterface() interface{} {\n")
				fmt.Fprint(bb, "switch ptr.Type() {\n")

				for _, v := range class.Enums[0].Values {
					if c, ok := parser.IsClass("Q" + v.Name); ok {
						if !parser.State.ClassMap[c].IsSupported() &&
							!(v.Name == "Map" ||
								v.Name == "String" ||
								v.Name == "StringList" ||
								v.Name == "Hash") {
							continue
						}
					}

					if f := class.GetFunction("to" + v.Name); f != nil && f.IsSupported() {
						fmt.Fprintf(bb, "case QVariant__%v:\n", v.Name)
						if len(f.Parameters) == 0 {
							fmt.Fprintf(bb, "return ptr.To%v()\n", v.Name)
						} else {
							fmt.Fprintf(bb, "return ptr.To%v(nil)\n", v.Name)
						}
					}
				}
				fmt.Fprint(bb, "\n}\nreturn ptr\n}\n")

				//

				fmt.Fprint(bb, `
func (ptr *QVariant) ToGoType(dst interface{}) {
v := reflect.ValueOf(dst)
if v.Kind() == reflect.Ptr {
	v = v.Elem()
}

switch ptr.Type() {
case QVariant__List:
	d := ptr.ToList()

	switch v.Kind() {
	case reflect.Slice:
		v.Set(reflect.MakeSlice(v.Type(), len(d), len(d)))

	case reflect.Array:
		v.Set(reflect.Indirect(reflect.New(v.Type())))
	}

	for i := 0; i < len(d); i++ {
		switch v.Type().Elem().Kind() {
		case reflect.Struct, reflect.Map, reflect.Slice, reflect.Array:
			s := reflect.New(v.Type().Elem())
			d[i].ToGoType(s.Interface())
			v.Index(i).Set(reflect.Indirect(s))

		default:
			v.Index(i).Set(reflect.ValueOf(d[i].ToInterface()).Convert(v.Type().Elem()))
		}
	}

case QVariant__Map:
	d := ptr.ToMap()

	if v.Kind() == reflect.Struct {
		for k, val := range d {
			realName := k
			toInt := false
			for id := 0; id < v.NumField(); id++ {
				if tag, ok := v.Type().Field(id).Tag.Lookup("json"); ok {
					switch {
					case strings.Count(tag, ",") == 0:
						if k == tag {
							realName = v.Type().Field(id).Name
						}
					default:
						tags := strings.Split(tag, ",")
						if k == tags[0] || (len(tags[0]) == 0 && k == v.Type().Field(id).Name) {
							realName = v.Type().Field(id).Name
							if tags[1] == "string" {
								toInt = true
							}
						}
					}
				}
			}
			field := v.FieldByName(realName)
			if !field.IsValid() {
				continue
			}
			if toInt {
				field.Set(reflect.ValueOf(val.ToLongLong(nil)).Convert(field.Type())) //TODO: pure go type conversion
			} else {
				val.ToGoType(field.Addr().Interface())
			}
		}
	} else {
		v.Set(reflect.MakeMapWithSize(v.Type(), len(d)))
		for k, val := range d {
			switch v.Type().Elem().Kind() {
			case reflect.Struct, reflect.Map, reflect.Slice, reflect.Array:
				s := reflect.New(v.Type().Elem())
				val.ToGoType(s.Interface())
				v.SetMapIndex(reflect.ValueOf(k).Convert(v.Type().Key()), reflect.Indirect(s))

			default:
				v.SetMapIndex(reflect.ValueOf(k).Convert(v.Type().Key()), reflect.ValueOf(val.ToInterface()).Convert(v.Type().Elem()))
			}
		}
	}

default:
	if v.Kind() == reflect.String {
		v.Set(reflect.ValueOf(ptr.ToString()))
	} else {
		v.Set(reflect.ValueOf(ptr.ToInterface()).Convert(v.Type()))
	}
}
}
`)

			} else if class.Name == "QObject" {
				fmt.Fprintln(bb, "\nfunc (ptr *QObject) ConnectSignal(f interface{}, a interface{}, t Qt__ConnectionType) {")
				fmt.Fprintln(bb, "\tfn := strings.TrimSuffix(strings.Split(runtime.FuncForPC(reflect.ValueOf(f).Pointer()).Name(), \".Connect\")[1], \"-fm\")")
				fmt.Fprintln(bb, "\tqt.RegisterConnectionType(ptr.Pointer(), strings.ToLower(fn[:1])+fn[1:], int64(t))")
				fmt.Fprintln(bb, "\treflect.ValueOf(f).Call([]reflect.Value{reflect.ValueOf(a)})")
				fmt.Fprintln(bb, "}\n")

				//TODO: invokeMethod
			}
		}
		cTemplate(bb, class, goEnum, goFunction, "\n\n", true)
	}

	if module == "QtQml" && !utils.QT_GEN_GO_WRAPPER() {
		fmt.Fprint(bb, "var (\n\thelper *core.QObject\n\thelperMutex sync.Mutex\n\thelperMap []string\n)\n")
	}

	fmt.Fprint(bb, "func init() {\n")

	if module == "QtQml" && !(utils.QT_STUB() || utils.QT_GEN_GO_WRAPPER()) {
		var free string
		if !UseJs() {
			free = `C.QJSValue_DestroyQJSValue(unsafe.Pointer(uintptr(ptr)))
			C.free(unsafe.Pointer(uintptr(ptr)))`
		} else {
			free = "qt.Module.Call(\"_QJSValue_DestroyQJSValue\", uintptr(ptr))"
		}
		fmt.Fprintf(bb, `
	helper = core.NewQObject(nil)
	helper.ConnectObjectNameChanged(func(pl string) {
		for _, p := range strings.Split(pl, "|") {
			ptr, err := strconv.ParseUint(p, 10, 64)
			if err != nil || ptr == 0 {
				return
			}
			%v
		}
	})
`, free) //TODO: replace with generic run on main thread helper function
	}

	if UseJs() {
		for _, l := range strings.Split(bb.String(), "\n") {
			if strings.HasPrefix(l, "//export") {
				if parser.UseWasm() {
					fmt.Fprintf(bb, "qt.Module.Set(\"_%[1]v\", js.FuncOf(%[1]v))\n", strings.TrimPrefix(l, "//export "))
				} else {
					fmt.Fprintf(bb, "qt.Module.Set(\"_%[1]v\", %[1]v)\n", strings.TrimPrefix(l, "//export "))
				}
			}
		}

		if parser.UseWasm() {
			//TODO:
		} else {
			fmt.Fprint(bb, "var module *js.Object\n")
			fmt.Fprintf(bb, "if m := js.Global.Get(\"%v\"); m == js.Undefined {\n", goModule(module))
			fmt.Fprint(bb, "\tmodule = new(js.Object)\n")
			fmt.Fprintf(bb, "\tjs.Global.Set(\"%v\", module)\n", goModule(module))
			fmt.Fprint(bb, "} else {\n")
			fmt.Fprint(bb, "\tmodule = m\n")
			fmt.Fprint(bb, "}\n")
		}
	}

	tsd := new(bytes.Buffer)
	defer tsd.Reset()
	var gosplitted []string

	switch {
	case utils.QT_GEN_TSD():
		gosplitted = strings.Split(bb.String(), "\n")

		fmt.Fprintf(tsd, "declare namespace %v {\n", goModule(func() string {
			if mode == MOC {
				return pkg
			}
			return module
		}()))

	case utils.QT_GEN_DART():
		gosplitted = strings.Split(bb.String(), "\n")

		fmt.Fprint(tsd, "import 'internal.dart';\n")

		fmt.Fprint(tsd, "bool inited = false;\nvoid initModule() {\nif (inited) { return; }\ninited = true;\n")
		for _, c := range parser.SortedClassesForModule(module, true) {
			if c.IsSupported() || (c.Export && mode == MINIMAL) {
				fmt.Fprintf(tsd, "constructorTable[\"%v.%v\"] = New%vFromPointer;\n", goModule(c.Module), c.Name, strings.Title(c.Name))
			}
		}
		fmt.Fprint(tsd, "\ninit();\n")
		fmt.Fprint(tsd, "\n___IMPORTPLACEHOLDER___\n}\n")

	case utils.QT_GEN_HAXE():
		gosplitted = strings.Split(bb.String(), "\n")

		fmt.Fprintf(tsd, "class %v {\n private static var inited = false;\npublic static function initModule() {\nif (inited) { return; }\ninited = true;\n", strings.TrimPrefix(module, "Qt"))
		for _, c := range parser.SortedClassesForModule(module, true) {
			if c.IsSupported() || (c.Export && mode == MINIMAL) {
				fmt.Fprintf(tsd, "Internal.constructorTable[\"%v.%v\"] = New%vFromPointer;\n", goModule(c.Module), c.Name, strings.Title(c.Name))
			}
		}
		fmt.Fprint(tsd, "\nInternal.init();\n")
		fmt.Fprint(tsd, "\n___IMPORTPLACEHOLDER___\n}\n}\n")

	case utils.QT_GEN_SWIFT():
		gosplitted = strings.Split(bb.String(), "\n")

		fmt.Fprintf(tsd, "class %v {\n private static var inited = false;\npublic static func initModule() {\nif (inited) { return; }\ninited = true;\n", strings.TrimPrefix(module, "Qt"))
		for _, c := range parser.SortedClassesForModule(module, true) {
			if c.IsSupported() || (c.Export && mode == MINIMAL) {
				fmt.Fprintf(tsd, "constructorTable[\"%v.%v\"] = New%vFromPointer;\n", goModule(c.Module), c.Name, strings.Title(c.Name))
			}
		}
		fmt.Fprint(tsd, "\nInit();\n")
		fmt.Fprint(tsd, "\n___IMPORTPLACEHOLDER___\n}\n}\n")
	}

	for _, c := range parser.SortedClassesForModule(module, true) {
		if utils.QT_GEN_HAXE() && strings.Title(c.Name) != c.Name {
			continue
		}

		if c.IsSupported() && utils.QT_GEN_GO_WRAPPER() && mode != MOC {
			bb.WriteString(fmt.Sprintf("internal.ConstructorTable[\"%v.%v\"] = New%vFromPointer\n", goModule(c.Module), c.Name, strings.Title(c.Name)))
		}

		if (c.IsSupported() && (cmd.ImportsQmlOrQuick() || cmd.ImportsInterop()) || mode == NONE) && !utils.QT_GEN_GO_WRAPPER() {
			bb.WriteString(fmt.Sprintf("qt.ItfMap[\"%[1]v.%[2]v_ITF\"] = %[2]v{}\n", goModule(func() string {
				if mode == MOC {
					return pkg
				}
				return module
			}()), c.Name))

			if utils.QT_GEN_TSD() || utils.QT_GEN_HAXE() || utils.QT_GEN_SWIFT() {

				var parents []string
				for _, parentClassName := range c.GetBases() {
					var parentClass, ok = parser.State.ClassMap[parentClassName]
					if !ok {
						continue
					}
					if parentClass.Module == c.Module || utils.QT_GEN_HAXE() || utils.QT_GEN_SWIFT() {
						parents = append(parents, fmt.Sprintf("%v_ITF", parentClassName))
					} else {
						parents = append(parents, fmt.Sprintf("%v.%v_ITF", goModule(parentClass.Module), parentClassName))
					}
				}

				if len(parents) != 0 {
					fmt.Fprintf(tsd, "\t%v %v_ITF %v %v {\n\t\t%v\t}\n\n",
						func() string {
							if utils.QT_GEN_SWIFT() {
								return "public protocol"
							}
							return "interface"
						}(),
						c.Name,
						func() string {
							if utils.QT_GEN_SWIFT() {
								return ":"
							}
							return "extends"
						}(),
						parents[len(parents)-1],
						func() string {
							o := fmt.Sprintf("%[1]v%[2]v_PTR()%[3]v%[2]v;\n",
								func() string {
									if utils.QT_GEN_HAXE() {
										return "public function "
									}
									if utils.QT_GEN_SWIFT() {
										return "func "
									}
									return ""
								}(),
								c.Name,
								func() string {
									if utils.QT_GEN_SWIFT() {
										return "->"
									}
									return ":"
								}())
							for i, v := range parents {
								if i == len(parents)-1 {
									continue
								}
								p := v
								if strings.Contains(v, ".") {
									p = strings.Split(v, ".")[1]
								}
								o += fmt.Sprintf("\t\t%v%v()%v%v;\n",
									func() string {
										if utils.QT_GEN_HAXE() {
											return "public function "
										}
										if utils.QT_GEN_SWIFT() {
											return "func "
										}
										return ""
									}(),
									strings.TrimSuffix(p, "_ITF")+"_PTR",
									func() string {
										if utils.QT_GEN_SWIFT() {
											return "->"
										}
										return ":"
									}(),
									strings.TrimSuffix(v, "_ITF"),
								)
							}
							return o
						}(),
					)
				} else {
					fmt.Fprintf(tsd, "\t%v %v_ITF {\n\t\t%v\t}\n\n",
						func() string {
							if utils.QT_GEN_SWIFT() {
								return "public protocol"
							}
							return "interface"
						}(),
						c.Name,
						fmt.Sprintf("%[1]v%[2]v_PTR()%[3]v%[2]v;\n",
							func() string {
								if utils.QT_GEN_HAXE() {
									return "public function "
								}
								if utils.QT_GEN_SWIFT() {
									return "func "
								}
								return ""
							}(),
							c.Name,
							func() string {
								if utils.QT_GEN_SWIFT() {
									return "->"
								}
								return ":"
							}(),
						),
					)
				}

				if len(parents) != 0 {
					fmt.Fprintf(tsd, "\t%vclass %v %v %v%v {\n",
						func() string {
							if utils.QT_GEN_SWIFT() {
								return "public "
							}
							return ""
						}(),
						c.Name,
						func() string {
							if utils.QT_GEN_SWIFT() {
								return ":"
							}
							return "extends"
						}(),
						strings.TrimSuffix(parents[len(parents)-1], "_ITF"),
						func() string {
							if utils.QT_GEN_HAXE() {
								return " implements " + c.Name + "_ITF"
							}
							if utils.QT_GEN_SWIFT() {
								return ", " + c.Name + "_ITF"
							}
							return ""
						}())

					for i, v := range parents {
						if i == len(parents)-1 {
							continue
						}
						p := v
						if strings.Contains(v, ".") {
							p = strings.Split(v, ".")[1]
						}
						fmt.Fprintf(tsd, "\t\t%v%v()%v%v%v\n",
							func() string {
								if utils.QT_GEN_HAXE() {
									return "public function "
								}
								if utils.QT_GEN_SWIFT() {
									return "public func "
								}
								return ""
							}(),
							strings.TrimSuffix(p, "_ITF")+"_PTR",
							func() string {
								if utils.QT_GEN_SWIFT() {
									return "->"
								}
								return ":"
							}(),
							strings.TrimSuffix(v, "_ITF"),
							func() string {
								if utils.QT_GEN_HAXE() {
									return fmt.Sprintf(" { return Internal.callLocalFunction([\"\", Pointer(), ___className, \"%v_PTR\"]); }", strings.TrimSuffix(p, "_ITF"))
								}
								if utils.QT_GEN_SWIFT() {
									return fmt.Sprintf(" { return callLocalFunction(l:[\"\", Pointer(), ___className, \"%[1]v_PTR\"]) as! %[1]v; }", strings.TrimSuffix(p, "_ITF"))
								}
								return ""
							}(),
						)
					}

				} else {
					if utils.QT_GEN_HAXE() {
						fmt.Fprintf(tsd, "\tclass %[1]v extends Internal implements %[1]v_ITF {\n", c.Name)
					} else if utils.QT_GEN_SWIFT() {
						fmt.Fprintf(tsd, "\tpublic class %[1]v : Internal, %[1]v_ITF {\n", c.Name)
					} else {
						fmt.Fprintf(tsd, "\tclass %v {\n", c.Name)
					}
				}

				if utils.QT_GEN_HAXE() {
					tsd.WriteString("\t\tpublic function new() { super(); }\n")
				} else if utils.QT_GEN_TSD() {
					tsd.WriteString("\t\t___pointer: number;\n")
				}

				if c.Name == "QJSEngine" {
					if utils.QT_GEN_HAXE() {
						//TODO: also NewQVariant1 + toInterface function
						fmt.Fprint(tsd, "\t\tpublic function NewGoType(i:Array<Dynamic>):QJSValue{ return Internal.callLocalFunction([\"\",Pointer(),___className,\"NewGoType\",i]); }\n") //TODO: embeed handwritten code
						//fmt.Fprint(tsd, "\t\tToGoType(jsval:QJSValue, dst:any):void\n")                                                                                                                 //TODO: possible with reflection
						//fmt.Fprint(tsd, "\t\tvoid NewJSType(QJSValue property, String name, dynamic i) { callLocalFunction([\"\",this.Pointer(),this.className,\"NewJSType\",property, name, i]); }\n") //TODO: only usefull inside qml?; check if one of the args is a function
					} else if utils.QT_GEN_SWIFT() {
						//TODO: also NewQVariant1 + toInterface function
						fmt.Fprint(tsd, "\t\tpublic func NewGoType(i:[Any])->QJSValue{ return callLocalFunction(l:[\"\",Pointer(),___className,\"NewGoType\",i]) as! QJSValue; };\n") //TODO: embeed handwritten code
						//fmt.Fprint(tsd, "\t\tToGoType(jsval:QJSValue, dst:any):void\n")                                                                                                                 //TODO: possible with reflection
						//fmt.Fprint(tsd, "\t\tvoid NewJSType(QJSValue property, String name, dynamic i) { callLocalFunction([\"\",this.Pointer(),this.className,\"NewJSType\",property, name, i]); }\n") //TODO: only usefull inside qml?; check if one of the args is a function
					} else {
						//TODO: also NewQVariant1 + toInterface function
						//fmt.Fprint(tsd, "\t\tNewGoType(i:any[]):QJSValue\n")            //TODO: only possible for the webbrowser js api?
						//fmt.Fprint(tsd, "\t\tToGoType(jsval:QJSValue, dst:any):void\n") //TODO: only possible for the webbrowser js api?
						fmt.Fprint(tsd, "\t\tNewJSType(property:QJSValue, name:string, i:any):void\n")
					}
				}

				for _, l := range gosplitted {
					var out string
					if utils.QT_GEN_TSD() {
						out = convertToTypeScriptDefinition(c.Name, l, true)
					} else if utils.QT_GEN_HAXE() {
						out = convertToHaxe(c.Name, l, true)
					} else {
						out = convertToSwift(c.Name, l, true)
					}
					if out != "" {
						tsd.WriteString("\t\t" + out + ";\n")
					}
				}

				fmt.Fprint(tsd, "\t}\n")

			} else if utils.QT_GEN_DART() {

				var parents []string
				for _, parentClassName := range c.GetBasesSorted() {
					var parentClass, ok = parser.State.ClassMap[parentClassName]
					if !ok {
						continue
					}
					if parentClass.Module == c.Module {
						parents = append(parents, fmt.Sprintf("%v_ITF", parentClassName))
					} else {
						parents = append(parents, fmt.Sprintf("%v.%v_ITF", goModule(parentClass.Module), parentClassName))
					}
				}

				if len(parents) != 0 && !(strings.HasPrefix(strings.Title(c.Name), "Generic") && c.Module == "QtSensors") {

					var with string
					if len(parents) > 1 {
						with = " with "
						for i, p := range parents {
							if i == 0 {
								continue
							}
							with += p + ", "
						}
						with = strings.TrimSuffix(with, ", ")
					}

					fmt.Fprintf(tsd, "\tabstract class %v_ITF extends %v%v {\n\t\t%v\t}\n\n",
						c.Name,
						parents[0],
						with,
						fmt.Sprintf("%[1]v %[1]v_PTR();\n", c.Name),
					)
				} else {
					fmt.Fprintf(tsd, "\tabstract class %v_ITF {\n\t\t%v\t}\n\n",
						c.Name,
						fmt.Sprintf("%[1]v %[1]v_PTR();\n", c.Name),
					)
				}

				if len(parents) != 0 && !(strings.HasPrefix(strings.Title(c.Name), "Generic") && c.Module == "QtSensors") {

					var with string
					if len(parents) > 1 {
						with = " with "
						for i, p := range parents {
							if i == 0 {
								continue
							}
							p = strings.TrimSuffix(p, "_ITF")
							with += p + ", "
						}
						with = strings.TrimSuffix(with, ", ")
					}

					fmt.Fprintf(tsd, "\tclass %[1]v extends %v%v implements %[1]v_ITF {\n",
						c.Name,
						strings.TrimSuffix(parents[0], "_ITF"),
						with,
					)
				} else {
					fmt.Fprintf(tsd, "\tclass %[1]v extends Internal implements %[1]v_ITF {\n",
						c.Name,
					)
				}

				if c.Name == "QJSEngine" { //TODO: also NewQVariant1 + toInterface function
					fmt.Fprint(tsd, "\t\tQJSValue NewGoType(List<dynamic> i) { return callLocalFunction([\"\",this.Pointer(),this.className,\"NewGoType\",i]); }\n") //TODO: check if one of the args is a function
					//fmt.Fprint(tsd, "\t\tToGoType(jsval:QJSValue, dst:any):void\n")                                                                                                                 //TODO: only possible with reflection in dart
					//fmt.Fprint(tsd, "\t\tvoid NewJSType(QJSValue property, String name, dynamic i) { callLocalFunction([\"\",this.Pointer(),this.className,\"NewJSType\",property, name, i]); }\n") //TODO: check if one of the args is a function
				}

				for _, l := range gosplitted {
					out := convertToDart(c.Name, l, true)
					if out != "" {
						tsd.WriteString("\t\t" + out + "\n")
					}
				}

				fmt.Fprint(tsd, "\t}\n")
			}
		}

		implemented := make(map[string]struct{})
		if !utils.QT_GEN_GO_WRAPPER() {
			for _, f := range c.Functions {
				if f.Meta != parser.CONSTRUCTOR && !f.Static {
					continue
				}
				if strings.Contains(f.Name, "RegisterMetaType") || strings.Contains(f.Name, "RegisterType") ||
					strings.Contains(f.Name, "QmlRegisterUncreatableType") || strings.Contains(f.Name, "QmlRegisterAnonymousType") { //TODO:
					continue
				}
				var _, e = implemented[fmt.Sprint(f.Name, f.OverloadNumber)]
				if e || !f.IsSupported() {
					continue
				}
				implemented[fmt.Sprint(f.Name, f.OverloadNumber)] = struct{}{}

				if parser.UseJs() {
					var ip string
					oldsm := f.SignalMode
					f.SignalMode = parser.CALLBACK
					f.FakeForJSCallback = true
					ip = converter.GoHeaderInput(f)
					ip = strings.TrimPrefix(ip, "ptr uintptr, ")
					f.SignalMode = oldsm
					f.FakeForJSCallback = false
					var out string
					if parser.UseWasm() {
						out = "" //TODO: export classes for jsinterop example
					} else {
						if converter.GoHeaderOutput(f) != "" {
							out = fmt.Sprintf("module.Set(\"%v\", func(%v) *js.Object { return qt.MakeWrapper(%v(%v)); })\n", converter.GoHeaderName(f), ip, converter.GoHeaderName(f), converter.GoInputParametersForCallback(f))
						} else {
							out = fmt.Sprintf("module.Set(\"%v\", func(%v) { %v(%v); })\n", converter.GoHeaderName(f), ip, converter.GoHeaderName(f), converter.GoInputParametersForCallback(f))
						}
					}
					if !strings.Contains(out, "unsupported_") && !strings.Contains(out, "C.") && strings.Contains(bb.String(), converter.GoHeaderName(f)+"(") {
						bb.WriteString(out)
					}
				}
				if cmd.ImportsQmlOrQuick() || cmd.ImportsInterop() || mode == NONE {
					out := fmt.Sprintf("qt.FuncMap[\"%[1]v.%[2]v\"] = %[2]v\n", goModule(func() string {
						if mode == MOC {
							return pkg
						}
						return module
					}()), converter.GoHeaderName(f))
					if !strings.Contains(out, "unsupported_") && strings.Contains(bb.String(), converter.GoHeaderName(f)+"(") {
						bb.WriteString(out)
					}
				}
			}

			if c.Name == "QVariant" {
				if parser.UseJs() {
					if parser.UseWasm() {
						//TODO:
					} else {
						fmt.Fprint(bb, "module.Set(\"NewQVariant1\", func(i interface{}) *js.Object { return qt.MakeWrapper(NewQVariant1(i)) })\n")
					}
				}
				if cmd.ImportsQmlOrQuick() || cmd.ImportsInterop() || mode == NONE {
					fmt.Fprint(bb, "qt.FuncMap[\"core.NewQVariant1\"] = NewQVariant1\n")
				}
			}
		}

		if (utils.QT_GEN_TSD() || utils.QT_GEN_DART() || utils.QT_GEN_HAXE() || utils.QT_GEN_SWIFT()) && c.IsSupported() {
			for _, l := range gosplitted {
				var out string
				if utils.QT_GEN_TSD() {
					out = convertToTypeScriptDefinition(c.Name, l, false)
					if out != "" {
						tsd.WriteString("\t" + out + ";\n")
					}
				} else if utils.QT_GEN_DART() {
					out = convertToDart(c.Name, l, false)
					if out != "" {
						tsd.WriteString("\t" + out + "\n")
					}
				} else if utils.QT_GEN_HAXE() {
					out = convertToHaxe(c.Name, l, false)
					if out != "" {
						tsd.WriteString("\t" + out + "\n")
					}
				} else {
					out = convertToSwift(c.Name, l, false)
					if out != "" {
						tsd.WriteString("\t" + out + "\n")
					}
				}
			}
		}

		if !utils.QT_GEN_GO_WRAPPER() {
			for _, e := range c.Enums {
				if e.ClassName() == "QColorSpace" { //TODO: 5.14.0
					continue
				}
				for _, v := range e.Values {
					if v.Name == "ByteOrder" {
						continue
					}
					if parser.UseWasm() {
						//TODO: same as for js ?
					} else if parser.UseJs() {
						fmt.Fprintf(bb, "module.Set(\"%v__%v\", int64(%v__%v))\n", strings.Split(e.Fullname, "::")[0], v.Name, strings.Split(e.Fullname, "::")[0], v.Name)
					}
					if cmd.ImportsQmlOrQuick() || cmd.ImportsInterop() || mode == NONE {
						bb.WriteString(fmt.Sprintf("qt.EnumMap[\"%[1]v.%[2]v__%[3]v\"] = int64(%[2]v__%[3]v)\n", goModule(module), strings.Split(e.Fullname, "::")[0], v.Name))

						if utils.QT_GEN_TSD() {
							fmt.Fprintf(tsd, "\tconst %v__%v: number;\n", strings.Split(e.Fullname, "::")[0], v.Name)
						}
					}
				}
			}
		}
	}
	if utils.QT_GEN_TSD() {
		tsd.WriteString("}")
	}

	fmt.Fprint(bb, "}\n")

	return preambleGo(module, goModule(module), bb.Bytes(), stub, mode, pkg, target, tags, tsd)
}

func preambleGo(oldModule string, module string, input []byte, stub bool, mode int, pkg, target, tags string, tsd *bytes.Buffer) []byte {
	bb := new(bytes.Buffer)
	defer bb.Reset()
	tsdbb := new(bytes.Buffer)
	defer tsdbb.Reset()

	if UseStub(stub, oldModule, mode) || UseJs() || utils.QT_GEN_GO_WRAPPER() {
		fmt.Fprintf(bb, `%v

package %v
`, buildTags(oldModule, stub, mode, tags),

			func() string {
				if mode == MOC {
					return pkg
				}
				return module
			}(),
		)

	} else {
		fmt.Fprintf(bb, `%v

package %v

//#include <stdint.h>
//#include <stdlib.h>
//#include <string.h>
//#include "%v.h"
import "C"
`,

			buildTags(oldModule, stub, mode, tags),

			func() string {
				if mode == MOC {
					return pkg
				}
				return module
			}(),

			func() string {
				switch module {
				case "androidextras":
					{
						return fmt.Sprintf("%v_android", module)
					}

				case "sailfish":
					{
						return fmt.Sprintf("%v_sailfish", module)
					}

				default:
					{
						if mode == MINIMAL {
							return fmt.Sprintf("%v-minimal", module)
						}

						if mode == MOC {
							return "moc"
						}

						return module
					}
				}
			}(),
		)
	}
	if utils.QT_GEN_HAXE() {
		fmt.Fprint(tsdbb, "package qt;\n")
	}

	inputString := string(input)
	if mode == MOC {
		for _, lib := range parser.GetLibs() {
			mlow := strings.ToLower(lib)
			for _, pre := range []string{" ", "=", "\t", "\r", "\n", "!", "*", "(", ")", "[", "]"} { //TODO: "=" shouldn't be needed, but there seemse to be some issue in the fmt or strings pkg in Go 1.14 (at least when the module mode is enabled); which causes the internal/runtime/moc.go file to come out invalid
				for _, past := range []string{"NewQ", "PointerFromQ", "Q"} {
					inputString = strings.Replace(inputString, fmt.Sprintf("%v%v.%v", pre, mlow, past), fmt.Sprintf("%vstd_%v.%v", pre, mlow, past), -1)
				}
			}
		}
	}

	var dartInput []string
	fmt.Fprint(bb, "import (\n")
	for _, m := range append(parser.GetLibs(), "qt", "strings", "unsafe", "log", "runtime", "fmt", "errors", "js", "time", "hex", "reflect", "math", "sync", "strconv", "internal", "gow") {
		mlow := strings.ToLower(m)
		if strings.Contains(inputString, fmt.Sprintf(" %v.", mlow)) ||
			strings.Contains(inputString, fmt.Sprintf("\t%v.", mlow)) ||
			strings.Contains(inputString, fmt.Sprintf("\r%v.", mlow)) ||
			strings.Contains(inputString, fmt.Sprintf("\n%v.", mlow)) ||
			strings.Contains(inputString, fmt.Sprintf("!%v.", mlow)) ||
			strings.Contains(inputString, fmt.Sprintf("*%v.", mlow)) ||
			strings.Contains(inputString, fmt.Sprintf("(%v.", mlow)) ||
			strings.Contains(inputString, fmt.Sprintf(")%v.", mlow)) ||
			strings.Contains(inputString, fmt.Sprintf("std_%v.", mlow)) {
			switch mlow {
			case "strings", "unsafe", "log", "runtime", "fmt", "errors", "time", "reflect", "math", "sync", "strconv":
				fmt.Fprintf(bb, "\"%v\"\n", mlow)

			case "hex":
				fmt.Fprintln(bb, "\"encoding/hex\"")

			case "qt":
				fmt.Fprintln(bb, "\"github.com/therecipe/qt\"")

			case "internal":
				fmt.Fprintln(bb, "\"github.com/therecipe/qt/internal\"")

			case "gow":
				fmt.Fprintln(bb, "\"github.com/therecipe/qt/interop/gow\"")

			case "js":
				if parser.UseWasm() {
					fmt.Fprintln(bb, "\"syscall/js\"")
				} else {
					fmt.Fprintln(bb, "\"github.com/gopherjs/gopherjs/js\"")
				}

			default:
				if mode == MOC {
					fmt.Fprintf(bb, "std_%[1]v \"github.com/therecipe/qt/%[1]v\"\n", mlow)
				} else {
					fmt.Fprintf(bb, "\"github.com/therecipe/qt/%v\"\n", mlow)
				}

				if utils.QT_GEN_TSD() {
					fmt.Fprintf(tsdbb, "/// <reference path=\"%v.d.ts\" />\n", mlow)
				} else if utils.QT_GEN_DART() {
					fmt.Fprintf(tsdbb, "import '%[1]v.dart' as %[1]v;\n", mlow)
					dartInput = append(dartInput, fmt.Sprintf("%v.initModule();", mlow))
				} else if utils.QT_GEN_HAXE() {
					fmt.Fprintf(tsdbb, "import qt.%v;\n", m)
					dartInput = append(dartInput, fmt.Sprintf("%v.initModule();", m))
				} else if utils.QT_GEN_SWIFT() {
					dartInput = append(dartInput, fmt.Sprintf("%v.initModule();", m))
				}

				if mode == MOC {
					parser.LibDeps[parser.MOC] = append(parser.LibDeps[parser.MOC], m)
				}

				//TODO: REVIEW
				if !UseJs() {
					if strings.HasPrefix(target, "ios") && mode == MINIMAL {
						oldModuleGo := strings.TrimPrefix(oldModule, "Qt")

						var (
							containsSub  bool
							containsSelf bool
						)

						for _, l := range parser.LibDeps["build_static"] {
							if l == m {
								containsSub = true
							}
							if l == oldModuleGo {
								containsSelf = true
							}
						}

						if !containsSelf || !containsSub {

							if !containsSelf {
								parser.LibDeps["build_static"] = append(parser.LibDeps["build_static"], oldModuleGo)

								switch oldModuleGo {
								case "Multimedia":
									parser.LibDeps["build_static"] = append(parser.LibDeps["build_static"], "MultimediaWidgets")
								case "Quick":
									parser.LibDeps["build_static"] = append(parser.LibDeps["build_static"], "QuickWidgets")
								}
							}

							if !containsSub {
								parser.LibDeps["build_static"] = append(parser.LibDeps["build_static"], m)

								switch m {
								case "Multimedia":
									parser.LibDeps["build_static"] = append(parser.LibDeps["build_static"], "MultimediaWidgets")
								case "Quick":
									parser.LibDeps["build_static"] = append(parser.LibDeps["build_static"], "QuickWidgets")
								}
							}

						}
					}
				}
				//TODO: REVIEW
			}
		}

		if module == "gui" && utils.QT_API_NUM(utils.QT_VERSION()) >= 5050 && !utils.QT_GEN_GO_WRAPPER() {
			if mode == NONE {
				fmt.Fprintln(bb, "_ \"github.com/therecipe/qt/internal/binding/runtime\"")
			} else if mode == MINIMAL && (cmd.ImportsQmlOrQuick() || cmd.ImportsInterop()) {
				fmt.Fprintln(bb, "_ \"github.com/therecipe/qt/internal/binding/runtime\"")
			}
		}
	}

	if mode == MOC {
		env, tagsEnv, _, _ := cmd.BuildEnv(target, "", "")
		if tags != "" {
			tagsEnv = append(tagsEnv, strings.Split(tags, " ")...)
		}
		for custom, m := range parser.GetCustomLibs(target, env, tagsEnv) {
			switch {
			case strings.Contains(m, "/vendor/") && strings.Contains(inputString, fmt.Sprintf("%v.", custom)):
				fmt.Fprintf(bb, "\"%v\"\n", custom)

			case strings.Contains(inputString, fmt.Sprintf("%v.", custom)):
				fmt.Fprintf(bb, "%v \"%v\"\n", custom, m)
			}
		}

		for i := range parser.State.MocImports {
			if strings.HasPrefix(i, "\"") { //TODO: golang.org/x/tools/imports seems to be buggy; advanced/qml_quick/tableview fails in module mode
				if strings.Contains(inputString, fmt.Sprintf("%v.", strings.Replace(i, "\"", "", -1))) {
					fmt.Fprintf(bb, "%v\n", i)
				}
			} else {
				fmt.Fprintf(bb, "%v\n", i)
			}

			if strings.HasPrefix(i, ".") {
				delete(parser.State.MocImports, i)
			}
		}
	}

	fmt.Fprintln(bb, ")")

	bb.WriteString(inputString)

	out, err := format.Source(renameSubClasses(bb.Bytes()))
	if err != nil {
		utils.Log.WithError(err).Errorln("failed to format:", pkg, module)
		out = bb.Bytes()
	}

	//TODO: regexp
	if mode == MOC {
		pre := string(out)
		libsm := make(map[string]struct{})
		for _, c := range parser.State.ClassMap {
			if c.Pkg != "" && c.IsSubClassOfQObject() {
				libsm[c.Module] = struct{}{}
			}
		}

		var libs []string
		for k := range libsm {
			libs = append(libs, k)
		}
		libs = append(libs, module)

		for _, c := range parser.SortedClassesForModule(strings.Join(libs, ","), true) {
			hName := c.Hash()
			sep := []string{"\"_", "\n", "(", "_", "callback", "C."}
			for _, p := range sep {
				for _, s := range sep {
					if s == "callback" || s == "C." || (p == "_" && s == "(" && UseJs()) {
						continue
					}
					pre = strings.Replace(pre, p+c.Name+s, p+c.Name+hName+s, -1)
				}
			}
		}
		out = []byte(pre)
	}

	if (utils.QT_GEN_TSD() || utils.QT_GEN_DART() || utils.QT_GEN_HAXE() || utils.QT_GEN_SWIFT()) && (mode == NONE || mode == MINIMAL) {
		tsdBytes := tsd.Bytes()
		if utils.QT_GEN_DART() || utils.QT_GEN_HAXE() || utils.QT_GEN_SWIFT() {
			tsdBytes = bytes.Replace(tsdBytes, []byte("\n___IMPORTPLACEHOLDER___"), []byte(strings.Join(dartInput, "\n")), -1)
		}
		tsdbb.Write(tsdBytes)

		var fn string
		switch {
		case utils.QT_GEN_TSD():
			fn = module + ".d.ts"
		case utils.QT_GEN_DART():
			fn = module + ".dart"
		case utils.QT_GEN_HAXE():
			fn = strings.TrimPrefix(oldModule, "Qt") + ".hx"
		case utils.QT_GEN_SWIFT():
			fn = module + ".swift"
		}
		utils.SaveBytes(utils.GoQtPkgPath(module, fn), tsdbb.Bytes())
	}

	return out
}

//TODO: regexp
func renameSubClasses(in []byte) []byte {
	for _, c := range parser.State.ClassMap {
		if c.Fullname != "" {
			sep := []string{"\n", ".", "\"", " ", "*", "(", ")", "{", "C.", "_ITF", "_PTR", " New", ".New", "(New", "\"New", "From", "Destroy", ",", "2"}
			for _, p := range sep {
				for _, s := range sep {
					in = bytes.Replace(in, []byte(p+c.Name+s), []byte(p+strings.Replace(c.Fullname, "::", "_", -1)+s), -1)
				}
			}
		}
	}
	return in
}
