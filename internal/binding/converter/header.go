package converter

import (
	"bytes"
	"fmt"
	"strings"

	"github.com/therecipe/qt/internal/binding/parser"
)

func GoHeaderName(f *parser.Function) string {

	if f.SignalMode == parser.CALLBACK {
		return fmt.Sprintf("callback%v_%v%v", f.ClassName(), strings.Replace(strings.Title(f.Name), parser.TILDE, "Destroy", -1), f.OverloadNumber)
	}

	var bb = new(bytes.Buffer)
	defer bb.Reset()

	if f.Static {
		fmt.Fprintf(bb, "%v_", strings.Split(f.Fullname, "::")[0])
	}

	fmt.Fprint(bb, f.SignalMode)

	switch {
	case f.Meta == parser.CONSTRUCTOR:
		{
			fmt.Fprint(bb, "New")
		}

	case f.Meta == parser.DESTRUCTOR, strings.HasPrefix(f.Name, parser.TILDE):
		{
			fmt.Fprint(bb, "Destroy")
		}
	}

	switch f.TemplateModeJNI {
	case "String", "Object":
		{
			if strings.Contains(f.Name, "Object") {
				if f.TemplateModeJNI == "String" {
					fmt.Fprintf(bb, "%v%v", strings.Replace(strings.Title(f.Name), "Object", "", -1), f.TemplateModeJNI)
				} else {
					fmt.Fprint(bb, strings.Title(f.Name))
				}
			}
		}

	default:
		{
			fmt.Fprintf(bb, "%v%v",

				func() string {
					if strings.HasSuffix(f.Name, "_atList") || strings.HasSuffix(f.Name, "_setList") ||
						strings.HasSuffix(f.Name, "_newList") || strings.HasSuffix(f.Name, "_keyList") {
						return f.Name
					}
					return strings.Title(f.Name)
				}(),

				f.TemplateModeJNI,
			)
		}
	}

	if f.Overload {
		fmt.Fprint(bb, f.OverloadNumber)
	}

	if f.Default {
		fmt.Fprint(bb, "Default")
	}

	if f.Exception {
		fmt.Fprint(bb, "Caught")
	}

	if strings.ContainsAny(bb.String(), "&<>=/!()[]{}^|*+-") || strings.Contains(bb.String(), "Operator") {
		f.Access = "unsupported_GoHeaderName"
		return f.Access
	}

	return strings.Replace(bb.String(), parser.TILDE, "", -1)
}

func CppHeaderName(f *parser.Function) string {
	return fmt.Sprintf("%v_%v", f.ClassName(), GoHeaderName(f))
}

func GoHeaderOutput(f *parser.Function) string {

	switch f.SignalMode {
	case parser.CALLBACK:
		{
			return cgoTypeOutput(f, f.Output)
		}

	case parser.CONNECT, parser.DISCONNECT:
		{
			return ""
		}
	}

	if f.PureGoOutput != "" {
		return f.PureGoOutput
	}

	var value = f.Output

	if f.Meta == parser.CONSTRUCTOR && f.Output == "" {
		value = f.Name
	}

	var o = goType(f, value)
	if isClass(o) {
		if !strings.HasPrefix(o, "[]") && !strings.HasPrefix(o, "map[") {
			o = fmt.Sprintf("*%v", o)
		}
	}

	if f.Exception {
		if o != "" {
			o += ", "
		}
		o += "error"
		o = fmt.Sprintf("(%v)", o)
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

func GoHeaderInput(f *parser.Function) string {
	var bb = new(bytes.Buffer)
	defer bb.Reset()

	if f.SignalMode == parser.DISCONNECT {
		return bb.String()
	}

	if f.SignalMode == parser.CALLBACK {
		fmt.Fprint(bb, "ptr unsafe.Pointer")
		for _, p := range f.Parameters {
			if v := cgoType(f, p.Value); v != "" {
				fmt.Fprintf(bb, ", %v %v", parser.CleanName(p.Name, p.Value), v)
			}
		}
		return bb.String()
	}

	if f.SignalMode == parser.CONNECT {
		fmt.Fprint(bb, "f func (")
	}

	if (f.Meta == parser.SIGNAL || strings.Contains(f.Virtual, parser.IMPURE)) && f.SignalMode != parser.CONNECT {
		if strings.Contains(f.Virtual, parser.IMPURE) && f.SignalMode == "" {
		} else {
			return bb.String()
		}
	}

	var tmp = make([]string, 0)
	for _, p := range f.Parameters {
		if p.PureGoType != "" {
			tmp = append(tmp, fmt.Sprintf("%v %v", parser.CleanName(p.Name, p.Value), p.PureGoType))
		} else {
			if v := goType(f, p.Value); v != "" {
				if isClass(v) && !parser.IsPackedList(parser.CleanValue(p.Value)) && !parser.IsPackedMap(parser.CleanValue(p.Value)) {
					if f.SignalMode == parser.CONNECT {
						tmp = append(tmp, fmt.Sprintf("%v *%v", parser.CleanName(p.Name, p.Value), v))
					} else {
						tmp = append(tmp, fmt.Sprintf("%v %v_ITF", parser.CleanName(p.Name, p.Value), v))
					}
				} else {
					tmp = append(tmp, fmt.Sprintf("%v %v", parser.CleanName(p.Name, p.Value), v))
				}
			} else {
				f.Access = "unsupported_GoHeaderInput"
				return f.Access
			}
		}
	}
	fmt.Fprint(bb, strings.Join(tmp, ", "))

	if f.SignalMode == parser.CONNECT {
		fmt.Fprint(bb, ")")

		if f.PureGoOutput != "" {
			fmt.Fprintf(bb, " %v", f.PureGoOutput)
		} else {
			if isClass(goType(f, f.Output)) && !parser.IsPackedList(parser.CleanValue(f.Output)) && !parser.IsPackedMap(parser.CleanValue(f.Output)) {
				fmt.Fprintf(bb, " *%v", goType(f, f.Output))
			} else {
				fmt.Fprintf(bb, " %v", goType(f, f.Output))
			}
		}
	}

	return bb.String()
}

//TODO: combine with above
func GoHeaderInputSignalFunction(f *parser.Function) string {
	var bb = new(bytes.Buffer)
	defer bb.Reset()

	fmt.Fprint(bb, "func (")

	var tmp = make([]string, 0)

	for _, p := range f.Parameters {
		if p.PureGoType != "" {
			tmp = append(tmp, fmt.Sprintf("%v", p.PureGoType))
		} else {
			if v := goType(f, p.Value); v != "" {
				if isClass(v) && !parser.IsPackedList(parser.CleanValue(p.Value)) && !parser.IsPackedMap(parser.CleanValue(p.Value)) {
					tmp = append(tmp, fmt.Sprintf("*%v", v))
				} else {
					tmp = append(tmp, v)
				}
			} else {
				f.Access = "unsupported_GoHeaderInputSignalFunction"
				return f.Access
			}
		}
	}

	fmt.Fprint(bb, strings.Join(tmp, ", "))

	fmt.Fprint(bb, ")")

	if f.SignalMode == parser.CALLBACK {
		if f.PureGoOutput != "" {
			fmt.Fprintf(bb, " %v", f.PureGoOutput)
		} else {
			if isClass(goType(f, f.Output)) && !parser.IsPackedList(parser.CleanValue(f.Output)) && !parser.IsPackedMap(parser.CleanValue(f.Output)) {
				fmt.Fprintf(bb, " *%v", goType(f, f.Output))
			} else {
				fmt.Fprintf(bb, " %v", goType(f, f.Output))
			}
		}
	}

	return bb.String()
}

func GoGoInput(f *parser.Function) string {
	var tmp = make([]string, 0)
	for _, p := range f.Parameters {
		tmp = append(tmp, parser.CleanName(p.Name, p.Value))
	}
	return strings.Join(tmp, ", ")
}

func CppHeaderInput(f *parser.Function) string {
	var tmp = make([]string, 0)

	if !(f.Static || f.Meta == parser.CONSTRUCTOR) {
		tmp = append(tmp, "void* ptr")
	}

	if f.Meta == parser.SIGNAL {
		return strings.Join(tmp, ", ")
	}

	for _, p := range f.Parameters {
		if v := cppTypeInput(f, p.Value); v != "" {
			tmp = append(tmp, fmt.Sprintf("%v %v", v, parser.CleanName(p.Name, p.Value)))
		} else {
			f.Access = "unsupported_CppHeaderInput"
			return f.Access
		}
	}

	return strings.Join(tmp, ", ")
}
