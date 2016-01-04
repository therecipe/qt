package templater

import (
	"fmt"
	"sort"
	"strings"

	"github.com/therecipe/qt/internal/binding/converter"
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
		var virtuals = make(map[string]bool)

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
								virtuals[f.Name+f.OverloadNumber] = true
								o += fmt.Sprintf("%v;\n", i)
							}
						}
						f.SignalMode = ""
					}

					var tmpMeta = f.Meta
					f.Meta = "plain"
					if i := cppFunctionHeader(f); isSupportedFunction(c, f) && !converter.IsPrivateSignal(f) {
						o += fmt.Sprintf("%v;\n", i)
					}
					f.Meta = tmpMeta

				case strings.Contains(f.Virtual, "impure") && !strings.Contains(f.Meta, "structor"):
					{
						if _, exists := virtuals[f.Name+f.OverloadNumber]; !exists {

							var tmpMeta = f.Meta
							f.Meta = "plain"
							if i := cppFunctionHeader(f); isSupportedFunction(c, f) {
								virtuals[f.Name+f.OverloadNumber] = true
								o += fmt.Sprintf("%v;\n", i)
								o += strings.Replace(fmt.Sprintf("%v;\n", i), "_"+strings.Title(f.Name), "_"+strings.Title(f.Name)+"Default", -1)
							}
							f.Meta = tmpMeta
						}
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
						if i := cppFunctionHeader(f); isSupportedFunction(c, f) {
							virtuals[f.Name+f.OverloadNumber] = true
							o += fmt.Sprintf("%v;\n", i)
						}
					}
				}
			}
		}

		for _, bcName := range c.GetAllBases([]string{}) {
			if bc, exists := parser.ClassMap[bcName]; exists {
				for _, f := range bc.Functions {

					if strings.Contains(f.Virtual, "impure") && f.Output == "void" {

						var tmpMeta = f.Meta
						f.Meta = "plain"
						if i := cppFunctionHeader(f); isSupportedFunction(bc, f) && isSupportedClass(bc) {
							if _, exists := virtuals[f.Name+f.OverloadNumber]; !exists {
								virtuals[f.Name+f.OverloadNumber] = true
								if !isBlockedVirtual(f.Name, c.Name) {
									o += strings.Replace(fmt.Sprintf("%v;\n", i), bc.Name, c.Name, -1)
									o += strings.Replace(strings.Replace(fmt.Sprintf("%v;\n", i), bc.Name, c.Name, -1), "_"+strings.Title(f.Name), "_"+strings.Title(f.Name)+"Default", -1)
								}
							}
						}
						f.Meta = tmpMeta
					}

				}
			}
		}

	}

	o += "\n#ifdef __cplusplus\n}\n#endif"
	return
}

func CppTemplate(module string) (o string) {

	var tmpArray = make([]string, 0)
	for _, c := range parser.ClassMap {
		if c.Module == module {
			tmpArray = append(tmpArray, c.Name)
		}
	}
	sort.Stable(sort.StringSlice(tmpArray))

	for _, cName := range tmpArray {
		var c = parser.ClassMap[cName]
		var virtuals = make(map[string]bool)

		if isSupportedClass(c) && (hasVirtualFunction(c) || hasSignalFunction(c)) {
			if !strings.Contains(c.Name, "tomic") {
				o += fmt.Sprintf("class My%v: public %v {\n", c.Name, c.Name)

				o += "public:\n"
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

				for _, f := range c.Functions {
					if (f.Meta == "signal" || strings.Contains(f.Virtual, "impure")) && f.Output == "void" {
						if i := cppFunctionSignal(f); isSupportedFunction(c, f) {
							if strings.Contains(f.Virtual, "impure") {
								if _, exists := virtuals[f.Name+f.OverloadNumber]; !exists {
									virtuals[f.Name+f.OverloadNumber] = true
									if !isBlockedVirtual(f.Name, c.Name) {
										o += fmt.Sprintf("\t%v;\n", i)
									}
								}
							} else {
								o += fmt.Sprintf("\t%v;\n", i)
							}
						}
					} else {
						virtuals[f.Name+f.OverloadNumber] = true
					}
				}

				for _, bcName := range c.GetAllBases([]string{}) {
					if bc, exists := parser.ClassMap[bcName]; exists {
						for _, f := range bc.Functions {
							if strings.Contains(f.Virtual, "impure") && f.Output == "void" {
								if i := cppFunctionSignal(f); isSupportedFunction(bc, f) && isSupportedClass(bc) {
									if _, exists := virtuals[f.Name+f.OverloadNumber]; !exists {
										virtuals[f.Name+f.OverloadNumber] = true
										if !isBlockedVirtual(f.Name, c.Name) {
											o += strings.Replace(fmt.Sprintf("\t%v;\n", i), bc.Name, c.Name, -1)
										}
									}
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

		virtuals = make(map[string]bool)
		if isSupportedClass(c) {
			for _, f := range c.Functions {

				switch {
				case f.Meta == "signal":
					{
						for _, signalMode := range []string{"Connect", "Disconnect"} {
							f.SignalMode = signalMode
							if i := cppFunction(f); isSupportedFunction(c, f) {
								virtuals[f.Name+f.OverloadNumber] = true
								o += fmt.Sprintf("%v\n\n", i)
							}
						}
						f.SignalMode = ""

						var tmpMeta = f.Meta
						f.Meta = "plain"
						if i := cppFunction(f); isSupportedFunction(c, f) && !converter.IsPrivateSignal(f) {
							o += fmt.Sprintf("%v\n\n", i)
						}
						f.Meta = tmpMeta
					}

				default:
					{
						if i := cppFunction(f); isSupportedFunction(c, f) {
							if _, exists := virtuals[f.Name+f.OverloadNumber]; !exists {
								virtuals[f.Name+f.OverloadNumber] = true

								var normal = fmt.Sprintf("%v\n\n", i)

								if strings.Contains(f.Virtual, "impure") && f.Output == "void" && (hasVirtualFunction(c) || hasSignalFunction(c)) {
									normal = strings.Replace(normal, "static_cast<"+c.Name+"*>(ptr)", "static_cast<My"+c.Name+"*>(ptr)", -1)
								}

								o += normal

								if strings.Contains(f.Virtual, "impure") && f.Output == "void" {
									var tmp = strings.Replace(fmt.Sprintf("%v\n\n", i), "_"+strings.Title(f.Name), "_"+strings.Title(f.Name)+"Default", -1)
									tmp = strings.Replace(tmp, "->"+f.Name, "->"+c.Name+"::"+f.Name, -1)

									o += tmp
								}
							}
						}
					}
				}
			}

			for _, bcName := range c.GetAllBases([]string{}) {
				if bc, exists := parser.ClassMap[bcName]; exists {
					for _, f := range bc.Functions {

						if strings.Contains(f.Virtual, "impure") && f.Output == "void" {

							if i := cppFunction(f); isSupportedFunction(bc, f) && isSupportedClass(bc) {
								if _, exists := virtuals[f.Name+f.OverloadNumber]; !exists {
									virtuals[f.Name+f.OverloadNumber] = true
									if !isBlockedVirtual(f.Name, c.Name) {

										var normal = strings.Replace(fmt.Sprintf("%v\n\n", i), bc.Name, c.Name, -1)
										if strings.Contains(f.Virtual, "impure") && f.Output == "void" && (hasVirtualFunction(c) || hasSignalFunction(c)) {
											normal = strings.Replace(normal, "static_cast<"+c.Name+"*>(ptr)", "static_cast<My"+c.Name+"*>(ptr)", -1)
										}

										o += normal

										var tmp = strings.Replace(strings.Replace(fmt.Sprintf("%v\n\n", i), bc.Name, c.Name, -1), "_"+strings.Title(f.Name), "_"+strings.Title(f.Name)+"Default", -1)
										tmp = strings.Replace(tmp, "->"+f.Name, "->"+c.Name+"::"+f.Name, -1)

										o += tmp
									}
								}
							}
						}
					}
				}
			}
		}
	}

	return managedImportsCpp(module, o)
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

	tmpI += "#define protected public\n\n"

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
