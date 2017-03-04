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

func UseStub(module string, mode int) bool {
	if mode == -1 {
		//TODO: minimal and moc detection
	}
	return utils.QT_STUB() && !parser.State.Minimal && module != parser.MOC && !(module == "QtAndroidExtras" || module == "QtSailfish")
}

func buildTags(module string, stub bool, mode int) string {
	switch {
	case stub:
		{
			if module == "QtAndroidExtras" || module == "androidextras" {
				return "// +build !android"
			}
			return "// +build !sailfish,!sailfish_emulator"
		}

	case mode == MINIMAL:
		{
			return "// +build minimal"
		}

	case mode == MOC:
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
