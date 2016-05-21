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
	value = cleanValue(value)

	switch value {
	case "QStringList":
		{
			return fmt.Sprintf("C.CString(strings.Join(%v, \"|\"))", name)
		}

	case "uchar", "char", "QString", "QByteArray":
		{
			if strings.Contains(vOld, "&") {
				if strings.Contains(vOld, "const") {
					return fmt.Sprintf("C.CString(%v)", name)
				}
				return fmt.Sprintf("C.CString(%v)", name)
			}

			if strings.Contains(vOld, "**") {
				return fmt.Sprintf("C.CString(strings.Join(%v, \"|\"))", name)
			}
			return fmt.Sprintf("C.CString(%v)", name)
		}

	case "bool":
		{
			return fmt.Sprintf("C.int(qt.GoBoolToInt(%v))", name)
		}

	case "int", "long":
		{
			return fmt.Sprintf("C.%v(%v)", value, name)
		}

	case "qreal":
		{
			if strings.Contains(vOld, "*") {
				f.Access = fmt.Sprintf("unsupported_goInput(%v)", value)
				return f.Access
			}

			return fmt.Sprintf("C.double(%v)", name)
		}

	case "qint64":
		{
			if strings.Contains(vOld, "*") {
				f.Access = fmt.Sprintf("unsupported_goInput(%v)", value)
				return f.Access
			}

			return fmt.Sprintf("C.longlong(%v)", name)
		}

	case "WId":
		{
			if strings.Contains(vOld, "*") {
				f.Access = fmt.Sprintf("unsupported_goInput(%v)", value)
				return f.Access
			}

			return fmt.Sprintf("C.ulonglong(%v)", name)
		}

	case "jclass", "jobject":
		{
			return name
		}

	case "...":
		{
			var tmp string
			for i := 0; i < 10; i++ {
				tmp += fmt.Sprintf("p%v, ", i)
			}
			return strings.TrimSuffix(tmp, ", ")
		}

	case "T":
		{
			switch f.TemplateMode {
			case "Int":
				{
					return fmt.Sprintf("C.int(%v)", name)
				}

			case "Boolean":
				{
					return fmt.Sprintf("C.int(qt.GoBoolToInt(%v))", name)
				}
			}

			if module(f) == "androidextras" {
				return "p0"
			}
		}

	case "void":
		{
			if strings.Contains(vOld, "*") {
				return name
			}
		}
	}

	switch {
	case isEnum(f.Class(), value):
		{
			return fmt.Sprintf("C.int(%v)", name)
		}

	case isClass(value):
		{
			if m := module(parser.ClassMap[value].Module); m != module(f) {
				if parser.ClassMap[f.Class()].WeakLink[parser.ClassMap[value].Module] {
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
	value = cleanValue(value)

	switch value {
	case "QStringList":
		{
			if strings.Contains(vOld, "*") {
				return fmt.Sprintf("new QStringList(%v.split(\"|\", QString::SkipEmptyParts))", cppInput(name, "QString", f))
			}
			return fmt.Sprintf("%v.split(\"|\", QString::SkipEmptyParts)", cppInput(name, "QString", f))
		}

	case "QString", "QByteArray":
		{
			if strings.Contains(vOld, "&") {
				if strings.Contains(vOld, "const") {
					return fmt.Sprintf("%v(%v)", value, name)
				}
				return fmt.Sprintf("*(new %v(%v))", value, name)
			}

			if strings.Contains(vOld, "*") {
				return fmt.Sprintf("new %v(%v)", value, name)
			}

			return fmt.Sprintf("%v(%v)", value, name)
		}

	case "bool":
		{
			if strings.Contains(vOld, "*") {
				return "NULL"
			}
			return fmt.Sprintf("%v != 0", name)
		}

	case "int", "long":
		{
			if strings.Contains(vOld, "*") {
				return fmt.Sprintf("&%v", name)
			}

			if strings.Contains(vOld, "&") && name == "argc" {
				return "argcs"
			}

			return name
		}

	case "char":
		{
			if strings.Contains(vOld, "const") {
				if strings.Contains(vOld, "*") {
					return fmt.Sprintf("const_cast<const %v*>(%v)", value, name)
				}
			}

			if strings.Contains(vOld, "**") {
				return "argvs"
				//return "const_cast<char **>(argvs.data())"
			}

			if strings.Contains(vOld, "*") {
				return name
			}

			return fmt.Sprintf("*%v", name)
		}

	case "qreal":
		{
			if strings.Contains(vOld, "*") {
				f.Access = fmt.Sprintf("unsupported_cppInput(%v)", value)
				return f.Access
			}

			return fmt.Sprintf("static_cast<double>(%v)", name)
		}

	case "qint64":
		{
			if strings.Contains(vOld, "*") {
				f.Access = fmt.Sprintf("unsupported_cppInput(%v)", value)
				return f.Access
			}

			return fmt.Sprintf("static_cast<long long>(%v)", name)
		}

	case "WId":
		{
			if strings.Contains(vOld, "*") {
				f.Access = fmt.Sprintf("unsupported_cppInput(%v)", value)
				return f.Access
			}

			return fmt.Sprintf("static_cast<unsigned long long>(%v)", name)
		}

	case "jclass", "jobject":
		{
			return fmt.Sprintf("static_cast<%v>(%v)", value, name)
		}

	case "...":
		{
			var tmp string
			for i := 0; i < 10; i++ {
				if i == 9 {
					tmp += fmt.Sprintf("static_cast<jobject>(%v), ", name)
				} else {
					tmp += fmt.Sprintf("static_cast<jobject>(%v%v), ", name, i)
				}
			}
			return strings.TrimSuffix(tmp, ", ")
		}

	case "T":
		{
			switch f.TemplateMode {
			case "Int", "Boolean":
				{
					return name
				}
			}

			if module(f) == "androidextras" {
				return fmt.Sprintf("static_cast<jobject>(%v)", name)
			}
		}

	case "void":
		{
			if strings.Contains(vOld, "*") {
				return name
			}
		}
	}

	switch {
	case isEnum(f.Class(), value):
		{
			if !strings.Contains(vOld, "*") {
				return fmt.Sprintf("static_cast<%v>(%v)", cppEnum(f, value, false), name)
			}
		}

	case isClass(value):
		{
			if strings.Contains(vOld, "*") {
				if strings.Contains(vOld, "&") {
					break
				}
				return fmt.Sprintf("static_cast<%v*>(%v)", value, name)
			}

			return fmt.Sprintf("*static_cast<%v*>(%v)", value, name)
		}
	}

	f.Access = fmt.Sprintf("unsupported_cppInput(%v)", value)
	return f.Access
}
