package templater

import (
	"strings"

	"github.com/therecipe/qt/internal/binding/parser"
	"github.com/therecipe/qt/internal/utils"
)

func GenModule(name string) {
	if ShouldBuild(name) {

		var (
			pkgName = strings.ToLower(name)
			suffix  string
		)

		if name == "AndroidExtras" {
			suffix = "_android"
		}

		//cleanup
		utils.RemoveAll(utils.GetQtPkgPath(pkgName))
		utils.RemoveAll(utils.GetQtPkgPath("internal", "binding", "dump", name))

		//prepare
		utils.MakeFolder(utils.GetQtPkgPath(pkgName))
		CopyCgo(name)
		if name == "AndroidExtras" {
			utils.Save(utils.GetQtPkgPath(pkgName, "androidextras_android.go"), utils.Load(utils.GetQtPkgPath("internal", "binding", "files", "androidextras_android.go")))
		}

		//dry run
		for _, c := range parser.ClassMap {
			if strings.TrimPrefix(c.Module, "Qt") == name {
				AddObjectNameFunctionIfNeeded(c)
				ManualWeakLink(c)
				GoTemplate(c)
			}
		}
		CppTemplate("Qt" + name)
		HTemplate("Qt" + name)

		//generate
		utils.Save(utils.GetQtPkgPath(pkgName, pkgName+suffix+".cpp"), CppTemplate("Qt"+name))
		utils.Save(utils.GetQtPkgPath(pkgName, pkgName+suffix+".h"), HTemplate("Qt"+name))

		for _, c := range parser.ClassMap {
			if strings.TrimPrefix(c.Module, "Qt") == name {

				if name == "AndroidExtras" {
					utils.Save(utils.GetQtPkgPath(pkgName, strings.ToLower(c.Name)+suffix+".go"), GoTemplate(c))
				}

				c.Stub = (name == "AndroidExtras")
				utils.Save(utils.GetQtPkgPath(pkgName, strings.ToLower(c.Name)+".go"), GoTemplate(c))
				c.Stub = false

				//c.Dump()
			}
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
				subclass.WeakLink["QtMultimedia"] = true
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
