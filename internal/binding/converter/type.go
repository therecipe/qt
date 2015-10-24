package converter

import "github.com/therecipe/qt/internal/binding/parser"

func goType(f *parser.Function, value string) string {

	value = cleanValue(value)

	switch value {
	case "uchar", "char", "QString", "QUrl", "QVariant":
		{
			return "string"
		}

	case "QStringList":
		{
			return "[]string"
		}

	case "void":
		{
			return ""
		}

	case "bool", "int", "":
		{
			return value
		}
	}

	switch {
	case isEnum(f.Class(), value):
		{
			if c, exists := parser.ClassMap[class(cppEnum(f, value, false))]; exists && module(c.Module) != module(f) && module(c.Module) != "" {
				return module(c.Module) + "." + goEnum(f, value)
			}
			return goEnum(f, value)
		}

	case isClass(value):
		{
			if m := module(parser.ClassMap[value].Module); m != module(f) {
				return m + "." + value
			}
			return value
		}
	}

	f.Access = "unsupported_goType"
	return f.Access
}

func cgoType(f *parser.Function, value string) string {

	value = cleanValue(value)

	switch value {
	case "uchar", "char", "QString", "QUrl", "QVariant", "QStringList":
		{
			return "*C.char"
		}

	case "bool", "int":
		{
			return "C.int"
		}

	case "void", "":
		{
			return ""
		}
	}

	switch {
	case isEnum(f.Class(), value):
		{
			return "C.int"
		}

	case isClass(value):
		{
			return "unsafe.Pointer"
		}
	}

	f.Access = "unsupported_cgoType"
	return f.Access
}

func cppType(f *parser.Function, value string) string {

	value = cleanValue(value)

	switch value {
	case "uchar", "char", "QString", "QStringList", "QUrl", "QVariant":
		{
			return "char*"
		}

	case "bool", "int":
		{
			return "int"
		}

	case "void", "":
		{
			return "void"
		}
	}

	switch {
	case isEnum(f.Class(), value):
		{
			return "int"
		}

	case isClass(value):
		{
			return "QtObjectPtr"
		}
	}

	f.Access = "unsupported_cppType"
	return f.Access
}
