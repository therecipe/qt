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
		fmt.Fprintf(bb, "func cGoUnpackString(s C.struct_%v_PackedString) string { if len := int(s.len); len == -1 {\n return C.GoString(s.data)\n }\n return C.GoStringN(s.data, C.int(s.len)) }\n", strings.Title(module))
	}

	if module == "QtAndroidExtras" {
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
func New%vFromPointer(ptr unsafe.Pointer) *%[2]v {
	var n *%[2]v
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
	return n
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
func New%vFromPointer(ptr unsafe.Pointer) *%[2]v {
	var n = new(%[2]v)
	n.SetPointer(ptr)
	return n
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

				fmt.Fprintf(bb, "gPtr := New%vFromPointer(ptr)\nqt.Register(ptr, gPtr)\n", strings.Title(class.Name))

				var lastModule string
				for _, bcn := range class.GetAllBases() {
					if bc := parser.State.ClassMap[bcn]; bc.Module != class.Module {
						if len(bc.Constructors) > 0 && lastModule != bc.Module {
							if strings.ToLower(bc.Constructors[0])[0] != bc.Constructors[0][0] {
								fmt.Fprintf(bb, "gPtr.%v.%v()\n", strings.Title(bc.Name), bc.Constructors[0])
							}
						}
						lastModule = bc.Module
					}
				}
				if len(class.Constructors) > 0 {
					fmt.Fprintf(bb, "gPtr.%v()\n", class.Constructors[0])
				}

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
			for _, pre := range []string{" ", "\t", "\r", "\n", "!", "*", "(", ")"} {
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
			}
		}
	}

	for custom, m := range parser.GetCustomLibs(target, tags) {
		mlows := strings.Split(m, "/")
		mlow := mlows[len(mlows)-1]
		switch {
		case strings.Contains(inputString, fmt.Sprintf("custom_%v.", mlow)):
			fmt.Fprintf(bb, "custom_%v \"%v\"\n", mlow, m)
		case strings.Contains(inputString, fmt.Sprintf("%v.", custom)):
			fmt.Fprintf(bb, "%v \"%v\"\n", custom, m)
		}
	}

	for i := range parser.State.MocImports {
		fmt.Fprintf(bb, "%v\n", i)
	}
	parser.State.MocImports = make(map[string]struct{})

	fmt.Fprintln(bb, ")")

	bb.WriteString(inputString)

	var out, err = format.Source(renameSubClasses(bb.Bytes(), "_"))
	if err != nil {
		utils.Log.WithError(err).Panicln("failed to format:", module)
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
