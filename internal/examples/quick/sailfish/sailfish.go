package main

import (
	"os"

	"github.com/StarAurryon/qt/core"
	"github.com/StarAurryon/qt/gui"
	"github.com/StarAurryon/qt/quick"
	"github.com/StarAurryon/qt/sailfish"
)

func main() {
	core.QCoreApplication_SetAttribute(core.Qt__AA_EnableHighDpiScaling, true)

	sailfish.SailfishApp_Application(len(os.Args), os.Args) //gui.NewQGuiApplication(len(os.Args), os.Args)

	var view = sailfish.SailfishApp_CreateView() //quick.NewQQuickView(nil)
	view.SetSource(core.NewQUrl3("qrc:/qml/sailfish.qml", 0))
	view.SetResizeMode(quick.QQuickView__SizeRootObjectToView)
	view.Show()

	gui.QGuiApplication_Exec()
}
