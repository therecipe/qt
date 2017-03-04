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
	if !function.IsSupported() || (UseStub(c.Module, -1) && function.SignalMode == parser.CALLBACK) {
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
			fmt.Fprintf(bb, "\nif !qt.ExistsSignal(fmt.Sprint(ptr.Pointer()), \"%v%v\") {\n",
				function.Name,
				function.OverloadNumber,
			)
		}

		var body = converter.GoOutputParametersFromC(function, fmt.Sprintf("C.%v(%v)", converter.CppHeaderName(function), converter.GoInputParametersForC(function)))
		fmt.Fprint(bb, func() string {
			if converter.GoHeaderOutput(function) == "" {
				return body
			}

			switch {
			case function.NeedsFinalizer && parser.State.ClassMap[parser.CleanValue(function.Output)].IsSupported() || function.Meta == parser.CONSTRUCTOR && !(parser.State.ClassMap[function.Name].HasCallbackFunctions() || parser.State.ClassMap[function.Name].IsSubClassOfQObject()):
				{
					var bb = new(bytes.Buffer)
					defer bb.Reset()

					fmt.Fprintf(bb, "var tmpValue = %v\n", body)

					if class.Name != "QAndroidJniObject" || class.Name == "QAndroidJniObject" && function.TemplateModeJNI == "String" {
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

					fmt.Fprintf(bb, "return tmpValue%v%v",
						func() string {
							if function.TemplateModeJNI == "String" {
								return ".ToString()"
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

					fmt.Fprintf(bb, "if !qt.ExistsSignal(fmt.Sprint(tmpValue.Pointer()), \"QObject::destroyed\") {\ntmpValue.ConnectDestroyed(func(%v){ tmpValue.SetPointer(nil) })\n}\n",
						func() string {
							if class.Module == "QtCore" {
								return "*QObject"
							}
							return "*core.QObject"
						}())

					var class, _ = function.Class()
					if class.Module == parser.MOC && function.Meta == parser.CONSTRUCTOR {
						if len(class.Constructors) > 0 {
							fmt.Fprintf(bb, "tmpValue.%v()\n", class.Constructors[0])
						}
					}

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
			fmt.Fprintf(bb, "if signal := qt.GetSignal(fmt.Sprint(ptr), \"%v%v\"); signal != nil {\n",
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
					fmt.Fprintf(bb, "return %v", converter.GoInput(fmt.Sprintf("signal.(%v)(%v)", converter.GoHeaderInputSignalFunction(function), converter.GoInputParametersForCallback(function)), function.Output, function))
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
						fmt.Fprintf(bb, "\nreturn %v", converter.GoInput(fmt.Sprintf("New%vFromPointer(ptr).%v%vDefault(%v)", strings.Title(class.Name), strings.Replace(strings.Title(function.Name), parser.TILDE, "Destroy", -1), function.OverloadNumber, converter.GoInputParametersForCallback(function)), function.Output, function))
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
						fmt.Fprintf(bb, "\nreturn %v", converter.GoInput(converter.GoOutputParametersFromCFailed(function), function.Output, function))
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
			fmt.Fprintf(bb, "\nif signal := qt.LendSignal(fmt.Sprint(ptr.Pointer()), \"%v%v\"); signal != nil {\n",
				function.Name,
				function.OverloadNumber,
			)

			fmt.Fprintf(bb, "\tqt.%vSignal(fmt.Sprint(ptr.Pointer()), \"%v%v\", %v)",
				function.SignalMode,
				function.Name,
				function.OverloadNumber,
				func() string {
					var bb = new(bytes.Buffer)
					defer bb.Reset()

					fmt.Fprintf(bb, "%v {\n", strings.TrimPrefix(converter.GoHeaderInput(function), "f "))

					fmt.Fprintf(bb, "signal.(%v)(%v)\n",
						converter.GoHeaderInputSignalFunction(function),
						converter.GoGoInput(function),
					)

					var f = *function
					f.SignalMode = ""
					if converter.GoHeaderOutput(&f) != "" {
						fmt.Fprint(bb, "return ")
					}

					fmt.Fprintf(bb, "f(%v)\n",
						converter.GoGoInput(function),
					)

					fmt.Fprint(bb, "}")

					return bb.String()
				}())

			fmt.Fprintf(bb, "} else {\n")
			fmt.Fprintf(bb, "\tqt.%vSignal(fmt.Sprint(ptr.Pointer()), \"%v%v\"%v)",
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
			fmt.Fprintf(bb, "\nqt.%vSignal(fmt.Sprint(ptr.Pointer()), \"%v%v\"%v)",
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

	if (function.Meta == parser.DESTRUCTOR || function.Name == "deleteLater" || function.Name == "destroyed" || strings.HasPrefix(function.Name, parser.TILDE)) && function.SignalMode == "" {
		if class.HasCallbackFunctions() || class.IsSubClassOfQObject() {
			fmt.Fprint(bb, "\nqt.DisconnectAllSignals(fmt.Sprint(ptr.Pointer()))")
		}
		fmt.Fprint(bb, "\nptr.SetPointer(nil)")
	}

	if !(function.Static || function.Meta == parser.CONSTRUCTOR || function.SignalMode == parser.CALLBACK || strings.Contains(function.Name, "_newList")) {
		fmt.Fprint(bb, "\n}")

		if converter.GoHeaderOutput(function) == "" {
			return bb.String()
		}

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

	return bb.String()
}
