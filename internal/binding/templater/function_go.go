package templater

import (
	"fmt"
	"strings"

	"github.com/therecipe/qt/internal/binding/converter"
	"github.com/therecipe/qt/internal/binding/parser"
)

func goFunction(f *parser.Function) string {
	if f.Meta == "signal" && !f.Overload {
		var tmp string
		for _, signalMode := range []string{"Connect", "Disconnect", "callback"} {
			f.SignalMode = signalMode
			if signalMode == "callback" {
				tmp += fmt.Sprintf("//export callback%v%v\n", f.Class(), strings.Title(f.Name))
			}
			tmp += fmt.Sprintf("%v{\n%v\n}\n\n", goFunctionHeader(f), goFunctionBody(f))
		}
		f.SignalMode = ""
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
	return fmt.Sprintf("%v(%v)%v", converter.GoHeaderName(f), converter.GoHeaderInput(f), converter.GoHeaderOutput(f))
}

func goFunctionBody(f *parser.Function) (o string) {

	if parser.ClassMap[f.Class()].Stub {
		if converter.GoHeaderOutput(f) != "" {
			return fmt.Sprintf("\nreturn %v", converter.GoBodyOutputFailed(f))
		}
		return ""
	}

	o += "defer func() {\n"
	o += "if recover() != nil {\n"
	o += fmt.Sprintf("log.Println(\"recovered in %v\")\n", f.Fullname)
	o += "}\n"
	o += "}()\n\n"

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
		o += fmt.Sprintf("qt.GetSignal(C.GoString(ptrName), \"%v\").(%v)(%v)", f.Name, converter.GoHeaderInputSignalFunction(f), converter.GoBodyInputSignalValues(f))
	} else {
		o += converter.GoBodyOutput(f, fmt.Sprintf("C.%v(%v)", converter.CppHeaderName(f), converter.GoBodyInput(f)))
	}

	if f.SignalMode == "Connect" {
		o += fmt.Sprintf("\nqt.ConnectSignal(ptr.ObjectName(), \"%v\", f)", f.Name)
	}

	if f.SignalMode == "Disconnect" {
		o += fmt.Sprintf("\nqt.DisconnectSignal(ptr.ObjectName(), \"%v\")", f.Name)
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
