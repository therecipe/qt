package templater

import (
	"bytes"
	"fmt"
	"go/format"
	"strings"

	"github.com/therecipe/qt/internal/binding/parser"
	"github.com/therecipe/qt/internal/utils"
)

func GoTemplate(module string, stub bool, mode int, pkg, target, tags string) []byte {
	utils.Log.WithField("module", module).Debug("generating go")
	parser.State.Target = target

	var bb = new(bytes.Buffer)
	defer bb.Reset()

	if mode != MOC {
		module = "Qt" + module
	}

	if !UseStub(stub, module, mode) {
		fmt.Fprintf(bb, "func cGoUnpackString(s C.struct_%v_PackedString) string { if int(s.len) == -1 {\n return C.GoString(s.data)\n }\n return C.GoStringN(s.data, C.int(s.len)) }\n", strings.Title(module))
	}

	if module == "QtAndroidExtras" && utils.QT_VERSION_NUM() >= 5060 {
		fmt.Fprint(bb, "func QAndroidJniEnvironment_ExceptionCatch() error {\n")
		if UseStub(stub, module, mode) {
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

		if UseStub(stub, module, mode) {
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
					fmt.Fprintf(bb, `
func (ptr *%[1]v) Destroy%[1]v() {
	if ptr != nil {
		C.free(ptr.Pointer())%v
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

			if mode == MOC {
				fmt.Fprintf(bb, `
//export callback%[1]v_Constructor
func callback%[1]v_Constructor(ptr unsafe.Pointer) {
`, class.Name)

				fmt.Fprintf(bb, "this := New%vFromPointer(ptr)\nqt.Register(ptr, this)\n", strings.Title(class.Name))

				var lastModule string
				for _, bcn := range class.GetAllBases() {
					if bc := parser.State.ClassMap[bcn]; bc.Module != class.Module {
						if len(bc.Constructors) > 0 && lastModule != bc.Module {
							if strings.ToLower(bc.Constructors[0])[0] != bc.Constructors[0][0] {
								fmt.Fprintf(bb, "this.%v.%v()\n", strings.Title(bc.Name), bc.Constructors[0])
							}
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
								fmt.Fprintf(bb, "qt.DisconnectSignal(ptr, \"%v\")\n", f.Name)
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
										t := f.Target
										if strings.Count(t, ".") != 2 {
											if !(len(strings.Split(f.Target, ".")) == 2 && strings.Split(f.Target, ".")[0] != "this" && strings.Split(f.Target, ".")[1][:1] == strings.ToLower(strings.Split(f.Target, ".")[1][:1])) {
												t = f.Target + "." + name
											}
										}
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
										t := f.Target
										if strings.Count(t, ".") != 2 {
											if !(len(strings.Split(f.Target, ".")) == 2 && strings.Split(f.Target, ".")[0] != "this" && strings.Split(f.Target, ".")[1][:1] == strings.ToLower(strings.Split(f.Target, ".")[1][:1])) {
												t = f.Target + "." + name
											}
										}
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

				fmt.Fprint(bb, "}\n\n")
			}
		}

		cTemplate(bb, class, goEnum, goFunction, "\n\n", true)
	}

	return preambleGo(module, goModule(module), bb.Bytes(), stub, mode, pkg, target, tags)
}

func preambleGo(oldModule string, module string, input []byte, stub bool, mode int, pkg, target, tags string) []byte {
	var bb = new(bytes.Buffer)
	defer bb.Reset()

	if UseStub(stub, oldModule, mode) {
		fmt.Fprintf(bb, `%v

package %v
`, buildTags(oldModule, stub, mode, tags), module)

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
	for _, m := range append(parser.GetLibs(), "qt", "strings", "unsafe", "log", "runtime", "fmt", "errors") {
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
			case "strings", "unsafe", "log", "runtime", "fmt", "errors":
				fmt.Fprintf(bb, "\"%v\"\n", mlow)

			case "qt":
				fmt.Fprintln(bb, "\"github.com/therecipe/qt\"")

			default:
				if mode == MOC {
					fmt.Fprintf(bb, "std_%[1]v \"github.com/therecipe/qt/%[1]v\"\n", mlow)
				} else {
					fmt.Fprintf(bb, "\"github.com/therecipe/qt/%v\"\n", mlow)
				}

				if mode == MOC {
					parser.LibDeps[parser.MOC] = append(parser.LibDeps[parser.MOC], m)
				}

				if strings.HasPrefix(target, "ios") && mode == MINIMAL {
					oldModuleGo := strings.TrimPrefix(oldModule, "Qt")

					var (
						containsSub  bool
						containsSelf bool
					)

					for _, l := range parser.LibDeps["build_ios"] {
						if l == m {
							containsSub = true
						}
						if l == oldModuleGo {
							containsSelf = true
						}
					}

					if !containsSelf || !containsSub {

						if !containsSelf {
							parser.LibDeps["build_ios"] = append(parser.LibDeps["build_ios"], oldModuleGo)

							switch oldModuleGo {
							case "Multimedia":
								parser.LibDeps["build_ios"] = append(parser.LibDeps["build_ios"], "MultimediaWidgets")
							case "Quick":
								parser.LibDeps["build_ios"] = append(parser.LibDeps["build_ios"], "QuickWidgets")
							}
						}

						if !containsSub {
							parser.LibDeps["build_ios"] = append(parser.LibDeps["build_ios"], m)

							switch m {
							case "Multimedia":
								parser.LibDeps["build_ios"] = append(parser.LibDeps["build_ios"], "MultimediaWidgets")
							case "Quick":
								parser.LibDeps["build_ios"] = append(parser.LibDeps["build_ios"], "QuickWidgets")
							}
						}

					}
				}
			}
		}
	}

	if mode == MOC {
		for custom, m := range parser.GetCustomLibs(target, tags) {
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

	out, err := format.Source(renameSubClasses(bb.Bytes(), "_"))
	if err != nil {
		utils.Log.WithError(err).Errorln("failed to format:", pkg, module)
		out = bb.Bytes()
	}

	//TODO: regexp
	if mode == MOC {
		pre := string(out)
		for _, c := range parser.SortedClassesForModule(module, true) {
			hName := c.Hash()
			sep := []string{"\n", "(", "_", "callback", "C."}
			for _, p := range sep {
				for _, s := range sep {
					if s == "callback" || s == "C." {
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

//TODO:
func renameSubClasses(in []byte, r string) []byte {
	for _, c := range parser.State.ClassMap {
		if c.Fullname != "" {
			in = []byte(strings.Replace(string(in), c.Name, strings.Replace(c.Fullname, "::", r, -1), -1))
			in = []byte(strings.Replace(string(in), "C."+strings.Replace(c.Fullname, "::", r, -1), "C."+c.Name, -1))
			in = []byte(strings.Replace(string(in), "_New"+strings.Replace(c.Fullname, "::", r, -1), "_New"+c.Name, -1))
		}
	}
	return in
}
