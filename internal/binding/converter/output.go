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

	case "T", "JavaVM", "jclass":
		{
			return fmt.Sprintf("unsafe.Pointer(%v)", name)
		}

	case "qreal":
		{
			return fmt.Sprintf("float64(%v)", name)
		}
	}

	switch {
	case isEnum(f.Class(), value):
		{
			if c, exists := parser.ClassMap[class(cppEnum(f, value, false))]; exists && module(c.Module) != module(f) && module(c.Module) != "" {
				return fmt.Sprintf("%v.%v(%v)", module(c.Module), goEnum(f, value), name)
			}
			return fmt.Sprintf("%v(%v)", goEnum(f, value), name)
		}

	case isClass(value):
		{
			if m := module(parser.ClassMap[value].Module); m != module(f) {
				return fmt.Sprintf("%v.New%vFromPointer(%v)", m, value, name)
			}
			return fmt.Sprintf("New%vFromPointer(%v)", value, name)
		}
	}

	f.Access = "unsupported_goOutput"
	return f.Access
}

func cgoOutput(name string, value string, f *parser.Function) string {

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
			return ""
		}
	}

	switch {
	case isEnum(f.Class(), value):
		{
			if c, exists := parser.ClassMap[class(cppEnum(f, value, false))]; exists && module(c.Module) != module(f) && module(c.Module) != "" {
				return fmt.Sprintf("%v.%v(%v)", module(c.Module), goEnum(f, value), name)
			}
			return fmt.Sprintf("%v(%v)", goEnum(f, value), name)
		}

	case isClass(value):
		{
			if m := module(parser.ClassMap[value].Module); m != module(f) {
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

	case "bool", "int", "void", "", "T", "JavaVM", "jclass":
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
			case "QModelIndex":
				{
					return fmt.Sprintf("%v.internalPointer()", name)
				}
			case "QJSValue", "QScriptValue", "QVariant", "QStringRef", "QDateTime", "QTimeZone", "QRegularExpressionMatchIterator", "QRegularExpressionMatch", "QRegularExpression", "QDir", "QByteArray", "QEasingCurve", "QCommandLineOption", "QRegExp", "QJsonObject", "QJsonArray", "QJsonDocument", "QRegion", "QBrush", "QColor":
				{
					return fmt.Sprintf("new %v(%v)", value, name)
				}

			case "QAndroidJniObject":
				{
					return name
				}
			}
		}
	}

	f.Access = "unsupported_cppOutput"
	return f.Access
}
