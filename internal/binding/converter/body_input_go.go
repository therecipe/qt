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
			if parameter.PureGoType != "" {
				input = append(input, GoInput(fmt.Sprintf("uintptr(unsafe.Pointer(%v%v))",
					func() string {
						if !strings.HasPrefix(parameter.PureGoType, "*") {
							return "&"
						}
						return ""
					}(), parser.CleanName(parameter.Name, parameter.Value)), parameter.Value, function))
			} else {
				var alloc = GoInput(parameter.Name, parameter.Value, function)
				if strings.Contains(alloc, "C.CString") {
					if parser.CleanValue(parameter.Value) == "QString" || parser.CleanValue(parameter.Value) == "QStringList" {
						input = append(input, fmt.Sprintf("C.struct_%v_PackedString{ data: %vC, len: %v }", strings.Title(parser.State.ClassMap[function.ClassName()].Module), parser.CleanName(parameter.Name, parameter.Value),
							func() string {
								if function.AsError {
									return "C.longlong(-1)"
								}
								if parser.CleanValue(parameter.Value) == "QStringList" {
									return fmt.Sprintf("C.longlong(len(strings.Join(%v, \"|\")))", parser.CleanName(parameter.Name, parameter.Value))
								}
								return fmt.Sprintf("C.longlong(len(%v))", parser.CleanName(parameter.Name, parameter.Value))
							}()))
					} else {
						input = append(input, fmt.Sprintf("%vC", parser.CleanName(parameter.Name, parameter.Value)))
					}
				} else {
					input = append(input, alloc)
				}
			}
		}
	}

	return strings.Join(input, ", ")
}

func GoInputParametersForCAlloc(function *parser.Function) []string {

	var input = make([]string, 0)

	if function.SignalMode == "" {
		for _, parameter := range function.Parameters {
			var (
				alloc = GoInput(parameter.Name, parameter.Value, function)
				name  = fmt.Sprintf("%vC", parser.CleanName(parameter.Name, parameter.Value))
			)
			switch goType(function, parameter.Value) {
			case "string":
				{
					input = append(input, fmt.Sprintf("var %v *C.char\nif %v != \"\" {\n %v = %v\ndefer C.free(unsafe.Pointer(%v))\n}\n", name, parser.CleanName(parameter.Name, parameter.Value), name, alloc, name))
				}

			case "*string", "[]string", "error":
				{
					input = append(input, fmt.Sprintf("var %v = %v\ndefer C.free(unsafe.Pointer(%v))\n", name, alloc, name))
				}
			}
		}
	}

	return input
}

func GoInputParametersForCallback(function *parser.Function) string {

	var input = make([]string, len(function.Parameters))

	for i, parameter := range function.Parameters {
		if parameter.PureGoType != "" {
			input[i] = fmt.Sprintf("%vD", parser.CleanName(parameter.Name, parameter.Value))
		} else {
			if function.Name == "readData" && strings.HasPrefix(cgoOutput(parameter.Name, parameter.Value, function), "cGoUnpackString") {
				input[i] = "&retS"
			} else {
				input[i] = cgoOutput(parameter.Name, parameter.Value, function)
			}
		}
	}

	return strings.Join(input, ", ")
}
