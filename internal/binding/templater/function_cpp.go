package templater

import (
	"fmt"
	"strings"

	"github.com/therecipe/qt/internal/binding/converter"
	"github.com/therecipe/qt/internal/binding/parser"
)

func cppFunctionSignal(f *parser.Function) string {

	//TODO: parse from docs
	var constP string
	if strings.Contains(strings.Split(f.Signature, ")")[1], "const") {
		constP = " const"
	}

	var originalInput string
	for _, p := range f.Parameters {
		if p.Name == "" {
			originalInput += "v, "
		} else {
			originalInput += fmt.Sprintf("%v, ", p.Name)
		}
	}
	originalInput = strings.TrimSuffix(originalInput, ", ")

	if strings.Contains(f.Virtual, "impure") && f.Output == "void" {
		if isDerivedFromPure(f) {
			return fmt.Sprintf("void %v(%v)%v { callback%v%v%v(%v); }", f.Name, converter.CppBodyInputCallback(f), constP, f.Class(), strings.Title(f.Name), cppFunctionSignalOverload(f), converter.CppBodyOutputCallback(f))
		} else {
			if f.Meta == "slot" || isDerivedFromSlot(f) {
				return fmt.Sprintf("void %v(%v)%v { if (!callback%v%v%v(%v)) { %v::%v(%v); }; }", f.Name, converter.CppBodyInputCallback(f), constP, f.Class(), strings.Title(f.Name), cppFunctionSignalOverload(f), converter.CppBodyOutputCallback(f), f.Class(), f.Name, originalInput)
			} else {
				return fmt.Sprintf("void %v(%v)%v { callback%v%v%v(%v); }", f.Name, converter.CppBodyInputCallback(f), constP, f.Class(), strings.Title(f.Name), cppFunctionSignalOverload(f), converter.CppBodyOutputCallback(f))
			}
		}
	}
	return fmt.Sprintf("void Signal_%v%v(%v) { callback%v%v%v(%v); }", strings.Title(f.Name), cppFunctionSignalOverload(f), converter.CppBodyInputCallback(f), f.Class(), strings.Title(f.Name), cppFunctionSignalOverload(f), converter.CppBodyOutputCallback(f))
}

func cppFunctionSignalOverload(f *parser.Function) string {
	if f.Overload {
		return f.OverloadNumber
	}
	return ""
}

func cppFunction(f *parser.Function) string {

	if isGeneric(f) {
		var tmp string
		for _, m := range jniGenericModes(f) {
			f.TemplateMode = m
			tmp += fmt.Sprintf("%v{\n\t%v\n}\n", cppFunctionHeader(f), cppFunctionBody(f))
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

	if f.Name == "objectNameAbs" || f.Name == "setObjectNameAbs" {
		o += fmt.Sprintf("if (dynamic_cast<My%v*>(static_cast<%v*>(ptr))) {\n\t\t", f.Class(), f.Class())
	}

	if converter.CppHeaderOutput(f) != "void" {
		o += "return "
	}

	var tmpMeta string
	if strings.Contains(f.Virtual, "impure") {
		if isDerivedFromSlot(f) {
			tmpMeta = f.Meta
			f.Meta = "slot"
		}
	}

	switch f.Meta {
	case "constructor":
		{
			if hasVirtualFunction(parser.ClassMap[f.Class()]) {
				o += fmt.Sprintf("new My%v(%v)", f.Class(), converter.CppBodyInput(f))
			} else {
				o += fmt.Sprintf("new %v(%v)", f.Class(), converter.CppBodyInput(f))
			}
		}
	case "slot":
		{
			if f.Static {
				o += fmt.Sprintf("QMetaObject::invokeMethod(%v::instance(), \"%v\"%v)", f.Class(), f.Name, converter.CppBodyInput(f))
			} else {
				o += fmt.Sprintf("QMetaObject::invokeMethod(static_cast<%v*>(ptr), \"%v\"%v)", f.Class(), f.Name, converter.CppBodyInput(f))
			}
		}
	case "plain", "destructor":
		{
			if f.Static {
				o += converter.CppBodyOutput(f, fmt.Sprintf("%v::%v%v(%v)", f.Class(), f.Name, converter.DeduceGeneric(f), converter.CppBodyInput(f)))
			} else {
				if f.Output == "T" && f.Class() == "QObject" {
					o += converter.CppBodyOutput(f, fmt.Sprintf("static_cast<%v*>(ptr)->%v<QObject*>(%v)", f.Class(), f.Name, converter.CppBodyInput(f)))
				} else if f.Output == "T" && f.Class() == "QMediaService" {
					o += converter.CppBodyOutput(f, fmt.Sprintf("static_cast<%v*>(ptr)->%v<QMediaControl*>(%v)", f.Class(), f.Name, converter.CppBodyInput(f)))
				} else {
					if parser.ClassMap[f.Class()].IsQObjectSubClass() {
						o += converter.CppBodyOutput(f, fmt.Sprintf("static_cast<%v*>(ptr)->%v%v(%v)", f.Class(), f.Name, converter.DeduceGeneric(f), converter.CppBodyInput(f)))
					} else {
						if f.Name == "objectNameAbs" || f.Name == "setObjectNameAbs" {
							o += converter.CppBodyOutput(f, fmt.Sprintf("static_cast<My%v*>(ptr)->%v%v(%v)", f.Class(), f.Name, converter.DeduceGeneric(f), converter.CppBodyInput(f)))
						} else {
							o += converter.CppBodyOutput(f, fmt.Sprintf("static_cast<%v*>(ptr)->%v%v(%v)", f.Class(), f.Name, converter.DeduceGeneric(f), converter.CppBodyInput(f)))
						}
					}
				}
			}
		}
	case "signal":
		{
			if converter.IsPrivateSignal(f) {
				o += fmt.Sprintf("QObject::%v(%v, &%v::%v, static_cast<My%v*>(ptr), static_cast<%v (My%v::*)(%v)>(&My%v::Signal_%v%v));", strings.ToLower(f.SignalMode), fmt.Sprintf("static_cast<%v*>(%v)", f.Class(), "ptr"), f.Class(), f.Name, f.Class(), f.Output, f.Class(), converter.CppBodyInput(f), f.Class(), strings.Title(f.Name), cppFunctionSignalOverload(f))
			} else {
				o += fmt.Sprintf("QObject::%v(%v, static_cast<void (%v::*)(%v)>(&%v::%v), static_cast<My%v*>(ptr), static_cast<%v (My%v::*)(%v)>(&My%v::Signal_%v%v));", strings.ToLower(f.SignalMode), fmt.Sprintf("static_cast<%v*>(%v)", f.Class(), "ptr"), f.Class(), converter.CppBodyInput(f), f.Class(), f.Name, f.Class(), f.Output, f.Class(), converter.CppBodyInput(f), f.Class(), strings.Title(f.Name), cppFunctionSignalOverload(f))
			}
		}
	default:
		{
			f.Access = "unsupported_CppFunctionBody"
			return f.Access
		}
	}

	if tmpMeta != "" {
		f.Meta = tmpMeta
	}

	if f.Name == "objectNameAbs" {
		o += fmt.Sprintf(";\n\t}\n\treturn QString(\"%v_BASE\").toUtf8().data()", f.Class())
	} else if f.Name == "setObjectNameAbs" {
		return o + ";\n\t}"
	}

	return o + ";"
}
