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

	if isTemplate(f) {
		var tmp string
		for _, m := range []string{"Int", "Boolean", "Void"} {
			if m == "Void" && (f.Fullname == "QAndroidJniObject::getField" || f.Fullname == "QAndroidJniObject::getStaticField") {
			} else {
				f.TemplateMode = m
				tmp += fmt.Sprintf("%v{\n\t%v\n}\n", cppFunctionHeader(f), cppFunctionBody(f))
			}
		}
		f.TemplateMode = ""
		return tmp
	}

	return fmt.Sprintf("%v{\n\t%v\n}", cppFunctionHeader(f), cppFunctionBody(f))
}

func cppFunctionHeader(f *parser.Function) string {
	return fmt.Sprintf("%v %v(%v)", converter.CppHeaderOutput(f), converter.CppHeaderName(f), converter.CppHeaderInput(f))
}

func cppFunctionBody(f *parser.Function) (o string) {

	/*
		for _, p := range f.Parameters {
			if strings.Contains(p.Value, "**") && p.Name == "argv" {
				o += "QList<QByteArray> aList = QByteArray(argv).split('|');\n"
				o += "\tQVarLengthArray<const char*> argvs(argc);\n"
				o += "\tstatic int argcs = argc;\n"
				o += "\tfor (int i = 0; i < argc; i++)\n"
				o += "\t\targvs[i] = static_cast<const char*>(aList[i].constData());\n\n\t"
			}
		}
	*/

	for _, p := range f.Parameters {
		if strings.Contains(p.Value, "**") && p.Name == "argv" {
			o += "QList<QByteArray> aList = QByteArray(argv).split('|');\n"
			o += "\tchar *argvs[argc];\n"
			o += "\tstatic int argcs = argc;\n"
			o += "\tfor (int i = 0; i < argc; i++)\n"
			o += "\t\targvs[i] = aList[i].data();\n\n\t"
		}
	}

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
			o += converter.CppBodyOutput(f, fmt.Sprintf("%v::%v%v(%v)", f.Class(), f.Name, converter.DeduceTemplate(f), converter.CppBodyInput(f)))
		} else {
			if f.Output == "T" && f.Class() == "QObject" {
				o += converter.CppBodyOutput(f, fmt.Sprintf("static_cast<%v*>(ptr)->%v<QObject*>(%v)", f.Class(), f.Name, converter.CppBodyInput(f)))
			} else if f.Output == "T" && f.Class() == "QMediaService" {
				o += converter.CppBodyOutput(f, fmt.Sprintf("static_cast<%v*>(ptr)->%v<QMediaControl*>(%v)", f.Class(), f.Name, converter.CppBodyInput(f)))
			} else {
				o += converter.CppBodyOutput(f, fmt.Sprintf("static_cast<%v*>(ptr)->%v%v(%v)", f.Class(), f.Name, converter.DeduceTemplate(f), converter.CppBodyInput(f)))
			}
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
