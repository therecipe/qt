//source: https://doc.qt.io/qt-5/qtwebview-minibrowser-example.html

package main

import (
	"os"

	"github.com/dev-drprasad/qt/core"
	"github.com/dev-drprasad/qt/gui"
	"github.com/dev-drprasad/qt/qml"
	"github.com/dev-drprasad/qt/webview"
)

func main() {
	core.QCoreApplication_SetAttribute(core.Qt__AA_EnableHighDpiScaling, true)

	gui.NewQGuiApplication(len(os.Args), os.Args)
	webview.QtWebView_Initialize()

	var view = qml.NewQQmlApplicationEngine(nil)
	view.Load(core.NewQUrl3("qrc:/main.qml", 0))

	gui.QGuiApplication_Exec()
}
