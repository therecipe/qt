package converter

import (
	"fmt"
	"strings"

	"github.com/therecipe/qt/internal/binding/parser"
)

func goType(f *parser.Function, value string) string {
	var vOld = value

	value = cleanValue(value)

	switch value {
	case "uchar", "char", "QString":
		{
			if strings.Contains(vOld, "**") {
				return "[]string"
			}

			return "string"
		}

	case "QStringList":
		{
			return "[]string"
		}

	case "void":
		{
			if strings.Contains(vOld, "*") {
				return "unsafe.Pointer"
			}
			return ""
		}

	case "bool", "int", "":
		{
			return value
		}

	case "T":
		{
			switch f.TemplateMode {
			case "Int":
				{
					return "int"
				}
			case "Boolean":
				{
					return "bool"
				}
			case "Void":
				return ""
			}

			return "unsafe.Pointer"
		}

	case "JavaVM", "jclass", "jobject":
		{
			return "unsafe.Pointer"
		}

	case "...":
		{
			if parser.ClassMap[f.Class()].Module == "QtAndroidExtras" {
				return "...interface{}"
			}
		}

	case "qreal":
		{
			return "float64"
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
	case "uchar", "char", "QString", "QStringList":
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
	var vOld = value

	value = cleanValue(value)

	switch value {
	case "uchar", "char", "QString", "QStringList":
		{
			return "char*"
		}

	case "bool", "int":
		{
			return "int"
		}

	case "void", "":
		{
			if strings.Contains(vOld, "*") {
				return "void*"
			}
			return "void"
		}

	case "T":
		{
			switch f.TemplateMode {
			case "Int":
				{
					return "int"
				}
			case "Boolean":
				{
					return "int"
				}
			case "Void":
				{
					return "void"
				}
			}

			return "void*"
		}

	case "JavaVM", "jclass", "jobject":
		{
			return "void*"
		}

	case "...":
		{
			var tmp string
			for i := 0; i < 10; i++ {
				if i == 9 {
					tmp += "void*"
				} else {
					tmp += fmt.Sprintf("void* %v%v, ", "v", i)
				}
			}
			return strings.TrimSuffix(tmp, ", ")
		}

	case "qreal":
		{
			return "double"
		}
	}

	switch {
	case isEnum(f.Class(), value):
		{
			return "int"
		}

	case isClass(value):
		{
			return "void*"
		}
	}

	f.Access = "unsupported_cppType"
	return f.Access
}
