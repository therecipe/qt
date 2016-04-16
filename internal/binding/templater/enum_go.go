package templater

import (
	"fmt"
	"strings"

	"github.com/therecipe/qt/internal/binding/converter"
	"github.com/therecipe/qt/internal/binding/parser"
)

func goEnum(e *parser.Enum) (o string) {
	var t string

	o += fmt.Sprintf("//%v\ntype %v int64\nconst (\n", e.Fullname, strings.Replace(e.Fullname, ":", "_", -1))

	for _, v := range e.Values {
		switch v.Name {
		case "ByteOrder":
			{

			}

		default:
			{
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
	}

	o += ")"

	if e.NoConst || strings.Contains(e.Name, "Style") {
		return strings.Replace(o, "const (", "var (", -1)
	}
	return
}
