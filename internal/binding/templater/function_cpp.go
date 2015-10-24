package templater

import (
	"fmt"
	"strings"

	"github.com/therecipe/qt/internal/binding/converter"
	"github.com/therecipe/qt/internal/binding/parser"
)

func cppFunctionSignal(f *parser.Function) string {
	return fmt.Sprintf("void Signal_%v(%v){callback%v%v(%v);}", strings.Title(f.Name), converter.CppBodyInputCallback(f), f.Class(), strings.Title(f.Name), converter.CppBodyOutputCallback(f))
}

func cppFunction(f *parser.Function) string {
	return fmt.Sprintf("%v{\n\t%v\n}", cppFunctionHeader(f), cppFunctionBody(f))
}

func cppFunctionHeader(f *parser.Function) string {
	return fmt.Sprintf("%v %v(%v)", converter.CppHeaderOutput(f), converter.CppHeaderName(f), converter.CppHeaderInput(f))
}

func cppFunctionBody(f *parser.Function) (o string) {

	if converter.CppHeaderOutput(f) != "void" {
		o += "return "
	}

	switch f.Meta {
	case "constructor":
		o += fmt.Sprintf("new %v(%v)", f.Class(), converter.CppBodyInput(f))

	case "slot":
		if f.Static {
			o += fmt.Sprintf("QMetaObject::invokeMethod(%v::instance(), \"%v\"%v)", f.Class(), f.Name, converter.CppBodyInput(f))
		} else {
			o += fmt.Sprintf("QMetaObject::invokeMethod(static_cast<%v*>(ptr), \"%v\"%v)", f.Class(), f.Name, converter.CppBodyInput(f))
		}

	case "plain", "destructor":
		if f.Static {
			o += converter.CppBodyOutput(f, fmt.Sprintf("%v::%v(%v)", f.Class(), f.Name, converter.CppBodyInput(f)))
		} else {
			o += converter.CppBodyOutput(f, fmt.Sprintf("static_cast<%v*>(ptr)->%v(%v)", f.Class(), f.Name, converter.CppBodyInput(f)))
		}

	case "signal":
		if converter.PrivateSignal(f) {
			o += fmt.Sprintf("QObject::%v(%v, &%v::%v, static_cast<My%v*>(ptr), static_cast<%v (My%v::*)(%v)>(&My%v::Signal_%v));", strings.ToLower(f.SignalMode), fmt.Sprintf("static_cast<%v*>(%v)", f.Class(), "ptr"), f.Class(), f.Name, f.Class(), f.Output, f.Class(), converter.CppBodyInput(f), f.Class(), strings.Title(f.Name))
		} else {
			o += fmt.Sprintf("QObject::%v(%v, static_cast<void (%v::*)(%v)>(&%v::%v), static_cast<My%v*>(ptr), static_cast<%v (My%v::*)(%v)>(&My%v::Signal_%v));", strings.ToLower(f.SignalMode), fmt.Sprintf("static_cast<%v*>(%v)", f.Class(), "ptr"), f.Class(), converter.CppBodyInput(f), f.Class(), f.Name, f.Class(), f.Output, f.Class(), converter.CppBodyInput(f), f.Class(), strings.Title(f.Name))
		}
	default:
		f.Access = "unsupported_CppFunctionBody"
		return f.Access
	}

	return fmt.Sprintf("%v;", o)
}
