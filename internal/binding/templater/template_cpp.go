package templater

import (
	"bytes"
	"fmt"
	"sort"
	"strings"

	"github.com/therecipe/qt/internal/binding/parser"
	"github.com/therecipe/qt/internal/utils"
)

func CppTemplate(module string) []byte {
	utils.Log.WithField("0_module", module).Debug("generating cpp")

	var bb = new(bytes.Buffer)
	defer bb.Reset()

	if !parser.State.Moc {
		module = "Qt" + module
	}

	if module == "QtCharts" || module == "QtDataVisualization" {
		for _, classname := range sortedClassNamesForModule(module) {
			fmt.Fprintf(bb, "typedef %v::%v %v;\n", module, classname, classname)
		}
		fmt.Fprint(bb, "\n")
	}

	for _, className := range sortedClassNamesForModule(module) {
		var class = parser.State.ClassMap[className]

		if class.IsSupported() {
			var implementedVirtuals = make(map[string]struct{})

			if class.HasCallbackFunctions() || parser.State.Moc {

				//TODO: split
				fmt.Fprintf(bb,
					`class %v%v: public %v
{
%vpublic:
`,
					func() string {
						if parser.State.Moc {
							return ""
						}
						return "My"
					}(),

					class.Name,

					func() string {
						if parser.State.Moc {
							return class.GetBases()[0]
						}
						return class.Name
					}(),

					func() string {
						if parser.State.Moc {
							var bb = new(bytes.Buffer)
							defer bb.Reset()
							fmt.Fprintln(bb, "Q_OBJECT")

							for _, p := range class.Properties {
								fmt.Fprintf(bb, "Q_PROPERTY(%v %v READ %v WRITE set%v NOTIFY %vChanged)\n", p.Output, p.Name,
									func() string {
										if p.Output == "bool" {
											return "is" + strings.Title(p.Name)
										}
										return p.Name
									}(), strings.Title(p.Name), p.Name)
							}

							return bb.String()
						}
						return ""
					}())

				if !hasUnimplementedPureVirtualFunctions(class.Name) {
					for _, function := range class.Functions {
						if function.Meta != parser.CONSTRUCTOR || !function.IsSupported() {
							continue
						}

						var input = make([]string, len(function.Parameters))
						for i, p := range function.Parameters {
							input[i] = p.Name
						}

						var out = fmt.Sprintf("\t%v%v(%v) : %v(%v) {};\n",
							func() string {
								if parser.State.Moc {
									return ""
								}
								return "My"
							}(),

							function.ClassName(),

							strings.Split(strings.Split(function.Signature, "(")[1], ")")[0],

							func() string {
								if parser.State.Moc {
									return class.GetBases()[0]
								}
								return function.ClassName()
							}(),

							strings.Join(input, ", "),
						)

						fmt.Fprint(bb, out)
					}
				}

				//TODO: integrate into cTemplate? -->

				//all class functions
				for _, function := range class.Functions {
					implementedVirtuals[fmt.Sprint(function.Fullname, function.OverloadNumber)] = struct{}{}
					if function.IsSupported() {
						if function.Virtual == parser.IMPURE || function.Virtual == parser.PURE || function.Meta == parser.SIGNAL || function.Meta == parser.SLOT {
							if !(parser.State.Moc && function.Meta == parser.SLOT) {
								var function = *function
								function.SignalMode = parser.CALLBACK
								fmt.Fprintf(bb, "\t%v\n", cppFunctionCallback(&function))
							}
						}
					}
				}

				//virtual parent functions
				for _, parentClassName := range class.GetAllBases() {
					var parentClass = parser.State.ClassMap[parentClassName]
					if parentClass.IsSupported() {

						for _, function := range parentClass.Functions {
							if _, exists := implementedVirtuals[fmt.Sprint(fmt.Sprintf("%v::%v", class.Name, function.Name), function.OverloadNumber)]; !exists {
								if function.IsSupported() && function.Meta != parser.DESTRUCTOR {

									var function = *function
									function.Fullname = fmt.Sprintf("%v::%v", class.Name, function.Name)
									function.SignalMode = parser.CALLBACK
									if function.Virtual == parser.IMPURE || function.Virtual == parser.PURE || function.Meta == parser.SLOT {
										implementedVirtuals[fmt.Sprint(fmt.Sprintf("%v::%v", class.Name, function.Name), function.OverloadNumber)] = struct{}{}
										fmt.Fprintf(bb, "\t%v\n", cppFunctionCallback(&function))
									}

								}
							}
						}

					}
				}
				//<--

				if parser.State.Moc {
					for _, p := range class.Properties {
						fmt.Fprintf(bb, "\t%v %v() { return _%v; };\n", p.Output, func() string {
							if p.Output == "bool" {
								return "is" + strings.Title(p.Name)
							}
							return p.Name
						}(), p.Name)
						fmt.Fprintf(bb, "\tvoid set%v(%v p) { if (p != _%v) { _%v = p; %vChanged(_%v); } };\n", strings.Title(p.Name), p.Output, p.Name, p.Name, p.Name, p.Name)
					}

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

					fmt.Fprintln(bb, "private:")
					for _, p := range class.Properties {
						fmt.Fprintf(bb, "\t%v _%v;\n", p.Output, p.Name)
					}
				}

				fmt.Fprint(bb, "};\n\n")
			}
			if parser.State.Moc {
				fmt.Fprintf(bb, "Q_DECLARE_METATYPE(%v*)\n\n", class.Name)
			}
		}

		cTemplate(bb, class, cppEnum, cppFunction, "\n\n", false)
	}

	if parser.State.Moc {
		fmt.Fprintln(bb, "#include \"moc_moc.h\"")
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
		buildTags(module, false),

		func() string {
			switch module {
			case "QtAndroidExtras":
				{
					return fmt.Sprintf("%v_android", goModule(module))
				}

			case "QtSailfish":
				{
					return fmt.Sprintf("%v_sailfish", goModule(module))
				}

			default:
				{
					if parser.State.Minimal {
						return fmt.Sprintf("%v-minimal", goModule(module))
					}

					if parser.State.Moc {
						return "moc"
					}

					return goModule(module)
				}
			}
		}(),
	)

	var classes = make([]string, 0)
	for _, class := range parser.State.ClassMap {
		if strings.Contains(string(input), class.Name) && !(strings.HasPrefix(class.Name, "Qt") || class.Module == parser.MOC) {
			classes = append(classes, class.Name)
		}
	}
	sort.Stable(sort.StringSlice(classes))

	for _, class := range classes {
		if class == "SailfishApp" {
			fmt.Fprint(bb, "#include <sailfishapp.h>\n")
		} else {
			if strings.HasPrefix(class, "Q") && !(class == "QBluetooth" || class == "QDBus" || class == "QCss" || class == "QPdf" || class == "QSsl" || class == "QPrint" || class == "QScript" || class == "QSql" || class == "QTest" || class == "QWebSocketProtocol") {
				fmt.Fprintf(bb, "#include <%v>\n", class)
			}
		}
	}
	fmt.Fprint(bb, "\n")

	bb.Write(input)

	return bb.Bytes()
}
