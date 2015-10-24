package templater

import (
	"fmt"
	"github.com/therecipe/qt/internal/binding/converter"
	"github.com/therecipe/qt/internal/binding/parser"
	"strings"
)

func cppEnum(e *parser.Enum, v *parser.Value) string {
	return fmt.Sprintf("%v{\n\t%v\n}", cppEnumHeader(e, v), cppEnumBody(e, v))
}

func cppEnumHeader(e *parser.Enum, v *parser.Value) string {
	return fmt.Sprintf("int %v_%v_Type()", strings.Split(e.Fullname, "::")[0], v.Name)
}

func cppEnumBody(e *parser.Enum, v *parser.Value) string {
	return fmt.Sprintf("return %v;", converter.CppEnum(v, e))
}
