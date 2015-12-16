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
				AddObjectNameFunctionIfNeeded(c)
				ManualWeakLink(c)
				for _, f := range map[string]func(*parser.Class) string{"go": GoTemplate} {
					f(c)
				}
			}
		}
		CppTemplate("Qt" + moduleT)
		HTemplate("Qt" + moduleT)

		utils.MakeFolder(utils.GetQtPkgPath(name))
		utils.Save(utils.GetQtPkgPath(name, fmt.Sprintf("%v%v.cpp", strings.ToLower(moduleT), suffix)), CppTemplate("Qt"+moduleT))
		utils.Save(utils.GetQtPkgPath(name, fmt.Sprintf("%v%v.h", strings.ToLower(moduleT), suffix)), HTemplate("Qt"+moduleT))

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

func AddObjectNameFunctionIfNeeded(c *parser.Class) {
	if !c.IsQObjectSubClass() && c.Name != "QMetaType" && hasVirtualFunction(c) {
		c.Functions = append(c.Functions, &parser.Function{
			Name:     "objectNameAbs",
			Fullname: c.Name + "::" + "objectNameAbs",
			Access:   "public",
			Meta:     "plain",
			Output:   "QString",
		})
		c.Functions = append(c.Functions, &parser.Function{
			Name:     "setObjectNameAbs",
			Fullname: c.Name + "::" + "setObjectNameAbs",
			Access:   "public",
			Meta:     "plain",
			Output:   "void",
			Parameters: []*parser.Parameter{&parser.Parameter{
				Name:  "name",
				Value: "QString",
			}},
		})
	}
}

func ManualWeakLink(c *parser.Class) {
	for _, subclass := range AllClassesUnderThatPkgDir(c.Module) {
		if subclass.WeakLink == nil {
			subclass.WeakLink = make(map[string]bool)
		}

		switch c.Module {
		case "QtCore":
			{
				subclass.WeakLink["QtWidgets"] = true
			}
		case "QtGui":
			{
				subclass.WeakLink["QtWidgets"] = true
			}
		case "QtMultimedia":
			{
				subclass.WeakLink["QtMultimediaWidgets"] = true
			}
		}
	}
}

func AllClassesUnderThatPkgDir(pkgdir string) []*parser.Class {
	var out = make([]*parser.Class, 0)
	for _, c := range parser.ClassMap {
		if c.Module == pkgdir {
			out = append(out, c)
		}
	}
	return out
}
