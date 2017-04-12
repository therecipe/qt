package subsub

import "github.com/therecipe/qt/core"

type SubSubTestStruct struct {
	core.QObject //TODO: gui.QWindow

	_ string `property:"subsubProperty"`
}
