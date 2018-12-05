package main

import (
	"os"

	"github.com/therecipe/qt/widgets"
)

func main() {
	widgets.NewQApplication(len(os.Args), os.Args)

	button := widgets.NewQPushButton2("check for updates", nil)
	button.ConnectClicked(func(bool) { sparkle_checkUpdates() })
	button.Show()

	widgets.QApplication_Exec()
}
