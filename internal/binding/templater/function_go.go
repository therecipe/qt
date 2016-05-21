package templater

import (
	"bytes"
	"fmt"
	"strings"

	"github.com/therecipe/qt/internal/binding/converter"
	"github.com/therecipe/qt/internal/binding/parser"
)

func goFunction(function *parser.Function) string {
	var output = fmt.Sprintf("%v{\n%v\n}", goFunctionHeader(function), goFunctionBody(function))
	if functionIsSupported(parser.ClassMap[function.Class()], function) {
		return output
	}
	return ""
}

func goFunctionHeader(function *parser.Function) string {
	return fmt.Sprintf("func %v %v%v(%v)%v",
		func() string {
			if function.Static || function.Meta == parser.CONSTRUCTOR || function.SignalMode == parser.CALLBACK {
				return ""
			}
			return fmt.Sprintf("(ptr *%v)", function.Class())
		}(),

		converter.GoHeaderName(function),

		func() string {
			if function.Default {
				return "Default"
			}
			return ""
		}(),

		converter.GoHeaderInput(function),

		converter.GoHeaderOutput(function),
	)
}

func goFunctionBody(function *parser.Function) string {
	var bb = new(bytes.Buffer)
	defer bb.Reset()

	if parser.ClassMap[function.Class()].Stub {
		if converter.GoHeaderOutput(function) != "" {
			return fmt.Sprintf("\nreturn %v", converter.GoOutputParametersFromCFailed(function))
		}
		return ""
	}

	fmt.Fprintf(bb, "defer qt.Recovering(\"%v%v\")\n\n",
		func() string {
			if function.SignalMode != "" {
				return strings.ToLower(function.SignalMode) + " "
			}
			return ""
		}(),

		function.Fullname,
	)

	if !(function.Static || function.Meta == parser.CONSTRUCTOR || function.SignalMode == parser.CALLBACK) {
		fmt.Fprintf(bb, "if ptr.Pointer() != nil {\n")
	}

	for _, parameter := range function.Parameters {
		if parameter.Value == "..." || (parameter.Value == "T" && parser.ClassMap[function.Class()].Module == "QtAndroidExtras" && function.TemplateMode == "") {
			for i := 0; i < 10; i++ {
				if parameter.Value == "T" {
					fmt.Fprintf(bb, "var p%v, d%v = assertion(%v, %v)\n", i, i, i, parameter.Name)
				} else {
					fmt.Fprintf(bb, "var p%v, d%v = assertion(%v, v...)\n", i, i, i)
				}
				fmt.Fprintf(bb, "if d%v != nil {\ndefer d%v()\n}\n", i, i)

				if parameter.Value == "T" {
					break
				}
			}
		}
	}

	if (function.Meta == parser.DESTRUCTOR || strings.Contains(function.Name, "deleteLater")) && function.SignalMode == "" {
		if needsCallbackFunctions(parser.ClassMap[function.Class()]) {
			fmt.Fprintf(bb, "qt.DisconnectAllSignals(ptr.ObjectName%v())\n",
				func() string {
					if !parser.ClassMap[function.Class()].IsQObjectSubClass() {
						return "Abs"
					}
					return ""
				}())
		}
	}

	if ((function.Meta == parser.PLAIN && function.SignalMode == "") ||
		(function.Meta == parser.SLOT && function.SignalMode == "") ||
		function.Meta == parser.CONSTRUCTOR || function.Meta == parser.DESTRUCTOR) ||
		(function.Meta == parser.SIGNAL && (function.SignalMode == "" || function.SignalMode == parser.CONNECT || function.SignalMode == parser.DISCONNECT)) ||
		(function.Meta == parser.GETTER || function.Meta == parser.SETTER) {

		//TODO:
		if functionIsSupported(parser.ClassMap[function.Class()], function) {
			cppFunction(function)
			if functionIsSupported(parser.ClassMap[function.Class()], function) {

				fmt.Fprintf(bb, "%v%v",

					func() string {
						if converter.GoHeaderOutput(function) != "" {
							return "return "
						}
						return ""
					}(),

					converter.GoOutputParametersFromC(function, fmt.Sprintf("C.%v%v(%v)",

						converter.CppHeaderName(function),

						func() string {
							if function.Default {
								return "Default"
							}
							return ""
						}(),

						converter.GoInputParametersForC(function))),
				)

			}
			function.Access = "public"
		}

	}

	switch function.SignalMode {
	case parser.CALLBACK:
		{

			fmt.Fprintf(bb, "%vif signal := qt.GetSignal(C.GoString(ptrName), \"%v%v\"); signal != nil {\n",
				func() string {
					if function.Meta != parser.SLOT {
						return "\n"
					}
					return ""
				}(), function.Name, function.OverloadNumber)

			if converter.GoHeaderOutput(function) == "" {
				fmt.Fprintf(bb, "signal.(%v)(%v)", converter.GoHeaderInputSignalFunction(function), converter.GoInputParametersForCallback(function))
			} else {
				fmt.Fprintf(bb, "return %v", converter.GoInput(fmt.Sprintf("signal.(%v)(%v)", converter.GoHeaderInputSignalFunction(function), converter.GoInputParametersForCallback(function)), function.Output, function))
			}

			fmt.Fprintf(bb, "\n}%v\n",
				func() string {
					if converter.GoHeaderOutput(function) == "" {
						if function.Virtual == parser.IMPURE {
							return "else{"
						}
					}
					return ""
				}(),
			)

			if converter.GoHeaderOutput(function) == "" {
				if function.Virtual == parser.IMPURE && functionIsSupportedDefault(function) {
					fmt.Fprintf(bb, "New%vFromPointer(ptr).%v%vDefault(%v)", strings.Title(function.Class()), strings.Title(function.Name), function.OverloadNumber, converter.GoInputParametersForCallback(function))
				}
			} else {
				if function.Virtual == parser.IMPURE && functionIsSupportedDefault(function) {
					fmt.Fprintf(bb, "\nreturn %v", converter.GoInput(fmt.Sprintf("New%vFromPointer(ptr).%v%vDefault(%v)", strings.Title(function.Class()), strings.Title(function.Name), function.OverloadNumber, converter.GoInputParametersForCallback(function)), function.Output, function))
				} else {
					fmt.Fprintf(bb, "\nreturn %v", converter.GoInput(converter.GoOutputParametersFromCFailed(function), function.Output, function))
				}
			}

			fmt.Fprintf(bb, "%v",
				func() string {
					if converter.GoHeaderOutput(function) == "" {
						if function.Virtual == parser.IMPURE {
							return "\n}"
						}
					}
					return ""
				}(),
			)

		}

	case parser.CONNECT, parser.DISCONNECT:
		{
			fmt.Fprintf(bb, "\nqt.%vSignal(ptr.ObjectName%v(), \"%v%v\"%v)",

				function.SignalMode,

				func() string {
					if !parser.ClassMap[function.Class()].IsQObjectSubClass() {
						return "Abs"
					}
					return ""
				}(),

				function.Name,

				function.OverloadNumber,

				func() string {
					if function.SignalMode == parser.CONNECT {
						return ", f"
					}
					return ""
				}(),
			)
		}
	}

	if (function.Meta == parser.DESTRUCTOR || strings.Contains(function.Name, "deleteLater")) && function.SignalMode == "" {
		fmt.Fprint(bb, "\nptr.SetPointer(nil)")
	}

	if !(function.Static || function.Meta == parser.CONSTRUCTOR || function.SignalMode == parser.CALLBACK) {
		fmt.Fprint(bb, "\n}")
		if converter.GoHeaderOutput(function) != "" {
			fmt.Fprintf(bb, "\nreturn %v", converter.GoOutputParametersFromCFailed(function))
		}
	}

	return bb.String()
}
