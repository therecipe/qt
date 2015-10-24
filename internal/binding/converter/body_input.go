package converter

import (
	"fmt"
	"path"
	"runtime"
	"strings"

	"github.com/therecipe/qt/internal/binding/parser"
	"github.com/therecipe/qt/internal/utils"
)

func GoBodyInput(f *parser.Function) (o string) {

	if !(f.Static || f.Meta == "constructor") {
		o += "C.QtObjectPtr(ptr.Pointer()), "
	}

	if f.Meta != "signal" {
		for _, p := range f.Parameters {
			o += fmt.Sprintf("%v, ", goInput(p.Name, p.Value, f))
		}
	}

	return strings.TrimSuffix(o, ", ")
}

func GoBodyInputSignalValues(f *parser.Function) (o string) {

	for _, p := range f.Parameters {
		o += fmt.Sprintf("%v, ", cgoOutput(p.Name, p.Value, f))
	}

	return strings.TrimSuffix(o, ", ")
}

func CppBodyInput(f *parser.Function) (o string) {

	if f.Meta == "slot" {
		for _, p := range f.Parameters {
			o += fmt.Sprintf(", Q_ARG(%v, %v)", CppBodyInputSlotValue(f, p), cppInput(p.Name, p.Value, f))
		}
		return
	}

	if f.Meta == "signal" {
		for _, p := range f.Parameters {

			if isEnum(f.Class(), cleanValue(p.Value)) {
				o += fmt.Sprintf("%v, ", cppEnum(f, cleanValue(p.Value), true))
			} else {
				o += fmt.Sprintf("%v, ", p.Value)
			}
		}
		return strings.TrimSuffix(o, ", ")
	}

	for _, p := range f.Parameters {
		o += fmt.Sprintf("%v, ", cppInput(p.Name, p.Value, f))
	}

	return strings.TrimSuffix(o, ", ")
}

func PrivateSignal(f *parser.Function) bool {

	var fData string

	switch runtime.GOOS {
	case "darwin":
		fData = utils.Load(fmt.Sprintf("/usr/local/Qt5.5.1/5.5/clang_64/lib/%v.framework/Versions/5/Headers/%v", strings.Title(parser.ClassMap[f.Class()].Module), path.Base(f.Filepath)))

	case "windows":
		fData = utils.Load(fmt.Sprintf("C:\\Qt\\Qt5.5.1\\5.5\\mingw492_32\\include\\%v\\%v", strings.Title(parser.ClassMap[f.Class()].Module), path.Base(f.Filepath)))

	case "linux":
		switch runtime.GOARCH {
		case "amd64":
			{
				fData = utils.Load(fmt.Sprintf("/usr/local/Qt5.5.1/5.5/gcc_64/include/%v/%v", strings.Title(parser.ClassMap[f.Class()].Module), path.Base(f.Filepath)))
			}
		case "386":
			{
				fData = utils.Load(fmt.Sprintf("/usr/local/Qt5.5.1/5.5/gcc/include/%v/%v", strings.Title(parser.ClassMap[f.Class()].Module), path.Base(f.Filepath)))
			}
		}
	}

	if fData != "" {
		return strings.Contains(strings.Split(strings.Split(fData, f.Name+"(")[1], ")")[0], "QPrivateSignal")
	} else {
		fmt.Println("converter.PrivateSignal", f.Class())
	}

	return false
}

func CppBodyInputCallback(f *parser.Function) (o string) {
	for _, p := range f.Parameters {

		if isEnum(f.Class(), cleanValue(p.Value)) {
			o += fmt.Sprintf("%v %v, ", cppEnum(f, cleanValue(p.Value), true), cleanName(p.Name))
		} else {
			o += fmt.Sprintf("%v %v, ", p.Value, cleanName(p.Name))
		}
	}

	return strings.TrimSuffix(o, ", ")
}

func CppBodyOutputCallback(f *parser.Function) (o string) {

	o += fmt.Sprintf("%v", cppOutput("this->objectName()", "QString", f))

	for _, p := range f.Parameters {
		o += fmt.Sprintf(", %v", cppOutput(p.Name, p.Value, f))
	}
	return
}

func CppBodyInputSlotValue(f *parser.Function, p *parser.Parameter) string {
	if strings.Contains(p.Value, "*") {
		return fmt.Sprintf("%v*", cleanValue(p.Value))
	}

	if isEnum(f.Class(), p.Value) {
		return cppEnum(f, p.Value, false)
	}

	return cleanValue(p.Value)
}
