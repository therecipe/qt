package converter

//TODO: GLchar, GLbyte

import (
	"fmt"
	"strings"

	"github.com/therecipe/qt/internal/binding/parser"
)

func goType(f *parser.Function, value string) string {
	var vOld = value

	value = parser.CleanValue(value)

	switch value {
	case "char", "qint8", "uchar", "quint8", "GLubyte", "QString", "QStringList":
		{
			if strings.Contains(vOld, "**") || value == "QStringList" {
				return "[]string"
			}

			if f.AsError {
				return "error"
			}

			if value == "char" && strings.Count(vOld, "*") == 1 && f.Name == "readData" {
				return "*string"
			}

			return "string"
		}

	case "void", "GLvoid", "":
		{
			if strings.Contains(vOld, "*") {
				return "unsafe.Pointer"
			}

			return ""
		}

	case "bool", "GLboolean":
		{
			return "bool"
		}

	case "short", "qint16", "GLshort":
		{
			return "int16"
		}

	case "ushort", "unsigned short", "quint16", "GLushort":
		{
			return "uint16"
		}

	case "int", "qint32", "GLint", "GLsizei", "GLintptrARB", "GLsizeiptrARB", "GLfixed", "GLclampx":
		{
			return "int"
		}

	case "uint", "unsigned int", "quint32", "GLenum", "GLbitfield", "GLuint":
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

	case "float", "GLfloat", "GLclampf":
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
			switch f.TemplateModeJNI {
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
			if parser.State.ClassMap[f.ClassName()].Module == "QtAndroidExtras" {
				return "...interface{}"
			}
		}
	}

	switch {
	case isEnum(f.ClassName(), value):
		{
			if c, ok := parser.State.ClassMap[class(cppEnum(f, value, false))]; ok && module(c.Module) != module(f) && module(c.Module) != "" {
				if _, ok := parser.State.ClassMap[f.ClassName()].WeakLink[c.Module]; ok {
					return "int64"
				}
				return fmt.Sprintf("%v.%v", module(c.Module), goEnum(f, value))
			}
			return goEnum(f, value)
		}

	case isClass(value):
		{
			if strings.Contains(value, ".") {
				value = strings.Split(value, ".")[1]
			}
			if m := module(parser.State.ClassMap[value].Module); m != module(f) {
				if _, ok := parser.State.ClassMap[f.ClassName()].WeakLink[parser.State.ClassMap[value].Module]; ok {
					return "unsafe.Pointer"
				}
				return fmt.Sprintf("%v.%v", m, value)
			}

			if f.TemplateModeJNI == "String" {
				return "string"
			}

			return value
		}

	case parser.IsPackedList(value):
		{
			return fmt.Sprintf("[]%v%v", func() string {
				if isClass(parser.UnpackedList(value)) && parser.UnpackedList(value) != "QString" && parser.UnpackedList(value) != "QStringList" {
					return "*"
				}
				return ""
			}(), goType(f, parser.UnpackedListDirty(value)))
		}

	case parser.IsPackedMap(value):
		{
			var key, value = parser.UnpackedMapDirty(value)
			return fmt.Sprintf("map[%v%v]%v%v",
				func() string {
					if isClass(parser.CleanValue(key)) && parser.CleanValue(key) != "QString" && parser.CleanValue(key) != "QStringList" {
						return "*"
					}
					return ""
				}(), goType(f, key),

				func() string {
					if isClass(parser.CleanValue(value)) && parser.CleanValue(value) != "QString" && parser.CleanValue(value) != "QStringList" {
						return "*"
					}
					return ""
				}(), goType(f, value))
		}
	}

	f.Access = fmt.Sprintf("unsupported_goType(%v)", value)
	return f.Access
}

func cgoTypeOutput(f *parser.Function, value string) string {
	switch parser.CleanValue(value) {
	case "char", "qint8", "uchar", "quint8", "GLubyte":
		{
			return "*C.char"
		}

	default:
		{
			if parser.IsPackedList(parser.CleanValue(value)) || parser.IsPackedMap(parser.CleanValue(value)) {
				return "unsafe.Pointer"
			}
			return cgoType(f, value)
		}
	}
}

func cgoType(f *parser.Function, value string) string {

	var vOld = value

	value = parser.CleanValue(value)

	switch value {
	case "char", "qint8", "uchar", "quint8", "GLubyte", "QString", "QStringList":
		{
			return fmt.Sprintf("C.struct_%v_PackedString", strings.Title(parser.State.ClassMap[f.ClassName()].Module))
		}

	case "void", "GLvoid", "":
		{
			if strings.Contains(vOld, "*") {
				return "unsafe.Pointer"
			}

			return ""
		}

	case "bool", "GLboolean":
		{
			return "C.char"
		}

	case "short", "qint16", "GLshort":
		{
			return "C.short"
		}

	case "ushort", "unsigned short", "quint16", "GLushort":
		{
			return "C.ushort"
		}

	case "int", "qint32", "GLint", "GLsizei", "GLintptrARB", "GLsizeiptrARB", "GLfixed", "GLclampx":
		{
			return "C.int"
		}

	case "uint", "unsigned int", "quint32", "GLenum", "GLbitfield", "GLuint":
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

	case "float", "GLfloat", "GLclampf":
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
	case isEnum(f.ClassName(), value):
		{
			return "C.longlong"
		}

	case isClass(value):
		{
			return "unsafe.Pointer"
		}

	case parser.IsPackedList(value) || parser.IsPackedMap(value):
		{
			return fmt.Sprintf("C.struct_%v_PackedList", strings.Title(parser.State.ClassMap[f.ClassName()].Module))
		}
	}

	f.Access = fmt.Sprintf("unsupported_cgoType(%v)", value)
	return f.Access
}

func cppTypeInput(f *parser.Function, value string) string {
	switch parser.CleanValue(value) {
	case "char", "qint8", "uchar", "quint8", "GLubyte":
		{
			return "char*"
		}

	default:
		{
			if parser.IsPackedList(parser.CleanValue(value)) || parser.IsPackedMap(parser.CleanValue(value)) {
				return "void*"
			}
			return cppType(f, value)
		}
	}
}

func cppType(f *parser.Function, value string) string {
	var vOld = value

	value = parser.CleanValue(value)

	switch value {
	case "char", "qint8", "uchar", "quint8", "GLubyte", "QString", "QStringList":
		{
			return fmt.Sprintf("struct %v_PackedString", strings.Title(parser.State.ClassMap[f.ClassName()].Module))
		}

	case "void", "GLvoid", "":
		{
			if strings.Contains(vOld, "*") {
				return "void*"
			}

			return "void"
		}

	case "bool", "GLboolean":
		{
			return "char"
		}

	case "short", "qint16", "GLshort":
		{
			return "short"
		}

	case "ushort", "unsigned short", "quint16", "GLushort":
		{
			return "unsigned short"
		}

	case "int", "qint32", "GLint", "GLsizei", "GLintptrARB", "GLsizeiptrARB", "GLfixed", "GLclampx":
		{
			return "int"
		}

	case "uint", "unsigned int", "quint32", "GLenum", "GLbitfield", "GLuint":
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

	case "float", "GLfloat", "GLclampf":
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
			switch f.TemplateModeJNI {
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
	case isEnum(f.ClassName(), value):
		{
			return "long long"
		}

	case isClass(value):
		{
			return "void*"
		}

	case parser.IsPackedList(value) || parser.IsPackedMap(value):
		{
			return fmt.Sprintf("struct %v_PackedList", strings.Title(parser.State.ClassMap[f.ClassName()].Module))
		}
	}

	f.Access = fmt.Sprintf("unsupported_cppType(%v)", value)
	return f.Access
}
