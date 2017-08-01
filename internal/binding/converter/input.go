package converter

//TODO: GLchar, GLbyte

import (
	"fmt"
	"strings"

	"github.com/therecipe/qt/internal/binding/parser"
)

func GoInput(name, value string, f *parser.Function) string {
	var vOld = value

	name = parser.CleanName(name, value)
	value = parser.CleanValue(value)

	switch value {
	case "char", "qint8":
		{
			if strings.Contains(vOld, "**") {
				return fmt.Sprintf("C.CString(strings.Join(%v, \"|\"))", name)
			}

			if value == "char" && strings.Count(vOld, "*") == 1 && f.Name == "readData" {
				return fmt.Sprintf("C.CString(strings.Repeat(\"0\", int(%v)))", parser.CleanName(f.Parameters[1].Name, f.Parameters[1].Value))
			}

			return fmt.Sprintf("C.CString(%v)", name)
		}

	case "uchar", "quint8", "GLubyte", "QString":
		{
			return fmt.Sprintf("C.CString(%v)", func() string {
				if f.AsError {
					return fmt.Sprintf("func() string { var tmp = %v\n if tmp != nil { return tmp.Error() }\n return \"\" }()", name)
				}
				return name
			}())
		}

	case "QStringList":
		{
			return fmt.Sprintf("C.CString(strings.Join(%v, \"|\"))", name)
		}

	case "void", "GLvoid" /*, ""*/ :
		{
			if strings.Contains(vOld, "*") {
				return name
			}
		}

	case "bool", "GLboolean":
		{
			return fmt.Sprintf("C.char(int8(qt.GoBoolToInt(%v)))", name)
		}

	case "short", "qint16", "GLshort":
		{
			return fmt.Sprintf("C.short(%v)", name)
		}

	case "ushort", "unsigned short", "quint16", "GLushort":
		{
			return fmt.Sprintf("C.ushort(%v)", name)
		}

	case "int", "qint32", "GLint", "GLsizei", "GLintptrARB", "GLsizeiptrARB", "GLfixed", "GLclampx":
		{
			return fmt.Sprintf("C.int(int32(%v))", name)
		}

	case "uint", "unsigned int", "quint32", "GLenum", "GLbitfield", "GLuint":
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

	case "float", "GLfloat", "GLclampf":
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
			if strings.Contains(value, ".") {
				value = strings.Split(value, ".")[1]
			}
			if m := module(parser.State.ClassMap[value].Module); m != module(f) {
				if _, ok := parser.State.ClassMap[f.ClassName()].WeakLink[parser.State.ClassMap[value].Module]; ok {
					return name
				}
				return fmt.Sprintf("%v.PointerFrom%v(%v)", m, strings.Title(value), name)
			}
			return fmt.Sprintf("PointerFrom%v(%v)", strings.Title(value), name)
		}

	case parser.IsPackedList(value):
		{
			if strings.ContainsAny(name, "*&()[]") {
				return fmt.Sprintf("func() unsafe.Pointer {\nvar tmpList = New%vFromPointer(New%vFromPointer(nil).__%v_newList%v())\nfor _,v := range %v{\ntmpList.__%v_setList%v(v)\n}\nreturn tmpList.Pointer()\n}()", strings.Title(f.ClassName()), strings.Title(f.ClassName()), f.Name, f.OverloadNumber, name, f.Name, f.OverloadNumber)
			}
			return fmt.Sprintf("func() unsafe.Pointer {\nvar tmpList = New%vFromPointer(New%vFromPointer(nil).__%v_%v_newList%v())\nfor _,v := range %v{\ntmpList.__%v_%v_setList%v(v)\n}\nreturn tmpList.Pointer()\n}()", strings.Title(f.ClassName()), strings.Title(f.ClassName()), f.Name, name, f.OverloadNumber, name, f.Name, name, f.OverloadNumber)
		}

	case parser.IsPackedMap(value):
		{
			if strings.ContainsAny(name, "*&()[]") {
				return fmt.Sprintf("func() unsafe.Pointer {\nvar tmpList = New%vFromPointer(New%vFromPointer(nil).__%v_newList%v())\nfor k,v := range %v{\ntmpList.__%v_setList%v(k,v)\n}\nreturn tmpList.Pointer()\n}()", strings.Title(f.ClassName()), strings.Title(f.ClassName()), f.Name, f.OverloadNumber, name, f.Name, f.OverloadNumber)
			}
			return fmt.Sprintf("func() unsafe.Pointer {\nvar tmpList = New%vFromPointer(New%vFromPointer(nil).__%v_%v_newList%v())\nfor k,v := range %v{\ntmpList.__%v_%v_setList%v(k,v)\n}\nreturn tmpList.Pointer()\n}()", strings.Title(f.ClassName()), strings.Title(f.ClassName()), f.Name, name, f.OverloadNumber, name, f.Name, name, f.OverloadNumber)
		}
	}

	f.Access = fmt.Sprintf("unsupported_goInput(%v)", value)
	return f.Access
}

func CppInput(name, value string, f *parser.Function) string {

	if (f.SignalMode == parser.CALLBACK || strings.HasPrefix(name, "callback")) && (parser.CleanValue(value) == "QString" || parser.CleanValue(value) == "QStringList") {
		return fmt.Sprintf("({ %v_PackedString tempVal = %v; %v ret = %v; free(tempVal.data); ret; })", strings.Title(parser.State.ClassMap[f.ClassName()].Module), name, parser.CleanValue(value), cppInput("tempVal", value, f))
	}

	return cppInput(name, value, f)
}

func cppInput(name, value string, f *parser.Function) string {
	var vOld = value

	name = parser.CleanName(name, value)
	value = parser.CleanValue(value)

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

	case "uchar", "quint8", "GLubyte":
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
				return fmt.Sprintf("new QString(QString::fromUtf8(%[1]v.data, %[1]v.len))", name)
			}

			if strings.Contains(vOld, "&") && !strings.Contains(vOld, "const") {
				return fmt.Sprintf("*(%v)", cppInput(name, "QString*", f))
			}

			return fmt.Sprintf("QString::fromUtf8(%[1]v.data, %[1]v.len)", name)
		}

	case "QStringList":
		{
			if strings.Contains(vOld, "*") {
				return fmt.Sprintf("new QStringList(%v)", cppInput(name, "QStringList", f))
			}

			if strings.Contains(vOld, "&") && !strings.Contains(vOld, "const") {
				return fmt.Sprintf("*(%v)", cppInput(name, "QStringList*", f))
			}

			return fmt.Sprintf("QString::fromUtf8(%[1]v.data, %[1]v.len).split(\"|\", QString::SkipEmptyParts)", name)
		}

	case "void", "GLvoid" /*, ""*/ :
		{
			if strings.Count(vOld, "*") == 2 && !strings.Contains(vOld, "**") {
				break
			}
			if strings.Contains(vOld, "**") {
				return fmt.Sprintf("&%v", name)
			}
			if strings.Contains(vOld, "*") {
				return name
			}
		}

	case "bool", "GLboolean":
		{
			if strings.Contains(vOld, "*") {
				return "NULL"
			}

			return fmt.Sprintf("%v != 0", name)
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
			if strings.Contains(value, ".") {
				value = strings.Split(value, ".")[1]
			}
			if strings.Contains(vOld, "*") && strings.Contains(vOld, "&") {
				break
			}

			if parser.State.ClassMap[value].Fullname != "" {
				value = parser.State.ClassMap[value].Fullname
			}

			if strings.Contains(vOld, "*") {
				return fmt.Sprintf("static_cast<%v*>(%v)", value, name)
			}

			return fmt.Sprintf("*static_cast<%v*>(%v)", value, name)
		}

	case parser.IsPackedList(value) || parser.IsPackedMap(value):
		{
			if strings.HasSuffix(vOld, "*") {
				return fmt.Sprintf("static_cast<%v*>(%v)", value, name)
			}

			return fmt.Sprintf("*static_cast<%v*>(%v)", value, name)
		}
	}

	f.Access = fmt.Sprintf("unsupported_cppInput(%v)", value)
	return f.Access
}
