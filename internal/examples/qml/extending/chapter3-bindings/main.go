//source: http://doc.qt.io/qt-5/qtqml-tutorials-extending-qml-example.html

package main

import (
	"os"

	"github.com/therecipe/qt/core"
	"github.com/therecipe/qt/gui"
	"github.com/therecipe/qt/quick"
)

func main() {
	gui.NewQGuiApplication(len(os.Args), os.Args)

	PieChart_QmlRegisterType2("Charts", 1, 0, "PieChart")

	view := quick.NewQQuickView(nil)
	view.SetResizeMode(quick.QQuickView__SizeRootObjectToView)
	view.RootContext().SetContextProperty("factory", NewPieChartFactory(nil))
	view.SetSource(core.NewQUrl3("qrc:///app.qml", 0))
	view.Show()

	gui.QGuiApplication_Exec()
}
