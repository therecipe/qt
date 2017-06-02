//author: https://github.com/longlongh4

package main

import (
	"fmt"
	"os"
	"time"

	"github.com/therecipe/qt/core"
	"github.com/therecipe/qt/gui"
	"github.com/therecipe/qt/quick"
)

type QmlBridge struct {
	core.QObject

	_ func(data string)        `signal:"sendToQml"`
	_ func(data string) string `slot:"sendToGo"`
}

func main() {

	gui.NewQGuiApplication(len(os.Args), os.Args)

	var view = quick.NewQQuickView(nil)
	view.SetResizeMode(quick.QQuickView__SizeRootObjectToView)

	var qmlBridge = NewQmlBridge(nil)
	qmlBridge.ConnectSendToGo(func(data string) string {
		fmt.Println("go:", data)
		return "hello from go"
	})

	view.RootContext().SetContextProperty("QmlBridge", qmlBridge)
	view.SetSource(core.NewQUrl3("qrc:/qml/bridge2.qml", 0))

	go func() {
		for t := range time.NewTicker(time.Second * 1).C {
			qmlBridge.SendToQml(t.Format(time.ANSIC))
		}
	}()

	view.Show()

	gui.QGuiApplication_Exec()
}
