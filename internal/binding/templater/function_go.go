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
	var o = fmt.Sprintf("%v{\n%v\n}", goFunctionHeader(function), goFunctionBody(function))
	var c, _ = function.Class()
	if !function.IsSupported() || (c.Stub && function.SignalMode == parser.CALLBACK) {
		return ""
	}
	return o
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

	var class, _ = function.Class()

	if class.Stub {
		//TODO: collect -->
		if converter.GoHeaderOutput(function) == "" {
			return bb.String()
		}

		fmt.Fprintf(bb, "\nreturn %v%v",
			converter.GoOutputParametersFromCFailed(function),
			func() string {
				if !function.Exception {
					return ""
				}
				if strings.Contains(converter.GoHeaderOutput(function), ",") {
					return ", nil"
				}
				return "nil"
			}(),
		)
		return bb.String()
		//<--
	}

	if utils.QT_DEBUG() {
		fmt.Fprintf(bb, "qt.Debug(\"\t%v%v%v(%v) %v\")\n",
			class.Name,
			strings.Repeat(" ", 45-len(class.Name)),
			converter.GoHeaderName(function),
			converter.GoHeaderInput(function),
			converter.GoHeaderOutput(function),
		)
	}

	if !(function.Static || function.Meta == parser.CONSTRUCTOR || function.SignalMode == parser.CALLBACK || strings.Contains(function.Name, "_newList")) {
		fmt.Fprintf(bb, "if ptr.Pointer() != nil {\n")
	}

	if class.Name == "QAndroidJniObject" {

		if strings.HasPrefix(function.Name, parser.TILDE) {
			fmt.Fprint(bb, "qt.DisconnectAllSignals(ptr.Pointer(), \"\")\n")
		}

		for _, parameter := range function.Parameters {
			if parameter.Value == "..." {
				for i := 0; i < 10; i++ {
					fmt.Fprintf(bb, "var p%v, d%v = assertion(%v, v...)\n", i, i, i)
					fmt.Fprintf(bb, "if d%v != nil {\ndefer d%v()\n}\n", i, i)
				}
			} else if parameter.Value == "T" && function.TemplateModeJNI == "" {
				fmt.Fprintf(bb, "var p0, d0 = assertion(0, %v)\n", parameter.Name)
				fmt.Fprint(bb, "if d0 != nil {\ndefer d0()\n}\n")
			}
		}
	}

	for _, alloc := range converter.GoInputParametersForCAlloc(function) {
		fmt.Fprint(bb, alloc)
	}

	if function.SignalMode == "" || (function.Meta == parser.SIGNAL && function.SignalMode != parser.CALLBACK) {

		//TODO: -->
		if function.SignalMode == parser.CONNECT {
			fmt.Fprintf(bb, "\nif !qt.ExistsSignal(ptr.Pointer(), \"%v%v\") {\n",
				function.Name,
				function.OverloadNumber,
			)
		}

		var body = converter.GoOutputParametersFromC(function, fmt.Sprintf("C.%v(%v)", converter.CppHeaderName(function), converter.GoInputParametersForC(function)))
		fmt.Fprint(bb, func() string {
			if function.IsMocFunction && function.SignalMode == "" {
				for _, p := range function.Parameters {
					if p.PureGoType != "" {
						fmt.Fprintf(bb, "qt.RegisterTemp(unsafe.Pointer(%[1]v%[2]v), %[2]v)\n",
							func() string {
								if !strings.HasPrefix(p.PureGoType, "*") {
									return "&"
								}
								return ""
							}(),
							parser.CleanName(p.Name, p.Value))
					}
				}

				if function.PureGoOutput != "" {
					bb := new(bytes.Buffer)
					defer bb.Reset()
					fmt.Fprintf(bb, "oP := unsafe.Pointer(%v)\n", body)
					fmt.Fprint(bb, "if oI, ok := qt.ReceiveTemp(oP); ok {\n")
					fmt.Fprintf(bb, "oD := oI.(%v)\n", function.PureGoOutput)
					if !function.IsMocProperty {
						fmt.Fprint(bb, "qt.UnregisterTemp(oP)\n")
					}
					fmt.Fprint(bb, "return oD\n")
					fmt.Fprint(bb, "}\n")
					return bb.String()
				}
			}

			if converter.GoHeaderOutput(function) == "" {
				return body
			}

			switch {
			case function.NeedsFinalizer && parser.State.ClassMap[parser.CleanValue(function.Output)].IsSupported() || function.Meta == parser.CONSTRUCTOR && !(parser.State.ClassMap[function.Name].HasCallbackFunctions() || parser.State.ClassMap[function.Name].IsSubClassOfQObject()):
				{
					var bb = new(bytes.Buffer)
					defer bb.Reset()

					fmt.Fprintf(bb, "var tmpValue = %v\n", body)

					if class.Name != "QAndroidJniObject" || class.Name == "QAndroidJniObject" && (function.TemplateModeJNI == "String" || function.Output == "QAndroidJniObject" || function.Meta == parser.CONSTRUCTOR) {
						if function.TemplateModeJNI == "String" {
							fmt.Fprint(bb, "tmpValueToString := tmpValue.ToString()\n")
							fmt.Fprint(bb, "tmpValue.DestroyQAndroidJniObject()\n")
						} else {
							class.HasFinalizer = true
							fmt.Fprintf(bb, "runtime.SetFinalizer(tmpValue, (%v).Destroy%v)\n",
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
							)
						}
					}

					fmt.Fprintf(bb, "return tmpValue%v%v",
						func() string {
							if function.TemplateModeJNI == "String" {
								return "ToString"
							}
							return ""
						}(),

						func() string {
							if function.Exception {
								return ", QAndroidJniEnvironment_ExceptionCatch()"
							}
							return ""
						}())

					return bb.String()
				}

			case parser.State.ClassMap[parser.CleanValue(function.Output)].IsSubClassOfQObject() && converter.GoHeaderOutput(function) != "unsafe.Pointer" || function.Meta == parser.CONSTRUCTOR && parser.State.ClassMap[parser.CleanValue(function.Name)].IsSubClassOfQObject():
				{
					var bb = new(bytes.Buffer)
					defer bb.Reset()

					fmt.Fprintf(bb, "var tmpValue = %v\n", body)

					fmt.Fprintf(bb, "if !qt.ExistsSignal(tmpValue.Pointer(), \"destroyed\") {\ntmpValue.ConnectDestroyed(func(%v){ tmpValue.SetPointer(nil) })\n}\n",
						func() string {
							if class.Module == "QtCore" {
								return "*QObject"
							}
							return "*core.QObject"
						}())

					/* TODO: re-implement for custom constructors
					var class, _ = function.Class()
					if class.Module == parser.MOC && function.Meta == parser.CONSTRUCTOR {
						if len(class.Constructors) > 0 {
							fmt.Fprintf(bb, "tmpValue.%v()\n", class.Constructors[0])
						}
					}
					*/

					fmt.Fprint(bb, "return tmpValue")

					return bb.String()
				}

			default:
				{
					if function.Name == "readData" && len(function.Parameters) == 2 {
						return fmt.Sprintf("var ret = %v\nif ret > 0 {\n*data = C.GoStringN(dataC, C.int(ret))\n}\nreturn ret", body)
					} else {
						if function.Exception && converter.GoHeaderOutput(function) == "(error)" {
							return fmt.Sprintf("%v\nreturn QAndroidJniEnvironment_ExceptionCatch()", body)
						}

						return fmt.Sprintf("return %v%v", body,
							func() string {
								if function.Exception {
									return ", QAndroidJniEnvironment_ExceptionCatch()"
								}
								return ""
							}())
					}
				}
			}

		}())

		if function.SignalMode == parser.CONNECT {
			fmt.Fprint(bb, "\n}\n")
		}
	}
	//<--

	switch function.SignalMode {
	case parser.CALLBACK:
		{

			if function.IsMocFunction {
				for _, p := range function.Parameters {
					if p.PureGoType != "" {
						fmt.Fprintf(bb, "var %vD %v\n", parser.CleanName(p.Name, p.Value), p.PureGoType)
						fmt.Fprintf(bb, "if  %[1]vI, ok := qt.ReceiveTemp(unsafe.Pointer(uintptr(%[1]v))); ok {\n", parser.CleanName(p.Name, p.Value))
						if !strings.HasSuffix(function.Name, "Changed") { //TODO: check if property instead
							fmt.Fprintf(bb, "qt.UnregisterTemp(unsafe.Pointer(uintptr(%v)))\n", parser.CleanName(p.Name, p.Value))
						}
						fmt.Fprintf(bb, "%[1]vD = %[1]vI.(%v)\n", parser.CleanName(p.Name, p.Value), p.PureGoType)
						fmt.Fprint(bb, "}\n")
					}
				}
			}

			fmt.Fprintf(bb, "if signal := qt.GetSignal(ptr, \"%v%v\"); signal != nil {\n",
				function.Name,
				function.OverloadNumber,
			)

			if converter.GoHeaderOutput(function) == "" {
				fmt.Fprintf(bb, "signal.(%v)(%v)", converter.GoHeaderInputSignalFunction(function), converter.GoInputParametersForCallback(function))
			} else {
				if function.Name == "readData" && len(function.Parameters) == 2 {
					fmt.Fprint(bb, "var retS = cGoUnpackString(data)\n")
					fmt.Fprintf(bb, "var ret = %v\n", converter.GoInput(fmt.Sprintf("signal.(%v)(%v)", converter.GoHeaderInputSignalFunction(function), converter.GoInputParametersForCallback(function)), function.Output, function))
					fmt.Fprint(bb, "if ret > 0 {\nC.memcpy(unsafe.Pointer(data.data), unsafe.Pointer(C.CString(retS)), C.size_t(ret))\n}\n")
					fmt.Fprint(bb, "return ret")
				} else {
					if function.IsMocFunction && function.PureGoOutput != "" {
						fmt.Fprintf(bb, "oP := %v\n", fmt.Sprintf("signal.(%v)(%v)", converter.GoHeaderInputSignalFunction(function), converter.GoInputParametersForCallback(function)))
						fmt.Fprintf(bb, "qt.RegisterTemp(unsafe.Pointer(%voP), oP)\n", func() string {
							if !strings.HasPrefix(function.PureGoOutput, "*") {
								return "&"
							}
							return ""
						}())
						fmt.Fprintf(bb, "return %v", converter.GoInput(fmt.Sprintf("uintptr(unsafe.Pointer(%voP))", func() string {
							if !strings.HasPrefix(function.PureGoOutput, "*") {
								return "&"
							}
							return ""
						}()), function.Output, function))
					} else {
						if parser.CleanValue(function.Output) == "QString" || parser.CleanValue(function.Output) == "QStringList" {
							fmt.Fprintf(bb, "tempVal := signal.(%v)(%v)\n", converter.GoHeaderInputSignalFunction(function), converter.GoInputParametersForCallback(function))
							fmt.Fprintf(bb, "return C.struct_%v_PackedString{ data: %v, len: %v }", strings.Title(parser.State.ClassMap[function.ClassName()].Module), converter.GoInput("tempVal", function.Output, function),
								func() string {
									if function.AsError {
										return "C.longlong(-1)"
									}
									if parser.CleanValue(function.Output) == "QStringList" {
										return "C.longlong(len(strings.Join(tempVal, \"|\")))"
									}
									return "C.longlong(len(tempVal))"
								}())
						} else {
							fmt.Fprintf(bb, "return %v", converter.GoInput(fmt.Sprintf("signal.(%v)(%v)", converter.GoHeaderInputSignalFunction(function), converter.GoInputParametersForCallback(function)), function.Output, function))
						}
					}
				}
			}

			fmt.Fprintf(bb, "\n}%v\n",
				func() string {
					if converter.GoHeaderOutput(function) == "" {
						if (!function.IsDerivedFromPure() || function.IsDerivedFromImpure() || function.Synthetic) && function.Meta != parser.SIGNAL {
							return "else{"
						}
					}
					return ""
				}(),
			)

			if converter.GoHeaderOutput(function) == "" {
				if (!function.IsDerivedFromPure() || function.IsDerivedFromImpure() || function.Synthetic) && function.Meta != parser.SIGNAL {
					fmt.Fprintf(bb, "New%vFromPointer(ptr).%v%vDefault(%v)", strings.Title(class.Name), strings.Replace(strings.Title(function.Name), parser.TILDE, "Destroy", -1), function.OverloadNumber, converter.GoInputParametersForCallback(function))
				}
			} else {
				if (!function.IsDerivedFromPure() || function.IsDerivedFromImpure() || function.Synthetic) && function.Meta != parser.SIGNAL {
					if function.Name == "readData" && len(function.Parameters) == 2 {
						fmt.Fprint(bb, "var retS = cGoUnpackString(data)\n")
						fmt.Fprintf(bb, "var ret = %v\n", converter.GoInput(fmt.Sprintf("New%vFromPointer(ptr).%v%vDefault(%v)", strings.Title(class.Name), strings.Replace(strings.Title(function.Name), parser.TILDE, "Destroy", -1), function.OverloadNumber, converter.GoInputParametersForCallback(function)), function.Output, function))
						fmt.Fprint(bb, "if ret > 0 {\nC.memcpy(unsafe.Pointer(data.data), unsafe.Pointer(C.CString(retS)), C.size_t(ret))\n}\n")
						fmt.Fprint(bb, "return ret")
					} else {
						if function.IsMocFunction && function.IsMocProperty && function.PureGoOutput != "" {
							fmt.Fprintf(bb, "oP := %v\n", fmt.Sprintf("New%vFromPointer(ptr).%v%vDefault(%v)", strings.Title(class.Name), strings.Replace(strings.Title(function.Name), parser.TILDE, "Destroy", -1), function.OverloadNumber, converter.GoInputParametersForCallback(function)))
							fmt.Fprintf(bb, "qt.RegisterTemp(unsafe.Pointer(%voP), oP)\n", func() string {
								if !strings.HasPrefix(function.PureGoOutput, "*") {
									return "&"
								}
								return ""
							}())
							fmt.Fprintf(bb, "return %v", converter.GoInput(fmt.Sprintf("uintptr(unsafe.Pointer(%voP))", func() string {
								if !strings.HasPrefix(function.PureGoOutput, "*") {
									return "&"
								}
								return ""
							}()), function.Output, function))
						} else {
							if parser.CleanValue(function.Output) == "QString" || parser.CleanValue(function.Output) == "QStringList" {
								fmt.Fprintf(bb, "tempVal := New%vFromPointer(ptr).%v%vDefault(%v)\n", strings.Title(class.Name), strings.Replace(strings.Title(function.Name), parser.TILDE, "Destroy", -1), function.OverloadNumber, converter.GoInputParametersForCallback(function))
								fmt.Fprintf(bb, "return C.struct_%v_PackedString{ data: %v, len: %v }", strings.Title(parser.State.ClassMap[function.ClassName()].Module), converter.GoInput("tempVal", function.Output, function),
									func() string {
										if function.AsError {
											return "C.longlong(-1)"
										}
										if parser.CleanValue(function.Output) == "QStringList" {
											return "C.longlong(len(strings.Join(tempVal, \"|\")))"
										}
										return "C.longlong(len(tempVal))"
									}())
							} else {
								fmt.Fprintf(bb, "\nreturn %v", converter.GoInput(fmt.Sprintf("New%vFromPointer(ptr).%v%vDefault(%v)", strings.Title(class.Name), strings.Replace(strings.Title(function.Name), parser.TILDE, "Destroy", -1), function.OverloadNumber, converter.GoInputParametersForCallback(function)), function.Output, function))
							}
						}
					}
				} else {
					if converter.GoOutputParametersFromCFailed(function) == "nil" {
						var (
							class, _ = function.Class()
							found    bool
							c, ok    = parser.State.ClassMap[parser.CleanValue(function.Output)]
						)
						if ok && c.IsSupported() && !hasUnimplementedPureVirtualFunctions(c.Name) {
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
						if parser.CleanValue(function.Output) == "QString" || parser.CleanValue(function.Output) == "QStringList" {
							fmt.Fprintf(bb, "tempVal := %v\n", converter.GoOutputParametersFromCFailed(function))
							fmt.Fprintf(bb, "return C.struct_%v_PackedString{ data: %v, len: %v }", strings.Title(parser.State.ClassMap[function.ClassName()].Module), converter.GoInput("tempVal", function.Output, function),
								func() string {
									if function.AsError {
										return "C.longlong(-1)"
									}
									if parser.CleanValue(function.Output) == "QStringList" {
										return "C.longlong(len(strings.Join(tempVal, \"|\")))"
									}
									return "C.longlong(len(tempVal))"
								}())
						} else {
							fmt.Fprintf(bb, "\nreturn %v", converter.GoInput(converter.GoOutputParametersFromCFailed(function), function.Output, function))
						}
					}
				}
			}

			fmt.Fprintf(bb, "%v",
				func() string {
					if converter.GoHeaderOutput(function) == "" {
						if (!function.IsDerivedFromPure() || function.IsDerivedFromImpure() || function.Synthetic) && function.Meta != parser.SIGNAL {
							return "\n}"
						}
					}
					return ""
				}(),
			)

		}

	case parser.CONNECT:
		{
			fmt.Fprintf(bb, "\nif signal := qt.LendSignal(ptr.Pointer(), \"%v%v\"); signal != nil {\n",
				function.Name,
				function.OverloadNumber,
			)

			fmt.Fprintf(bb, "\tqt.%vSignal(ptr.Pointer(), \"%v%v\", %v)",
				function.SignalMode,
				function.Name,
				function.OverloadNumber,
				func() string {
					var bb = new(bytes.Buffer)
					defer bb.Reset()

					fmt.Fprintf(bb, "%v {\n", strings.TrimPrefix(converter.GoHeaderInput(function), "f "))

					var f = *function
					f.SignalMode = ""

					fmt.Fprintf(bb, "signal.(%v%v)(%v)\n",
						converter.GoHeaderInputSignalFunction(&f),
						converter.GoHeaderOutput(&f),
						converter.GoGoInput(&f),
					)

					if converter.GoHeaderOutput(&f) != "" {
						fmt.Fprint(bb, "return ")
					}

					fmt.Fprintf(bb, "f(%v)\n",
						converter.GoGoInput(&f),
					)

					fmt.Fprint(bb, "}")

					return bb.String()
				}())

			fmt.Fprintf(bb, "} else {\n")
			fmt.Fprintf(bb, "\tqt.%vSignal(ptr.Pointer(), \"%v%v\"%v)",
				function.SignalMode,
				function.Name,
				function.OverloadNumber,
				func() string {
					if function.SignalMode == parser.CONNECT {
						return ", f"
					}
					return ""
				}(),
			)

			fmt.Fprintf(bb, "}")
		}

	case parser.DISCONNECT:
		{
			fmt.Fprintf(bb, "\nqt.%vSignal(ptr.Pointer(), \"%v%v\"%v)",
				function.SignalMode,
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

	if (function.Name == "deleteLater" || strings.HasPrefix(function.Name, parser.TILDE)) && function.SignalMode == "" {
		fmt.Fprint(bb, "\nptr.SetPointer(nil)")
		if class.HasFinalizer {
			fmt.Fprint(bb, "\nruntime.SetFinalizer(ptr, nil)")
		}
	}

	if !(function.Static || function.Meta == parser.CONSTRUCTOR || function.SignalMode == parser.CALLBACK || strings.Contains(function.Name, "_newList")) {
		fmt.Fprint(bb, "\n}")

		if converter.GoHeaderOutput(function) == "" {
			return bb.String()
		}

		if function.IsMocFunction && function.PureGoOutput != "" && function.SignalMode == "" {
			fmt.Fprintf(bb, "\nvar out %v\nreturn out", function.PureGoOutput)
		} else {
			//TODO: collect -->
			fmt.Fprintf(bb, "\nreturn %v%v",
				converter.GoOutputParametersFromCFailed(function),
				func() string {
					if !function.Exception {
						return ""
					}
					if strings.Contains(converter.GoHeaderOutput(function), ",") {
						return ", errors.New(\"*.Pointer() == nil\")"
					}
					return "errors.New(\"*.Pointer() == nil\")"
				}(),
			)

			return bb.String()
			//<--
		}
	}

	return bb.String()
}
