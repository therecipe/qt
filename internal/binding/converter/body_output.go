package converter

import "github.com/therecipe/qt/internal/binding/parser"

func GoBodyOutput(f *parser.Function, name string) string {

	var value = f.Output

	if f.Meta == "constructor" && f.Output == "" {
		value = f.Name
	}

	return goOutput(name, value, f)
}

func GoBodyOutputFailed(f *parser.Function) string {

	var value = f.Output

	if f.Meta == "constructor" && f.Output == "" {
		value = f.Name
	}

	return goOutputFailed(value, f)
}

func CppBodyOutput(f *parser.Function, name string) string {
	var value = f.Output

	if f.Meta == "constructor" && f.Output == "" {
		value = f.Name
	}

	return cppOutput(name, value, f)
}

func DeduceGeneric(f *parser.Function) string {

	switch f.TemplateMode {
	case "Int":
		{
			return "<jint>"
		}

	case "Boolean":
		{
			return "<jboolean>"
		}

	case "Void":
		{
			return "<void>"
		}
	}

	if f.Fullname == "QAndroidJniObject::getStaticObjectField" || f.Fullname == "QAndroidJniObject::getObjectField" || f.Fullname == "QAndroidJniObject::callStaticObjectMethod" || f.Fullname == "QAndroidJniObject::callObjectMethod" {

		if (f.Fullname == "QAndroidJniObject::callObjectMethod" || f.Fullname == "QAndroidJniObject::callStaticObjectMethod") && containsT(f) {
			return ""
		}

		return "<jobject>"
	}

	return ""
}

func containsT(f *parser.Function) bool {

	for _, p := range f.Parameters {
		if p.Value == "..." {
			return true
		}
	}

	return false
}
