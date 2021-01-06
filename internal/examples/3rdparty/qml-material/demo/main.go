//source: https://github.com/papyros/qml-material

package main

import (
	"os"

	"github.com/dev-drprasad/qt/core"
	"github.com/dev-drprasad/qt/gui"
	"github.com/dev-drprasad/qt/qml"
	"github.com/dev-drprasad/qt/quickcontrols2"

	_ "github.com/dev-drprasad/qt/internal/examples/3rdparty/qml-material/demo/icons"
)

func main() {

	core.QCoreApplication_SetAttribute(core.Qt__AA_EnableHighDpiScaling, true)

	gui.NewQGuiApplication(len(os.Args), os.Args)

	if quickcontrols2.QQuickStyle_Name() == "" {
		quickcontrols2.QQuickStyle_SetStyle("Material")
	}

	engine := qml.NewQQmlApplicationEngine(nil)
	engine.Load(core.NewQUrl3("qrc:/main.qml", 0))

	gui.QGuiApplication_Exec()
}
