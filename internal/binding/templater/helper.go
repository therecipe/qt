package templater

import (
	"strings"

	"github.com/therecipe/qt/internal/binding/parser"
	"github.com/therecipe/qt/internal/utils"
)

func hasUnimplementedPureVirtualFunctions(className string) bool {
	for _, f := range parser.State.ClassMap[className].Functions {

		if !f.Checked {
			cppFunction(f)
			goFunction(f)
			f.Checked = true
		}

		if f.Virtual == parser.PURE && !f.IsSupported() {
			return true
		}
	}
	return false
}

func goModule(module string) string {
	return strings.ToLower(strings.TrimPrefix(module, "Qt"))
}

func UseStub() bool {
	return utils.QT_STUB() && !parser.State.Minimal && !parser.State.Moc && !(parser.State.Module == "AndroidExtras" || parser.State.Module == "Sailfish")
}

func buildTags(module string, stub bool) string {
	switch {
	case stub:
		{
			if module == "QtAndroidExtras" || module == "androidextras" {
				return "// +build !android"
			}
			return "// +build !sailfish,!sailfish_emulator"
		}

	case parser.State.Minimal:
		{
			return "// +build minimal"
		}

	case parser.State.Moc:
		{
			return ""
		}

	case module == "QtAndroidExtras", module == "androidextras":
		{
			return "// +build android"
		}

	case module == "QtSailfish", module == "sailfish":
		{
			return "// +build sailfish sailfish_emulator"
		}

	default:
		{
			return "// +build !minimal"
		}
	}
}
