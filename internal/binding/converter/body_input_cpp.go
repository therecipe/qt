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
			return fmt.Sprintf("%v*", parser.CleanValue(parameter.Value))
		}

	case isEnum(function.ClassName(), parameter.Value):
		{
			if function.Meta == parser.SLOT && function.SignalMode == "" && parser.CleanValue(parameter.Value) == "Qt::Alignment" {
				return parser.CleanValue(parameter.Value)
			}
			return cppEnum(function, parameter.Value, false)
		}

	default:
		{
			return parser.CleanValue(parameter.Value)
		}
	}
}

func CppInputParametersForSignalConnect(function *parser.Function) string {

	var input = make([]string, len(function.Parameters))

	for i, parameter := range function.Parameters {
		if isEnum(function.ClassName(), parameter.Value) {
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
		if isEnum(function.ClassName(), parameter.Value) {
			input[i] = fmt.Sprintf("%v %v", cppEnum(function, parameter.Value, true), parser.CleanName(parameter.Name, parameter.Value))
		} else {
			input[i] = fmt.Sprintf("%v %v", parameter.Value, parser.CleanName(parameter.Name, parameter.Value))
		}
	}

	return strings.Join(input, ", ")
}

func CppInputParametersForCallbackBody(function *parser.Function) string {
	var input = make([]string, len(function.Parameters)+1)

	if strings.Contains(strings.Split(function.Signature, ")")[1], "const") {

		input[0] = fmt.Sprintf("const_cast<%v%v*>(this)",

			func() string {
				if parser.State.Moc {
					return ""
				}
				return "My"
			}(),

			function.ClassName())

	} else {
		input[0] = "this"
	}

	for i, parameter := range function.Parameters {
		input[i+1] = cppOutputPacked(parameter.Name, parameter.Value, function)
	}

	return strings.Join(input, ", ")
}

func CppInputParametersForCallbackBodyPrePack(function *parser.Function) string {
	var input = make([]string, 0)

	for _, parameter := range function.Parameters {
		if packed := cppOutputPack(parameter.Name, parameter.Value, function); packed != "" {
			input = append(input, packed)
		}
	}

	return strings.Join(input, "")
}
