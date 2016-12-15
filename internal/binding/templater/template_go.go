package templater

import (
	"bytes"
	"fmt"
	"go/format"
	"strings"

	"github.com/therecipe/qt/internal/binding/converter"
	"github.com/therecipe/qt/internal/binding/parser"
	"github.com/therecipe/qt/internal/utils"
)

func GoTemplate(module string, stub bool) []byte {
	var bb = new(bytes.Buffer)
	defer bb.Reset()

	if !(UseStub() || stub) {
		fmt.Fprintf(bb, "func cGoUnpackString(s C.struct_%v_PackedString) string { if len := int(s.len); len == -1 {\n return C.GoString(s.data)\n }\n return C.GoStringN(s.data, C.int(s.len)) }\n", strings.Title(module))
	}

	for _, class := range getSortedClassesForModule(module) {

		//all class enums
		for _, enum := range class.Enums {
			fmt.Fprintf(bb, "%v\n\n", goEnum(enum))
		}

		class.Stub = UseStub() || stub

		if !parser.CurrentState.Minimal || (parser.CurrentState.Minimal && class.Export) {

			if module != parser.MOC {
				fmt.Fprintf(bb, "type %v struct {\n%v\n}\n\n",

					class.Name,

					func() string {
						if class.Bases == "" {
							return "ptr unsafe.Pointer"
						}

						var bb = new(bytes.Buffer)
						defer bb.Reset()

						for _, parentClassName := range class.GetBases() {
							var parentClass = parser.CurrentState.ClassMap[parentClassName]
							if parentClass.Module == class.Module {
								fmt.Fprintf(bb, "%v\n", parentClassName)
							} else {
								fmt.Fprintf(bb, "%v.%v\n", shortModule(parentClass.Module), parentClassName)
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
						var parentClass = parser.CurrentState.ClassMap[parentClassName]
						if parentClass.Module == class.Module {
							fmt.Fprintf(bb, "%v_ITF\n", parentClassName)
						} else {
							fmt.Fprintf(bb, "%v.%v_ITF\n", shortModule(parentClass.Module), parentClassName)
						}
					}

					return bb.String()
				}(),

				fmt.Sprintf("%v_PTR() *%v\n", class.Name, class.Name),
			)

			fmt.Fprintf(bb, "func (p *%v) %v_PTR() *%v {\nreturn p\n}\n\n", class.Name, class.Name, class.Name)

			if class.Bases == "" {
				fmt.Fprintf(bb, "func (p *%v)Pointer() unsafe.Pointer {\nif p != nil {\nreturn p.ptr\n}\nreturn nil\n}\n\n", class.Name)
				fmt.Fprintf(bb, "func (p *%v)SetPointer(ptr unsafe.Pointer) {\nif p != nil {\np.ptr = ptr\n}\n}\n\n", class.Name)
			} else {
				fmt.Fprintf(bb, "func (p *%v)Pointer() unsafe.Pointer {\nif p != nil {\nreturn p.%v_PTR().Pointer()\n}\nreturn nil\n}\n\n", class.Name, class.GetBases()[0])

				fmt.Fprintf(bb, "func (p *%v)SetPointer(ptr unsafe.Pointer) {\nif p != nil{\n%v}\n}\n",

					class.Name,

					func() string {
						var bb = new(bytes.Buffer)
						defer bb.Reset()

						for _, parentClassName := range class.GetBases() {
							fmt.Fprintf(bb, "p.%v_PTR().SetPointer(ptr)\n", parentClassName)
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

			if classNeedsDestructor(class) {
				if UseStub() {
					fmt.Fprintf(bb, `
			func (ptr *%v) Destroy%v() {}
			`, class.Name, class.Name)
				} else {
					fmt.Fprintf(bb, `
func (ptr *%v) Destroy%v() {
C.free(ptr.Pointer())%v
ptr.SetPointer(nil)
}

`, class.Name, class.Name, func() string {
						if classNeedsCallbackFunctions(class) || class.IsSubClassOfQObject() {
							return "\nqt.DisconnectAllSignals(fmt.Sprint(ptr.Pointer()))"
						}
						return ""
					}())
				}
			}
		}

		if classIsSupported(class) {

			var implementedVirtuals = make(map[string]struct{})

			//all class functions
			for _, function := range class.Functions {

				if functionIsSupported(class, function) {

					implementedVirtuals[fmt.Sprint(function.Name, function.OverloadNumber)] = struct{}{}
					switch {
					case (function.Virtual == parser.IMPURE || function.Virtual == parser.PURE || function.Meta == parser.SIGNAL || function.Meta == parser.SLOT):
						{
							for _, signalMode := range []string{parser.CALLBACK, parser.CONNECT, parser.DISCONNECT} {
								var function = *function
								function.SignalMode = signalMode

								fmt.Fprintf(bb, "%v%v\n\n",
									func() string {
										if signalMode == parser.CALLBACK {
											return fmt.Sprintf("//export %v\n", converter.GoHeaderName(&function))
										}
										return ""
									}(),

									goFunction(&function),
								)
							}

							if !converter.IsPrivateSignal(function) {
								var function = *function
								function.Meta = parser.PLAIN
								fmt.Fprintf(bb, "%v\n\n", goFunction(&function))

								if function.Virtual == parser.IMPURE {
									function.Default = true
									fmt.Fprintf(bb, "%v\n\n", goFunction(&function))
								}
							}
						}

					case isJNIGeneric(function):
						{
							for _, mode := range converter.CppOutputParametersJNIGenericModes(function) {
								var function = *function
								function.TemplateModeJNI = mode

								fmt.Fprintf(bb, "%v\n\n", goFunction(&function))
							}
						}

					default:
						{
							if !(function.Meta == parser.CONSTRUCTOR && hasUnimplementedPureVirtualFunctions(class.Name)) {
								fmt.Fprintf(bb, "%v\n\n", goFunction(function))

								if function.Static {
									fmt.Fprintf(bb, "%v{\n%v\n}\n\n",
										func() string {
											var function = *function
											function.Static = false
											return goFunctionHeader(&function)
										}(),
										goFunctionBody(function),
									)
								}
							}
						}
					}
				}
			}

			//virtual parent functions
			for _, parentClassName := range class.GetAllBases() {
				var parentClass = parser.CurrentState.ClassMap[parentClassName]
				if classIsSupported(parentClass) {

					for _, function := range parentClass.Functions {
						if _, exists := implementedVirtuals[fmt.Sprint(function.Name, function.OverloadNumber)]; !exists {

							if functionIsSupported(parentClass, function) {
								if function.Meta != parser.SIGNAL && (function.Virtual == parser.IMPURE || function.Virtual == parser.PURE || function.Meta == parser.SLOT) && function.Meta != parser.DESTRUCTOR {
									implementedVirtuals[fmt.Sprint(function.Name, function.OverloadNumber)] = struct{}{}

									for _, signalMode := range []string{parser.CALLBACK, parser.CONNECT, parser.DISCONNECT} {
										var pbf = function.Virtual == parser.PURE

										var function = *function
										function.Fullname = fmt.Sprintf("%v::%v", class.Name, function.Name)
										function.Virtual = parser.IMPURE
										function.SignalMode = signalMode
										function.PureBaseFunction = pbf

										fmt.Fprintf(bb, "%v%v\n\n",
											func() string {
												if signalMode == parser.CALLBACK {
													return fmt.Sprintf("//export %v\n", converter.GoHeaderName(&function))
												}
												return ""
											}(),

											goFunction(&function),
										)
									}

									var function = *function
									function.Fullname = fmt.Sprintf("%v::%v", class.Name, function.Name)
									if function.Meta != parser.SLOT {
										function.Meta = parser.PLAIN
									}
									fmt.Fprintf(bb, "%v\n\n", goFunction(&function))

									var class, exist = function.Class()
									if !exist || (class.Module == parser.MOC && function.Virtual == parser.PURE) {
										continue
									}

									function.Meta = parser.PLAIN
									function.Default = true
									fmt.Fprintf(bb, "%v\n\n", goFunction(&function))

								}
							}

						}
					}

				}
			}

		}

	}

	return preambleGo(shortModule(module), bb.Bytes(), stub)
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
			if parser.CurrentState.MocModule != "" {
				return parser.CurrentState.MocModule
			}
			return module
		}())

	} else {
		fmt.Fprintf(bb, `%v

package %v

//#include <stdint.h>
//#include <stdlib.h>
//#include "%v.h"
import "C"
import (
`,

			func() string {
				switch {
				case stub:
					{
						if module == "androidextras" {
							return "// +build !android"
						}
						return "// +build !sailfish,!sailfish_emulator"
					}

				case parser.CurrentState.Minimal:
					{
						return "// +build minimal"
					}

				case module == parser.MOC:
					{
						return ""
					}

				case module == "androidextras":
					{
						return "// +build android"
					}

				case module == "sailfish":
					{
						return "// +build sailfish sailfish_emulator"
					}

				default:
					{
						return "// +build !minimal"
					}
				}
			}(),

			func() string {
				if parser.CurrentState.MocModule != "" {
					return parser.CurrentState.MocModule
				}
				return module
			}(),

			func() string {
				if parser.CurrentState.Minimal {
					return fmt.Sprintf("%v-minimal", module)
				}

				switch module {
				case parser.MOC:
					{
						return "moc"
					}

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

					if module == parser.MOC {
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

func renameSubClasses(in []byte, r string) []byte {
	for _, c := range parser.CurrentState.ClassMap {
		if c.Fullname != "" {
			in = []byte(strings.Replace(string(in), c.Name, strings.Replace(c.Fullname, "::", r, -1), -1))
			in = []byte(strings.Replace(string(in), "C."+strings.Replace(c.Fullname, "::", r, -1), "C."+c.Name, -1))
			in = []byte(strings.Replace(string(in), "_New"+strings.Replace(c.Fullname, "::", r, -1), "_New"+c.Name, -1))
		}
	}
	return in
}
