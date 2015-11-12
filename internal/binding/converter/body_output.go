package converter

import (
	"github.com/therecipe/qt/internal/binding/parser"
	"strings"
)

func GoBodyOutput(f *parser.Function, name string) string {

	var value = f.Output

	if f.Meta == "constructor" && f.Output == "" {
		value = f.Name
	}

	return goOutput(name, value, f)
}

func GoBodyOutputFailed(value string, f *parser.Function) string {
	var vOld = value

	value = cleanValue(value)

	switch value {
	case "bool":
		return "false"

	case "int", "qreal":
		return "0"

	case "uchar", "char", "QString":
		return "\"\""

	case "QStringList":
		return "make([]string, 0)"

	case "void", "":
		if strings.Contains(vOld, "*") {
			return "nil"
		}
		return ""

	case "T", "JavaVM", "jclass":
		return "nil"
	}

	switch {
	case isEnum(f.Class(), value):
		return "0"

	case isClass(value):
		return "nil"

	default:
		f.Access = "unsupported_GoBodyOutputFailed"
		return f.Access
	}
}

func CppBodyOutput(f *parser.Function, name string) string {
	var value = f.Output

	if f.Meta == "constructor" && f.Output == "" {
		value = f.Name
	}

	return cppOutput(name, value, f)
}
