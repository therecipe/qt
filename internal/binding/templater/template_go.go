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

	if module == "AndroidExtras" {
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
	}

	utils.Log.WithField("module", module).Debug("generating go")
	parser.State.Target = target

	bb := new(bytes.Buffer)
	defer bb.Reset()

	if mode != MOC {
		module = "Qt" + module
	}

	if !(UseStub(stub, module, mode) || UseJs()) {
		fmt.Fprintf(bb, "func cGoUnpackString(s C.struct_%v_PackedString) string { if int(s.len) == -1 {\n return C.GoString(s.data)\n }\n return C.GoStringN(s.data, C.int(s.len)) }\n", strings.Title(module))
		fmt.Fprintf(bb, "func cGoUnpackBytes(s C.struct_%v_PackedString) []byte { if int(s.len) == -1 {\n gs := C.GoString(s.data)\n return *(*[]byte)(unsafe.Pointer(&gs))\n }\n return C.GoBytes(unsafe.Pointer(s.data), C.int(s.len)) }\n", strings.Title(module))
	}

	if UseJs() {
		fmt.Fprint(bb, "func jsGoUnpackString(s string) string { dec, _ := hex.DecodeString(s)\n return string(dec)\n }\n") //TODO: calling it cGoUnpackString won't work, bug in go wasm ?
		fmt.Fprint(bb, "func jsGoUnpackBytes(s string) []byte { dec, _ := hex.DecodeString(s)\n return dec\n }\n")
	}

	fmt.Fprint(bb, "func unpackStringList(s string) []string {\nif len(s) == 0 {\nreturn make([]string, 0)\n}\nreturn strings.Split(s, \"¡¦!\")\n}\n")

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
				fmt.Fprintf(bb, "func (ptr *%v) Pointer() unsafe.Pointer {\nif ptr != nil {\nreturn ptr.ptr\n}\nreturn nil\n}\n\n", class.Name)
				fmt.Fprintf(bb, "func (ptr *%v) SetPointer(p unsafe.Pointer) {\nif ptr != nil {\nptr.ptr = p\n}\n}\n\n", class.Name)
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
func PointerFrom%v(ptr %[2]v_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.%[2]v_PTR().Pointer()
	}
	return nil
}
`, strings.Title(class.Name), class.Name)

			if class.Module == parser.MOC {
				fmt.Fprintf(bb, `
func New%vFromPointer(ptr unsafe.Pointer) (n *%[2]v) {
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
	}
	return
}
`, strings.Title(class.Name), class.Name,
					func() string {
						bc := class.GetBases()[0]
						if class.Module == parser.State.ClassMap[bc].Module {
							return bc
						}
						return fmt.Sprintf("%v.%v", strings.ToLower(strings.TrimPrefix(parser.State.ClassMap[bc].Module, "Qt")), bc)
					}(), class.GetBases()[0])
			} else {
				fmt.Fprintf(bb, `
func New%vFromPointer(ptr unsafe.Pointer) (n *%[2]v) {
	n = new(%[2]v)
	n.SetPointer(ptr)
	return
}
`, strings.Title(class.Name), class.Name)
			}

			if !class.HasDestructor() {
				if UseStub(stub, module, mode) {
					fmt.Fprintf(bb, "\nfunc (ptr *%v) Destroy%v() {}\n\n", class.Name, strings.Title(class.Name))
				} else if !class.IsSubClassOfQObject() {
					if UseJs() { //TODO: free c ptr in js/wasm ?
						fmt.Fprintf(bb, `
func (ptr *%[1]v) Destroy%[1]v() {
	if ptr != nil {
		%v
		ptr.SetPointer(nil)
		runtime.SetFinalizer(ptr, nil)
	}
}

`, class.Name, func() string {
							if class.HasCallbackFunctions() {
								return "\nqt.DisconnectAllSignals(ptr.Pointer(), \"\")"
							}
							return ""
						}())
					} else {
						fmt.Fprintf(bb, `
func (ptr *%[1]v) Destroy%[1]v() {
	if ptr != nil {
		%v
		C.free(ptr.Pointer())
		ptr.SetPointer(nil)
		runtime.SetFinalizer(ptr, nil)
	}
}

`, class.Name, func() string {
							if class.HasCallbackFunctions() {
								return "\nqt.DisconnectAllSignals(ptr.Pointer(), \"\")"
							}
							return ""
						}())
					}
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
					fmt.Fprintf(bb, "//export callback%[1]v_Constructor\nfunc callback%[1]v_Constructor(ptr unsafe.Pointer) {", class.Name)
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

				if UseJs() {
					if parser.UseWasm() {
						bb.WriteString("\nreturn nil\n")
					}
				}

				fmt.Fprint(bb, "}\n\n")
			}

			if class.Name == "QVariant" {
				fmt.Fprint(bb, "type qVariant_ITF interface { ToVariant() *QVariant }\n")
				fmt.Fprint(bb, "func NewQVariant1(i interface{}) *QVariant {\n")
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

				fmt.Fprint(bb, "func (v *QVariant) ToInterface() interface{} {\n")
				fmt.Fprint(bb, "switch v.Type() {\n")

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
							fmt.Fprintf(bb, "return v.To%v()\n", v.Name)
						} else {
							fmt.Fprintf(bb, "return v.To%v(nil)\n", v.Name)
						}
					}
				}
				fmt.Fprint(bb, "\n}\nreturn v\n}\n")

				//

				fmt.Fprint(bb, `
func (src *QVariant) ToGoType(dst interface{}) {
v := reflect.ValueOf(dst)
if v.Kind() == reflect.Ptr {
	v = v.Elem()
}

switch src.Type() {
case QVariant__List:
	d := src.ToList()

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
	d := src.ToMap()

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
	v.Set(reflect.ValueOf(src.ToInterface()).Convert(v.Type()))
}
}
`)

			} else if class.Name == "QObject" {
				fmt.Fprintln(bb, "\nfunc (ptr *QObject) ConnectSignal(f, a interface{}, t Qt__ConnectionType) {")
				fmt.Fprintln(bb, "\tfn := strings.TrimSuffix(strings.Split(runtime.FuncForPC(reflect.ValueOf(f).Pointer()).Name(), \".Connect\")[1], \"-fm\")")
				fmt.Fprintln(bb, "\tqt.RegisterConnectionType(ptr.Pointer(), strings.ToLower(fn[:1])+fn[1:], int64(t))")
				fmt.Fprintln(bb, "\treflect.ValueOf(f).Call([]reflect.Value{reflect.ValueOf(a)})")
				fmt.Fprintln(bb, "}\n")
			}
		}
		cTemplate(bb, class, goEnum, goFunction, "\n\n", true)
	}

	fmt.Fprint(bb, "func init() {\n")
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

	for _, c := range parser.SortedClassesForModule(module, true) {

		if !parser.UseJs() {
			if c.IsSupported() {
				bb.WriteString(fmt.Sprintf("qt.ItfMap[\"%[1]v.%[2]v_ITF\"] = %[2]v{}\n", goModule(func() string {
					if mode == MOC {
						return pkg
					}
					return module
				}()), c.Name))
			}
		}

		implemented := make(map[string]struct{})
		for _, f := range c.Functions {
			if f.Meta != parser.CONSTRUCTOR && !f.Static {
				continue
			}
			if strings.Contains(f.Name, "RegisterMetaType") || strings.Contains(f.Name, "RegisterType") { //TODO:
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
			} else {
				out := fmt.Sprintf("qt.FuncMap[\"%[1]v.%[2]v\"] = %[2]v\n", goModule(func() string {
					if mode == MOC {
						return pkg
					}
					return module
				}()), converter.GoHeaderName(f)) //TODO: use setter and getter
				if !strings.Contains(out, "unsupported_") && strings.Contains(bb.String(), converter.GoHeaderName(f)+"(") {
					bb.WriteString(out)
				}
			}
		}

		for _, e := range c.Enums {
			for _, v := range e.Values {
				if v.Name == "ByteOrder" {
					continue
				}
				if parser.UseWasm() {
					//TODO: same as for js ?
				} else if parser.UseJs() {
					fmt.Fprintf(bb, "module.Set(\"%v__%v\", int64(%v__%v))\n", strings.Split(e.Fullname, "::")[0], v.Name, strings.Split(e.Fullname, "::")[0], v.Name)
				} else {
					bb.WriteString(fmt.Sprintf("qt.EnumMap[\"%[1]v.%[2]v__%[3]v\"] = int64(%[2]v__%[3]v)\n", goModule(module), strings.Split(e.Fullname, "::")[0], v.Name))
				}
			}
		}
	}

	fmt.Fprint(bb, "}\n")

	return preambleGo(module, goModule(module), bb.Bytes(), stub, mode, pkg, target, tags)
}

func preambleGo(oldModule string, module string, input []byte, stub bool, mode int, pkg, target, tags string) []byte {
	var bb = new(bytes.Buffer)
	defer bb.Reset()

	if UseStub(stub, oldModule, mode) || UseJs() {
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

	inputString := string(input)
	if mode == MOC {
		for _, lib := range parser.GetLibs() {
			mlow := strings.ToLower(lib)
			for _, pre := range []string{" ", "\t", "\r", "\n", "!", "*", "(", ")", "[", "]"} {
				for _, past := range []string{"NewQ", "PointerFromQ", "Q"} {
					inputString = strings.Replace(inputString, fmt.Sprintf("%v%v.%v", pre, mlow, past), fmt.Sprintf("%vstd_%v.%v", pre, mlow, past), -1)
				}
			}
		}
	}

	fmt.Fprint(bb, "import (\n")
	for _, m := range append(parser.GetLibs(), "qt", "strings", "unsafe", "log", "runtime", "fmt", "errors", "js", "time", "hex", "reflect", "math") {
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
			case "strings", "unsafe", "log", "runtime", "fmt", "errors", "time", "reflect", "math":
				fmt.Fprintf(bb, "\"%v\"\n", mlow)

			case "hex":
				fmt.Fprintln(bb, "\"encoding/hex\"")

			case "qt":
				fmt.Fprintln(bb, "\"github.com/therecipe/qt\"")

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

		if module == "gui" && utils.QT_API_NUM(utils.QT_VERSION()) >= 5050 {
			if mode == NONE {
				fmt.Fprintln(bb, "_ \"github.com/therecipe/qt/internal/binding/runtime\"")
			} else if mode == MINIMAL && cmd.ImportsQmlOrQuick() {
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
			case strings.Contains(m, "/vendor/"):
				fmt.Fprintf(bb, "\"%v\"\n", custom)

			case strings.Contains(inputString, fmt.Sprintf("%v.", custom)):
				fmt.Fprintf(bb, "%v \"%v\"\n", custom, m)
			}
		}

		for i := range parser.State.MocImports {
			fmt.Fprintf(bb, "%v\n", i)

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
