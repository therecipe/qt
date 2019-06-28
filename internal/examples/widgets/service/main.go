package main

import (
	"os"

	"github.com/therecipe/qt/core"
	"github.com/therecipe/qt/gui"
	"github.com/therecipe/qt/widgets"
)

func main() {
	app := widgets.NewQApplication(len(os.Args), os.Args)

	app.ConnectEvent(func(e *core.QEvent) bool {
		if e.Type() == core.QEvent__FileOpen {
			widgets.QMessageBox_Information(nil, "OK", gui.NewQFileOpenEventFromPointer(e.Pointer()).Url().ToString(0), widgets.QMessageBox__Ok, widgets.QMessageBox__Ok)
		}
		return app.EventDefault(e)
	})

	button := widgets.NewQPushButton2("click me!", nil)
	button.ConnectClicked(func(bool) {
		gui.QDesktopServices_OpenUrl(core.NewQUrl3("some://myapplication/mytopic", 0))
	})

	window := widgets.NewQMainWindow(nil, 0)
	window.SetCentralWidget(button)
	window.Show()

	app.Exec()
}
