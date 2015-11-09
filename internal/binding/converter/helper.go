package converter

import (
	"fmt"
	"strings"

	"github.com/therecipe/qt/internal/binding/parser"
)

func module(input interface{}) string {

	switch input.(type) {
	case *parser.Enum, *parser.Function:
		{
			return module(parser.ClassMap[class(input)].Module)
		}

	case string:
		{
			return strings.ToLower(strings.TrimPrefix(input.(string), "Qt"))
		}
	}
	return ""
}

func class(input interface{}) string {

	switch input.(type) {
	case *parser.Function:
		{
			return class(input.(*parser.Function).Fullname)
		}

	case *parser.Enum:
		{
			return class(input.(*parser.Enum).Fullname)
		}

	case string:
		{
			if strings.Contains(input.(string), "::") {
				return strings.Split(input.(string), "::")[0]
			}
			if strings.Contains(input.(string), "__") {
				return strings.Split(input.(string), "__")[0]
			}
		}
	}

	return ""
}

func cleanValue(value string) string {
	for _, b := range []string{"*", "const", "&"} {
		value = strings.Replace(value, b, "", -1)
	}
	return strings.TrimSpace(value)
}

func cleanName(name string) string {
	switch name {
	case
		"type",
		"func",
		"range",
		"string",
		"int",
		"map",
		"const",
		"interface",
		"select",
		"strings",
		"new":
		{
			return name[:len(name)-2]
		}

	case "":
		{
			return "v"
		}
	}

	return name
}

func isClass(value string) bool {
	if strings.Contains(value, ".") {
		return isClass(strings.Split(value, ".")[1])
	}

	var _, exists = parser.ClassMap[value]
	return exists
}

func isEnum(class, value string) bool {
	if strings.Contains(value, "::") {
		return true
	}

	var outE, _ = findEnum(class, value, false)
	return outE != ""
}

func findEnum(className, value string, byValue bool) (string, string) {

	//Look in given class
	if c, exists := parser.ClassMap[class(value)]; exists {
		for _, e := range c.Enums {
			if outE, outT := findEnumH(e, value, byValue); outE != "" {
				return outE, outT
			}
		}
	}

	//Look in same class
	if c, exists := parser.ClassMap[className]; exists {
		for _, e := range c.Enums {
			if outE, outT := findEnumH(e, value, byValue); outE != "" {
				return outE, outT
			}
		}
	}

	//Look in super classes
	if c, exists := parser.ClassMap[className]; exists {
		for _, s := range c.GetAllBases([]string{}) {
			if sc, exists := parser.ClassMap[s]; exists {
				for _, e := range sc.Enums {
					if outE, outT := findEnumH(e, value, byValue); outE != "" {
						return outE, outT
					}
				}
			}
		}
	}

	//Look in namespaces
	for m := range parser.SubnamespaceMap {
		if c, exists := parser.ClassMap[m]; exists {
			for _, e := range c.Enums {
				if outE, outT := findEnumH(e, value, byValue); outE != "" {
					return outE, outT
				}
			}
		}
	}

	return "", ""
}

func findEnumH(e *parser.Enum, value string, byValue bool) (string, string) {

	if byValue {
		for _, v := range e.Values {
			if outE, _ := findEnumHelper(value, class(e)+"::"+v.Name, ""); outE != "" {
				return outE, ""
			}
		}
	} else {
		if outE, outT := findEnumHelper(value, e.Fullname, e.Typedef); outE != "" {
			return outE, outT
		}
	}

	return "", ""
}

func findEnumHelper(value, name, typedef string) (string, string) {

	var fullName = name

	if strings.Contains(value, "::") {
		value = strings.Split(value, "::")[1]
	}

	if strings.Contains(name, "::") {
		name = strings.Split(name, "::")[1]
	}

	if strings.Contains(typedef, "::") {
		typedef = strings.Split(typedef, "::")[1]
	}

	switch value {
	case name, typedef:
		return fullName, typedef
	}
	return "", ""
}

func goEnum(inter interface{}, value string) string {

	var findByValue bool

	switch inter.(type) {
	case *parser.Enum:
		findByValue = true
	}

	if outE, _ := findEnum(class(inter), value, findByValue); outE != "" {
		return strings.Replace(outE, ":", "_", -1)
	}

	switch inter.(type) {
	case *parser.Function:
		{
			inter.(*parser.Function).Access = "unsupported_goEnum"
		}

	case *parser.Enum:
		{
			inter.(*parser.Enum).Access = "unsupported_goEnum"
		}
	}

	return "unsupported_goEnum"
}

func cppEnum(f *parser.Function, value string, exact bool) string {

	if outE, outT := findEnum(class(f), value, false); outE != "" {
		if exact {

			if outT == "" {
				return outE
			}

			if !strings.Contains(outT, "::") {
				outT = fmt.Sprintf("%v::%v", class(outE), outT)
			}

			return cppEnumExact(value, outE, outT)
		}

		return outE
	}

	f.Access = "unsupported_cppEnum"
	return f.Access
}

func cppEnumExact(value, outE, outT string) string {
	var trimedValue = value

	if strings.Contains(value, "::") {
		trimedValue = strings.Split(value, "::")[1]
	}

	if trimedValue == strings.Split(outT, "::")[1] {
		return outT
	}
	return outE
}
