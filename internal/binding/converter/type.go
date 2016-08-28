package converter

import (
	"fmt"
	"strings"

	"github.com/therecipe/qt/internal/binding/parser"
)

func goType(f *parser.Function, value string) string {
	var vOld = value

	value = CleanValue(value)

	switch value {
	case "char", "qint8", "uchar", "quint8", "QString", "QByteArray", "QStringList":
		{
			if strings.Contains(vOld, "**") || value == "QStringList" {
				return "[]string"
			}

			return "string"
		}

	case "void", "":
		{
			if strings.Contains(vOld, "*") {
				return "unsafe.Pointer"
			}

			return ""
		}

	case "bool":
		{
			return "bool"
		}

	case "short", "qint16":
		{
			return "int16"
		}

	case "ushort", "unsigned short", "quint16":
		{
			return "uint16"
		}

	case "int", "qint32":
		{
			return "int"
		}

	case "uint", "unsigned int", "quint32":
		{
			return "uint"
		}

	case "long":
		{
			return "int"
		}

	case "ulong", "unsigned long":
		{
			return "uint"
		}

	case "longlong", "long long", "qlonglong", "qint64":
		{
			return "int64"
		}

	case "ulonglong", "unsigned long long", "qulonglong", "quint64":
		{
			return "uint64"
		}

	case "float":
		{
			return "float32"
		}

	case "double", "qreal":
		{
			return "float64"
		}

	case "uintptr_t", "uintptr", "quintptr", "WId":
		{
			return "uintptr"
		}

		//non std types

	case "T":
		{
			switch f.TemplateMode {
			case "Boolean":
				{
					return "bool"
				}

			case "Int":
				{
					return "int"
				}

			case "Void":
				{
					return ""
				}
			}

			if module(f) == "androidextras" && f.Name != "object" {
				return fmt.Sprintf("interface{}")
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
	}

	switch {
	case isEnum(f.Class(), value):
		{
			if c, exists := parser.ClassMap[class(cppEnum(f, value, false))]; exists && module(c.Module) != module(f) && module(c.Module) != "" {
				if parser.ClassMap[f.Class()].WeakLink[c.Module] {
					return "int64"
				}
				return fmt.Sprintf("%v.%v", module(c.Module), goEnum(f, value))
			}
			return goEnum(f, value)
		}

	case isClass(value):
		{
			if m := module(parser.ClassMap[value].Module); m != module(f) {
				if parser.ClassMap[f.Class()].WeakLink[parser.ClassMap[value].Module] {
					return "unsafe.Pointer"
				}
				return fmt.Sprintf("%v.%v", m, value)
			}

			if f.TemplateMode == "String" {
				return "string"
			}

			return value
		}
	}

	f.Access = fmt.Sprintf("unsupported_goType(%v)", value)
	return f.Access
}

func cgoType(f *parser.Function, value string) string {

	var vOld = value

	value = CleanValue(value)

	switch value {
	case "char", "qint8", "uchar", "quint8", "QString", "QByteArray", "QStringList":
		{
			return "*C.char"
		}

	case "void", "":
		{
			if strings.Contains(vOld, "*") {
				return "unsafe.Pointer"
			}

			return ""
		}

	case "bool":
		{
			return "C.char"
		}

	case "short", "qint16":
		{
			return "C.short"
		}

	case "ushort", "unsigned short", "quint16":
		{
			return "C.ushort"
		}

	case "int", "qint32":
		{
			return "C.int"
		}

	case "uint", "unsigned int", "quint32":
		{
			return "C.uint"
		}

	case "long":
		{
			return "C.long"
		}

	case "ulong", "unsigned long":
		{
			return "C.ulong"
		}

	case "longlong", "long long", "qlonglong", "qint64":
		{
			return "C.longlong"
		}

	case "ulonglong", "unsigned long long", "qulonglong", "quint64":
		{
			return "C.ulonglong"
		}

	case "float":
		{
			return "C.float"
		}

	case "double", "qreal":
		{
			return "C.double"
		}

	case "uintptr_t", "uintptr", "quintptr", "WId":
		{
			return "C.uintptr_t"
		}
	}

	switch {
	case isEnum(f.Class(), value):
		{
			return "C.longlong"
		}

	case isClass(value):
		{
			return "unsafe.Pointer"
		}
	}

	f.Access = fmt.Sprintf("unsupported_cgoType(%v)", value)
	return f.Access
}

func cppType(f *parser.Function, value string) string {
	var vOld = value

	value = CleanValue(value)

	switch value {
	case "char", "qint8", "uchar", "quint8", "QString", "QByteArray", "QStringList":
		{
			return "char*"
		}

	case "void", "":
		{
			if strings.Contains(vOld, "*") {
				return "void*"
			}

			return "void"
		}

	case "bool":
		{
			return "char"
		}

	case "short", "qint16":
		{
			return "short"
		}

	case "ushort", "unsigned short", "quint16":
		{
			return "unsigned short"
		}

	case "int", "qint32":
		{
			return "int"
		}

	case "uint", "unsigned int", "quint32":
		{
			return "unsigned int"
		}

	case "long":
		{
			return "long"
		}

	case "ulong", "unsigned long":
		{
			return "unsigned long"
		}

	case "longlong", "long long", "qlonglong", "qint64":
		{
			return "long long"
		}

	case "ulonglong", "unsigned long long", "qulonglong", "quint64":
		{
			return "unsigned long long"
		}

	case "float":
		{
			return "float"
		}

	case "double", "qreal":
		{
			return "double"
		}

	case "uintptr_t", "uintptr", "quintptr", "WId":
		{
			return "uintptr_t"
		}

		//non std types

	case "T":
		{
			switch f.TemplateMode {
			case "Boolean":
				{
					return "char"
				}

			case "Int":
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
			var tmp = make([]string, 10)
			for i := 0; i < 10; i++ {
				if i == 9 {
					tmp[i] = "void*"
				} else {
					tmp[i] = fmt.Sprintf("void* v%v", i)
				}
			}
			return strings.Join(tmp, ", ")
		}
	}

	switch {
	case isEnum(f.Class(), value):
		{
			return "long long"
		}

	case isClass(value):
		{
			return "void*"
		}
	}

	f.Access = fmt.Sprintf("unsupported_cppType(%v)", value)
	return f.Access
}
