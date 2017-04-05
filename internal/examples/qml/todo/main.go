//source: https://github.com/benlau/quickflux

//you may need to modify ./quickflux/cgo_stub.go to point to your qmake binary
//go:generate go generate ./quickflux

package main

import (
	"os"

	"github.com/therecipe/qt/core"
	"github.com/therecipe/qt/gui"
	"github.com/therecipe/qt/qml"

	_ "github.com/therecipe/qt/internal/examples/qml/todo/quickflux"
)

type Bridge struct {
	core.QObject

	_ func(action, msg string) `slot:"sendToGo"`
}

func main() {

	gui.NewQGuiApplication(len(os.Args), os.Args)

	var app = qml.NewQQmlApplicationEngine(nil)

	bridge := NewBridge(nil)
	bridge.ConnectSendToGo(func(action, msg string) { println("go:", action, msg) })
	app.RootContext().SetContextProperty("bridge", bridge)

	app.Load(core.NewQUrl3("qrc:///main.qml", 0))

	gui.QGuiApplication_Exec()
}
