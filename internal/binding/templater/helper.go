package templater

import (
	"sort"
	"strings"

	"github.com/therecipe/qt/internal/binding/parser"
	"github.com/therecipe/qt/internal/utils"
)

func hasUnimplementedPureVirtualFunctions(className string) bool {
	for _, f := range parser.State.ClassMap[className].Functions {
		var f = *f
		cppFunction(&f)

		if f.Virtual == parser.PURE && (!(&f).IsSupported() || parser.IsPackedList(f.Output)) {
			return true
		}
	}
	return false
}

func goModule(module string) string {
	return strings.ToLower(strings.TrimPrefix(module, "Qt"))
}

func sortedClassNamesForModule(module string) []string {
	var output = make([]string, 0)
	for _, class := range parser.State.ClassMap {
		if class.Module == module {
			output = append(output, class.Name)
		}
	}
	sort.Stable(sort.StringSlice(output))

	if parser.State.Moc {
		var items = make(map[string]string)

		for _, cn := range output {
			if class, exist := parser.State.ClassMap[cn]; exist {
				items[cn] = class.Bases
			}
		}

		var tmpOutput = make([]string, 0)

		for len(items) > 0 {
			for item, dep := range items {

				var c, exist = parser.State.ClassMap[dep]
				if exist && c.Module != parser.MOC {
					tmpOutput = append(tmpOutput, item)
					delete(items, item)
					continue
				}

				for _, key := range tmpOutput {
					if key == dep {
						tmpOutput = append(tmpOutput, item)
						delete(items, item)
						break
					}
				}

			}
		}
		output = tmpOutput
	}

	return output
}

func sortedClassesForModule(module string) []*parser.Class {
	var (
		classNames = sortedClassNamesForModule(module)
		output     = make([]*parser.Class, len(classNames))
	)
	for i, name := range classNames {
		output[i] = parser.State.ClassMap[name]
	}
	return output
}

func UseStub() bool {
	return utils.QT_STUB() && !parser.State.Minimal && !parser.State.Moc && !(parser.State.Module == "AndroidExtras" || parser.State.Module == "Sailfish")
}

func buildTags(module string) string {
	switch {
	case parser.State.Minimal:
		{
			return "// +build minimal"
		}

	case parser.State.Moc:
		{
			return ""
		}

	case module == "QtAndroidExtras":
		{
			return "// +build android"
		}

	case module == "QtSailfish":
		{
			return "// +build sailfish sailfish_emulator"
		}

	default:
		{
			return "// +build !minimal"
		}
	}
}
