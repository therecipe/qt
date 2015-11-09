package converter

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/therecipe/qt/internal/binding/parser"
)

func GoEnum(n string, v string, e *parser.Enum) string {
	var _, err = strconv.Atoi(v)
	switch {
	case strings.ContainsAny(v, "()<>~+") || v == "":
		{
			e.NoConst = true
			return fmt.Sprintf("C.%v_%v_Type()", strings.Split(e.Fullname, "::")[0], n)
		}

	case strings.Contains(v, "0x"):
		{
			return v
		}

	case err != nil:
		{
			if c, exists := parser.ClassMap[class(goEnum(e, v))]; exists && module(c.Module) != module(e) && module(c.Module) != "" {
				return fmt.Sprintf("%v.%v", module(c.Module), goEnum(e, v))
			}
			return goEnum(e, v)
		}
	}

	return v
}

func CppEnum(v *parser.Value, e *parser.Enum) string {
	if strings.ContainsAny(v.Value, "()<>~+") || v.Value == "" {
		return fmt.Sprintf("%v::%v", class(e), v.Name)
	}

	return v.Value
}
