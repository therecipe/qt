package converter

import "github.com/therecipe/qt/internal/binding/parser"

func GoBodyOutput(f *parser.Function, name string) string {

	var value = f.Output

	if f.Meta == "constructor" && f.Output == "" {
		value = f.Name
	}

	return goOutput(name, value, f)
}

func GoBodyOutputFailed(value string, f *parser.Function) string {
	value = cleanValue(value)

	switch value {
	case "bool":
		return "false"

	case "int":
		return "0"

	case "uchar", "char", "QString", "QUrl", "QVariant":
		return "\"\""

	case "QStringList":
		return "make([]string, 0)"

	case "void", "":
		return ""
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
