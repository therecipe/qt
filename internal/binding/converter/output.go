package converter

//TODO: GLchar, GLbyte

import (
	"crypto/sha1"
	"encoding/hex"
	"fmt"
	"strings"

	"github.com/therecipe/qt/internal/binding/parser"
)

func goOutput(name, value string, f *parser.Function) string {
	var vOld = value

	name = parser.CleanName(name, value)
	value = parser.CleanValue(value)

	switch value {
	case "char", "qint8", "uchar", "quint8", "GLubyte", "QString":
		{
			return func() string {
				var out = fmt.Sprintf("cGoUnpackString(%v)", name)
				if f.AsError {
					return fmt.Sprintf("errors.New(%v)", out)
				}
				return out
			}()
		}

	case "QStringList":
		{
			return fmt.Sprintf("strings.Split(cGoUnpackString(%v), \"|\")", name)
		}

	case "void", "GLvoid", "":
		{
			if strings.Contains(vOld, "*") {
				return fmt.Sprintf("unsafe.Pointer(%v)", name)
			}

			return name
		}

	case "bool", "GLboolean":
		{
			return fmt.Sprintf("%v != 0", name)
		}

	case "short", "qint16", "GLshort":
		{
			return fmt.Sprintf("int16(%v)", name)
		}

	case "ushort", "unsigned short", "quint16", "GLushort":
		{
			return fmt.Sprintf("uint16(%v)", name)
		}

	case "int", "qint32", "GLint", "GLsizei", "GLintptrARB", "GLsizeiptrARB", "GLfixed", "GLclampx":
		{
			return fmt.Sprintf("int(int32(%v))", name)
		}

	case "uint", "unsigned int", "quint32", "GLenum", "GLbitfield", "GLuint":
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

	case "float", "GLfloat", "GLclampf":
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
	case isEnum(f.ClassName(), value):
		{
			if c, ok := parser.State.ClassMap[class(cppEnum(f, value, false))]; ok && module(c.Module) != module(f) && module(c.Module) != "" {
				if _, ok := parser.State.ClassMap[f.ClassName()].WeakLink[c.Module]; ok {
					return fmt.Sprintf("int64(%v)", name)
				}
				return fmt.Sprintf("%v.%v(%v)", module(c.Module), goEnum(f, value), name)
			}
			return fmt.Sprintf("%v(%v)", goEnum(f, value), name)
		}

	case isClass(value):
		{
			if strings.Contains(value, ".") {
				value = strings.Split(value, ".")[1]
			}
			if m := module(parser.State.ClassMap[value].Module); m != module(f) {
				if _, ok := parser.State.ClassMap[f.ClassName()].WeakLink[parser.State.ClassMap[value].Module]; ok {
					return fmt.Sprintf("unsafe.Pointer(%v)", name)
				}
				return fmt.Sprintf("%v.New%vFromPointer(%v)", m, strings.Title(value), name)
			}
			return fmt.Sprintf("New%vFromPointer(%v)", strings.Title(value), name)
		}

	case parser.IsPackedList(value):
		{
			return fmt.Sprintf("func(l C.struct_%v_PackedList)%v{var out = make(%v, int(l.len))\nfor i:=0;i<int(l.len);i++{ out[i] = New%vFromPointer(l.data).__%v_atList%v(i) }\nreturn out}(%v)", strings.Title(parser.State.ClassMap[f.ClassName()].Module), goType(f, value), goType(f, value), strings.Title(f.ClassName()), f.Name, f.OverloadNumber, name)
		}

	case parser.IsPackedMap(value):
		{
			return fmt.Sprintf("func(l C.struct_%v_PackedList)%v{var out = make(%v, int(l.len))\nfor _,i:=range New%vFromPointer(l.data).__%v_keyList(){ out[i] = New%vFromPointer(l.data).__%v_atList%v(i) }\nreturn out}(%v)", strings.Title(parser.State.ClassMap[f.ClassName()].Module), goType(f, value), goType(f, value), strings.Title(f.ClassName()), f.Name, strings.Title(f.ClassName()), f.Name, f.OverloadNumber, name)
		}
	}

	f.Access = fmt.Sprintf("unsupported_goOutput(%v)", value)
	return f.Access
}

func GoOutputFailed(value string, f *parser.Function) string { return goOutputFailed(value, f) }
func goOutputFailed(value string, f *parser.Function) string {
	var vOld = value

	value = parser.CleanValue(value)

	switch value {
	case "char", "qint8", "uchar", "quint8", "GLubyte", "QString":
		{
			if f.AsError {
				return "errors.New(\"\")"
			}
			return "\"\""
		}

	case "QStringList":
		{
			return "make([]string, 0)"
		}

	case "void", "GLvoid", "":
		{
			if strings.Contains(vOld, "*") {
				return "nil"
			}

			return ""
		}

	case "bool", "GLboolean":
		{
			return "false"
		}

	case
		"short", "qint16", "GLshort",
		"ushort", "unsigned short", "quint16", "GLushort",

		"int", "qint32", "GLint", "GLsizei", "GLintptrARB", "GLsizeiptrARB", "GLfixed", "GLclampx",
		"uint", "unsigned int", "quint32", "GLenum", "GLbitfield", "GLuint",

		"long",
		"ulong", "unsigned long",

		"longlong", "long long", "qlonglong", "qint64",
		"ulonglong", "unsigned long long", "qulonglong", "quint64",

		"float", "GLfloat", "GLclampf",
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
	case isEnum(f.ClassName(), value):
		{
			return "0"
		}

	case isClass(value):
		{
			if strings.Contains(value, ".") {
				value = strings.Split(value, ".")[1]
			}
			if f.TemplateModeJNI == "String" {
				return "\"\""
			}

			return "nil"
		}

	case parser.IsPackedList(value) || parser.IsPackedMap(value):
		{
			return fmt.Sprintf("make(%v, 0)", goType(f, value))
		}
	}

	f.Access = fmt.Sprintf("unsupported_goOutputFailed(%v)", value)
	return f.Access
}

func cgoOutput(name, value string, f *parser.Function) string {

	var vOld = value

	name = parser.CleanName(name, value)
	value = parser.CleanValue(value)

	switch value {
	case "char", "qint8", "uchar", "quint8", "GLubyte", "QString":
		{
			return func() string {
				var out = fmt.Sprintf("cGoUnpackString(%v)", name)
				if f.AsError {
					return fmt.Sprintf("errors.New(%v)", out)
				}
				return out
			}()
		}

	case "QStringList":
		{
			return fmt.Sprintf("strings.Split(cGoUnpackString(%v), \"|\")", name)
		}

	case "void", "GLvoid", "":
		{
			if strings.Contains(vOld, "*") {
				return name
			}

			return ""
		}

	case "bool", "GLboolean":
		{
			return fmt.Sprintf("int8(%v) != 0", name)
		}

	case "short", "qint16", "GLshort":
		{
			return fmt.Sprintf("int16(%v)", name)
		}

	case "ushort", "unsigned short", "quint16", "GLushort":
		{
			return fmt.Sprintf("uint16(%v)", name)
		}

	case "int", "qint32", "GLint", "GLsizei", "GLintptrARB", "GLsizeiptrARB", "GLfixed", "GLclampx":
		{
			return fmt.Sprintf("int(int32(%v))", name)
		}

	case "uint", "unsigned int", "quint32", "GLenum", "GLbitfield", "GLuint":
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

	case "float", "GLfloat", "GLclampf":
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
	case isEnum(f.ClassName(), value):
		{
			if c, ok := parser.State.ClassMap[class(cppEnum(f, value, false))]; ok && module(c.Module) != module(f) && module(c.Module) != "" {
				if _, ok := parser.State.ClassMap[f.ClassName()].WeakLink[c.Module]; ok {
					return fmt.Sprintf("int64%v)", name)
				}
				return fmt.Sprintf("%v.%v(%v)", module(c.Module), goEnum(f, value), name)
			}
			return fmt.Sprintf("%v(%v)", goEnum(f, value), name)
		}

	case isClass(value):
		{
			if strings.Contains(value, ".") {
				value = strings.Split(value, ".")[1]
			}
			if m := module(parser.State.ClassMap[value].Module); m != module(f) {
				if _, ok := parser.State.ClassMap[f.ClassName()].WeakLink[parser.State.ClassMap[value].Module]; ok {
					return fmt.Sprintf("unsafe.Pointer(%v)", name)
				}
				return fmt.Sprintf("%v.New%vFromPointer(%v)", m, strings.Title(value), name)
			}
			return fmt.Sprintf("New%vFromPointer(%v)", strings.Title(value), name)
		}

	case parser.IsPackedList(value):
		{
			return fmt.Sprintf("func(l C.struct_%v_PackedList)%v{var out = make(%v, int(l.len))\nfor i:=0;i<int(l.len);i++{ out[i] = New%vFromPointer(l.data).__%v_%v_atList%v(i) }\nreturn out}(%v)", strings.Title(parser.State.ClassMap[f.ClassName()].Module), goType(f, value), goType(f, value), strings.Title(f.ClassName()), f.Name, name, f.OverloadNumber, name)
		}

	case parser.IsPackedMap(value):
		{
			return fmt.Sprintf("func(l C.struct_%v_PackedList)%v{var out = make(%v, int(l.len))\nfor _,i:=range New%vFromPointer(l.data).__%v_keyList(){ out[i] = New%vFromPointer(l.data).__%v_%v_atList%v(i) }\nreturn out}(%v)", strings.Title(parser.State.ClassMap[f.ClassName()].Module), goType(f, value), goType(f, value), strings.Title(f.ClassName()), f.Name, strings.Title(f.ClassName()), f.Name, name, f.OverloadNumber, name)
		}
	}

	f.Access = fmt.Sprintf("unsupported_cgoOutput(%v)", value)
	return f.Access
}

func CppOutput(name, value string, f *parser.Function) string {
	if strings.HasSuffix(f.Name, "_atList") {
		if f.IsMap {
			return cppOutput(fmt.Sprintf("%v->value%v", strings.Split(name, "->")[0], strings.Split(name, "_atList")[1]), value, f)
		}
		return cppOutput(fmt.Sprintf("%v->at%v", strings.Split(name, "->")[0], strings.Split(name, "_atList")[1]), value, f)
	}
	if strings.HasSuffix(f.Name, "_setList") {
		if len(f.Parameters) == 2 {
			return cppOutput(fmt.Sprintf("%v->insert%v", strings.Split(name, "->")[0], strings.Split(name, "_setList")[1]), value, f)
		}
		return cppOutput(fmt.Sprintf("%v->append%v", strings.Split(name, "->")[0], strings.Split(name, "_setList")[1]), value, f)
	}
	if strings.HasSuffix(f.Name, "_newList") {
		return fmt.Sprintf("new %v", parser.CleanValue(f.Container))
	}
	if strings.HasSuffix(f.Name, "_keyList") {
		return cppOutput(fmt.Sprintf("static_cast<%v*>(ptr)->keys()", f.Container), value, f)
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
		return strings.Replace(out, "_PackedString", fmt.Sprintf("_PackedString %vPacked =", parser.CleanName(name, value)), -1)
	}

	return ""
}

func cppOutputPacked(name, value string, f *parser.Function) string {
	var out = CppOutput(name, value, f)

	if strings.Contains(out, "_PackedString") {
		return fmt.Sprintf("%vPacked", parser.CleanName(name, value))
	}

	return out
}

func cppOutput(name, value string, f *parser.Function) string {
	var vOld = value

	var tHash = sha1.New()
	tHash.Write([]byte(name))
	var tHashName = hex.EncodeToString(tHash.Sum(nil)[:3])

	name = parser.CleanName(name, value)
	value = parser.CleanValue(value)

	switch value {
	case "char", "qint8":
		{
			if strings.Contains(vOld, "*") {

				var fSizeVariable string
				for _, p := range f.Parameters {
					if strings.Contains(p.Value, "int") {
						fSizeVariable = parser.CleanName(p.Name, p.Value)
						break
					}
				}

				if fSizeVariable == "" && strings.Contains(strings.ToLower(f.Name), "data") && parser.State.ClassMap[f.ClassName()].HasFunctionWithName("size") {
					fSizeVariable = fmt.Sprintf("static_cast<%v*>(ptr)->size()", f.ClassName())
				}

				if strings.Contains(vOld, "const") {
					if fSizeVariable != "" {
						return fmt.Sprintf("%v_PackedString { const_cast<char*>(%v), %v }", strings.Title(parser.State.ClassMap[f.ClassName()].Module), name, fSizeVariable)
					}
					return fmt.Sprintf("%v_PackedString { const_cast<char*>(%v), -1 }", strings.Title(parser.State.ClassMap[f.ClassName()].Module), name)
				} else {
					if fSizeVariable != "" {
						return fmt.Sprintf("%v_PackedString { %v, %v }", strings.Title(parser.State.ClassMap[f.ClassName()].Module), name, fSizeVariable)
					}
					return fmt.Sprintf("%v_PackedString { %v, -1 }", strings.Title(parser.State.ClassMap[f.ClassName()].Module), name)
				}
			}

			return fmt.Sprintf("({ char t%v = %v; %v_PackedString { &t%v, -1 }; })", tHashName, name, strings.Title(parser.State.ClassMap[f.ClassName()].Module), tHashName)
		}

	case "uchar", "quint8", "GLubyte":
		{
			if strings.Contains(vOld, "*") {
				if strings.Contains(vOld, "const") {
					return fmt.Sprintf("({ char* t%v = static_cast<char*>(static_cast<void*>(const_cast<%v*>(%v))); %v_PackedString { t%v, -1 }; })", tHashName, value, name, strings.Title(parser.State.ClassMap[f.ClassName()].Module), tHashName)
				}
				return fmt.Sprintf("({ char* t%v = static_cast<char*>(static_cast<void*>(%v)); %v_PackedString { t%v, -1 }; })", tHashName, name, strings.Title(parser.State.ClassMap[f.ClassName()].Module), tHashName)
			}

			if strings.Contains(vOld, "const") {
				return fmt.Sprintf("({ %v pret%v = %v; char* t%v = static_cast<char*>(static_cast<void*>(const_cast<%v*>(&pret%v))); %v_PackedString { t%v, -1 }; })", vOld, tHashName, name, tHashName, value, tHashName, strings.Title(parser.State.ClassMap[f.ClassName()].Module), tHashName)
			}
			return fmt.Sprintf("({ %v pret%v = %v; char* t%v = static_cast<char*>(static_cast<void*>(&pret%v)); %v_PackedString { t%v, -1 }; })", vOld, tHashName, name, tHashName, tHashName, strings.Title(parser.State.ClassMap[f.ClassName()].Module), tHashName)
		}

	case "QString":
		{
			if strings.Contains(vOld, "*") {
				return fmt.Sprintf("({ QByteArray t%v = %v->toUtf8(); %v_PackedString { const_cast<char*>(t%v.prepend(\"WHITESPACE\").constData()+10), t%v.size()-10 }; })", tHashName, name, strings.Title(parser.State.ClassMap[f.ClassName()].Module), tHashName, tHashName)
			}
			return fmt.Sprintf("({ QByteArray t%v = %v.toUtf8(); %v_PackedString { const_cast<char*>(t%v.prepend(\"WHITESPACE\").constData()+10), t%v.size()-10 }; })", tHashName, name, strings.Title(parser.State.ClassMap[f.ClassName()].Module), tHashName, tHashName)
		}

	case "QStringList":
		{
			if strings.Contains(vOld, "*") {
				return fmt.Sprintf("({ QByteArray t%v = %v->join(\"|\").toUtf8(); %v_PackedString { const_cast<char*>(t%v.prepend(\"WHITESPACE\").constData()+10), t%v.size()-10 }; })", tHashName, name, strings.Title(parser.State.ClassMap[f.ClassName()].Module), tHashName, tHashName)
			}
			return fmt.Sprintf("({ QByteArray t%v = %v.join(\"|\").toUtf8(); %v_PackedString { const_cast<char*>(t%v.prepend(\"WHITESPACE\").constData()+10), t%v.size()-10 }; })", tHashName, name, strings.Title(parser.State.ClassMap[f.ClassName()].Module), tHashName, tHashName)
		}

	case
		"bool", "GLboolean",

		"short", "qint16", "GLshort",
		"ushort", "unsigned short", "quint16", "GLushort",

		"int", "qint32", "GLint", "GLsizei", "GLintptrARB", "GLsizeiptrARB", "GLfixed", "GLclampx",
		"uint", "unsigned int", "quint32", "GLenum", "GLbitfield", "GLuint",

		"long",
		"ulong", "unsigned long",

		"longlong", "long long", "qlonglong", "qint64",
		"ulonglong", "unsigned long long", "qulonglong", "quint64",

		"float", "GLfloat", "GLclampf",
		"double", "qreal",

		"uintptr_t", "uintptr", "quintptr", "WId":
		{
			if strings.Contains(vOld, "*") {
				return fmt.Sprintf("*%v", name)
			}

			return name
		}

		//non std types

	case "void", "GLvoid", "", "T", "JavaVM", "jclass", "jobject":
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
	case isEnum(f.ClassName(), value):
		{
			return name
		}

	case isClass(value):
		{
			if strings.Contains(value, ".") {
				value = strings.Split(value, ".")[1]
			}
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
				if f.SignalMode == parser.CALLBACK {
					return fmt.Sprintf("static_cast<%v*>(&%v)", value, name)
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
			case "QColor::toVariant", "QFont::toVariant", "QImage::toVariant", "QObject::toVariant", "QIcon::toVariant":
				{
					if f.Fullname == "QObject::toVariant" {
						return fmt.Sprintf("new %v(QVariant::fromValue(%v))", value, strings.Split(name, "->")[0])
					}
					return fmt.Sprintf("new %v(*%v)", value, strings.Split(name, "->")[0])
				}

			case "QVariant::toColor", "QVariant::toFont", "QVariant::toImage", "QVariant::toObject", "QVariant::toIcon":
				{
					f.NeedsFinalizer = false

					if f.Fullname == "QVariant::toObject" {
						return fmt.Sprintf("qvariant_cast<%v*>(*%v)", value, strings.Split(name, "->")[0])
					}
					return fmt.Sprintf("new %v(qvariant_cast<%v>(*%v))", value, value, strings.Split(name, "->")[0])
				}
			}

			for _, f := range parser.State.ClassMap[value].Functions {
				if f.Meta == parser.CONSTRUCTOR {
					switch len(f.Parameters) {
					case 0:
						{
							if value == "QDataStream" {

							} else {
								return fmt.Sprintf("new %v(%v)", value, name)
							}
						}

					case 1:
						{
							if parser.CleanValue(f.Parameters[0].Value) == value {
								return fmt.Sprintf("new %v(%v)", value, name)
							}
						}
					}
				}
			}
		}

	case parser.IsPackedList(value) || parser.IsPackedMap(value):
		{
			if strings.HasSuffix(vOld, "*") {
				if strings.Contains(vOld, "const") {
					return fmt.Sprintf("({ %v* tmpValue = const_cast<%v*>(%v); %v_PackedList { tmpValue, tmpValue->size() }; })", value, value, name, strings.Title(parser.State.ClassMap[f.ClassName()].Module))
				}
				return fmt.Sprintf("({ %v* tmpValue = %v; %v_PackedList { tmpValue, tmpValue->size() }; })", value, name, strings.Title(parser.State.ClassMap[f.ClassName()].Module))
			}

			if strings.HasSuffix(vOld, "&") {
				if strings.Contains(vOld, "const") {
					return fmt.Sprintf("({ %v* tmpValue = const_cast<%v*>(&%v); %v_PackedList { tmpValue, tmpValue->size() }; })", value, value, name, strings.Title(parser.State.ClassMap[f.ClassName()].Module))

				}
				if f.SignalMode == parser.CALLBACK {
					return fmt.Sprintf("({ %v* tmpValue = static_cast<%v*>(&%v); %v_PackedList { tmpValue, tmpValue->size() }; })", value, value, name, strings.Title(parser.State.ClassMap[f.ClassName()].Module))
				}
			}

			return fmt.Sprintf("({ %v* tmpValue = new %v(%v); %v_PackedList { tmpValue, tmpValue->size() }; })", value, value, name, strings.Title(parser.State.ClassMap[f.ClassName()].Module))
		}
	}

	f.Access = fmt.Sprintf("unsupported_cppOutput(%v)", value)
	return f.Access
}
