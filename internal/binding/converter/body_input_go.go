package converter

import (
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
			input = append(input, goInput(parameter.Name, parameter.Value, function))
		}
	}

	return strings.Join(input, ", ")
}

func GoInputParametersForCallback(function *parser.Function) string {

	var input = make([]string, len(function.Parameters))

	for i, parameter := range function.Parameters {
		input[i] = cgoOutput(parameter.Name, parameter.Value, function)
	}

	return strings.Join(input, ", ")
}
