package converter

import (
	"fmt"
	"strings"

	"github.com/therecipe/qt/internal/binding/parser"
)

func GoInput(name, value string, f *parser.Function) string {
	return goInput(name, value, f)
}

func goInput(name, value string, f *parser.Function) string {
	var vOld = value

	name = cleanName(name, value)
	value = CleanValue(value)

	switch value {
	case "char", "qint8":
		{
			if strings.Contains(vOld, "**") {
				return fmt.Sprintf("C.CString(strings.Join(%v, \"|\"))", name)
			}

			return fmt.Sprintf("C.CString(%v)", name)
		}

	case "uchar", "quint8", "QString":
		{
			return fmt.Sprintf("C.CString(%v)", name)
		}

	case "QStringList":
		{
			return fmt.Sprintf("C.CString(strings.Join(%v, \"|\"))", name)
		}

	case "void" /*, ""*/ :
		{
			if strings.Contains(vOld, "*") {
				return name
			}
		}

	case "bool":
		{
			return fmt.Sprintf("C.char(int8(qt.GoBoolToInt(%v)))", name)
		}

	case "short", "qint16":
		{
			return fmt.Sprintf("C.short(%v)", name)
		}

	case "ushort", "unsigned short", "quint16":
		{
			return fmt.Sprintf("C.ushort(%v)", name)
		}

	case "int", "qint32":
		{
			return fmt.Sprintf("C.int(int32(%v))", name)
		}

	case "uint", "unsigned int", "quint32":
		{
			return fmt.Sprintf("C.uint(uint32(%v))", name)
		}

	case "long":
		{
			return fmt.Sprintf("C.long(int32(%v))", name)
		}

	case "ulong", "unsigned long":
		{
			return fmt.Sprintf("C.ulong(uint32(%v))", name)
		}

	case "longlong", "long long", "qlonglong", "qint64":
		{
			return fmt.Sprintf("C.longlong(%v)", name)
		}

	case "ulonglong", "unsigned long long", "qulonglong", "quint64":
		{
			return fmt.Sprintf("C.ulonglong(%v)", name)
		}

	case "float":
		{
			return fmt.Sprintf("C.float(%v)", name)
		}

	case "double", "qreal":
		{
			return fmt.Sprintf("C.double(%v)", name)
		}

	case "uintptr_t", "uintptr", "quintptr", "WId":
		{
			return fmt.Sprintf("C.uintptr_t(%v)", name)
		}

		//non std types

	case "T":
		{
			switch f.TemplateModeJNI {
			case "Boolean":
				{
					return fmt.Sprintf("C.char(int8(qt.GoBoolToInt(%v)))", name)
				}

			case "Int":
				{
					return fmt.Sprintf("C.int(int32(%v))", name)
				}
			}

			if module(f) == "androidextras" {
				return "p0"
			}
		}

	case "JavaVM", "jclass", "jobject":
		{
			return name
		}

	case "...":
		{
			var tmp = make([]string, 10)
			for i := range tmp {
				tmp[i] = fmt.Sprintf("p%v", i)
			}
			return strings.Join(tmp, ", ")
		}
	}

	switch {
	case isEnum(f.ClassName(), value):
		{
			return fmt.Sprintf("C.longlong(%v)", name)
		}

	case isClass(value):
		{
			if m := module(parser.CurrentState.ClassMap[value].Module); m != module(f) {
				if _, exists := parser.CurrentState.ClassMap[f.ClassName()].WeakLink[parser.CurrentState.ClassMap[value].Module]; exists {
					return name
				}
				return fmt.Sprintf("%v.PointerFrom%v(%v)", m, strings.Title(value), name)
			}
			return fmt.Sprintf("PointerFrom%v(%v)", strings.Title(value), name)
		}
	}

	f.Access = fmt.Sprintf("unsupported_goInput(%v)", value)
	return f.Access
}

func CppInput(name, value string, f *parser.Function) string {
	return cppInput(name, value, f)
}

func cppInput(name, value string, f *parser.Function) string {
	var vOld = value

	name = cleanName(name, value)
	value = CleanValue(value)

	switch value {
	case "char", "qint8":
		{
			if strings.Contains(vOld, "**") && name == "argv" {
				return "argvs"
			}

			if strings.Contains(vOld, "*") {
				if strings.Contains(vOld, "const") {
					return fmt.Sprintf("const_cast<const %v*>(%v)", value, name)
				}
				return name
			}

			return fmt.Sprintf("*%v", name)
		}

	case "uchar", "quint8":
		{
			if strings.Contains(vOld, "*") {
				if strings.Contains(vOld, "const") {
					return fmt.Sprintf("const_cast<const %v*>(static_cast<%v*>(static_cast<void*>(%v)))", value, value, name)
				}
				return fmt.Sprintf("static_cast<%v*>(static_cast<void*>(%v))", value, name)
			}

			return fmt.Sprintf("*static_cast<%v*>(static_cast<void*>(%v))", value, name)
		}

	case "QString":
		{
			if strings.Contains(vOld, "*") {
				return fmt.Sprintf("new QString(%v)", name)
			}

			if strings.Contains(vOld, "&") && !strings.Contains(vOld, "const") {
				return fmt.Sprintf("*(%v)", cppInput(name, "QString*", f))
			}

			return fmt.Sprintf("QString(%v)", name)
		}

	case "QStringList":
		{
			if strings.Contains(vOld, "*") {
				return fmt.Sprintf("new QStringList(%v)", cppInput(name, "QStringList", f))
			}

			if strings.Contains(vOld, "&") && !strings.Contains(vOld, "const") {
				return fmt.Sprintf("*(%v)", cppInput(name, "QStringList*", f))
			}

			return fmt.Sprintf("QString(%v).split(\"|\", QString::SkipEmptyParts)", name)
		}

	case "void" /*, ""*/ :
		{
			if strings.Contains(vOld, "*") {
				return name
			}
		}

	case "bool":
		{
			if strings.Contains(vOld, "*") {
				return "NULL"
			}

			return fmt.Sprintf("%v != 0", name)
		}

	case
		"short", "qint16",
		"ushort", "unsigned short", "quint16",

		"int", "qint32",
		"uint", "unsigned int", "quint32",

		"long",
		"ulong", "unsigned long",

		"longlong", "long long", "qlonglong", "qint64",
		"ulonglong", "unsigned long long", "qulonglong", "quint64",

		"float",
		"double", "qreal",

		"uintptr_t", "uintptr", "quintptr", "WId":
		{
			if strings.Contains(vOld, "&") && name == "argc" {
				return "argcs"
			}

			if strings.Contains(vOld, "*") {
				if strings.Contains(vOld, "const") {
					return fmt.Sprintf("const_cast<const %v*>(&%v)", value, name)
				}
				return fmt.Sprintf("&%v", name)
			}

			return name
		}

		//non std types

	case "T":
		{
			switch f.TemplateModeJNI {
			case "Boolean", "Int":
				{
					return name
				}
			}

			if module(f) == "androidextras" {
				return fmt.Sprintf("static_cast<jobject>(%v)", name)
			}
		}

	case "JavaVM", "jclass", "jobject":
		{
			return fmt.Sprintf("static_cast<%v>(%v)", value, name)
		}

	case "...":
		{
			var tmp = make([]string, 10)
			for i := range tmp {
				if i == 9 {
					tmp[i] = fmt.Sprintf("static_cast<jobject>(%v)", name)
				} else {
					tmp[i] = fmt.Sprintf("static_cast<jobject>(%v%v)", name, i)
				}
			}
			return strings.Join(tmp, ", ")
		}
	}

	switch {
	case isEnum(f.ClassName(), value):
		{
			if !strings.Contains(vOld, "*") {
				if f.Meta == parser.SLOT && f.SignalMode == "" && value == "Qt::Alignment" {
					return fmt.Sprintf("static_cast<%v>(static_cast<%v>(%v))", value, cppEnum(f, value, false), name)
				}
				return fmt.Sprintf("static_cast<%v>(%v)", cppEnum(f, value, false), name)
			}
		}

	case isClass(value):
		{
			if strings.Contains(vOld, "*") && strings.Contains(vOld, "&") {
				break
			}

			if parser.CurrentState.ClassMap[value].Fullname != "" {
				value = parser.CurrentState.ClassMap[value].Fullname
			}

			if strings.Contains(vOld, "*") {
				return fmt.Sprintf("static_cast<%v*>(%v)", value, name)
			}

			return fmt.Sprintf("*static_cast<%v*>(%v)", value, name)
		}
	}

	f.Access = fmt.Sprintf("unsupported_cppInput(%v)", value)
	return f.Access
}
