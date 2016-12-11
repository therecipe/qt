package templater

import (
	"bytes"
	"fmt"
	"sort"
	"strings"

	"github.com/therecipe/qt/internal/binding/converter"
	"github.com/therecipe/qt/internal/binding/parser"
)

func CppTemplate(module string) []byte {
	var bb = new(bytes.Buffer)
	defer bb.Reset()

	var tmpArray = getSortedClassNamesForModule(module)

	if module == parser.MOC {
		var items = make(map[string]string)

		for _, class := range tmpArray {
			items[class] = parser.CurrentState.ClassMap[class].Bases
		}

		var provided = make([]string, 0)

		for len(items) > 0 {
			for item, dependency := range items {
				var existsInOtherModule bool
				if parser.CurrentState.ClassMap[dependency].Module != parser.MOC {
					existsInOtherModule = true
				}
				var existsInCurrentOrder bool
				for _, providedItem := range provided {
					if dependency == providedItem {
						existsInCurrentOrder = true
						break
					}
				}

				if existsInOtherModule || existsInCurrentOrder {
					provided = append(provided, item)
					delete(items, item)
				}
			}
		}
		tmpArray = provided
	}

	for _, className := range tmpArray {
		var class = parser.CurrentState.ClassMap[className]

		//all class enums
		for _, enum := range class.Enums {
			for _, value := range enum.Values {
				if converter.EnumNeedsCppGlue(value.Value) {
					fmt.Fprintf(bb, "%v\n\n", cppEnum(enum, value))
				}
			}
		}

		if classIsSupported(class) {
			var implementedVirtuals = make(map[string]struct{})

			if classNeedsCallbackFunctions(class) || class.Module == parser.MOC {

				fmt.Fprintf(bb,
					`class %v%v: public %v
{
%vpublic:
`,
					func() string {
						if module == parser.MOC {
							return ""
						}
						return "My"
					}(),

					class.Name,

					func() string {
						if module == parser.MOC {
							return class.GetBases()[0]
						}
						return class.Name
					}(),

					func() string {
						if module == parser.MOC {
							return "Q_OBJECT\n"
						}
						return ""
					}())

				if !hasUnimplementedPureVirtualFunctions(class.Name) {
					for _, function := range class.Functions {
						if function.Meta == parser.CONSTRUCTOR {
							if functionIsSupported(class, function) {

								var input = make([]string, len(function.Parameters))
								for i, p := range function.Parameters {
									input[i] = func() string {
										if p.Name == "" {
											return "v"
										}
										return p.Name
									}()
								}

								fmt.Fprintf(bb, "\t%v%v(%v) : %v(%v) {};\n",
									func() string {
										if module == parser.MOC {
											return ""
										}
										return "My"
									}(),

									function.ClassName(),

									strings.Split(strings.Split(function.Signature, "(")[1], ")")[0],

									func() string {
										if module == parser.MOC {
											return class.GetBases()[0]
										}
										return function.ClassName()
									}(),

									strings.Join(input, ", "),
								)

							}
						}
					}
				}

				//all class functions
				for _, function := range class.Functions {
					implementedVirtuals[fmt.Sprint(function.Fullname, function.OverloadNumber)] = struct{}{}
					if functionIsSupported(class, function) && !strings.Contains(function.Meta, "constructor") {
						if function.Virtual == parser.IMPURE || function.Virtual == parser.PURE || function.Meta == parser.SIGNAL || function.Meta == parser.SLOT {
							if !(module == parser.MOC && function.Meta == parser.SLOT) {
								fmt.Fprintf(bb, "\t%v\n", cppFunctionCallback(function))
							}
						}
					}
				}

				//virtual parent functions
				for _, parentClassName := range class.GetAllBases() {
					var parentClass = parser.CurrentState.ClassMap[parentClassName]
					if classIsSupported(parentClass) {

						for _, function := range parentClass.Functions {
							if _, exists := implementedVirtuals[fmt.Sprint(fmt.Sprintf("%v::%v", class.Name, function.Name), function.OverloadNumber)]; !exists {
								implementedVirtuals[fmt.Sprint(fmt.Sprintf("%v::%v", class.Name, function.Name), function.OverloadNumber)] = struct{}{}
								if functionIsSupported(class, function) && !strings.Contains(function.Meta, "structor") {

									var function = *function
									function.Fullname = fmt.Sprintf("%v::%v", class.Name, function.Name)
									if function.Virtual == parser.IMPURE || function.Virtual == parser.PURE || function.Meta == parser.SLOT {
										fmt.Fprintf(bb, "\t%v\n", cppFunctionCallback(&function))
									}

								}
							}
						}

					}
				}

				if module == parser.MOC {
					fmt.Fprintln(bb, "signals:")
					for _, function := range class.Functions {
						if function.Meta == parser.SIGNAL {
							var function = *function
							function.Meta = parser.SLOT
							fmt.Fprintf(bb, "\t%v;\n", cppFunctionCallbackHeader(&function))
						}
					}

					fmt.Fprintln(bb, "public slots:")
					for _, function := range class.Functions {
						if function.Meta == parser.SLOT {
							fmt.Fprintf(bb, "\t%v\n", cppFunctionCallback(function))
						}
					}
				}

				fmt.Fprint(bb, "};\n\n")
			}
			if class.Module == parser.MOC {
				fmt.Fprintf(bb, "Q_DECLARE_METATYPE(%v*)\n\n", class.Name)
			}

			test(bb, class, cppFunction, "\n\n")

		}

	}

	return preambleCpp(module, bb.Bytes())
}

func preambleCpp(module string, input []byte) []byte {
	var bb = new(bytes.Buffer)
	defer bb.Reset()

	fmt.Fprintf(bb, `%v

#define protected public
#define private public

#include "%v.h"
#include "_cgo_export.h"

`,
		func() string {
			switch {
			case parser.CurrentState.Minimal:
				{
					return "// +build minimal"
				}

			case module == parser.MOC:
				{
					return ""
				}

			case module == "QtAndroidExtras":
				{
					return "// +build android"
				}

			case module == "QtSailfish":
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
			switch module {
			case parser.MOC:
				{
					return "moc"
				}

			case "QtAndroidExtras":
				{
					return fmt.Sprintf("%v_android", shortModule(module))
				}

			case "QtSailfish":
				{
					return fmt.Sprintf("%v_sailfish", shortModule(module))
				}

			default:
				{
					if parser.CurrentState.Minimal {
						return fmt.Sprintf("%v-minimal", shortModule(module))
					}
					return shortModule(module)
				}
			}
		}(),
	)

	var classes = make([]string, 0)
	for _, class := range parser.CurrentState.ClassMap {
		if strings.Contains(string(input), class.Name) && !(strings.HasPrefix(class.Name, "Qt") || class.Module == parser.MOC) {
			classes = append(classes, class.Name)
		}
	}
	sort.Stable(sort.StringSlice(classes))

	for _, class := range classes {
		if class == "SailfishApp" {
			fmt.Fprintln(bb, "#include <sailfishapp.h>")
		} else {
			if strings.HasPrefix(class, "Q") && !(class == "QBluetooth" || class == "QDBus" || class == "QCss" || class == "QPdf" || class == "QSsl" || class == "QPrint" || class == "QScript" || class == "QSql" || class == "QTest" || class == "QWebSocketProtocol") {
				fmt.Fprintf(bb, "#include <%v>\n", class)
			}
		}
	}
	fmt.Fprint(bb, "\n")

	bb.Write(input)

	if module == parser.MOC {
		fmt.Fprintln(bb, "#include \"moc_moc.h\"")
	}

	return bb.Bytes()
}
