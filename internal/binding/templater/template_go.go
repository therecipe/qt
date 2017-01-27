package templater

import (
	"bytes"
	"fmt"
	"go/format"
	"strings"

	"github.com/therecipe/qt/internal/binding/parser"
	"github.com/therecipe/qt/internal/utils"
)

func GoTemplate(module string, stub bool) []byte {
	utils.Log.WithField("0_module", module).Debug("generating go")

	var bb = new(bytes.Buffer)
	defer bb.Reset()

	if !parser.State.Moc {
		module = "Qt" + module
	}

	if !(UseStub() || stub) {
		fmt.Fprintf(bb, "func cGoUnpackString(s C.struct_%v_PackedString) string { if len := int(s.len); len == -1 {\n return C.GoString(s.data)\n }\n return C.GoStringN(s.data, C.int(s.len)) }\n", strings.Title(module))
	}

	if module == "QtAndroidExtras" {
		fmt.Fprint(bb, "func QAndroidJniEnvironment_ExceptionCatch() error {\n")
		if !(UseStub() || stub) {
			fmt.Fprint(bb, "var err error\n")
			fmt.Fprint(bb, "if QAndroidJniEnvironment_ExceptionCheck() {\nvar tmpExcPtr = QAndroidJniEnvironment_ExceptionOccurred()\nQAndroidJniEnvironment_ExceptionClear()\n")
			fmt.Fprint(bb, "var tmpExc = NewQAndroidJniObject6(tmpExcPtr)\n")
			fmt.Fprint(bb, "err = errors.New(tmpExc.CallMethodString2(\"toString\", \"()Ljava/lang/String;\"))\n")
			fmt.Fprint(bb, "tmpExc.DestroyQAndroidJniObject()\n")
			fmt.Fprint(bb, "}\nreturn err\n")
		} else {
			fmt.Fprint(bb, "return nil\n")
		}
		fmt.Fprint(bb, "}\n\n")

		if !(UseStub() || stub) {
			fmt.Fprint(bb, "func (e *QAndroidJniEnvironment) ExceptionCatch() error { return QAndroidJniEnvironment_ExceptionCatch() }\n")
		} else {
			fmt.Fprint(bb, "func (e *QAndroidJniEnvironment) ExceptionCatch() error { return nil }\n")
		}
	}

	for _, class := range sortedClassesForModule(module) {

		class.Stub = UseStub() || stub

		if !parser.State.Minimal || (parser.State.Minimal && class.Export) {

			if !parser.State.Moc {
				fmt.Fprintf(bb, "type %v struct {\n%v\n}\n\n",

					class.Name,

					func() string {
						if class.Bases == "" {
							return "ptr unsafe.Pointer"
						}

						var bb = new(bytes.Buffer)
						defer bb.Reset()

						for _, parentClassName := range class.GetBases() {
							var parentClass = parser.State.ClassMap[parentClassName]
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
						var parentClass = parser.State.ClassMap[parentClassName]
						if parentClass.Module == class.Module {
							fmt.Fprintf(bb, "%v_ITF\n", parentClassName)
						} else {
							fmt.Fprintf(bb, "%v.%v_ITF\n", goModule(parentClass.Module), parentClassName)
						}
					}

					return bb.String()
				}(),

				fmt.Sprintf("%v_PTR() *%v\n", class.Name, class.Name),
			)

			fmt.Fprintf(bb, "func (ptr *%v) %v_PTR() *%v {\nreturn ptr\n}\n\n", class.Name, class.Name, class.Name)

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
func PointerFrom%v(ptr %v_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.%v_PTR().Pointer()
	}
	return nil
}
`, strings.Title(class.Name), class.Name, class.Name)

			fmt.Fprintf(bb, `
func New%vFromPointer(ptr unsafe.Pointer) *%v {
	var n = new(%v)
	n.SetPointer(ptr)
	return n
}
`, strings.Title(class.Name), class.Name, class.Name)

			if class.NeedsDestructor() {
				if UseStub() {
					fmt.Fprintf(bb, `
			func (ptr *%v) Destroy%v() {}
			`, class.Name, class.Name)
				} else {
					fmt.Fprintf(bb, `
func (ptr *%v) Destroy%v() {
	if ptr != nil {
		C.free(ptr.Pointer())%v
		ptr.SetPointer(nil)
	}
}

`, class.Name, class.Name, func() string {
						if class.HasCallbackFunctions() || class.IsSubClassOfQObject() {
							return "\nqt.DisconnectAllSignals(fmt.Sprint(ptr.Pointer()))"
						}
						return ""
					}())
				}
			}
		}

		cTemplate(bb, class, goEnum, goFunction, "\n\n", true)
	}

	return preambleGo(goModule(module), bb.Bytes(), stub)
}

func preambleGo(module string, input []byte, stub bool) []byte {
	var bb = new(bytes.Buffer)
	defer bb.Reset()

	if UseStub() {
		fmt.Fprintf(bb, `// +build !minimal

package %v

import "C"
import (
`, func() string {
			if parser.State.Moc {
				return parser.State.Module
			}
			return module
		}())

	} else {
		fmt.Fprintf(bb, `%v

package %v

//#include <stdint.h>
//#include <stdlib.h>
//#include <string.h>
//#include "%v.h"
import "C"
import (
`,

			buildTags(module, stub),

			func() string {
				if parser.State.Moc {
					return parser.State.Module
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
						if parser.State.Minimal {
							return fmt.Sprintf("%v-minimal", module)
						}

						if parser.State.Moc {
							return "moc"
						}

						return module
					}
				}
			}(),
		)
	}

	for _, m := range append(parser.Libs, "qt", "strings", "unsafe", "log", "runtime", "fmt", "errors") {
		m = strings.ToLower(m)
		if strings.Contains(string(input), fmt.Sprintf("%v.", m)) {
			switch m {
			case "strings", "unsafe", "log", "runtime", "fmt", "errors":
				{
					fmt.Fprintf(bb, "\"%v\"\n", m)
				}

			case "qt":
				{
					fmt.Fprintln(bb, "\"github.com/therecipe/qt\"")
				}

			default:
				{
					fmt.Fprintf(bb, "\"github.com/therecipe/qt/%v\"\n", m)

					if parser.State.Moc {
						parser.LibDeps[parser.MOC] = append(parser.LibDeps[parser.MOC], strings.Title(m))
					}
				}
			}
		}
	}

	fmt.Fprintln(bb, ")")

	bb.Write(input)

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
