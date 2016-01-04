package templater

import (
	"fmt"
	"strings"

	"github.com/therecipe/qt/internal/binding/converter"
	"github.com/therecipe/qt/internal/binding/parser"
)

func goFunction(f *parser.Function) string {
	if f.Meta == "signal" || strings.Contains(f.Virtual, "impure") && f.Output == "void" {
		var tmp string
		for _, signalMode := range []string{"Connect", "Disconnect", "callback"} {
			f.SignalMode = signalMode
			if signalMode == "callback" {
				tmp += fmt.Sprintf("//export callback%v%v", f.Class(), strings.Replace(strings.Title(f.Name), "~", "Destroy", -1))
				if f.Overload {
					tmp += f.OverloadNumber
				}
				tmp += "\n"
			}
			tmp += fmt.Sprintf("%v{\n%v\n}\n\n", goFunctionHeader(f), goFunctionBody(f))
		}
		f.SignalMode = ""

		var isPrivate bool
		if f.Meta == "signal" {
			isPrivate = converter.IsPrivateSignal(f)
		}

		var tmpMeta = f.Meta
		f.Meta = "plain"
		if !isPrivate {
			tmp += fmt.Sprintf("%v{\n%v\n}\n\n", goFunctionHeader(f), goFunctionBody(f))
			if tmpMeta != "signal" {
				var tmpTmp = strings.Replace(fmt.Sprintf("%v{\n%v\n}\n\n", goFunctionHeader(f), goFunctionBody(f)), "_"+strings.Title(f.Name)+"(", "_"+strings.Title(f.Name)+"Default(", -1)
				tmp += strings.Replace(tmpTmp, ")"+strings.Title(f.Name)+"(", ")"+strings.Title(f.Name)+"Default(", -1)
			}
		}
		f.Meta = tmpMeta

		return tmp
	}

	if isGeneric(f) {
		var tmp string
		for _, m := range jniGenericModes(f) {
			f.TemplateMode = m
			tmp += fmt.Sprintf("%v{\n%v\n}\n", goFunctionHeader(f), goFunctionBody(f))
		}
		f.TemplateMode = ""
		return tmp
	}

	return fmt.Sprintf("%v{\n%v\n}", goFunctionHeader(f), goFunctionBody(f))
}

func goFunctionHeader(f *parser.Function) string {
	if f.Static || f.Meta == "constructor" || f.SignalMode == "callback" {
		return fmt.Sprintf("func %v", goInterface(f))
	}
	return fmt.Sprintf("func (ptr *%v)%v", f.Class(), goInterface(f))
}

func goInterface(f *parser.Function) string {
	if f.Virtual == "impure" && f.SignalMode == "callback" {
		if f.Meta == "slot" {
			return fmt.Sprintf("%v(%v)%v", converter.GoHeaderName(f), converter.GoHeaderInput(f), "bool")
		} else {
			return fmt.Sprintf("%v(%v)", converter.GoHeaderName(f), converter.GoHeaderInput(f))
		}
	}
	return fmt.Sprintf("%v(%v)%v", converter.GoHeaderName(f), converter.GoHeaderInput(f), converter.GoHeaderOutput(f))
}

func goFunctionBody(f *parser.Function) (o string) {

	if parser.ClassMap[f.Class()].Stub {
		if converter.GoHeaderOutput(f) != "" {
			return fmt.Sprintf("\nreturn %v", converter.GoBodyOutputFailed(f))
		}
		return ""
	}

	if f.SignalMode != "" {
		o += fmt.Sprintf("defer qt.Recovering(\"%v %v\")\n\n", strings.ToLower(f.SignalMode), f.Fullname)
	} else {
		o += fmt.Sprintf("defer qt.Recovering(\"%v\")\n\n", f.Fullname)
	}

	if !f.Static && f.Meta != "constructor" && f.SignalMode != "callback" {
		o += fmt.Sprintf("if ptr.Pointer() != nil {\n")
	}

	if converter.GoHeaderOutput(f) != "" {
		o += "return "
	}

	/*
		if (f.Meta == "destructor" && isObjectSubClass(f.Class())) || (f.Meta == "slot" && strings.Contains(f.Name, "deleteLater")) {
			o += "getSignal(ptr.ObjectName(), \"destroyed\").(func(QObject, QObject))(ptr, ptr)\n"
		}
	*/

	if f.SignalMode == "callback" {

		o += fmt.Sprintf("if signal := qt.GetSignal(C.GoString(ptrName), \"%v%v\"); signal != nil {\n", f.Name, cppFunctionSignalOverload(f))

		o += fmt.Sprintf("\tsignal.(%v)(%v)\n", converter.GoHeaderInputSignalFunction(f), converter.GoBodyInputSignalValues(f))

		if f.Virtual == "impure" {
			if f.Meta == "slot" {
				o += fmt.Sprintf("\treturn true\n}\nreturn false\n")
			} else {
				o += fmt.Sprintf("} else {\n\tNew%vFromPointer(ptr).%vDefault(%v)\n}", f.Class(), strings.Title(f.Name), converter.GoBodyInputSignalValues(f))
			}
		} else {
			o += "}\n"
		}

	} else {
		if strings.Contains(f.Virtual, "impure") && f.SignalMode != "" {
		} else {
			o += converter.GoBodyOutput(f, fmt.Sprintf("C.%v(%v)", converter.CppHeaderName(f), converter.GoBodyInput(f)))
		}
	}

	if f.SignalMode == "Connect" {
		if parser.ClassMap[f.Class()].IsQObjectSubClass() {
			o += fmt.Sprintf("\nqt.ConnectSignal(ptr.ObjectName(), \"%v%v\", f)", f.Name, cppFunctionSignalOverload(f))
		} else {
			o += fmt.Sprintf("\nqt.ConnectSignal(ptr.ObjectNameAbs(), \"%v%v\", f)", f.Name, cppFunctionSignalOverload(f))
		}
	}

	if f.SignalMode == "Disconnect" {
		if parser.ClassMap[f.Class()].IsQObjectSubClass() {
			o += fmt.Sprintf("\nqt.DisconnectSignal(ptr.ObjectName(), \"%v%v\")", f.Name, cppFunctionSignalOverload(f))
		} else {
			o += fmt.Sprintf("\nqt.DisconnectSignal(ptr.ObjectNameAbs(), \"%v%v\")", f.Name, cppFunctionSignalOverload(f))
		}
	}

	if (f.Meta == "destructor" && isObjectSubClass(f.Class())) || (f.Meta == "slot" && strings.Contains(f.Name, "deleteLater")) {
		o += "\nptr.SetPointer(nil)"
	}

	if !f.Static && f.Meta != "constructor" && f.SignalMode != "callback" {
		o += "\n}"
		if converter.GoHeaderOutput(f) != "" {
			o += fmt.Sprintf("\nreturn %v", converter.GoBodyOutputFailed(f))
		}
	}

	return o
}
