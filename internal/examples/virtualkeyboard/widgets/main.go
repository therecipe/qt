package main

import (
	"os"

	"github.com/therecipe/qt/widgets"
)

func main() {
	os.Setenv("QT_IM_MODULE", "qtvirtualkeyboard")

	widgets.NewQApplication(0, nil)

	window := widgets.NewQMainWindow(nil, 0)
	window.SetMinimumSize2(250, 200)

	widget := widgets.NewQWidget(nil, 0)
	widget.SetLayout(widgets.NewQVBoxLayout())
	window.SetCentralWidget(widget)

	input := widgets.NewQLineEdit(nil)
	input.SetPlaceholderText("Write something ...")
	widget.Layout().AddWidget(input)

	window.Show()

	widgets.QApplication_Exec()
}
