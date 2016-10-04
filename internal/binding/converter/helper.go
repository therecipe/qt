package converter

import (
	"fmt"
	"path/filepath"
	"runtime"
	"strings"

	"github.com/therecipe/qt/internal/binding/parser"
	"github.com/therecipe/qt/internal/utils"
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

func CleanValue(value string) string {
	for _, b := range []string{"*", "const", "&amp", "&", ";"} {
		value = strings.Replace(value, b, "", -1)
	}
	return strings.TrimSpace(value)
}

func cleanName(name, value string) string {
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
		"new",
		"signal",
		"ptr",
		"register":
		{
			return name[:len(name)-2]
		}

	case "":
		{
			return fmt.Sprintf("v%v", strings.Replace(strings.ToLower(CleanValue(value)[:2]), ".", "", -1))
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

	//look in given class
	if c, exists := parser.ClassMap[class(value)]; exists {
		for _, e := range c.Enums {
			if outE, outT := findEnumH(e, value, byValue); outE != "" {
				return outE, outT
			}
		}
	}

	//look in same class
	if c, exists := parser.ClassMap[className]; exists {
		for _, e := range c.Enums {
			if outE, outT := findEnumH(e, value, byValue); outE != "" {
				return outE, outT
			}
		}
	}

	//look in super classes
	if c, exists := parser.ClassMap[className]; exists {
		for _, s := range c.GetAllBases() {
			if sc, exists := parser.ClassMap[s]; exists {
				for _, e := range sc.Enums {
					if outE, outT := findEnumH(e, value, byValue); outE != "" {
						return outE, outT
					}
				}
			}
		}
	}

	//look in namespaces
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
			if outE, _ := findEnumHelper(value, fmt.Sprintf("%v::%v", class(e), v.Name), ""); outE != "" {
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
		{
			return fullName, typedef
		}
	}
	return "", ""
}

func goEnum(inter interface{}, value string) string {

	var findByValue bool

	switch inter.(type) {
	case *parser.Enum:
		{
			findByValue = true
		}
	}

	if outE, _ := findEnum(class(inter), value, findByValue); outE != "" {
		return strings.Replace(outE, ":", "_", -1)
	}

	switch deduced := inter.(type) {
	case *parser.Function:
		{
			deduced.Access = "unsupported_goEnum"
		}

	case *parser.Enum:
		{
			deduced.Access = "unsupported_goEnum"
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

	f.Access = fmt.Sprintf("unsupported_cppEnum(%v)", value)
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

func IsPrivateSignal(f *parser.Function) bool {

	if parser.ClassMap[f.Class()].Module == "QtCore" {

		var (
			fData string
			fPath = strings.Replace(filepath.Base(f.Filepath), ".cpp", ".h", -1)
		)
		fPath = strings.Replace(fPath, ".mm", ".h", -1)

		if strings.HasSuffix(fPath, "_win.h") {
			fPath = strings.Replace(fPath, "_win.h", ".h", -1)
		}

		switch runtime.GOOS {
		case "darwin":
			{
				fData = utils.Load(filepath.Join(utils.QT_DIR(), "5.7", "clang_64", "lib", fmt.Sprintf("%v.framework", strings.Title(parser.ClassMap[f.Class()].DocModule)), "Versions", "5", "Headers", fPath))
			}

		case "windows":
			{
				fData = utils.Load(filepath.Join(utils.QT_DIR(), "5.7", "mingw53_32", "include", strings.Title(parser.ClassMap[f.Class()].DocModule), fPath))
			}

		case "linux":
			{
				fData = utils.Load(filepath.Join(utils.QT_DIR(), "5.7", "gcc_64", "include", strings.Title(parser.ClassMap[f.Class()].DocModule), fPath))
			}
		}

		if fData != "" {
			if strings.Contains(fData, fmt.Sprintf("%v(", f.Name)) {

				return strings.Contains(strings.Split(strings.Split(fData, fmt.Sprintf("%v(", f.Name))[1], ")")[0], "QPrivateSignal")
			}

			if strings.Contains(fData, fmt.Sprintf("%v (", f.Name)) {
				return strings.Contains(strings.Split(strings.Split(fData, fmt.Sprintf("%v (", f.Name))[1], ")")[0], "QPrivateSignal")
			}
		}

		fmt.Println("converter.IsPrivateSignal", f.Class())
	}

	return false
}
