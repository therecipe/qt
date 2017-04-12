package converter

import (
	"fmt"

	"github.com/therecipe/qt/internal/binding/parser"
)

func CppOutputParameters(function *parser.Function, name string) string {
	if function.Meta == parser.CONSTRUCTOR {
		return CppOutput(name, function.Name, function)
	}
	return CppOutput(name, function.Output, function)
}

func CppOutputParametersFailed(function *parser.Function) string {
	var output = GoOutputParametersFromCFailed(function)
	if output == "nil" {
		output = "NULL"
	}
	return output
}

func CppOutputParametersDeducedFromGeneric(function *parser.Function) string {

	if function.TemplateModeGo != "" {
		return fmt.Sprintf("<%v>", function.TemplateModeGo)
	}

	switch function.TemplateModeJNI {
	case "Int":
		{
			return "<jint>"
		}

	case "Boolean":
		{
			return "<jboolean>"
		}

	case "Void":
		{
			return "<void>"
		}

	case "Object", "String":
		{
			if function.Name == "callObjectMethod" || function.Name == "callStaticObjectMethod" {
				if function.OverloadNumber == "2" || function.OverloadNumber == "4" {
					return ""
				}
			}

			return "<jobject>"
		}
	}

	return ""
}

func CppOutputParametersJNIGenericModes(function *parser.Function) []string {

	switch function.Name {
	case "callMethod", "callStaticMethod":
		{
			return []string{"Int", "Boolean", "Void"} //TODO: more primitives
		}

	case "getField", "setField", "getStaticField", "setStaticField":
		{
			return []string{"Int", "Boolean"} //TODO: more primitives
		}

	case "getObjectField", "getStaticObjectField", "callObjectMethod", "callStaticObjectMethod":
		{
			return []string{"Object", "String"} //TODO: add []string, []int, []object, ...
		}
	}

	return make([]string, 0)
}
