package main

import (
	"fmt"
	"os"

	"github.com/StarAurryon/qt/core"
	"github.com/StarAurryon/qt/qml"
	"github.com/StarAurryon/qt/quick"
	"github.com/StarAurryon/qt/widgets"
)

//this example demonstrates one possible way to deal with qml errors/warnings
//therefore the qml file does contain errors and will NOT work properly

func main() {
	core.QCoreApplication_SetAttribute(core.Qt__AA_EnableHighDpiScaling, true)

	widgets.NewQApplication(len(os.Args), os.Args)

	view := quick.NewQQuickView(nil)
	view.Engine().ConnectWarnings(func(warnings []*qml.QQmlError) {
		for _, w := range warnings {
			fmt.Println("warning:", w.ToString())
		}
		os.Exit(1)
	})

	view.SetResizeMode(quick.QQuickView__SizeViewToRootObject)
	view.SetSource(core.NewQUrl3("qrc:/qml/view.qml", 0))
	view.Show()

	if errors := view.Errors(); len(errors) != 0 {
		for _, e := range errors {
			fmt.Println("error:", e.ToString())
		}
		os.Exit(1)
	}

	widgets.QApplication_Exec()
}
