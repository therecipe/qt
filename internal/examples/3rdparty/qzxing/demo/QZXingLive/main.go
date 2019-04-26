//source: https://github.com/ftylitak/qzxing/tree/master/examples/QZXingLive

package main

import (
	"os"

	"github.com/therecipe/qt/core"
	"github.com/therecipe/qt/gui"

	"github.com/therecipe/qt/internal/examples/3rdparty/qzxing"
)

func main() {

	// enable high dpi scaling
	// useful for devices with high pixel density displays
	// such as smartphones, retina displays, ...
	core.QCoreApplication_SetAttribute(core.Qt__AA_EnableHighDpiScaling, true)

	// needs to be called once before you can start using QML
	gui.NewQGuiApplication(len(os.Args), os.Args)

	qzxing.RegisterQMLTypes()

	customApp := NewApplication(nil)
	customApp.checkPermissions()

	// start the main Qt event loop
	// and block until app.Exit() is called
	// or the window is closed by the user
	gui.QGuiApplication_Exec()
}
