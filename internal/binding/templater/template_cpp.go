package templater

import (
	"fmt"
	"sort"
	"strings"

	"github.com/therecipe/qt/internal/binding/parser"
)

func HTemplate(m string) (o string) {
	o += "#ifdef __cplusplus\nextern \"C\" {\n#endif\n\n"

	var tmpArray = make([]string, 0)
	for _, c := range parser.ClassMap {
		if c.Module == m {
			tmpArray = append(tmpArray, c.Name)
		}
	}
	sort.Stable(sort.StringSlice(tmpArray))

	for _, cName := range tmpArray {
		var c = parser.ClassMap[cName]

		for _, e := range c.Enums {
			if isSupportedEnum(e) {
				for _, v := range e.Values {
					if needsCppValueGlue(v) {
						o += fmt.Sprintf("%v;\n", cppEnumHeader(e, v))
					}
				}
			}
		}

		if isSupportedClass(c) {
			for _, f := range c.Functions {
				switch {
				case f.Meta == "signal":
					{
						for _, signalMode := range []string{"Connect", "Disconnect"} {
							f.SignalMode = signalMode
							if i := cppFunctionHeader(f); isSupportedFunction(c, f) {
								o += fmt.Sprintf("%v;\n", i)
							}
						}
						f.SignalMode = ""
					}

				case isGeneric(f):
					{
						for _, m := range jniGenericModes(f) {
							f.TemplateMode = m
							if i := cppFunctionHeader(f); isSupportedFunction(c, f) {
								o += fmt.Sprintf("%v;\n", i)
							}
							f.TemplateMode = ""
						}
					}

				default:
					{
						if i := cppFunctionHeader(f); isSupportedFunction(c, f) && f.Access == "public" {
							o += fmt.Sprintf("%v;\n", i)
						}
					}
				}
			}
		}
	}

	o += "\n#ifdef __cplusplus\n}\n#endif"
	return
}

func CppTemplate(m string) (o string) {

	var tmpArray = make([]string, 0)
	for _, c := range parser.ClassMap {
		if c.Module == m {
			tmpArray = append(tmpArray, c.Name)
		}
	}
	sort.Stable(sort.StringSlice(tmpArray))

	for _, cName := range tmpArray {
		var c = parser.ClassMap[cName]

		if isSupportedClass(c) && (hasVirtualFunction(c) || hasSignalFunction(c)) {
			if !strings.Contains(c.Name, "tomic") {
				o += fmt.Sprintf("class My%v: public %v {\n", c.Name, c.Name)
				for _, m := range []string{"public", "protected"} {
					o += m + ":\n"
					if m == "public" {
						if !c.IsQObjectSubClass() {
							if c.Name != "QMetaType" {
								if hasVirtualFunction(c) {
									o += "\tQString _objectName;\n"
									o += "\tQString objectNameAbs() const { return this->_objectName; };\n"
									o += "\tvoid setObjectNameAbs(const QString &name) { this->_objectName = name; };\n"
								}
							}
						}

						if hasVirtualFunction(c) {
							for _, f := range c.Functions {
								if f.Meta == "constructor" && isSupportedFunction(c, f) {

									var originalInput string
									for _, p := range f.Parameters {
										if p.Name == "" {
											originalInput += "v, "
										} else {
											originalInput += fmt.Sprintf("%v, ", p.Name)
										}
									}
									originalInput = strings.TrimSuffix(originalInput, ", ")

									o += fmt.Sprintf("\tMy%v(%v) : %v(%v) {};\n", f.Class(), strings.Split(strings.Split(f.Signature, "(")[1], ")")[0], f.Class(), originalInput)
								}
							}
						}
					}

					for _, f := range c.Functions {
						if f.Access == m {
							if (f.Meta == "signal" || strings.Contains(f.Virtual, "impure")) && f.Output == "void" {
								if i := cppFunctionSignal(f); isSupportedFunction(c, f) {
									o += fmt.Sprintf("\t%v;\n", i)
								}
							}
						}
					}
				}
				o += "};\n\n"
			}
		}

		for _, e := range c.Enums {
			if isSupportedEnum(e) {
				for _, v := range e.Values {
					if needsCppValueGlue(v) {
						o += fmt.Sprintf("%v\n\n", cppEnum(e, v))
					}
				}
			}
		}

		if isSupportedClass(c) {
			for _, f := range c.Functions {
				if f.Meta == "signal" {
					for _, signalMode := range []string{"Connect", "Disconnect"} {
						f.SignalMode = signalMode
						if i := cppFunction(f); isSupportedFunction(c, f) {
							o += fmt.Sprintf("%v\n\n", i)
						}
					}
					f.SignalMode = ""
				} else {
					if i := cppFunction(f); isSupportedFunction(c, f) && f.Access == "public" {
						o += fmt.Sprintf("%v\n\n", i)
					}
				}
			}
		}

	}

	return managedImportsCpp(m, o)
}

func managedImportsCpp(module, input string) string {
	var tmpIM = make([]string, 0)

	for m := range parser.ClassMap {
		if strings.Contains(input, m) && strings.HasPrefix(m, "Q") && !strings.HasPrefix(m, "Qt") {
			tmpIM = append(tmpIM, m)
		}
	}

	sort.Stable(sort.StringSlice(tmpIM))

	var tmpI string

	if strings.Contains(module, "droid") {
		tmpI += fmt.Sprintf("#include \"%v_android.h\"\n", shortModule(module))
	} else {
		tmpI += fmt.Sprintf("#include \"%v.h\"\n", shortModule(module))
	}

	tmpI += "#include \"_cgo_export.h\"\n\n"

	for _, i := range tmpIM {
		tmpI += fmt.Sprintf("#include <%v>\n", i)
	}

	return tmpI + "\n" + input
}
