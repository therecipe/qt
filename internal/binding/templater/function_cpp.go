package templater

import (
	"bytes"
	"fmt"
	"strings"

	"github.com/therecipe/qt/internal/binding/converter"
	"github.com/therecipe/qt/internal/binding/parser"
)

func cppFunctionCallback(function *parser.Function) string {
	var output = fmt.Sprintf("%v { %v };", cppFunctionCallbackHeader(function), cppFunctionCallbackBody(function))
	if functionIsSupported(parser.ClassMap[function.Class()], function) {
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
			if strings.HasPrefix(function.Name, parser.TILDE) && parser.ClassMap[function.Class()].Module != parser.MOC {
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
	return fmt.Sprintf("%v%v;",

		func() string {
			if converter.CppHeaderOutput(function) != parser.VOID {
				return "return "
			}
			return ""
		}(),

		func() string {
			var output = fmt.Sprintf("callback%v_%v%v(%v)", function.Class(), strings.Replace(strings.Title(function.Name), parser.TILDE, "Destroy", -1), function.OverloadNumber, converter.CppInputParametersForCallbackBody(function))

			if converter.CppHeaderOutput(function) != parser.VOID {
				output = converter.CppInput(output, function.Output, function)
			}

			return output
		}(),
	)
}

func cppFunction(function *parser.Function) string {
	var output = fmt.Sprintf("%v\n{\n%v\n}", cppFunctionHeader(function), cppFunctionBodyWithGuards(function))
	if functionIsSupported(nil, function) {
		return output
	}
	return ""
}

func cppFunctionHeader(function *parser.Function) string {
	var output = fmt.Sprintf("%v %v(%v)", converter.CppHeaderOutput(function), converter.CppHeaderName(function), converter.CppHeaderInput(function))
	if functionIsSupported(nil, function) {
		return output
	}
	return ""
}

func cppFunctionBodyWithGuards(function *parser.Function) string {

	if function.Default {
		switch {
		case
			strings.HasPrefix(function.Class(), "QMac") && !strings.HasPrefix(parser.ClassMap[function.Class()].Module, "QMac"):
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

	switch function.Meta {
	case parser.CONSTRUCTOR:
		{
			return fmt.Sprintf("%v\treturn new %v%v(%v);",
				func() string {
					if function.Name == "QCoreApplication" || function.Name == "QGuiApplication" || function.Name == "QApplication" {
						return `	QList<QByteArray> aList = QByteArray(argv).split('|');
	char *argvs[argc];
	static int argcs = argc;
	for (int i = 0; i < argc; i++)
		argvs[i] = const_cast<char*>(aList[i].constData());

`
					}
					return ""
				}(),

				func() string {
					var class = parser.ClassMap[function.Class()]
					if class.Module != parser.MOC {
						if classNeedsCallbackFunctions(class) {
							return "My"
						}
					}
					return ""
				}(),

				function.Class(),

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
				if function.Output != "void*" && !parser.ClassMap[strings.TrimSuffix(functionOutputType, "*")].IsQObjectSubClass() {
					functionOutputType = strings.TrimSuffix(functionOutputType, "*")
				}
				fmt.Fprintf(bb, "%v returnArg;\n\t", functionOutputType)
			}

			fmt.Fprintf(bb, "QMetaObject::invokeMethod(static_cast<%v*>(ptr), \"%v\"%v%v);",

				function.Class(),

				function.Name,

				func() string {
					if converter.CppHeaderOutput(function) != parser.VOID {
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
				return fmt.Sprintf(`	QList<QByteArray> aList = QByteArray(argv).split('|');
	char *argvs[argc];
	static int argcs = argc;
	for (int i = 0; i < argc; i++)
	argvs[i] = const_cast<char*>(aList[i].constData());

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

				converter.CppOutputParameters(function, fmt.Sprintf("%v%v%v(%v)",

					func() string {
						if function.Static {
							return fmt.Sprintf("%v::", function.Class())
						}
						return fmt.Sprintf("static_cast<%v*>(ptr)->", function.Class())
					}(),

					func() string {
						if function.Default {
							if parser.ClassMap[function.Class()].Module == parser.MOC {
								return fmt.Sprintf("%v::%v", parser.ClassMap[function.Class()].GetBases()[0], function.Name)
							} else {
								return fmt.Sprintf("%v::%v", function.Class(), function.Name)
							}
						}
						return function.Name
					}(),

					converter.CppOutputParametersDeducedFromGeneric(function), converter.CppInputParameters(function))))
		}

	case parser.GETTER:
		{
			return fmt.Sprintf("\treturn %v;", converter.CppOutputParameters(function,
				func() string {
					if function.Static {
						return function.Fullname
					}
					return fmt.Sprintf("static_cast<%v*>(ptr)->%v", function.Class(), function.Name)
				}()))
		}

	case parser.SETTER:
		{
			var function = *function
			function.Name = function.TmpName
			function.Fullname = fmt.Sprintf("%v::%v", function.Class(), function.Name)

			return fmt.Sprintf("\t%v = %v;", converter.CppOutputParameters(&function,
				func() string {
					if function.Static {
						return function.Fullname
					}
					return fmt.Sprintf("static_cast<%v*>(ptr)->%v", function.Class(), function.Name)
				}()),

				converter.CppInputParameters(&function),
			)
		}

	case parser.SIGNAL:
		{
			var my string
			if parser.ClassMap[function.Class()].Module != parser.MOC {
				my = "My"
			}
			if converter.IsPrivateSignal(function) {
				return fmt.Sprintf("\tQObject::%v(static_cast<%v*>(ptr), &%v::%v, static_cast<%v%v*>(ptr), static_cast<%v (%v%v::*)(%v)>(&%v%v::Signal_%v%v));", strings.ToLower(function.SignalMode), function.Class(), function.Class(), function.Name, my, function.Class(), function.Output, my, function.Class(), converter.CppInputParametersForSignalConnect(function), my, function.Class(), strings.Title(function.Name), function.OverloadNumber)
			}
			return fmt.Sprintf("\tQObject::%v(static_cast<%v*>(ptr), static_cast<%v (%v::*)(%v)>(&%v::%v), static_cast<%v%v*>(ptr), static_cast<%v (%v%v::*)(%v)>(&%v%v::Signal_%v%v));",
				strings.ToLower(function.SignalMode),

				function.Class(), function.Output, function.Class(), converter.CppInputParametersForSignalConnect(function), function.Class(), function.Name,

				my, function.Class(), function.Output, my, function.Class(), converter.CppInputParametersForSignalConnect(function), my, function.Class(), strings.Title(function.Name), function.OverloadNumber)
		}
	}

	function.Access = "unsupported_cppFunctionBody"
	return function.Access
}
