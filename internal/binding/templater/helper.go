package templater

import (
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
	return parser.SortedClassNamesForModule(module)
}

func sortedClassesForModule(module string) []*parser.Class {
	return parser.SortedClassesForModule(module)
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
