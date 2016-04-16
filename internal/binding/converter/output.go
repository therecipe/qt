package converter

import (
	"fmt"
	"strings"

	"github.com/therecipe/qt/internal/binding/parser"
)

func goOutput(name string, value string, f *parser.Function) string {
	var vOld = value

	name = cleanName(name)
	value = cleanValue(value)

	switch value {
	case "QStringList":
		{
			return fmt.Sprintf("strings.Split(%v, \"|\")", goOutput(name, "QString", f))
		}

	case "uchar", "char", "QString":
		{
			return fmt.Sprintf("C.GoString(%v)", name)
		}

	case "int":
		{
			return fmt.Sprintf("int(%v)", name)
		}

	case "bool":
		{
			return fmt.Sprintf("%v != 0", name)
		}

	case "void", "":
		{
			if strings.Contains(vOld, "*") {
				return fmt.Sprintf("unsafe.Pointer(%v)", name)
			}
			return name
		}

	case "T", "JavaVM", "jclass", "jobject":
		{
			switch f.TemplateMode {
			case "Int":
				{
					return fmt.Sprintf("int(%v)", name)
				}

			case "Boolean":
				{
					return fmt.Sprintf("int(%v) != 0", name)
				}

			case "Void":
				{
					return name
				}
			}
			return fmt.Sprintf("unsafe.Pointer(%v)", name)
		}

	case "qreal":
		{
			return fmt.Sprintf("float64(%v)", name)
		}

	case "qint64":
		{
			return fmt.Sprintf("int64(%v)", name)
		}

	case "WId":
		{
			return fmt.Sprintf("uintptr(%v)", name)
		}
	}

	switch {
	case isEnum(f.Class(), value):
		{
			if c, exists := parser.ClassMap[class(cppEnum(f, value, false))]; exists && module(c.Module) != module(f) && module(c.Module) != "" {
				if parser.ClassMap[f.Class()].WeakLink[c.Module] {
					return fmt.Sprintf("int64(%v)", name)
				}
				return fmt.Sprintf("%v.%v(%v)", module(c.Module), goEnum(f, value), name)
			}
			return fmt.Sprintf("%v(%v)", goEnum(f, value), name)
		}

	case isClass(value):
		{
			if m := module(parser.ClassMap[value].Module); m != module(f) {
				if parser.ClassMap[f.Class()].WeakLink[parser.ClassMap[value].Module] {
					return fmt.Sprintf("unsafe.Pointer(%v)", name)
				}
				return fmt.Sprintf("%v.New%vFromPointer(%v)", m, value, name)
			}

			if f.Meta == "constructor" {
				return fmt.Sprintf("new%vFromPointer(%v)", value, name)
			}
			return fmt.Sprintf("New%vFromPointer(%v)", value, name)
		}
	}

	f.Access = "unsupported_goOutput"
	return f.Access
}

func goOutputFailed(value string, f *parser.Function) string {
	var vOld = value

	value = cleanValue(value)

	switch value {
	case "bool":
		{
			return "false"
		}

	case "int", "qreal", "qint64", "WId":
		{
			return "0"
		}

	case "uchar", "char", "QString":
		{
			return "\"\""
		}

	case "QStringList":
		{
			return "make([]string, 0)"
		}

	case "void", "":
		{
			if strings.Contains(vOld, "*") {
				return "nil"
			}
			return ""
		}

	case "T", "JavaVM", "jclass", "jobject":
		{
			switch f.TemplateMode {
			case "Int":
				{
					return "0"
				}

			case "Boolean":
				{
					return "false"
				}

			case "Void":
				{
					return ""
				}
			}

			return "nil"
		}
	}

	switch {
	case isEnum(f.Class(), value):
		{
			return "0"
		}

	case isClass(value):
		{
			return "nil"
		}

	default:
		{
			f.Access = "unsupported_GoBodyOutputFailed"
			return f.Access
		}
	}
}

func cgoOutput(name string, value string, f *parser.Function) string {

	var vOld = value

	name = cleanName(name)
	value = cleanValue(value)

	switch value {
	case "QStringList":
		{
			return fmt.Sprintf("strings.Split(%v, \"|\")", cgoOutput(name, "QString", f))
		}

	case "uchar", "char", "QString":
		{
			return fmt.Sprintf("C.GoString(%v)", name)
		}

	case "int":
		{
			return fmt.Sprintf("int(%v)", name)
		}

	case "bool":
		{
			return fmt.Sprintf("%v != 0", cgoOutput(name, "int", f))
		}

	case "void", "":
		{
			if strings.Contains(vOld, "*") {
				return name
			}
			return ""
		}

	case "qreal":
		{
			return fmt.Sprintf("float64(%v)", name)
		}

	case "qint64":
		{
			return fmt.Sprintf("int64(%v)", name)
		}
	}

	switch {
	case isEnum(f.Class(), value):
		{
			if c, exists := parser.ClassMap[class(cppEnum(f, value, false))]; exists && module(c.Module) != module(f) && module(c.Module) != "" {
				if parser.ClassMap[f.Class()].WeakLink[c.Module] {
					return fmt.Sprintf("unsafe.Pointer(%v)", name)
				}
				return fmt.Sprintf("%v.%v(%v)", module(c.Module), goEnum(f, value), name)
			}
			return fmt.Sprintf("%v(%v)", goEnum(f, value), name)
		}

	case isClass(value):
		{
			if m := module(parser.ClassMap[value].Module); m != module(f) {
				if parser.ClassMap[f.Class()].WeakLink[parser.ClassMap[value].Module] {
					return fmt.Sprintf("unsafe.Pointer(%v)", name)
				}
				return fmt.Sprintf("%v.New%vFromPointer(%v)", m, value, name)
			}
			return fmt.Sprintf("New%vFromPointer(%v)", value, name)
		}
	}

	f.Access = "unsupported_cgoOutput"
	return f.Access
}

func cppOutput(name string, value string, f *parser.Function) string {

	var vOld = value

	name = cleanName(name)
	value = cleanValue(value)

	switch value {
	case "QStringList":
		{
			return cppOutput(fmt.Sprintf("%v.join(\"|\")", name), "QString", f)
		}

	case "QString":
		{
			if strings.Contains(vOld, "*") {
				return fmt.Sprintf("%v->toUtf8().data()", name)
			}
			return fmt.Sprintf("%v.toUtf8().data()", name)
		}

	case "bool", "int", "void", "", "T", "JavaVM", "jclass", "jobject":
		{
			if value == "void" {
				if strings.Contains(vOld, "*") {
					if strings.Contains(vOld, "const") {
						return fmt.Sprintf("const_cast<%v*>(%v)", value, name)
					}
					return name
				}
			}

			return name
		}

	case "qreal":
		{
			return fmt.Sprintf("static_cast<double>(%v)", name)
		}

	case "qint64":
		{
			return fmt.Sprintf("static_cast<long long>(%v)", name)
		}

	case "WId":
		{
			return fmt.Sprintf("static_cast<unsigned long long>(%v)", name)
		}
	}

	switch {
	case isEnum(f.Class(), value):
		{
			return name
		}

	case isClass(value):
		{
			if strings.Contains(vOld, "*") {
				if strings.Contains(vOld, "const") {
					return fmt.Sprintf("const_cast<%v*>(%v)", value, name)
				}
				return name
			}

			switch value {
			case "QModelIndex", "QMediaContent", "QIcon", "QUrl", "QJSValue", "QScriptValue", "QVariant", "QStringRef", "QDateTime", "QTimeZone", "QRegularExpressionMatchIterator", "QRegularExpressionMatch", "QRegularExpression", "QDir", "QByteArray", "QEasingCurve", "QCommandLineOption", "QRegExp", "QJsonObject", "QJsonArray", "QJsonDocument", "QRegion", "QBrush", "QColor":
				{
					return fmt.Sprintf("new %v(%v)", value, name)
				}

			case "QAndroidJniObject":
				{
					return fmt.Sprintf("new %v(%v.object())", value, name)
				}

			case "QPoint":
				{
					return fmt.Sprintf("new %v(static_cast<%v>(%v).x(), static_cast<%v>(%v).y())", value, value, name, value, name)
				}

			case "QSize":
				{
					return fmt.Sprintf("new %v(static_cast<%v>(%v).width(), static_cast<%v>(%v).height())", value, value, name, value, name)
				}

			case "QRect":
				{
					return fmt.Sprintf("new %v(static_cast<%v>(%v).x(), static_cast<%v>(%v).y(), static_cast<%v>(%v).width(), static_cast<%v>(%v).height())", value, value, name, value, name, value, name, value, name)
				}
			}
		}
	}

	f.Access = "unsupported_cppOutput"
	return f.Access
}
