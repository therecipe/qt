package converter

import (
	"fmt"
	"strings"

	"github.com/therecipe/qt/internal/binding/parser"
)

func CppInputParameters(function *parser.Function) string {

	var input = make([]string, len(function.Parameters))

	for i, parameter := range function.Parameters {
		input[i] = cppInput(parameter.Name, parameter.Value, function)
	}

	return strings.Join(input, ", ")
}

func CppInputParametersForSlotInvoke(function *parser.Function) string {

	var input = make([]string, len(function.Parameters))

	for i, parameter := range function.Parameters {
		input[i] = fmt.Sprintf("Q_ARG(%v, %v)", CppInputParametersForSlotArguments(function, parameter), cppInput(parameter.Name, parameter.Value, function))
	}

	return fmt.Sprintf("%v%v",

		func() string {
			if len(input) > 0 {
				return ", "
			}
			return ""
		}(),

		strings.Join(input, ", "),
	)
}

func CppInputParametersForSlotArguments(function *parser.Function, parameter *parser.Parameter) string {
	switch {
	case strings.Contains(parameter.Value, "*"):
		{
			return fmt.Sprintf("%v*", CleanValue(parameter.Value))
		}

	case isEnum(function.Class(), parameter.Value):
		{
			return cppEnum(function, parameter.Value, false)
		}

	default:
		{
			return CleanValue(parameter.Value)
		}
	}
}

func CppInputParametersForSignalConnect(function *parser.Function) string {

	var input = make([]string, len(function.Parameters))

	for i, parameter := range function.Parameters {
		if isEnum(function.Class(), parameter.Value) {
			input[i] = cppEnum(function, parameter.Value, true)
		} else {
			input[i] = parameter.Value
		}
	}

	return strings.Join(input, ", ")
}

func CppInputParametersForCallbackHeader(function *parser.Function) string {

	var input = make([]string, len(function.Parameters))

	for i, parameter := range function.Parameters {
		if isEnum(function.Class(), parameter.Value) {
			input[i] = fmt.Sprintf("%v %v", cppEnum(function, parameter.Value, true), cleanName(parameter.Name, parameter.Value))
		} else {
			input[i] = fmt.Sprintf("%v %v", parameter.Value, cleanName(parameter.Name, parameter.Value))
		}
	}

	return strings.Join(input, ", ")
}

func CppInputParametersForCallbackBody(function *parser.Function) (o string) {

	var input = make([]string, len(function.Parameters)+2)

	if strings.Contains(strings.Split(function.Signature, ")")[1], "const") {

		input[0] = fmt.Sprintf("const_cast<%v%v*>(this)",

			func() string {
				if parser.ClassMap[function.Class()].Module == parser.MOC {
					return ""
				}
				return "My"
			}(),

			function.Class())

	} else {
		input[0] = "this"
	}

	input[1] = cppOutput(
		fmt.Sprintf("this->objectName%v()",

			func() string {
				if parser.ClassMap[function.Class()].IsQObjectSubClass() {
					return ""
				}
				return "Abs"
			}(),
		),

		"QString", function)

	for i, parameter := range function.Parameters {
		input[i+2] = cppOutput(parameter.Name, parameter.Value, function)
	}

	return strings.Join(input, ", ")
}
