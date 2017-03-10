package converter

import (
	"fmt"
	"os/exec"
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
			return module(parser.State.ClassMap[class(input)].Module)
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

func isClass(value string) bool {
	if strings.Contains(value, ".") {
		return isClass(strings.Split(value, ".")[1])
	}

	var _, ok = parser.State.ClassMap[value]
	return ok
}

func isEnum(class, value string) bool {
	if strings.ContainsAny(value, "<>") {
		return false
	}
	if strings.Contains(value, "::") {
		return true
	}

	var outE, _ = findEnum(class, value, false)
	return outE != ""
}

func findEnum(className, value string, byValue bool) (string, string) {

	//look in given class
	if c, ok := parser.State.ClassMap[class(value)]; ok {
		for _, e := range c.Enums {
			if outE, outT := findEnumH(e, value, byValue); outE != "" {
				return outE, outT
			}
		}
	}

	//look in same class
	if c, ok := parser.State.ClassMap[className]; ok {
		for _, e := range c.Enums {
			if outE, outT := findEnumH(e, value, byValue); outE != "" {
				return outE, outT
			}
		}
	}

	//look in super classes
	if c, ok := parser.State.ClassMap[className]; ok {
		for _, s := range c.GetAllBases() {
			if sc, ok := parser.State.ClassMap[s]; ok {
				for _, e := range sc.Enums {
					if outE, outT := findEnumH(e, value, byValue); outE != "" {
						return outE, outT
					}
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
	var fc, ok = f.Class()
	if !ok {
		return false
	}

	if fc.Module == "QtCore" {

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
				if utils.QT_HOMEBREW() {
					fData = utils.LoadOptional(filepath.Join(utils.QT_DARWIN_DIR(), "lib", fmt.Sprintf("%v.framework", strings.Title(parser.State.ClassMap[f.ClassName()].DocModule)), "Versions", "5", "Headers", fPath))
				} else {
					fData = utils.Load(filepath.Join(utils.QT_DARWIN_DIR(), "lib", fmt.Sprintf("%v.framework", strings.Title(parser.State.ClassMap[f.ClassName()].DocModule)), "Versions", "5", "Headers", fPath))
				}
			}

		case "windows":
			{
				if utils.QT_MSYS2() {
					fData = utils.LoadOptional(filepath.Join(utils.QT_MSYS2_DIR(), "include", strings.Title(parser.State.ClassMap[f.ClassName()].DocModule), fPath))
				} else {
					fData = utils.Load(filepath.Join(utils.QT_DIR(), utils.QT_VERSION_MAJOR(), "mingw53_32", "include", strings.Title(parser.State.ClassMap[f.ClassName()].DocModule), fPath))
				}
			}

		case "linux":
			{
				if utils.QT_PKG_CONFIG() {
					fData = utils.LoadOptional(filepath.Join(strings.TrimSpace(utils.RunCmd(exec.Command("pkg-config", "--variable=includedir", "Qt5Core"), "convert.IsPrivateSignal_includeDir")), strings.Title(parser.State.ClassMap[f.ClassName()].DocModule), fPath))
				} else {
					fData = utils.Load(filepath.Join(utils.QT_DIR(), utils.QT_VERSION_MAJOR(), "gcc_64", "include", strings.Title(parser.State.ClassMap[f.ClassName()].DocModule), fPath))
				}
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

		utils.Log.Debugln("converter.IsPrivateSignal", f.ClassName())
	}

	return false
}
