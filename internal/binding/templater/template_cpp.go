package templater

import (
	"bytes"
	"crypto/sha1"
	"encoding/hex"
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
	} else {
		for _, c := range parser.SortedClassNamesForModule(module, true) {
			var class, e = parser.State.ClassMap[c]
			if !e {
				continue
			}

			var typeMap = make(map[string]string)
			for _, f := range class.Functions {
				if parser.IsPackedMap(f.Output) {
					var tHash = sha1.New()
					tHash.Write([]byte(f.Output))
					typeMap[f.Output] = hex.EncodeToString(tHash.Sum(nil)[:3])
				}
				for _, p := range f.Parameters {
					if parser.IsPackedMap(p.Value) {
						var tHash = sha1.New()
						tHash.Write([]byte(p.Value))
						typeMap[p.Value] = hex.EncodeToString(tHash.Sum(nil)[:3])
					}
				}
			}

			for _, p := range class.Properties {
				if parser.IsPackedMap(p.Output) {
					var tHash = sha1.New()
					tHash.Write([]byte(p.Output))
					typeMap[p.Output] = hex.EncodeToString(tHash.Sum(nil)[:3])
				}
			}

			for typ, hash := range typeMap {
				fmt.Fprintf(bb, "typedef %v type%v;\n", typ, hash)
			}
		}
	}

	if module == "QtCharts" || module == "QtDataVisualization" {
		for _, classname := range parser.SortedClassNamesForModule(module, true) {
			fmt.Fprintf(bb, "typedef %v::%v %v;\n", module, classname, classname)
		}
		fmt.Fprint(bb, "\n")
	}

	for _, className := range parser.SortedClassNamesForModule(module, true) {
		var class = parser.State.ClassMap[className]

		if class.IsSupported() {

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

								var ty = p.Output
								if parser.IsPackedMap(p.Output) {
									var tHash = sha1.New()
									tHash.Write([]byte(p.Output))
									ty = fmt.Sprintf("type%v", hex.EncodeToString(tHash.Sum(nil)[:3]))
								}

								fmt.Fprintf(bb, "Q_PROPERTY(%v %v READ %v WRITE set%v NOTIFY %vChanged)\n", ty, p.Name,
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

				//callback functions
				var implementedVirtuals = make(map[string]struct{})
				for i, parentClassName := range append([]string{class.Name}, class.GetAllBases()...) {
					var parentClass, e = parser.State.ClassMap[parentClassName]
					if !e || !parentClass.IsSupported() {
						continue
					}

					for _, f := range parentClass.Functions {
						var _, e = implementedVirtuals[f.Name+f.OverloadNumber]
						if e || !f.IsSupported() {
							continue
						}

						if parentClass.Module == parser.MOC && f.Meta == parser.SLOT {
							continue
						}

						if i > 0 && (f.Meta == parser.CONSTRUCTOR || f.Meta == parser.DESTRUCTOR) {
							continue
						}

						var f = *f
						f.SignalMode = parser.CALLBACK
						f.Fullname = fmt.Sprintf("%v::%v", class.Name, f.Name)
						f.Fullname = fmt.Sprintf("%v::%v", f.FindDeepestImplementation(), f.Name)

						if f.Meta == parser.SLOT || f.Meta == parser.SIGNAL || f.Virtual == parser.IMPURE || f.Virtual == parser.PURE {
							implementedVirtuals[f.Name+f.OverloadNumber] = struct{}{}
							fmt.Fprintf(bb, "\t%v\n", cppFunctionCallback(&f))
						}
					}
				}

				if parser.State.Moc {
					for _, p := range class.Properties {

						var ty = p.Output
						if parser.IsPackedMap(p.Output) {
							var tHash = sha1.New()
							tHash.Write([]byte(p.Output))
							ty = fmt.Sprintf("type%v", hex.EncodeToString(tHash.Sum(nil)[:3]))
						}

						fmt.Fprintf(bb, "\t%v %v() { return _%v; };\n", ty, func() string {
							if p.Output == "bool" {
								return "is" + strings.Title(p.Name)
							}
							return p.Name
						}(), p.Name)
						fmt.Fprintf(bb, "\tvoid set%v(%v p) { if (p != _%v) { _%v = p; %vChanged(_%v); } };\n", strings.Title(p.Name), ty, p.Name, p.Name, p.Name, p.Name)
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
						var ty = p.Output
						if parser.IsPackedMap(p.Output) {
							var tHash = sha1.New()
							tHash.Write([]byte(p.Output))
							ty = fmt.Sprintf("type%v", hex.EncodeToString(tHash.Sum(nil)[:3]))
						}

						fmt.Fprintf(bb, "\t%v _%v;\n", ty, p.Name)
					}
				}

				fmt.Fprint(bb, "};\n\n")
			}
			if parser.State.Moc {
				fmt.Fprintf(bb, "Q_DECLARE_METATYPE(%v*)\n\n", class.Name)
			}
		}

		if !parser.State.Moc {
			cTemplate(bb, class, cppEnum, cppFunction, "\n\n", false)
		}
	}

	if parser.State.Moc {
		for _, className := range parser.SortedClassNamesForModule(module, true) {
			var class = parser.State.ClassMap[className]

			if class.IsSupported() {
				cTemplate(bb, class, cppEnum, cppFunction, "\n\n", false)
			}
		}

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
		if strings.Contains(string(input), class.Name) && class.Module != parser.MOC {
			classes = append(classes, class.Name)
		}
	}
	sort.Stable(sort.StringSlice(classes))

	for _, class := range classes {
		if class == "SailfishApp" {
			fmt.Fprint(bb, "#include <sailfishapp.h>\n")
		} else {
			var c, _ = parser.State.ClassMap[class]
			switch c.Name {
			case
				"Qt",
				"QPdf",
				"QDBus",
				"QAudio",
				"QMultimedia",
				"QSsl",
				"QPrint",
				"QScript",
				"QSql",
				"QTest",
				"QWebSocketProtocol",
				"OSXBluetooth",
				"QBluetooth",
				"PaintContext",
				"QPlatformGraphicsBuffer",
				"QDBusPendingReplyTypes":
				{
					continue
				}
			}
			fmt.Fprintf(bb, "#include <%v>\n", class)
		}
	}

	if module == "QtCore" {
		if !strings.Contains(bb.String(), "QTextDocument") {
			fmt.Fprint(bb, "#include <QTextDocument>\n")
		}
	}

	if parser.State.Minimal {
		if module == "QtCore" {
			fmt.Fprint(bb, "#include <QObject>\n")
		} else if module == "QtNetwork" {
			fmt.Fprint(bb, "#include <QSsl>\n")
		}

		if !strings.Contains(bb.String(), "QStringList") {
			fmt.Fprint(bb, "#include <QStringList>\n")
		}
	}

	fmt.Fprint(bb, "\n")

	bb.Write(input)

	return bb.Bytes()
}
