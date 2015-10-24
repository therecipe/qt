package templater

import (
	"fmt"
	"github.com/therecipe/qt/internal/binding/converter"
	"github.com/therecipe/qt/internal/binding/parser"
	"strings"
)

func goEnum(e *parser.Enum) (o string) {
	var t string

	o += fmt.Sprintf("//%v\ntype %v int\nvar (\n", e.Fullname, strings.Replace(e.Fullname, ":", "_", -1))

	for _, v := range e.Values {
		switch v.Name {
		case "ByteOrder":

		default:
			if strings.Contains(v.Value, " | ") {
				var tArray = make([]string, 0)
				for _, s := range strings.Split(v.Value, " | ") {
					tArray = append(tArray, converter.GoEnum(v.Name, s, e))
				}
				t = strings.Join(tArray, " | ")
			} else {
				t = converter.GoEnum(v.Name, v.Value, e)
			}
			o += fmt.Sprintf("%v__%v = %v(%v)\n", strings.Split(e.Fullname, "::")[0], v.Name, strings.Replace(e.Fullname, ":", "_", -1), t)
		}
	}

	o += ")"

	return
}
