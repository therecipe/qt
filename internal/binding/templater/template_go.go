package templater

import (
	"fmt"
	"strings"

	"github.com/therecipe/qt/internal/binding/parser"
)

func GoTemplate(c *parser.Class) (o string) {

	if c.Name != "Qt" {

		o += fmt.Sprintf("type %v struct {\n", c.Name)
		if c.Bases == "" {
			o += "ptr unsafe.Pointer"
		} else {
			for _, b := range c.GetBases() {
				if bC, exists := parser.ClassMap[b]; exists {
					if m := shortModule(bC.Module); m != shortModule(c.Module) {
						o += fmt.Sprintf("%v.%v\n", m, b)
					} else {
						o += fmt.Sprintf("%v\n", b)
					}
				}
			}
		}
		o += "}\n\n"

		o += fmt.Sprintf("type %v_ITF interface {\n", c.Name)

		if c.Bases != "" {
			for _, b := range c.GetBases() {
				if bC, exists := parser.ClassMap[b]; exists {
					if m := shortModule(bC.Module); m != shortModule(c.Module) {
						o += fmt.Sprintf("%v.%v_ITF\n", m, b)
					} else {
						o += fmt.Sprintf("%v_ITF\n", b)
					}
				}
			}
		}
		o += fmt.Sprintf("%v_PTR() *%v\n", c.Name, c.Name)
		o += "\n}\n\n"

		if c.Bases == "" {
			o += fmt.Sprintf("func (p *%v)Pointer() unsafe.Pointer {\nreturn p.ptr\n}\n\n", c.Name)
			o += fmt.Sprintf("func (p *%v)SetPointer(ptr unsafe.Pointer) {\np.ptr = ptr\n}\n\n", c.Name)
		}

		if len(c.GetBases()) > 1 {
			o += fmt.Sprintf("func (p *%v)Pointer() unsafe.Pointer {\nreturn p.%v_PTR().Pointer()\n}\n\n", c.Name, c.GetBases()[0])

			o += fmt.Sprintf("func (p *%v)SetPointer(ptr unsafe.Pointer) {\n", c.Name)
			for _, b := range c.GetBases() {
				o += fmt.Sprintf("p.%v_PTR().SetPointer(ptr)\n", b)
			}
			o += "}\n\n"
		}

		o += fmt.Sprintf("func PointerFrom%v(ptr %v_ITF) unsafe.Pointer {\n", c.Name, c.Name)
		o += fmt.Sprintf("if ptr != nil {\nreturn ptr.%v_PTR().Pointer()\n}\nreturn nil\n}\n\n", c.Name)

		o += fmt.Sprintf("func New%vFromPointer(ptr unsafe.Pointer) *%v {\n", c.Name, c.Name)
		o += fmt.Sprintf("var n = new(%v)\n", c.Name)
		o += "n.SetPointer(ptr)\n"

		if isObjectSubClass(c.Name) {
			o += fmt.Sprintf("for len(n.ObjectName()) < len(\"%v_\") {\n", c.Name)
			o += fmt.Sprintf("n.SetObjectName(\"%v_\" + qt.Identifier())\n", c.Name)
			o += "}\n"
		} else if c.Name != "QMetaType" && hasVirtualFunction(c) && isSupportedClass(c) {
			o += fmt.Sprintf("for len(n.ObjectNameAbs()) < len(\"%v_\") {\n", c.Name)
			o += fmt.Sprintf("n.SetObjectNameAbs(\"%v_\" + qt.Identifier())\n", c.Name)
			o += "}\n"
		}

		o += "return n\n}\n\n"

		o += fmt.Sprintf("func (ptr *%v) %v_PTR() *%v {\n", c.Name, c.Name, c.Name)
		o += "return ptr\n}\n\n"

	}

	for _, e := range c.Enums {
		if isSupportedEnum(e) {
			o += fmt.Sprintf("%v\n\n", goEnum(e))
		}
	}

	if isSupportedClass(c) {
		for _, f := range c.Functions {
			if i := goFunction(f); isSupportedFunction(c, f) {
				o += fmt.Sprintf("%v\n\n", i)
			}
		}
	}

	return preambleGo(c.Name, o)
}

func preambleGo(className, input string) string {

	var tmp string

	if parser.ClassMap[className].Stub {
		tmp += "\n// +build !android\n\n"
	}

	tmp += fmt.Sprintf("package %v\n", shortModule(parser.ClassMap[className].Module))

	if strings.Contains(parser.ClassMap[className].Module, "droid") {
		if !parser.ClassMap[className].Stub {
			tmp += fmt.Sprintf("//#include \"%v_android.h\"\n", shortModule(parser.ClassMap[className].Module))
		}
	} else {
		tmp += fmt.Sprintf("//#include \"%v.h\"\n", shortModule(parser.ClassMap[className].Module))
	}

	tmp += "import \"C\"\n"
	tmp += "import (\n"

	for _, i := range append(Libs, "qt", "strings", "unsafe", "log") {
		i = strings.ToLower(i)

		if strings.Contains(input, fmt.Sprintf("%v.", i)) && i != shortModule(parser.ClassMap[className].Module) {
			switch i {
			case "qt":
				tmp += fmt.Sprintf("\"github.com/therecipe/%v\"\n", i)

			case "strings", "unsafe", "log":
				tmp += fmt.Sprintf("\"%v\"\n", i)

			default:
				tmp += fmt.Sprintf("\"github.com/therecipe/qt/%v\"\n", i)
			}
		}

	}

	tmp += ")\n"

	return tmp + input
}

func shortModule(module string) string {
	return strings.ToLower(strings.TrimPrefix(module, "Qt"))
}
