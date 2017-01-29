package templater

import (
	"bytes"
	"crypto/sha1"
	"encoding/hex"
	"fmt"
	"strings"

	"github.com/therecipe/qt/internal/binding/converter"
	"github.com/therecipe/qt/internal/binding/parser"
)

func cppFunctionCallback(function *parser.Function) string {
	var output = fmt.Sprintf("%v { %v };", cppFunctionCallbackHeader(function), cppFunctionCallbackBody(function))
	if function.IsSupported() {
		return cppFunctionCallbackWithGuards(function, output)
	}
	return ""
}

func cppFunctionCallbackWithGuards(function *parser.Function, output string) string {
	switch {
	case
		function.Fullname == "QProcess::nativeArguments", function.Fullname == "QProcess::setNativeArguments",
		function.Fullname == "QAbstractEventDispatcher::registerEventNotifier", function.Fullname == "QAbstractEventDispatcher::unregisterEventNotifier":
		{
			return fmt.Sprintf("#ifdef Q_OS_WIN\n\t\t%v\n\t#endif", output)
		}

	case
		function.Fullname == "QSensorGesture::detected":
		{
			return fmt.Sprintf("#ifdef Q_QDOC\n\t\t%v\n\t#endif", output)
		}
	}

	return output
}

func cppFunctionCallbackHeader(function *parser.Function) string {
	return fmt.Sprintf("%v %v%v(%v)%v",

		function.Output,

		func() string {
			if function.Meta == parser.SIGNAL {
				return fmt.Sprintf("Signal_%v", strings.Title(function.Name))
			}
			if strings.HasPrefix(function.Name, parser.TILDE) && !parser.State.Moc {
				return strings.Replace(function.Name, parser.TILDE, fmt.Sprintf("%vMy", parser.TILDE), -1)
			}
			return function.Name
		}(),

		func() string {
			if function.Meta == parser.SIGNAL {
				return function.OverloadNumber
			}
			return ""
		}(),

		converter.CppInputParametersForCallbackHeader(function),

		func() string {
			if strings.Contains(function.Signature, ") const") {
				return " const"
			}
			return ""
		}(),
	)
}

func cppFunctionCallbackBody(function *parser.Function) string {
	return fmt.Sprintf("%v%v%v;",

		converter.CppInputParametersForCallbackBodyPrePack(function),

		func() string {
			if converter.CppHeaderOutput(function) != parser.VOID {
				return "return "
			}
			return ""
		}(),

		func() string {
			var output = fmt.Sprintf("callback%v_%v%v(%v)", function.ClassName(), strings.Replace(strings.Title(function.Name), parser.TILDE, "Destroy", -1), function.OverloadNumber, converter.CppInputParametersForCallbackBody(function))

			if converter.CppHeaderOutput(function) != parser.VOID {
				output = converter.CppInput(output, function.Output, function)
			}

			return output
		}(),
	)
}

func cppFunction(function *parser.Function) string {
	var output = fmt.Sprintf("%v\n{\n%v\n}", cppFunctionHeader(function), cppFunctionBodyWithGuards(function))
	if function.IsSupported() {
		return output
	}
	return ""
}

func cppFunctionHeader(function *parser.Function) string {
	var output = fmt.Sprintf("%v %v(%v)", converter.CppHeaderOutput(function), converter.CppHeaderName(function), converter.CppHeaderInput(function))
	if function.IsSupported() {
		return output
	}
	return ""
}

func cppFunctionBodyWithGuards(function *parser.Function) string {

	if function.Default {
		switch {
		case
			strings.HasPrefix(function.ClassName(), "QMac") && !strings.HasPrefix(parser.State.ClassMap[function.ClassName()].Module, "QMac"):
			{
				return fmt.Sprintf("#ifdef Q_OS_OSX\n%v%v\n#endif", cppFunctionBody(function), cppFunctionBodyFailed(function))
			}
		}
	} else {
		switch {
		case
			function.Fullname == "QMenu::setAsDockMenu":
			{
				return fmt.Sprintf("#ifdef Q_OS_OSX\n%v%v\n#endif", cppFunctionBody(function), cppFunctionBodyFailed(function))
			}

		case
			function.Fullname == "QProcess::nativeArguments", function.Fullname == "QProcess::setNativeArguments",
			function.Fullname == "QAbstractEventDispatcher::registerEventNotifier", function.Fullname == "QAbstractEventDispatcher::unregisterEventNotifier":
			{
				return fmt.Sprintf("#ifdef Q_OS_WIN\n%v%v\n#endif", cppFunctionBody(function), cppFunctionBodyFailed(function))
			}

		case
			function.Fullname == "QApplication::navigationMode", function.Fullname == "QApplication::setNavigationMode",
			function.Fullname == "QWidget::hasEditFocus", function.Fullname == "QWidget::setEditFocus":
			{
				return fmt.Sprintf("#ifdef QT_KEYPAD_NAVIGATION\n%v%v\n#endif", cppFunctionBody(function), cppFunctionBodyFailed(function))
			}

		case
			function.Fullname == "QMenuBar::defaultAction", function.Fullname == "QMenuBar::setDefaultAction":
			{
				return fmt.Sprintf("#ifdef Q_OS_WINCE\n%v%v\n#endif", cppFunctionBody(function), cppFunctionBodyFailed(function))
			}

		case
			function.Fullname == "QWidget::setupUi", function.Fullname == "QSensorGesture::detected":
			{
				return fmt.Sprintf("#ifdef Q_QDOC\n%v%v\n#endif", cppFunctionBody(function), cppFunctionBodyFailed(function))
			}

		case
			function.Fullname == "QTextDocument::print", function.Fullname == "QPlainTextEdit::print", function.Fullname == "QTextEdit::print":
			{
				return fmt.Sprintf("#ifndef Q_OS_IOS\n%v%v\n#endif", cppFunctionBody(function), cppFunctionBodyFailed(function))
			}
		}
	}

	return cppFunctionBody(function)
}

func cppFunctionBodyFailed(function *parser.Function) string {
	if converter.CppHeaderOutput(function) != parser.VOID {
		return fmt.Sprintf("\n#else\n\treturn %v;", converter.CppOutputParametersFailed(function))
	}
	return ""
}

func cppFunctionBody(function *parser.Function) string {

	var polyinputs, polyName = function.PossiblePolymorphic(false)
	var polyinputsSelf, _ = function.PossiblePolymorphic(true)

	if (len(polyinputsSelf) == 0 && len(polyinputs) == 0) ||
		!(function.Meta == parser.PLAIN || function.Meta == parser.GETTER || function.Meta == parser.SETTER || function.Meta == parser.SLOT) ||
		strings.HasPrefix(function.Name, parser.TILDE) {

		return cppFunctionBodyInternal(function)
	}

	var bb = new(bytes.Buffer)
	defer bb.Reset()

	bb.WriteString("\t")

	var deduce = func(input []string, polyName string, inner bool, body string) string {
		var bb = new(bytes.Buffer)
		defer bb.Reset()
		for _, polyType := range input {
			if polyType == "QObject" || polyType == input[len(input)-1] {
				continue
			}
			fmt.Fprintf(bb, "if (dynamic_cast<%v*>(static_cast<QObject*>(%v))) {\n", polyType, polyName)
			fmt.Fprintf(bb, "\t%v\n", func() string {
				var ibody = strings.Replace(body, "static_cast<"+input[len(input)-1]+"*>("+polyName+")", "static_cast<"+polyType+"*>("+polyName+")", -1)
				if inner {
					return ibody
				}
				if strings.Count(ibody, "\n") > 1 {
					return "\t" + strings.Replace(ibody, "\n", "\n\t", -1)
				}
				return ibody
			}())
			fmt.Fprint(bb, "\t} else ")
		}
		if bb.String() == "" {
			return body
		}
		fmt.Fprintf(bb, "{\n\t%v\n\t}", func() string {
			if strings.Count(body, "\n") > 1 {
				return "\t" + strings.Replace(body, "\n", "\n\t", -1)
			}
			return body
		}())
		return bb.String()
	}

	if function.Static {
		fmt.Fprint(bb, deduce(polyinputs, polyName, true, cppFunctionBodyInternal(function)))
	} else if function.Meta == parser.GETTER || function.Meta == parser.SETTER || function.Meta == parser.SLOT {
		fmt.Fprint(bb, deduce(polyinputsSelf, "ptr", false, cppFunctionBodyInternal(function)))
	} else {
		fmt.Fprint(bb, deduce(polyinputsSelf, "ptr", false, deduce(polyinputs, polyName, true, cppFunctionBodyInternal(function))))
	}

	return bb.String()
}

func cppFunctionBodyInternal(function *parser.Function) string {

	switch function.Meta {
	case parser.CONSTRUCTOR:
		{
			return fmt.Sprintf("%v\treturn new %v%v(%v);",
				func() string {
					if parser.State.ClassMap[function.ClassName()].IsSubClassOf("QCoreApplication") {
						return `	static int argcs = argc;
	static char** argvs = static_cast<char**>(malloc(argcs * sizeof(char*)));

	QList<QByteArray> aList = QByteArray(argv).split('|');
	for (int i = 0; i < argcs; i++)
		argvs[i] = (new QByteArray(aList.at(i)))->data();

`
					}
					return ""
				}(),

				func() string {
					var class = parser.State.ClassMap[function.ClassName()]
					if !parser.State.Moc {
						if class.HasCallbackFunctions() {
							return "My"
						}
					}
					return ""
				}(),

				func() string {
					if c := parser.State.ClassMap[function.ClassName()]; c != nil && c.Fullname != "" {
						return c.Fullname
					}
					return function.ClassName()
				}(),

				converter.CppInputParameters(function),
			)
		}

	case parser.SLOT:
		{
			var (
				functionOutputType string
				bb                 = new(bytes.Buffer)
			)
			defer bb.Reset()

			fmt.Fprint(bb, "\t")

			if converter.CppHeaderOutput(function) != parser.VOID {
				functionOutputType = converter.CppInputParametersForSlotArguments(function, &parser.Parameter{Name: "returnArg", Value: function.Output})
				if function.Output != "void*" && !parser.State.ClassMap[strings.TrimSuffix(functionOutputType, "*")].IsSubClassOfQObject() {
					functionOutputType = strings.TrimSuffix(functionOutputType, "*")
				}
				fmt.Fprintf(bb, "%v returnArg;\n\t", functionOutputType)
			}

			fmt.Fprintf(bb, "QMetaObject::invokeMethod(static_cast<%v*>(ptr), \"%v\"%v%v);",

				function.ClassName(),

				function.Name,

				func() string {
					if converter.CppHeaderOutput(function) != parser.VOID {

						if c, _ := function.Class(); c.Module == parser.MOC && parser.IsPackedMap(function.Output) {
							var tHash = sha1.New()
							tHash.Write([]byte(function.Output))
							return fmt.Sprintf(", Q_RETURN_ARG(%v, returnArg)", strings.Replace(functionOutputType, parser.CleanValue(function.Output), fmt.Sprintf("type%v", hex.EncodeToString(tHash.Sum(nil)[:3])), -1))
						}

						return fmt.Sprintf(", Q_RETURN_ARG(%v, returnArg)", functionOutputType)
					}
					return ""
				}(),

				converter.CppInputParametersForSlotInvoke(function),
			)

			if converter.CppHeaderOutput(function) != parser.VOID {
				fmt.Fprintf(bb, "\n\treturn %v;", converter.CppOutput("returnArg", functionOutputType, function))
			}

			return bb.String()
		}

	case parser.PLAIN, parser.DESTRUCTOR:
		{
			if (function.Meta == parser.DESTRUCTOR || strings.HasPrefix(function.Name, parser.TILDE)) && function.Default {
				return ""
			}

			if function.Fullname == "SailfishApp::application" || function.Fullname == "SailfishApp::main" {
				return fmt.Sprintf(`	static int argcs = argc;
	static char** argvs = static_cast<char**>(malloc(argcs * sizeof(char*)));

	QList<QByteArray> aList = QByteArray(argv).split('|');
	for (int i = 0; i < argcs; i++)
		argvs[i] = (new QByteArray(aList.at(i)))->data();

	return %v(%v);`,

					function.Fullname,

					converter.CppInputParameters(function),
				)
			}

			return fmt.Sprintf("\t%v%v;",

				func() string {
					if converter.CppHeaderOutput(function) != parser.VOID {
						return "return "
					}
					return ""
				}(),

				converter.CppOutputParameters(function, fmt.Sprintf("%v%v%v(%v)%v",

					func() string {
						var c, _ = function.Class()
						//TODO:
						if c.Name == "QAndroidJniEnvironment" && function.Meta == parser.PLAIN && strings.HasPrefix(function.Name, "Exception") {
							return "({ QAndroidJniEnvironment env; env->"
						}
						if function.NonMember {
							return ""
						}
						if function.Static {
							return fmt.Sprintf("%v::", function.ClassName())
						}
						return fmt.Sprintf("static_cast<%v*>(ptr)->",
							func() string {
								if c.Fullname != "" {
									return c.Fullname
								}
								if strings.HasSuffix(function.Name, "_atList") {
									if function.IsMap {
										return fmt.Sprintf("%v<%v,%v>", parser.CleanValue(function.Container), function.Parameters[0].Value, strings.TrimPrefix(function.Output, "const "))
									}
									return fmt.Sprintf("%v<%v>", parser.CleanValue(function.Container), strings.TrimPrefix(function.Output, "const "))
								}
								if strings.HasSuffix(function.Name, "_setList") {
									if len(function.Parameters) == 2 {
										return fmt.Sprintf("%v<%v,%v>", parser.CleanValue(function.Container), function.Parameters[0].Value, strings.TrimPrefix(function.Parameters[1].Value, "const "))
									}
									return fmt.Sprintf("%v<%v>", parser.CleanValue(function.Container), strings.TrimPrefix(function.Parameters[0].Value, "const "))
								}
								if strings.HasSuffix(function.Name, "_newList") {
									//will be overriden
								}
								if strings.HasSuffix(function.Name, "_keyList") {
									//will be overriden
								}
								return function.ClassName()
							}(),
						)
					}(),

					func() string {
						if function.Default {
							if parser.State.Moc {
								return fmt.Sprintf("%v::%v", parser.State.ClassMap[function.ClassName()].GetBases()[0], function.Name)
							} else {
								return fmt.Sprintf("%v::%v", function.ClassName(), function.Name)
							}
						}
						return function.Name
					}(),

					converter.CppOutputParametersDeducedFromGeneric(function), converter.CppInputParameters(function),
					//TODO:
					func() string {
						var c, _ = function.Class()
						if c.Name == "QAndroidJniEnvironment" && function.Meta == parser.PLAIN && strings.HasPrefix(function.Name, "Exception") {
							return "; })"
						}
						return ""
					}())))
		}

	case parser.GETTER:
		{
			return fmt.Sprintf("\treturn %v;", converter.CppOutputParameters(function,
				func() string {
					if function.Static {
						return function.Fullname
					}
					return fmt.Sprintf("static_cast<%v*>(ptr)->%v", func() string {
						if c := parser.State.ClassMap[function.ClassName()]; c != nil && c.Fullname != "" {
							return c.Fullname
						}
						return function.ClassName()
					}(), function.Name)
				}()))
		}

	case parser.SETTER:
		{
			var function = *function
			function.Name = function.TmpName
			function.Fullname = fmt.Sprintf("%v::%v", function.ClassName(), function.Name)

			return fmt.Sprintf("\t%v = %v;", converter.CppOutputParameters(&function,
				func() string {
					if function.Static {
						return function.Fullname
					}
					return fmt.Sprintf("static_cast<%v*>(ptr)->%v", func() string {
						if c := parser.State.ClassMap[function.ClassName()]; c != nil && c.Fullname != "" {
							return c.Fullname
						}
						return function.ClassName()
					}(), function.Name)
				}()),

				converter.CppInputParameters(&function),
			)
		}

	case parser.SIGNAL:
		{
			var my string
			if !parser.State.Moc {
				my = "My"
			}
			if converter.IsPrivateSignal(function) {
				return fmt.Sprintf("\tQObject::%v(static_cast<%v*>(ptr), &%v::%v, static_cast<%v%v*>(ptr), static_cast<%v (%v%v::*)(%v)>(&%v%v::Signal_%v%v));", strings.ToLower(function.SignalMode), function.ClassName(), function.ClassName(), function.Name, my, function.ClassName(), function.Output, my, function.ClassName(), converter.CppInputParametersForSignalConnect(function), my, function.ClassName(), strings.Title(function.Name), function.OverloadNumber)
			}
			return fmt.Sprintf("\tQObject::%v(static_cast<%v*>(ptr), static_cast<%v (%v::*)(%v)>(&%v::%v), static_cast<%v%v*>(ptr), static_cast<%v (%v%v::*)(%v)>(&%v%v::Signal_%v%v));",
				strings.ToLower(function.SignalMode),

				function.ClassName(), function.Output, function.ClassName(), converter.CppInputParametersForSignalConnect(function), function.ClassName(), function.Name,

				my, function.ClassName(), function.Output, my, function.ClassName(), converter.CppInputParametersForSignalConnect(function), my, function.ClassName(), strings.Title(function.Name), function.OverloadNumber)
		}
	}

	function.Access = "unsupported_cppFunctionBody"
	return function.Access
}
