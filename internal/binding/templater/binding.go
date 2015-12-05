package templater

import (
	"fmt"
	"strings"

	"github.com/therecipe/qt/internal/binding/parser"
	"github.com/therecipe/qt/internal/utils"
)

func terminalSeperator(module string) string {
	var sep = ""
	if len(module) < 8 {
		sep += "\t\t"
	}
	if len(module) >= 8 && len(module) < 17 {
		sep += "\t"
	}
	return sep
}

func GenModule(name string) {
	var moduleT = name

	name = strings.ToLower(name)

	parser.GetModule(name)

	if ShouldBuild(moduleT) {

		var suffix string
		if strings.TrimPrefix(name, "Qt") == "androidextras" {
			suffix = "_android"
		}

		utils.RemoveAll(utils.GetQtPkgPath(name))
		utils.RemoveAll(utils.GetQtPkgPath("internal", "binding", "dump", moduleT))

		//Dry Run
		for _, c := range parser.ClassMap {
			if moduleT == strings.TrimPrefix(c.Module, "Qt") {
				for _, f := range map[string]func(*parser.Class) string{"go": GoTemplate} {
					f(c)
				}
			}
		}
		CppTemplateAIO("Qt" + moduleT)
		HTemplateAIO("Qt" + moduleT)

		utils.MakeFolder(utils.GetQtPkgPath(name))
		utils.Save(utils.GetQtPkgPath(name, fmt.Sprintf("%v%v.cpp", strings.ToLower(moduleT), suffix)), CppTemplateAIO("Qt"+moduleT))
		utils.Save(utils.GetQtPkgPath(name, fmt.Sprintf("%v%v.h", strings.ToLower(moduleT), suffix)), HTemplateAIO("Qt"+moduleT))

		var fcount = 0
		for _, c := range parser.ClassMap {

			if moduleT == strings.TrimPrefix(c.Module, "Qt") {

				for e, f := range map[string]func(*parser.Class) string{"go": GoTemplate} {
					CopyCgo(moduleT)

					if moduleT == "AndroidExtras" {
						utils.Save(utils.GetQtPkgPath(name, fmt.Sprintf("%v_android.%v", strings.ToLower(c.Name), e)), f(c))

						c.Stub = true
						utils.Save(utils.GetQtPkgPath(name, fmt.Sprintf("%v.%v", strings.ToLower(c.Name), e)), f(c))
						c.Stub = false
					} else {
						utils.Save(utils.GetQtPkgPath(name, fmt.Sprintf("%v.%v", strings.ToLower(c.Name), e)), f(c))
					}

				}

				//c.Dump()

				fcount += len(c.Functions)
			}

		}

		fmt.Printf("%v%v\tfunctions:%v", name, terminalSeperator(name), fcount)

		if name == "androidextras" {
			utils.Save(utils.GetQtPkgPath(name, "androidextras_android.go"), `package androidextras

			import (
				"C"
				"strings"
				"unsafe"

				"github.com/therecipe/qt"
			)

			func assertion(key int, input ...interface{}) unsafe.Pointer {
				if len(input) > key {
					switch input[key].(type) {
					case string:
						{
							return QAndroidJniObject_FromString(input[key].(string)).Object()
						}
					case []string:
						{
							return QAndroidJniObject_FromString(strings.Join(input[key].([]string), ",,,")).CallObjectMethod2("split", "(Ljava/lang/String;)[Ljava/lang/String;", ",,,").Object()
						}
					case int:
						{
							return unsafe.Pointer(uintptr(C.int(input[key].(int))))
						}
					case bool:
						{
							return unsafe.Pointer(uintptr(C.int(qt.GoBoolToInt(input[key].(bool)))))
						}
					case unsafe.Pointer:
						{
							return input[key].(unsafe.Pointer)
						}
					case *QAndroidJniObject:
						{
							return input[key].(*QAndroidJniObject).Object()
						}
					}
				}
				return nil
			}
`)
		}

	}
}
