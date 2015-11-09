package templater

import (
	"fmt"
	"strings"

	"github.com/therecipe/qt/internal/binding/parser"
)

func HTemplate(c *parser.Class) (o string) {
	o += "#ifdef __cplusplus\nextern \"C\" {\n#endif\n\n"

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
			if f.Meta == "signal" && !f.Overload {
				for _, signalMode := range []string{"Connect", "Disconnect"} {
					f.SignalMode = signalMode
					if i := cppFunctionHeader(f); isSupportedFunction(c, f) {
						o += fmt.Sprintf("%v;\n", i)
					}
				}
				f.SignalMode = ""
			} else {
				if i := cppFunctionHeader(f); isSupportedFunction(c, f) {
					o += fmt.Sprintf("%v;\n", i)
				}
			}
		}
	}

	o += "\n#ifdef __cplusplus\n}\n#endif"
	return
}

func CppTemplate(c *parser.Class) (o string) {
	for _, h := range []string{strings.ToLower(c.Name)} {
		o += fmt.Sprintf("#include \"%v.h\"\n", h)
	}

	if strings.Contains(c.Name, "List") {
		o += "#include <QList>\n"
	}

	o += "PLACEHOLDER_IMPORTS"

	if isSupportedFile(c) {
		o += fmt.Sprintf("#include <%v>\n", c.Name)
	}

	o += "#include \"_cgo_export.h\"\n\n"

	if isSupportedClass(c) {
		if !strings.Contains(c.Name, "tomic") {
			o += fmt.Sprintf("class My%v: public %v {\npublic:\n", c.Name, c.Name)
			for _, f := range c.Functions {
				if f.Meta == "signal" && !f.Overload {
					if i := cppFunctionSignal(f); isSupportedFunction(c, f) {
						o += fmt.Sprintf("%v;\n", i)
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
			if f.Meta == "signal" && !f.Overload {
				for _, signalMode := range []string{"Connect", "Disconnect"} {
					f.SignalMode = signalMode
					if i := cppFunction(f); isSupportedFunction(c, f) {
						o += fmt.Sprintf("%v\n\n", i)
					}
				}
				f.SignalMode = ""
			} else {
				if i := cppFunction(f); isSupportedFunction(c, f) {
					o += fmt.Sprintf("%v\n\n", i)
				}
			}
		}
	}

	return managedImportsCpp(c.Name, o)
}

func managedImportsCpp(className, input string) string {
	var tmpIM = map[string]bool{"QString": true, "QVariant": true, "QUrl": true, "QModelIndex": true}

	for m := range parser.ClassMap {
		if strings.Contains(input, m) && m != className && strings.HasPrefix(m, "Q") && !strings.HasPrefix(m, "Qt") {
			tmpIM[m] = true
		}
	}

	var tmpI string

	for i := range tmpIM {
		tmpI += fmt.Sprintf("#include <%v>\n", i)
	}

	return strings.Replace(input, "PLACEHOLDER_IMPORTS", tmpI, -1)
}
