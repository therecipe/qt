package templater

import (
	"bytes"
	"crypto/sha1"
	"encoding/hex"
	"fmt"
	"sort"
	"strings"

	"github.com/therecipe/qt/internal/binding/converter"
	"github.com/therecipe/qt/internal/binding/parser"
	"github.com/therecipe/qt/internal/utils"
)

func CppTemplate(module string, mode int, target, tags string) []byte {
	utils.Log.WithField("module", module).Debug("generating cpp")
	parser.State.Target = target

	var bb = new(bytes.Buffer)
	defer bb.Reset()

	if mode != MOC {
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

			if class.HasCallbackFunctions() || mode == MOC {

				//TODO: split
				fmt.Fprintf(bb,
					`class %v%v: public %v
{
%vpublic:
`,
					func() string {
						if mode == MOC {
							return ""
						}
						return "My"
					}(),

					class.Name,

					func() string {
						if mode == MOC {
							return class.GetBases()[0]
						}
						return class.Name
					}(),

					func() string {
						if mode == MOC {
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
										if p.Output == "bool" && !strings.HasPrefix(strings.ToLower(p.Name), "is") {
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

						var out = fmt.Sprintf("\t%v%v(%v) : %v(%v) {%v};\n",
							func() string {
								if mode == MOC {
									return ""
								}
								return "My"
							}(),

							function.ClassName(),

							func() string {
								var input []string
								for _, p := range function.OgParameters {
									name := parser.CleanName(p.Name, p.Value)
									if p.Default != "" {
										if strings.HasSuffix(p.Value, "*") || strings.HasSuffix(p.Value, "&") {
											input = append(input, p.Value+name+" = "+p.Default)
										} else {
											input = append(input, p.Value+" "+name+" = "+p.Default)
										}
									} else {
										if strings.HasSuffix(p.Value, "*") || strings.HasSuffix(p.Value, "&") {
											input = append(input, p.Value+name)
										} else {
											input = append(input, p.Value+" "+name)
										}
									}
								}
								return strings.Join(input, ", ")
							}(),

							func() string {
								if mode == MOC {
									return class.GetBases()[0]
								}
								return function.ClassName()
							}(),

							func() string {
								input := make([]string, len(function.Parameters))
								for i, p := range function.Parameters {
									input[i] = parser.CleanName(p.Name, p.Value)
								}
								return strings.Join(input, ", ")
							}(),

							func() string {
								var pre string
								if class.IsSubClassOfQObject() {
									pre = fmt.Sprintf("%[1]v_%[1]v_QRegisterMetaType();", className)
								}
								if mode != MOC {
									return pre
								}
								return fmt.Sprintf("qRegisterMetaType<quintptr>(\"quintptr\");%v%[2]v_%[2]v_QRegisterMetaTypes();callback%[2]v_Constructor(this);", pre, className)
							}(),
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

						if (parentClass.Module == parser.MOC || parentClass.Pkg != "") && f.Meta == parser.SLOT {
							continue
						}

						if i > 0 && (f.Meta == parser.CONSTRUCTOR || f.Meta == parser.DESTRUCTOR) {
							continue
						}

						implementedVirtuals[f.Name+f.OverloadNumber] = struct{}{}

						var f = *f
						f.SignalMode = parser.CALLBACK
						f.Fullname = fmt.Sprintf("%v::%v", class.Name, f.Name)
						f.Fullname = fmt.Sprintf("%v::%v", f.FindDeepestImplementation(), f.Name)

						if f.Meta == parser.SLOT || f.Meta == parser.SIGNAL || f.Virtual == parser.IMPURE || f.Virtual == parser.PURE {
							fmt.Fprintf(bb, "\t%v\n", cppFunctionCallback(&f))
						}
					}
				}

				if mode == MOC {
					for _, p := range class.Properties {

						var ty = p.Output
						if parser.IsPackedMap(p.Output) {
							var tHash = sha1.New()
							tHash.Write([]byte(p.Output))
							ty = fmt.Sprintf("type%v", hex.EncodeToString(tHash.Sum(nil)[:3]))
						}

						fmt.Fprintf(bb, "\t%v %v%v() { return _%v; };\n",
							ty,
							func() string {
								if p.Output == "bool" && !strings.HasPrefix(strings.ToLower(p.Name), "is") {
									return "is" + strings.Title(p.Name)
								}
								return p.Name
							}(),
							func() string {
								if p.IsMocSynthetic {
									return ""
								}
								return "Default"
							}(),
							p.Name,
						)
						fmt.Fprintf(bb, "\tvoid set%v%v(%v p) { if (p != _%v) { _%v = p; %vChanged(_%v); } };\n",
							strings.Title(p.Name),
							func() string {
								if p.IsMocSynthetic {
									return ""
								}
								return "Default"
							}(),
							ty,
							p.Name,
							p.Name,
							p.Name,
							p.Name,
						)
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
			if class.IsSubClassOfQObject() {
				fmt.Fprintf(bb, "Q_DECLARE_METATYPE(%v%v*)\n\n",
					func() string {
						if mode != MOC {
							return "My"
						}
						return ""
					}(), class.Name)

				if mode != MOC {
					if strings.HasPrefix(class.Name, "QMac") && !strings.HasPrefix(parser.State.ClassMap[class.Name].Module, "QtMac") {
						fmt.Fprintf(bb, "int %[1]v_%[1]v_QRegisterMetaType(){\n\t#ifdef Q_OS_OSX\n\t\tqRegisterMetaType<%[1]v*>(); return qRegisterMetaType<My%[1]v*>();\n\t#else\n\t\treturn 0;\n\t#endif\n}\n\n", class.Name)
					} else {
						fmt.Fprintf(bb, "int %[1]v_%[1]v_QRegisterMetaType(){qRegisterMetaType<%[1]v*>(); return qRegisterMetaType<My%[1]v*>();}\n\n", class.Name)
					}
				} else {
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

					propTypes := make(map[string]struct{})
					for _, p := range class.Properties {
						if parser.IsPackedMap(p.Output) {
							var tHash = sha1.New()
							tHash.Write([]byte(p.Output))
							typeMap[p.Output] = hex.EncodeToString(tHash.Sum(nil)[:3])
						}
						if o := converter.CppRegisterMetaTypeProp(p); o != "" {
							propTypes[o] = struct{}{}
						}
					}

					for _, hash := range typeMap {
						if hash == "30021d" || //QHash<quintptr, quintptr>
							hash == "95ad14" || //QHash<qint32, quintptr>
							hash == "d01680" { //QHash<qint32, QByteArray>
							continue
						}
						fmt.Fprintf(bb, "Q_DECLARE_METATYPE(type%v)\n", hash)
					}

					fmt.Fprintf(bb, "\nvoid %[1]v_%[1]v_QRegisterMetaTypes() {\n", class.Name)
					for _, hash := range typeMap {
						fmt.Fprintf(bb, "\tqRegisterMetaType<type%v>(\"type%v\");\n", hash, hash)
					}
					for t := range propTypes {
						fmt.Fprintf(bb, "\tqRegisterMetaType<%v>();\n", t)
					}
					fmt.Fprint(bb, "}\n\n")
				}
			}
		}

		if mode != MOC {
			cTemplate(bb, class, cppEnum, cppFunction, "\n\n", false)
		}
	}

	if mode == MOC {
		for _, className := range parser.SortedClassNamesForModule(module, true) {
			var class = parser.State.ClassMap[className]

			if class.IsSupported() {
				cTemplate(bb, class, cppEnum, cppFunction, "\n\n", false)
			}
		}

		fmt.Fprintln(bb, "#include \"moc_moc.h\"")
	}

	return preambleCpp(module, bb.Bytes(), mode, tags)
}

func preambleCpp(module string, input []byte, mode int, tags string) []byte {
	var bb = new(bytes.Buffer)
	defer bb.Reset()

	if mode == MOC {
		libsm := make(map[string]struct{}, 0)
		for _, c := range parser.State.ClassMap {
			if c.Pkg != "" && c.IsSubClassOfQObject() {
				libsm[c.Module] = struct{}{}
			}
		}

		var libs []string
		for k := range libsm {
			libs = append(libs, k)
		}

		for _, c := range parser.SortedClassesForModule(strings.Join(libs, ","), true) {
			if c.Pkg == "" || !strings.Contains(string(input), c.Name) /*|| !c.HasConstructor()*/ {
				continue
			}

			fmt.Fprintf(bb, "class %v: public %v{\npublic:\n", c.Name, c.GetBases()[0])

			for _, function := range c.Functions {
				if function.Meta != parser.CONSTRUCTOR || !function.IsSupported() {
					continue
				}

				var input = make([]string, len(function.Parameters))
				for i, p := range function.Parameters {
					input[i] = p.Name
				}

				fmt.Fprintf(bb, "\t%v%v(%v) : %v(%v) {};\n",
					func() string {
						if mode == MOC {
							return ""
						}
						return "My"
					}(),

					function.ClassName(),

					strings.Split(strings.Split(function.Signature, "(")[1], ")")[0],

					func() string {
						if mode == MOC {
							return c.GetBases()[0]
						}
						return function.ClassName()
					}(),

					strings.Join(input, ", "),
				)
			}
			fmt.Fprint(bb, "\n};\n")
		}

		fmt.Fprint(bb, "\n")

		bb.Write(input)
		input = []byte(bb.String())
		bb.Reset()
	}

	fmt.Fprintf(bb, `%v

#define protected public
#define private public

#include "%v.h"
#include "_cgo_export.h"

`,
		buildTags(module, false, mode, tags),

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
					if mode == MINIMAL {
						return fmt.Sprintf("%v-minimal", goModule(module))
					}

					if mode == MOC {
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
			if strings.HasPrefix(c.Module, "custom_") {
				continue
			}
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
				"QDBusPendingReplyTypes",
				"QRemoteObjectPackets",
				"QRemoteObjectStringLiterals":
				{
					continue
				}
			}
			fmt.Fprintf(bb, "#include <%v>\n", class)
		}

		c, ok := parser.State.ClassMap[class]
		if ok && !strings.Contains(strings.Join(parser.LibDeps[strings.TrimPrefix(module, "Qt")], " "), strings.TrimPrefix(c.Module, "Qt")) {
			if strings.HasPrefix(c.Module, "custom_") {
				continue
			}

			utils.Log.Debugf("%v add dependency: %v", module, c.Module)
			parser.LibDeps[strings.TrimPrefix(module, "Qt")] = append(parser.LibDeps[strings.TrimPrefix(module, "Qt")], strings.TrimPrefix(c.Module, "Qt"))
			switch c.Module {
			case "QtMultimedia":
				parser.LibDeps[strings.TrimPrefix(module, "Qt")] = append(parser.LibDeps[strings.TrimPrefix(module, "Qt")], "MultimediaWidgets")
			case "QtWebEngine":
				parser.LibDeps[strings.TrimPrefix(module, "Qt")] = append(parser.LibDeps[strings.TrimPrefix(module, "Qt")], "WebEngineWidgets")
			case "QtQuick":
				parser.LibDeps[strings.TrimPrefix(module, "Qt")] = append(parser.LibDeps[strings.TrimPrefix(module, "Qt")], "QuickWidgets")
			}
		}
	}

	if module == "QtCore" {
		if !strings.Contains(bb.String(), "QTextDocument") {
			fmt.Fprint(bb, "#include <QTextDocument>\n")
		}
	}

	if mode == MINIMAL {
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
