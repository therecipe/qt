package templater

import (
	"bytes"
	"fmt"
	"strings"

	"github.com/therecipe/qt/internal/binding/converter"
	"github.com/therecipe/qt/internal/binding/parser"
	"github.com/therecipe/qt/internal/utils"
)

func goFunction(function *parser.Function) string {
	var output = fmt.Sprintf("%v{\n%v\n}", goFunctionHeader(function), goFunctionBody(function))
	if function.IsSupported() {
		if UseStub() {
			if function.SignalMode != parser.CALLBACK {
				return output
			}
		} else {
			return output
		}
	}
	return ""
}

func goFunctionHeader(function *parser.Function) string {
	return fmt.Sprintf("func %v %v(%v)%v",
		func() string {
			if function.Static || function.Meta == parser.CONSTRUCTOR || function.SignalMode == parser.CALLBACK {
				return ""
			}
			return fmt.Sprintf("(ptr *%v)", function.ClassName())
		}(),

		converter.GoHeaderName(function),

		converter.GoHeaderInput(function),

		converter.GoHeaderOutput(function),
	)
}

func goFunctionBody(function *parser.Function) string {
	var bb = new(bytes.Buffer)
	defer bb.Reset()

	if utils.QT_DEBUG() {
		fmt.Fprintf(bb, "defer qt.Recover(\"\t%v%v%v(%v) %v\")\n",
			function.ClassName(),

			func() string {
				return strings.Repeat(" ", 45-len(function.ClassName()))
			}(),

			converter.GoHeaderName(function), converter.GoHeaderInput(function), converter.GoHeaderOutput(function))

		fmt.Fprintf(bb, "qt.Debug(\"\t%v%v%v(%v) %v\")\n",

			function.ClassName(),

			func() string {
				return strings.Repeat(" ", 45-len(function.ClassName()))
			}(),

			converter.GoHeaderName(function), converter.GoHeaderInput(function), converter.GoHeaderOutput(function))
	}

	if parser.State.ClassMap[function.ClassName()].Stub {
		if converter.GoHeaderOutput(function) != "" {
			return fmt.Sprintf("\nreturn %v", converter.GoOutputParametersFromCFailed(function))
		}
		return ""
	}

	if !(function.Static || function.Meta == parser.CONSTRUCTOR || function.SignalMode == parser.CALLBACK) {
		fmt.Fprintf(bb, "if ptr.Pointer() != nil {\n")
	}

	for _, parameter := range function.Parameters {
		if parameter.Value == "..." || (parameter.Value == "T" && parser.State.ClassMap[function.ClassName()].Module == "QtAndroidExtras" && function.TemplateModeJNI == "") {
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

	if function.SignalMode == "" || (function.Meta == parser.SIGNAL && function.SignalMode != parser.CALLBACK) {

		for _, alloc := range converter.GoInputParametersForCAlloc(function) {
			fmt.Fprint(bb, alloc)
		}

		var body = converter.GoOutputParametersFromC(function, fmt.Sprintf("C.%v(%v)", converter.CppHeaderName(function), converter.GoInputParametersForC(function)))
		fmt.Fprint(bb, func() string {
			if converter.GoHeaderOutput(function) != "" {
				switch {
				case function.NeedsFinalizer && parser.State.ClassMap[parser.CleanValue(function.Output)].IsSupported() || function.Meta == parser.CONSTRUCTOR && !(parser.State.ClassMap[function.Name].HasCallbackFunctions() || parser.State.ClassMap[function.Name].IsSubClassOfQObject()):
					{
						return fmt.Sprintf("var tmpValue = %v\nruntime.SetFinalizer(tmpValue, (%v).Destroy%v)\nreturn tmpValue%v",

							body,

							func() string {
								if function.TemplateModeJNI != "" {
									return fmt.Sprintf("*%v", parser.CleanValue(function.Output))
								}
								return converter.GoHeaderOutput(function)
							}(),

							func() string {
								if function.Meta == parser.CONSTRUCTOR {
									return function.Name
								}
								return parser.CleanValue(function.Output)
							}(),

							func() string {
								if function.TemplateModeJNI == "String" {
									return ".ToString()"
								}
								return ""
							}())
					}

				case parser.State.ClassMap[parser.CleanValue(function.Output)].IsSubClassOfQObject() && converter.GoHeaderOutput(function) != "unsafe.Pointer" || function.Meta == parser.CONSTRUCTOR && parser.State.ClassMap[parser.CleanValue(function.Name)].IsSubClassOfQObject():
					{
						return fmt.Sprintf("var tmpValue = %v\nif !qt.ExistsSignal(fmt.Sprint(tmpValue.Pointer()), \"QObject::destroyed\") {\ntmpValue.ConnectDestroyed(func(%v){ tmpValue.SetPointer(nil) })\n}\nreturn tmpValue",

							body,

							func() string {
								if parser.State.ClassMap[function.ClassName()].Module == "QtCore" {
									return "*QObject"
								}
								return "*core.QObject"
							}())
					}

				default:
					{
						if function.Name == "readData" && len(function.Parameters) == 2 {
							return fmt.Sprintf("var ret = %v\nif ret > 0 {\n*data = C.GoStringN(dataC, C.int(ret))\n}\nreturn ret", body)
						} else {
							return fmt.Sprintf("return %v", body)
						}
					}
				}
			}
			return body
		}())
	}

	switch function.SignalMode {
	case parser.CALLBACK:
		{
			fmt.Fprintf(bb, "%vif signal := qt.GetSignal(fmt.Sprint(ptr), \"%v::%v%v\"); signal != nil {\n",
				func() string {
					if function.Meta != parser.SLOT {
						return "\n"
					}
					return ""
				}(), function.ClassName(), function.Name, function.OverloadNumber)

			if converter.GoHeaderOutput(function) == "" {
				fmt.Fprintf(bb, "signal.(%v)(%v)", converter.GoHeaderInputSignalFunction(function), converter.GoInputParametersForCallback(function))
			} else {
				if function.Name == "readData" && len(function.Parameters) == 2 {
					fmt.Fprint(bb, "var retS = cGoUnpackString(data)\n")
					fmt.Fprintf(bb, "var ret = %v\n", converter.GoInput(fmt.Sprintf("signal.(%v)(%v)", converter.GoHeaderInputSignalFunction(function), converter.GoInputParametersForCallback(function)), function.Output, function))
					fmt.Fprint(bb, "if ret > 0 {\nC.memcpy(unsafe.Pointer(data.data), unsafe.Pointer(C.CString(retS)), C.size_t(ret))\n}\n")
					fmt.Fprint(bb, "return ret")
				} else {
					fmt.Fprintf(bb, "return %v", converter.GoInput(fmt.Sprintf("signal.(%v)(%v)", converter.GoHeaderInputSignalFunction(function), converter.GoInputParametersForCallback(function)), function.Output, function))
				}
			}

			fmt.Fprintf(bb, "\n}%v\n",
				func() string {
					if converter.GoHeaderOutput(function) == "" {
						if (function.Virtual != parser.PURE || (function.Synthetic && function.IsDerivedFromImpure())) && function.Meta != parser.SIGNAL {
							return "else{"
						}
					}
					return ""
				}(),
			)

			if converter.GoHeaderOutput(function) == "" {
				if (function.Virtual != parser.PURE || (function.Synthetic && function.IsDerivedFromImpure())) && function.Meta != parser.SIGNAL {
					fmt.Fprintf(bb, "New%vFromPointer(ptr).%v%vDefault(%v)", strings.Title(function.ClassName()), strings.Replace(strings.Title(function.Name), parser.TILDE, "Destroy", -1), function.OverloadNumber, converter.GoInputParametersForCallback(function))
				}
			} else {
				if (function.Virtual != parser.PURE || (function.Synthetic && function.IsDerivedFromImpure())) && function.Meta != parser.SIGNAL {
					if function.Name == "readData" && len(function.Parameters) == 2 {
						fmt.Fprint(bb, "var retS = cGoUnpackString(data)\n")
						fmt.Fprintf(bb, "var ret = %v\n", converter.GoInput(fmt.Sprintf("New%vFromPointer(ptr).%v%vDefault(%v)", strings.Title(function.ClassName()), strings.Replace(strings.Title(function.Name), parser.TILDE, "Destroy", -1), function.OverloadNumber, converter.GoInputParametersForCallback(function)), function.Output, function))
						fmt.Fprint(bb, "if ret > 0 {\nC.memcpy(unsafe.Pointer(data.data), unsafe.Pointer(C.CString(retS)), C.size_t(ret))\n}\n")
						fmt.Fprint(bb, "return ret")
					} else {
						fmt.Fprintf(bb, "\nreturn %v", converter.GoInput(fmt.Sprintf("New%vFromPointer(ptr).%v%vDefault(%v)", strings.Title(function.ClassName()), strings.Replace(strings.Title(function.Name), parser.TILDE, "Destroy", -1), function.OverloadNumber, converter.GoInputParametersForCallback(function)), function.Output, function))
					}
				} else {
					if converter.GoOutputParametersFromCFailed(function) == "nil" && !parser.State.Minimal {
						var (
							class, _ = function.Class()
							found    bool
							c, exist = parser.State.ClassMap[parser.CleanValue(function.Output)]
						)
						if exist && c.IsSupported() && !hasUnimplementedPureVirtualFunctions(c.Name) {
							for _, f := range c.Functions {
								if f.Meta == parser.CONSTRUCTOR && len(f.Parameters) == 0 {
									if c.Module == class.Module {
										fmt.Fprintf(bb, "\nreturn %v", converter.GoInput(fmt.Sprintf("%v()", converter.GoHeaderName(f)), function.Output, function))
									} else {
										fmt.Fprintf(bb, "\nreturn %v", converter.GoInput(fmt.Sprintf("%v.%v()", goModule(c.Module), converter.GoHeaderName(f)), function.Output, function))
									}
									found = true
									break
								}
							}
							if !found {
								for _, f := range c.Functions {
									if f.Meta == parser.CONSTRUCTOR && len(f.Parameters) == 1 {
										if c.Module == class.Module {
											fmt.Fprintf(bb, "\nreturn %v", converter.GoInput(fmt.Sprintf("%v(%v)", converter.GoHeaderName(f), converter.GoOutputFailed(f.Parameters[0].Value, f)), function.Output, function))
										} else {
											fmt.Fprintf(bb, "\nreturn %v", converter.GoInput(fmt.Sprintf("%v.%v(%v)", goModule(c.Module), converter.GoHeaderName(f), converter.GoOutputFailed(f.Parameters[0].Value, f)), function.Output, function))
										}
										found = true
										break
									}
								}
							}
						}

						if !found {
							//TODO:
							//function.Access = fmt.Sprintf("unsupported_FailedNilOutputInCallbacksForPureVirtualFunctions(%v)", function.Output)
							fmt.Fprintf(bb, "\nreturn %v", converter.GoInput(converter.GoOutputParametersFromCFailed(function), function.Output, function))
						}
					} else {
						fmt.Fprintf(bb, "\nreturn %v", converter.GoInput(converter.GoOutputParametersFromCFailed(function), function.Output, function))
					}
				}
			}

			fmt.Fprintf(bb, "%v",
				func() string {
					if converter.GoHeaderOutput(function) == "" {
						if (function.Virtual != parser.PURE || (function.Synthetic && function.IsDerivedFromImpure())) && function.Meta != parser.SIGNAL {
							return "\n}"
						}
					}
					return ""
				}(),
			)

		}

	case parser.CONNECT, parser.DISCONNECT:
		{
			fmt.Fprintf(bb, "\nqt.%vSignal(fmt.Sprint(ptr.Pointer()), \"%v::%v%v\"%v)",
				function.SignalMode,

				function.ClassName(),

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

	if (function.Meta == parser.DESTRUCTOR || function.Name == "deleteLater" || function.Name == "destroyed" || strings.HasPrefix(function.Name, parser.TILDE)) && function.SignalMode == "" {
		if parser.State.ClassMap[function.ClassName()].HasCallbackFunctions() || parser.State.ClassMap[function.ClassName()].IsSubClassOfQObject() {
			fmt.Fprint(bb, "\nqt.DisconnectAllSignals(fmt.Sprint(ptr.Pointer()))")
		}
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
