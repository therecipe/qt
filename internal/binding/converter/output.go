package converter

import (
	"crypto/sha1"
	"encoding/hex"
	"fmt"
	"strings"

	"github.com/therecipe/qt/internal/binding/parser"
)

func goOutput(name, value string, f *parser.Function) string {
	var vOld = value

	name = cleanName(name, value)
	value = CleanValue(value)

	switch value {
	case "char", "qint8", "uchar", "quint8", "QString":
		{
			return fmt.Sprintf("cGoUnpackString(%v)", name)
		}

	case "QStringList":
		{
			return fmt.Sprintf("strings.Split(cGoUnpackString(%v), \"|\")", name)
		}

	case "void", "":
		{
			if strings.Contains(vOld, "*") {
				return fmt.Sprintf("unsafe.Pointer(%v)", name)
			}

			return name
		}

	case "bool":
		{
			return fmt.Sprintf("%v != 0", name)
		}

	case "short", "qint16":
		{
			return fmt.Sprintf("int16(%v)", name)
		}

	case "ushort", "unsigned short", "quint16":
		{
			return fmt.Sprintf("uint16(%v)", name)
		}

	case "int", "qint32":
		{
			return fmt.Sprintf("int(int32(%v))", name)
		}

	case "uint", "unsigned int", "quint32":
		{
			return fmt.Sprintf("uint(uint32(%v))", name)
		}

	case "long":
		{
			return fmt.Sprintf("int(int32(%v))", name)
		}

	case "ulong", "unsigned long":
		{
			return fmt.Sprintf("uint(uint32(%v))", name)
		}

	case "longlong", "long long", "qlonglong", "qint64":
		{
			return fmt.Sprintf("int64(%v)", name)
		}

	case "ulonglong", "unsigned long long", "qulonglong", "quint64":
		{
			return fmt.Sprintf("uint64(%v)", name)
		}

	case "float":
		{
			return fmt.Sprintf("float32(%v)", name)
		}

	case "double", "qreal":
		{
			return fmt.Sprintf("float64(%v)", name)
		}

	case "uintptr_t", "uintptr", "quintptr", "WId":
		{
			return fmt.Sprintf("uintptr(%v)", name)
		}

		//non std types

	case "T", "JavaVM", "jclass", "jobject":
		{
			switch f.TemplateModeJNI {
			case "Boolean":
				{
					return fmt.Sprintf("int8(%v) != 0", name)
				}

			case "Int":
				{
					return fmt.Sprintf("int(int32(%v))", name)
				}

			case "Void":
				{
					return name
				}
			}

			return fmt.Sprintf("unsafe.Pointer(%v)", name)
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
				return fmt.Sprintf("%v.New%vFromPointer(%v)", m, strings.Title(value), name)
			}

			return fmt.Sprintf("New%vFromPointer(%v)", strings.Title(value), name)
		}

	case parser.IsPackedList(value):
		{
			return fmt.Sprintf("func(l C.struct_%v_PackedList)%v{var out = make(%v, int(l.len))\nfor i:=0;i<int(l.len);i++{ out[i] = New%vFromPointer(l.data).%v_atList(i) }\nreturn out}(%v)", strings.Title(parser.ClassMap[f.Class()].Module), goType(f, value), goType(f, value), f.Class(), f.Name, name)
		}
	}

	f.Access = fmt.Sprintf("unsupported_goOutput(%v)", value)
	return f.Access
}

func goOutputFailed(value string, f *parser.Function) string {
	var vOld = value

	value = CleanValue(value)

	switch value {
	case "char", "qint8", "uchar", "quint8", "QString":
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

	case "bool":
		{
			return "false"
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
			return "0"
		}

		//non std types

	case "T", "JavaVM", "jclass", "jobject":
		{
			switch f.TemplateModeJNI {
			case "Boolean":
				{
					return "false"
				}

			case "Int":
				{
					return "0"
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
			if f.TemplateModeJNI == "String" {
				return "\"\""
			}

			return "nil"
		}

	case parser.IsPackedList(value):
		{
			return "nil"
		}
	}

	f.Access = fmt.Sprintf("unsupported_goOutputFailed(%v)", value)
	return f.Access
}

func cgoOutput(name, value string, f *parser.Function) string {

	var vOld = value

	name = cleanName(name, value)
	value = CleanValue(value)

	switch value {
	case "char", "qint8", "uchar", "quint8", "QString":
		{
			return fmt.Sprintf("cGoUnpackString(%v)", name)
		}

	case "QStringList":
		{
			return fmt.Sprintf("strings.Split(cGoUnpackString(%v), \"|\")", name)
		}

	case "void", "":
		{
			if strings.Contains(vOld, "*") {
				return name
			}

			return ""
		}

	case "bool":
		{
			return fmt.Sprintf("int8(%v) != 0", name)
		}

	case "short", "qint16":
		{
			return fmt.Sprintf("int16(%v)", name)
		}

	case "ushort", "unsigned short", "quint16":
		{
			return fmt.Sprintf("uint16(%v)", name)
		}

	case "int", "qint32":
		{
			return fmt.Sprintf("int(int32(%v))", name)
		}

	case "uint", "unsigned int", "quint32":
		{
			return fmt.Sprintf("uint(uint32(%v))", name)
		}

	case "long":
		{
			return fmt.Sprintf("int(int32(%v))", name)
		}

	case "ulong", "unsigned long":
		{
			return fmt.Sprintf("uint(uint32(%v))", name)
		}

	case "longlong", "long long", "qlonglong", "qint64":
		{
			return fmt.Sprintf("int64(%v)", name)
		}

	case "ulonglong", "unsigned long long", "qulonglong", "quint64":
		{
			return fmt.Sprintf("uint64(%v)", name)
		}

	case "float":
		{
			return fmt.Sprintf("float32(%v)", name)
		}

	case "double", "qreal":
		{
			return fmt.Sprintf("float64(%v)", name)
		}

	case "uintptr_t", "uintptr", "quintptr", "WId":
		{
			return fmt.Sprintf("uintptr(%v)", name)
		}
	}

	switch {
	case isEnum(f.Class(), value):
		{
			if c, exists := parser.ClassMap[class(cppEnum(f, value, false))]; exists && module(c.Module) != module(f) && module(c.Module) != "" {
				if parser.ClassMap[f.Class()].WeakLink[c.Module] {
					return fmt.Sprintf("int64%v)", name)
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
				return fmt.Sprintf("%v.New%vFromPointer(%v)", m, strings.Title(value), name)
			}
			return fmt.Sprintf("New%vFromPointer(%v)", strings.Title(value), name)
		}
	}

	f.Access = fmt.Sprintf("unsupported_cgoOutput(%v)", value)
	return f.Access
}

func CppOutput(name, value string, f *parser.Function) string {
	if strings.HasSuffix(f.Name, "_atList") {
		return cppOutput(fmt.Sprintf("%v->at%v", strings.Split(name, "->")[0], strings.Split(name, "_atList")[1]), value, f)
	}
	return cppOutput(name, value, f)
}

func cppOutputPack(name, value string, f *parser.Function) string {
	var out = CppOutput(name, value, f)

	if strings.Contains(out, "_PackedString") {
		var out = strings.Replace(out, "({ ", "", -1)
		out = strings.Replace(out, " })", "", -1)
		if !strings.HasSuffix(out, ";") {
			out = fmt.Sprintf("%v;", out)
		}
		return strings.Replace(out, "_PackedString", fmt.Sprintf("_PackedString %vPacked =", cleanName(name, value)), -1)
	}

	return ""
}

func cppOutputPacked(name, value string, f *parser.Function) string {
	var out = CppOutput(name, value, f)

	if strings.Contains(out, "_PackedString") {
		return fmt.Sprintf("%vPacked", cleanName(name, value))
	}

	return out
}

func cppOutput(name, value string, f *parser.Function) string {
	var vOld = value

	var tHash = sha1.New()
	tHash.Write([]byte(name))
	var tHashName = hex.EncodeToString(tHash.Sum(nil)[:3])

	name = cleanName(name, value)
	value = CleanValue(value)

	switch value {
	case "char", "qint8":
		{
			if strings.Contains(vOld, "*") {

				var fSizeVariable string
				for _, p := range f.Parameters {
					if strings.Contains(p.Value, "int") {
						fSizeVariable = p.Name
						break
					}
				}

				if fSizeVariable == "" && strings.Contains(strings.ToLower(f.Name), "data") && parser.ClassMap[f.Class()].HasFunctionWithName("size") {
					fSizeVariable = fmt.Sprintf("static_cast<%v*>(ptr)->size()", f.Class())
				}

				if strings.Contains(vOld, "const") {
					if fSizeVariable != "" {
						return fmt.Sprintf("%v_PackedString { const_cast<char*>(%v), %v }", strings.Title(parser.ClassMap[f.Class()].Module), name, fSizeVariable)
					}
					return fmt.Sprintf("%v_PackedString { const_cast<char*>(%v), -1 }", strings.Title(parser.ClassMap[f.Class()].Module), name)
				} else {
					if fSizeVariable != "" {
						return fmt.Sprintf("%v_PackedString { %v, %v }", strings.Title(parser.ClassMap[f.Class()].Module), name, fSizeVariable)
					}
					return fmt.Sprintf("%v_PackedString { %v, -1 }", strings.Title(parser.ClassMap[f.Class()].Module), name)
				}
			}

			return fmt.Sprintf("({ char t%v = %v; %v_PackedString { &t%v, -1 }; })", tHashName, name, strings.Title(parser.ClassMap[f.Class()].Module), tHashName)
		}

	case "uchar", "quint8":
		{
			if strings.Contains(vOld, "*") {
				if strings.Contains(vOld, "const") {
					return fmt.Sprintf("({ char* t%v = static_cast<char*>(static_cast<void*>(const_cast<%v*>(%v))); %v_PackedString { t%v, -1 }; })", tHashName, value, name, strings.Title(parser.ClassMap[f.Class()].Module), tHashName)
				}
				return fmt.Sprintf("({ char* t%v = static_cast<char*>(static_cast<void*>(%v)); %v_PackedString { t%v, -1 }; })", tHashName, name, strings.Title(parser.ClassMap[f.Class()].Module), tHashName)
			}

			return fmt.Sprintf("({ %v pret%v = %v; char* t%v = static_cast<char*>(static_cast<void*>(&pret%v)); %v_PackedString { t%v, -1 }; })", vOld, tHashName, name, tHashName, tHashName, strings.Title(parser.ClassMap[f.Class()].Module), tHashName)
		}

	case "QString":
		{
			if strings.Contains(vOld, "*") {
				return fmt.Sprintf("({ QByteArray t%v = %v->toUtf8(); %v_PackedString { const_cast<char*>(t%v.prepend(\"WHITESPACE\").constData()+10), t%v.size()-10 }; })", tHashName, name, strings.Title(parser.ClassMap[f.Class()].Module), tHashName, tHashName)
			}
			return fmt.Sprintf("({ QByteArray t%v = %v.toUtf8(); %v_PackedString { const_cast<char*>(t%v.prepend(\"WHITESPACE\").constData()+10), t%v.size()-10 }; })", tHashName, name, strings.Title(parser.ClassMap[f.Class()].Module), tHashName, tHashName)
		}

	case "QStringList":
		{
			if strings.Contains(vOld, "*") {
				return fmt.Sprintf("({ QByteArray t%v = %v->join(\"|\").toUtf8(); %v_PackedString { const_cast<char*>(t%v.prepend(\"WHITESPACE\").constData()+10), t%v.size()-10 }; })", tHashName, name, strings.Title(parser.ClassMap[f.Class()].Module), tHashName, tHashName)
			}
			return fmt.Sprintf("({ QByteArray t%v = %v.join(\"|\").toUtf8(); %v_PackedString { const_cast<char*>(t%v.prepend(\"WHITESPACE\").constData()+10), t%v.size()-10 }; })", tHashName, name, strings.Title(parser.ClassMap[f.Class()].Module), tHashName, tHashName)
		}

	case
		"bool",

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
			if strings.Contains(vOld, "*") {
				return fmt.Sprintf("*%v", name)
			}

			return name
		}

		//non std types

	case "void", "", "T", "JavaVM", "jclass", "jobject":
		{
			if value == "void" || value == "T" {
				if strings.Contains(vOld, "*") && strings.Contains(vOld, "const") {
					return fmt.Sprintf("const_cast<void*>(%v)", name)
				}
			}

			return name
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

			if strings.Contains(vOld, "&") {
				if strings.Contains(vOld, "const") {
					return fmt.Sprintf("const_cast<%v*>(&%v)", value, name)
				}
			}

			f.NeedsFinalizer = true

			switch value {
			case "QModelIndex", "QMetaMethod", "QItemSelection":
				{
					return fmt.Sprintf("new %v(%v)", value, name)
				}

			case "QAndroidJniObject":
				{
					return fmt.Sprintf("new %v(%v.object())", value, name)
				}

			case "QPoint", "QPointF":
				{
					return fmt.Sprintf("({ %v tmpValue = %v; new %v(tmpValue.x(), tmpValue.y()); })", value, name, value)
				}

			case "QSize", "QSizeF":
				{
					return fmt.Sprintf("({ %v tmpValue = %v; new %v(tmpValue.width(), tmpValue.height()); })", value, name, value)
				}

			case "QRect", "QRectF":
				{
					return fmt.Sprintf("({ %v tmpValue = %v; new %v(tmpValue.x(), tmpValue.y(), tmpValue.width(), tmpValue.height()); })", value, name, value)
				}

			case "QLine", "QLineF":
				{
					return fmt.Sprintf("({ %v tmpValue = %v; new %v(tmpValue.p1(), tmpValue.p2()); })", value, name, value)
				}

			case "QMargins", "QMarginsF":
				{
					return fmt.Sprintf("({ %v tmpValue = %v; new %v(tmpValue.left(), tmpValue.top(), tmpValue.right(), tmpValue.bottom()); })", value, name, value)
				}
			}

			switch f.Fullname {
			case "QColor::toVariant", "QFont::toVariant", "QImage::toVariant":
				{
					return fmt.Sprintf("new %v(*%v)", value, strings.Split(name, "->")[0])
				}

			case "QVariant::toColor", "QVariant::toFont", "QVariant::toImage":
				{
					f.NeedsFinalizer = false
					return fmt.Sprintf("new %v(qvariant_cast<%v>(*%v))", value, value, strings.Split(name, "->")[0])
				}
			}

			for _, f := range parser.ClassMap[value].Functions {
				if f.Meta == parser.CONSTRUCTOR {
					if len(f.Parameters) == 1 {
						if CleanValue(f.Parameters[0].Value) == value {
							return fmt.Sprintf("new %v(%v)", value, name)
						}
					}
				}
			}
		}

	case parser.IsPackedList(value):
		{
			vOld = strings.Replace(vOld, " &", "", -1)
			return fmt.Sprintf("({ %v* tmpValue = new %v(%v); %v_PackedList { tmpValue, tmpValue->size() }; })", vOld, vOld, name, strings.Title(parser.ClassMap[f.Class()].Module))
		}
	}

	f.Access = fmt.Sprintf("unsupported_cppOutput(%v)", value)
	return f.Access
}
