//source: https://doc.qt.io/qt-5/qtsensors-accelbubble-example.html

package main

import (
	"os"

	"github.com/dev-drprasad/qt/core"
	"github.com/dev-drprasad/qt/gui"
	"github.com/dev-drprasad/qt/qml"
	_ "github.com/dev-drprasad/qt/sensors"
)

func main() {
	core.QCoreApplication_SetAttribute(core.Qt__AA_EnableHighDpiScaling, true)

	gui.NewQGuiApplication(len(os.Args), os.Args)

	var app = qml.NewQQmlApplicationEngine(nil)
	app.Load(core.NewQUrl3("qrc:///accelbubble.qml", 0))

	gui.QGuiApplication_Exec()
}
