package converter

import (
	"fmt"
	"strings"

	"github.com/therecipe/qt/internal/binding/parser"
)

func GoHeaderName(f *parser.Function) (o string) {

	if f.SignalMode == parser.CALLBACK {
		return fmt.Sprintf("callback%v_%v%v", f.Class(), strings.Replace(strings.Title(f.Name), parser.TILDE, "Destroy", -1), f.OverloadNumber)
	}

	if f.Static {
		o += fmt.Sprintf("%v_", strings.Split(f.Fullname, "::")[0])
	}

	o += f.SignalMode

	switch f.Meta {
	case parser.CONSTRUCTOR:
		{
			o += "New"
		}
	}

	if f.Meta == parser.DESTRUCTOR || strings.HasPrefix(f.Name, parser.TILDE) {
		o += "Destroy"
	}

	if f.TemplateMode == "String" || f.TemplateMode == "Object" {

		if strings.Contains(f.Name, "Object") {
			if f.TemplateMode == "String" {
				o += strings.Replace(strings.Title(f.Name), "Object", "", -1) + f.TemplateMode
			} else {
				o += strings.Title(f.Name)
			}
		}

	} else {
		o += strings.Title(f.Name) + f.TemplateMode
	}

	if f.Overload {
		o += f.OverloadNumber
	}

	if strings.ContainsAny(o, "&<>=/!()[]{}^|*+-") || strings.Contains(o, "Operator") {
		f.Access = "unsupported_GoHeaderName"
		return f.Access
	}

	return strings.Replace(o, parser.TILDE, "", -1)
}

func CppHeaderName(f *parser.Function) string {
	return fmt.Sprintf("%v_%v", strings.Split(f.Fullname, "::")[0], GoHeaderName(f))
}

func GoHeaderOutput(f *parser.Function) string {

	if f.SignalMode == parser.CALLBACK {
		return cgoType(f, f.Output)
	}

	if f.SignalMode == parser.CONNECT || f.SignalMode == parser.DISCONNECT {
		return ""
	}

	var value = f.Output

	if f.Meta == parser.CONSTRUCTOR && f.Output == "" {
		value = f.Name
	}

	var o = goType(f, value)
	if isClass(o) {
		return fmt.Sprintf("*%v", o)
	}
	return o
}

func CppHeaderOutput(f *parser.Function) string {

	var value = f.Output

	if f.Meta == parser.CONSTRUCTOR && f.Output == "" {
		value = f.Name
	}

	return cppType(f, value)
}

func GoHeaderInput(f *parser.Function) (o string) {

	if f.SignalMode == parser.CALLBACK {
		o += "ptr unsafe.Pointer"
		for _, p := range f.Parameters {
			if v := cgoType(f, p.Value); v != "" {
				o += fmt.Sprintf(", %v %v", cleanName(p.Name, p.Value), v)
			}
		}
		return strings.TrimSuffix(o, ", ")
	}

	if f.SignalMode == parser.CONNECT {
		o += "f func ("
	}

	if (f.Meta == parser.SIGNAL || strings.Contains(f.Virtual, parser.IMPURE)) && f.SignalMode != parser.CONNECT {
		if strings.Contains(f.Virtual, parser.IMPURE) && f.SignalMode == "" {
		} else {
			return
		}
	}

	for _, p := range f.Parameters {
		if v := goType(f, p.Value); v != "" {
			if isClass(v) {
				if f.SignalMode == parser.CONNECT {
					o += fmt.Sprintf("%v *%v, ", cleanName(p.Name, p.Value), v)
				} else {
					o += fmt.Sprintf("%v %v_ITF, ", cleanName(p.Name, p.Value), v)
				}
			} else {
				o += fmt.Sprintf("%v %v, ", cleanName(p.Name, p.Value), v)
			}
		} else {
			f.Access = "unsupported_GoHeaderInput"
			return f.Access
		}
	}

	o = strings.TrimSuffix(o, ", ")

	if f.SignalMode == parser.CONNECT {
		o += ")"
	}

	if f.SignalMode == parser.CONNECT {
		if isClass(goType(f, f.Output)) {
			o += " *" + goType(f, f.Output)
		} else {
			o += " " + goType(f, f.Output)
		}
	}

	return
}

func GoHeaderInputSignalFunction(f *parser.Function) (o string) {

	o += "func ("

	for _, p := range f.Parameters {
		if v := goType(f, p.Value); v != "" {
			if isClass(v) {
				o += fmt.Sprintf("*%v, ", v)
			} else {
				o += fmt.Sprintf("%v, ", v)
			}
		} else {
			f.Access = "unsupported_GoHeaderInputSignalFunction"
			return f.Access
		}
	}

	o = strings.TrimSuffix(o, ", ")

	o += ")"

	if f.SignalMode == parser.CALLBACK {
		if isClass(goType(f, f.Output)) {
			o += " *" + goType(f, f.Output)
		} else {
			o += " " + goType(f, f.Output)
		}
	}

	return
}

func CppHeaderInput(f *parser.Function) (o string) {
	if !(f.Static || f.Meta == parser.CONSTRUCTOR) {
		o += "void* ptr, "
	}

	if f.Meta == parser.SIGNAL {
		return strings.TrimSuffix(o, ", ")
	}

	for _, p := range f.Parameters {
		if v := cppType(f, p.Value); !(v == "") {
			o += fmt.Sprintf("%v %v, ", v, cleanName(p.Name, p.Value))
		} else {
			f.Access = "unsupported_CppHeaderInput"
			return f.Access
		}
	}

	return strings.TrimSuffix(o, ", ")
}
