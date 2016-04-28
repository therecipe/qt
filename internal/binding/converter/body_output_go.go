package converter

import "github.com/therecipe/qt/internal/binding/parser"

func GoOutputParametersFromC(function *parser.Function, name string) string {
	if function.Meta == parser.CONSTRUCTOR {
		return goOutput(name, function.Name, function)
	}
	return goOutput(name, function.Output, function)
}

func GoOutputParametersFromCFailed(function *parser.Function) string {
	if function.Meta == parser.CONSTRUCTOR {
		return goOutputFailed(function.Name, function)
	}
	return goOutputFailed(function.Output, function)
}
