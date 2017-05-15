// +build qml

package main

import (
	"os"
	"path/filepath"

	"github.com/therecipe/qt/core"
	"github.com/therecipe/qt/qml"
	"github.com/therecipe/qt/widgets"
)

var qApp *widgets.QApplication

func main() {
	path := filepath.Join(os.Getenv("GOPATH"), "src", "github.com", "therecipe", "qt", "internal", "examples", "sql", "masterdetail", "qml", "main.qml")

	qApp = widgets.NewQApplication(len(os.Args), os.Args)

	app := initQmlApplication(path)

	app.RootContext().SetContextProperty("listModel", NewListModel())
	app.RootContext().SetContextProperty("viewModel", NewSortFilterModel(NewViewModel()))

	controller := NewController(nil)
	controller.initWith(core.NewQFile2(":///albumdetails.xml"))
	app.RootContext().SetContextProperty("controller", controller)

	app.Load(core.NewQUrl3(path, 0))

	widgets.QApplication_Exec()
}

func initQmlApplication(path string) *qml.QQmlApplicationEngine {

	app := qml.NewQQmlApplicationEngine(nil)

	watcher := core.NewQFileSystemWatcher2([]string{filepath.Dir(path)}, nil)

	reload := func(p string) {
		println("changed:", p)
		app.Load(core.NewQUrl())
		app.ClearComponentCache()
		app.Load(core.NewQUrl3(path, 0))
	}

	//watcher.ConnectFileChanged(reload)
	watcher.ConnectDirectoryChanged(reload)

	return app
}
