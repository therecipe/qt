package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"os"
	"time"

	"github.com/therecipe/qt/core"
	"github.com/therecipe/qt/quick"
	"github.com/therecipe/qt/widgets"
)

var (
	qmlObjects = make(map[string]*core.QObject)

	qmlBridge          *QmlBridge
	manipulatedFromQml *widgets.QWidget

	colors = []string{"red", "green", "blue", "cyan", "magenta", "yellow", "gray"}
)

func main() {

	widgets.NewQApplication(len(os.Args), os.Args)

	var layout = widgets.NewQHBoxLayout()
	layout.AddWidget(newCppWidget(), 0, 0)
	layout.AddWidget(newSeperator(), 0, 0)
	layout.AddWidget(newQmlWidget(), 0, 0)

	var window = widgets.NewQMainWindow(nil, 0)

	var centralWidget = widgets.NewQWidget(window, 0)
	centralWidget.SetLayout(layout)
	window.SetCentralWidget(centralWidget)

	window.Show()

	widgets.QApplication_Exec()
}

func newCppWidget() *widgets.QWidget {

	var button = widgets.NewQPushButton2("Call Qml Function", nil)
	button.ConnectClicked(func(_ bool) {
		rand.Seed(time.Now().UnixNano())
		qmlBridge.SendToQml("GoButton", "click", colors[rand.Intn(len(colors))])
	})

	manipulatedFromQml = widgets.NewQWidget(nil, 0)

	var layout = widgets.NewQVBoxLayout()
	layout.AddWidget(button, 0, 0)
	layout.AddWidget(manipulatedFromQml, 0, 0)

	var widget = widgets.NewQWidget(nil, 0)
	widget.SetLayout(layout)

	return widget
}

func newSeperator() *widgets.QFrame {
	var frame = widgets.NewQFrame(nil, 0)
	frame.SetFrameShape(widgets.QFrame__VLine)
	return frame
}

func newQmlWidget() *quick.QQuickWidget {
	var quickWidget = quick.NewQQuickWidget(nil)
	quickWidget.SetResizeMode(quick.QQuickWidget__SizeRootObjectToView)

	initQmlContext(quickWidget)
	initQmlBridge(quickWidget)

	quickWidget.SetSource(core.NewQUrl3("qrc:/qml/bridge.qml", 0))

	return quickWidget
}

func initQmlContext(quickWidget *quick.QQuickWidget) {

	var m = map[string]map[string]string{
		"QmlButton": {
			"color":        "lightGray",
			"pressedColor": "darkGray",
			"text":         "Call Go Function",
		},
	}

	var b, err = json.Marshal(m)
	if err != nil {
		log.Println("initQmlContext", err)
	}
	quickWidget.RootContext().SetContextProperty2("qmlInitContext", core.NewQVariant14(string(b)))
}

type QmlBridge struct {
	core.QObject

	_ func(source, action, data string) `signal:"sendToQml"`
	_ func(source, action, data string) `slot:"sendToGo"`

	_ func(object *core.QObject) `slot:"registerToGo"`
	_ func(objectName string)    `slot:"deregisterToGo"`
}

func initQmlBridge(quickWidget *quick.QQuickWidget) {

	qmlBridge = NewQmlBridge(nil)
	quickWidget.RootContext().SetContextProperty("qmlBridge", qmlBridge)

	qmlBridge.ConnectSendToGo(func(source, action, data string) {
		if source == "QmlButton" && action == "click" {
			rand.Seed(time.Now().UnixNano())
			manipulatedFromQml.SetStyleSheet(fmt.Sprintf("background-color:%v;", colors[rand.Intn(len(colors))]))
		}
	})

	qmlBridge.ConnectRegisterToGo(func(object *core.QObject) {
		qmlObjects[object.ObjectName()] = object
	})

	qmlBridge.ConnectDeregisterToGo(func(objectName string) {
		qmlObjects[objectName] = nil
	})
}
