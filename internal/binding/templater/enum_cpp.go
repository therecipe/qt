package templater

import (
	"fmt"

	"github.com/therecipe/qt/internal/binding/parser"
)

func cppEnum(enum *parser.Enum, value *parser.Value) string {
	return fmt.Sprintf("%v\n{\n\t%v\n}", cppEnumHeader(enum, value), cppEnumBody(enum, value))
}

func cppEnumHeader(enum *parser.Enum, value *parser.Value) string {
	return fmt.Sprintf("int %v_%v_Type()", enum.Class(), value.Name)
}

func cppEnumBody(enum *parser.Enum, value *parser.Value) string {
	return fmt.Sprintf("return %v::%v;", enum.Class(), value.Name)
}
