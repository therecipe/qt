package main

import (
	"fmt"
	"os"
	"time"

	"github.com/therecipe/qt/core"
	"github.com/therecipe/qt/gui"
	"github.com/therecipe/qt/widgets"
)

type QmlBridge struct {
	core.QObject

	_ func(data string)        `signal:sendToQml`
	_ func(data string) string `slot:sendToGo` //only slots can return something
}

func main() {
	var qmlBridge *QmlBridge

	widgets.NewQApplication(len(os.Args), os.Args)

	//create a label, which is also later showed (instead of a QMainWindow)
	var label = widgets.NewQLabel(nil, 0)
	label.SetMinimumSize2(320, 240)
	label.SetStyleSheet("QLabel { background-color: black; color: white; font-size: 16px }")
	label.SetAlignment(core.Qt__AlignCenter)

	//used this because QLabel got no clicked signal
	label.ConnectMousePressEvent(func(ev *gui.QMouseEvent) {
		//in main thread

		println(qmlBridge.SendToGo("hello from qml"))
	})

	//create a instance of QmlBridge and connect the slot and signal
	qmlBridge = NewQmlBridge(nil)
	qmlBridge.ConnectSendToGo(func(data string) string {
		//in main thread

		fmt.Println("go:", data)
		return "hello from go"
	})
	qmlBridge.ConnectSendToQml(func(data string) {
		//in main thread

		label.SetText(data)
	})

	//timer in another goroutine (and thread) that emits the signal to update the label
	//you need to use slots or signals if you want to change visual Qt elements from another thread
	go func() {
		//some other thread

		for t := range time.NewTicker(time.Second * 1).C {
			qmlBridge.SendToQml(t.Format(time.ANSIC))
		}
	}()

	//show label (which will act as a window)
	label.Show()

	widgets.QApplication_Exec()
}
