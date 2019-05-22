package converter

import (
	"fmt"
	"strings"

	"github.com/therecipe/qt/internal/binding/parser"
)

func GoInputParametersForC(function *parser.Function) string {

	var input = make([]string, 0)

	if !(function.Static || function.Meta == parser.CONSTRUCTOR) {
		input = append(input, "ptr.Pointer()")
	}

	if function.SignalMode == "" {
		for _, parameter := range function.Parameters {
			if parameter.PureGoType != "" && !(parameter.Value == "QMap<QString, QVariant>" || parameter.Value == "QList<QVariant>") && !parser.IsBlackListedPureGoType(parameter.PureGoType) {
				input = append(input, GoInput(fmt.Sprintf("uintptr(unsafe.Pointer(%v%v))",
					func() string {
						if !strings.HasPrefix(parameter.PureGoType, "*") {
							return "&"
						}
						return ""
					}(), parser.CleanName(parameter.Name, parameter.Value)), parameter.Value, function, parameter.PureGoType))
			} else {
				var alloc = GoInput(parameter.Name, parameter.Value, function, parameter.PureGoType)
				if strings.Contains(alloc, "C.CString") || strings.Contains(alloc, "qt.GoBoolToInt(*") || strings.Contains(alloc, "*C.char") || strings.Contains(alloc, "C.long(int32(*") {
					if parser.CleanValue(parameter.Value) == "QString" || parser.CleanValue(parameter.Value) == "QStringList" {
						input = append(input, fmt.Sprintf("C.struct_%v_PackedString{data: %vC, len: %v}", strings.Title(parser.State.ClassMap[function.ClassName()].Module), parser.CleanName(parameter.Name, parameter.Value),
							func() string {
								if parser.IsBlackListedPureGoType(parameter.PureGoType) {
									return "C.longlong(-1)"
								}
								if parser.CleanValue(parameter.Value) == "QStringList" {
									return fmt.Sprintf("C.longlong(len(strings.Join(%v, \"¡¦!\")))", parser.CleanName(parameter.Name, parameter.Value))
								}
								return fmt.Sprintf("C.longlong(len(%v))", parser.CleanName(parameter.Name, parameter.Value))
							}()))
					} else {
						if strings.Contains(alloc, "qt.GoBoolToInt(*") || strings.Contains(alloc, "C.long(int32(*") {
							input = append(input, fmt.Sprintf("&%vC", parser.CleanName(parameter.Name, parameter.Value)))
						} else {
							input = append(input, fmt.Sprintf("%vC", parser.CleanName(parameter.Name, parameter.Value)))
						}
					}
				} else {
					input = append(input, alloc)
				}
			}
		}
	}

	return strings.Join(input, ", ")
}

func GoInputParametersForJS(function *parser.Function) string {

	input := make([]string, 0)

	if !(function.Static || function.Meta == parser.CONSTRUCTOR) {
		input = append(input, "uintptr(ptr.Pointer())")
	}

	if function.SignalMode == "" {
		for _, parameter := range function.Parameters {
			if parameter.PureGoType != "" && !(parameter.Value == "QMap<QString, QVariant>" || parameter.Value == "QList<QVariant>") && !parser.IsBlackListedPureGoType(parameter.PureGoType) {
				if parser.UseWasm() {
					if parameter.PureGoType != "" && !(parameter.Value == "QMap<QString, QVariant>" || parameter.Value == "QList<QVariant>") && !parser.IsBlackListedPureGoType(parameter.PureGoType) {
						input = append(input, GoInput(fmt.Sprintf("unsafe.Pointer(%v%v)",
							func() string {
								if !strings.HasPrefix(parameter.PureGoType, "*") {
									return "&"
								}
								return ""
							}(), parser.CleanName(parameter.Name, parameter.Value)), parameter.Value, function, parameter.PureGoType))
					} else {
						input = append(input, GoInputJS(parser.CleanName(parameter.Name, parameter.Value), parameter.Value, function, parameter.PureGoType))
					}
				} else {
					input = append(input, GoInputJS(fmt.Sprintf("%vTID", parser.CleanName(parameter.Name, parameter.Value)), parameter.Value, function, parameter.PureGoType))
				}
			} else {
				alloc := GoInputJS(parameter.Name, parameter.Value, function, parameter.PureGoType)
				if gType := GoType(function, parameter.Value, parameter.PureGoType); (parser.UseWasm() && strings.Contains(alloc, "js.TypedArrayOf(")) || gType == "*bool" || gType == "*int" {
					input = append(input, fmt.Sprintf("%vC", parser.CleanName(parameter.Name, parameter.Value)))
				} else {
					input = append(input, alloc)
				}
			}
		}
	}

	return strings.Join(input, ", ")
}

func GoInputParametersForJSAlloc(function *parser.Function) []string {

	input := make([]string, 0)

	if function.SignalMode == "" {
		for _, parameter := range function.Parameters {
			var (
				alloc = GoInputJS(parameter.Name, parameter.Value, function, parameter.PureGoType)
				name  = fmt.Sprintf("%vC", parser.CleanName(parameter.Name, parameter.Value))
			)
			switch goType(function, parameter.Value, parameter.PureGoType) {
			case "string":
				{
					if !parser.UseWasm() {
						continue
					}
					//TODO: make it possible to pass nil strings; fix this on C side instead
					input = append(input, fmt.Sprintf("var %v js.Value\nif %v != \"\" || true {\n%v = %v\ndefer (*js.TypedArray)(unsafe.Pointer(uintptr(%v.Get(\"data_ptr\").Int()))).Release()\n}\n", name, parser.CleanName(parameter.Name, parameter.Value), name, alloc, name))
				}

			case "[]byte":
				{
					if !parser.UseWasm() {
						continue
					}
					//TODO: make it possible to pass nil []bytes; fix this on C side instead
					input = append(input, fmt.Sprintf("var %v js.Value\nif len(%v) != 0 || true {\n%v = %v\ndefer (*js.TypedArray)(unsafe.Pointer(uintptr(%v.Get(\"data_ptr\").Int()))).Release()\n}\n", name, parser.CleanName(parameter.Name, parameter.Value), name, alloc, name))
				}

			case "*string", "[]string", "error":
				{
					if !parser.UseWasm() {
						continue
					}
					input = append(input, fmt.Sprintf("%v := %v\ndefer (*js.TypedArray)(unsafe.Pointer(uintptr(%v.Get(\"data_ptr\").Int()))).Release()\n", name, alloc, name))
				}

			case "*bool":
				{
					input = append(input, fmt.Sprintf("var %v %v\nif %v != nil {\n%v = qt.WASM.Call(\"_malloc\", 1)\nqt.WASM.Call(\"setValue\", %v, qt.GoBoolToInt(*%v), \"i8\")\ndefer func(){*%v = int8(qt.WASM.Call(\"getValue\", %v, \"i8\").Int()) != 0\nqt.WASM.Call(\"_free\", %v)\n}()\n}\n", name, func() string {
						if parser.UseWasm() {
							return "js.Value"
						}
						return "*js.Object"
					}(), strings.TrimSuffix(name, "C"), name, name, parser.CleanName(parameter.Name, parameter.Value), parser.CleanName(parameter.Name, parameter.Value), name, name))
				}

			case "*int":
				{
					input = append(input, fmt.Sprintf("var %v %v\nif %v != nil {\n%v = qt.WASM.Call(\"_malloc\", 4)\nqt.WASM.Call(\"setValue\", %v, *%v, \"i32\")\ndefer func(){*%v = int(int32(qt.WASM.Call(\"getValue\", %v, \"i32\").Int()))\nqt.WASM.Call(\"_free\", %v)\n}()\n}\n", name,
						func() string {
							if parser.UseWasm() {
								return "js.Value"
							}
							return "*js.Object"
						}(), strings.TrimSuffix(name, "C"), name, name, parser.CleanName(parameter.Name, parameter.Value), parser.CleanName(parameter.Name, parameter.Value), name, name))
				}
			}
		}
	}

	return input
}

func GoInputParametersForCAlloc(function *parser.Function) []string {

	input := make([]string, 0)

	if function.SignalMode == "" {
		for _, parameter := range function.Parameters {
			var (
				alloc = GoInput(parameter.Name, parameter.Value, function, parameter.PureGoType)
				name  = fmt.Sprintf("%vC", parser.CleanName(parameter.Name, parameter.Value))
			)
			switch gType := goType(function, parameter.Value, parameter.PureGoType); gType {
			case "string":
				{
					input = append(input, fmt.Sprintf("var %v *C.char\nif %v != \"\" {\n%v = %v\ndefer C.free(unsafe.Pointer(%v))\n}\n", name, parser.CleanName(parameter.Name, parameter.Value), name, alloc, name))
				}

			case "[]byte":
				{
					input = append(input, fmt.Sprintf("var %v *C.char\nif len(%v) != 0 {\n%v = %v\n}\n", name, parser.CleanName(parameter.Name, parameter.Value), name, alloc))
				}

			case "*string", "[]string", "error":
				{
					input = append(input, fmt.Sprintf("%v := %v\ndefer C.free(unsafe.Pointer(%v))\n", name, alloc, name))
				}

			case "*bool", "*int":
				{
					input = append(input, fmt.Sprintf("var %v %v\nif %v != nil {\n%v = %v\ndefer func(){*%v = %v}()\n}\n", name, strings.TrimPrefix(cgoType(function, parameter.Value), "*"), strings.TrimSuffix(name, "C"), name, alloc, parser.CleanName(parameter.Name, parameter.Value), goOutput(name, parameter.Value, function, parameter.PureGoType)))
				}
			}
		}
	}

	return input
}

func GoInputParametersForCallback(function *parser.Function) string {

	input := make([]string, len(function.Parameters))

	for i, parameter := range function.Parameters {
		if parameter.PureGoType != "" && !(parameter.Value == "QMap<QString, QVariant>" || parameter.Value == "QList<QVariant>") && !parser.IsBlackListedPureGoType(parameter.PureGoType) {
			input[i] = fmt.Sprintf("%vD", parser.CleanName(parameter.Name, parameter.Value))
		} else {
			if function.Name == "readData" && strings.HasPrefix(cgoOutput(parameter.Name, parameter.Value, function, parameter.PureGoType), "cGoUnpack") {
				input[i] = "&retS"
			} else if pType := goType(function, parameter.Value, parameter.PureGoType); strings.Contains(pType, "*bool") || strings.Contains(pType, "*int") {
				if function.SignalMode != parser.CALLBACK {
					input[i] = "nil" //TODO: make *bool and *int usable from pure js
				} else {
					input[i] = fmt.Sprintf("&%vR", parser.CleanName(parameter.Name, parameter.Value))
				}
			} else {
				input[i] = cgoOutput(parameter.Name, parameter.Value, function, parameter.PureGoType)
			}
		}
	}

	return strings.Join(input, ", ")
}
