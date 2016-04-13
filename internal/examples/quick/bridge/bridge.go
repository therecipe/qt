package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"os"
	"runtime"
	"sync"
	"time"

	"github.com/therecipe/qt/core"
	"github.com/therecipe/qt/quick"
	"github.com/therecipe/qt/widgets"
)

var (
	quickWidget   *quick.QQuickWidget
	messageObject *widgets.QWidget
	qmlObjects    = make(map[string]*core.QObject)

	manipulateFromQml *widgets.QWidget

	colors = []string{"red", "green", "blue", "cyan", "magenta", "yellow", "gray"}
)

func main() {

	widgets.NewQApplication(len(os.Args), os.Args)

	var window = widgets.NewQMainWindow(nil, 0)
	if runtime.GOOS != "android" {
		window.SetMinimumSize2(500, 250)
	}

	var layout = widgets.NewQHBoxLayout()
	layout.AddWidget(newCppWidget(), 0, 0)
	layout.AddWidget(newSeperator(), 0, 0)
	layout.AddWidget(newQmlWidget(), 0, 0)

	window.Layout().DestroyQObject()
	window.SetLayout(layout)

	if runtime.GOOS == "android" {
		window.ShowFullScreen()
	} else {
		window.Show()
	}

	widgets.QApplication_Exec()
}

func newCppWidget() *widgets.QWidget {
	var (
		widget = widgets.NewQWidget(nil, 0)
		layout = widgets.NewQVBoxLayout()
		button = widgets.NewQPushButton2("Call Qml Function", nil)
	)

	button.ConnectClicked(func(_ bool) {
		rand.Seed(time.Now().UnixNano())

		var msg = &message{Sender: "MainButton", Action: "click", Data: colors[rand.Intn(len(colors))]}
		var b, err = json.Marshal(msg)
		if err != nil {
			log.Println("newCppWidget.button.clicked", err)
		}

		qmlObjects["ManipulateFromGo"].SetProperty("messageFromGo", core.NewQVariant14(string(b)))
	})
	layout.AddWidget(button, 0, 0)

	manipulateFromQml = widgets.NewQWidget(nil, 0)
	layout.AddWidget(manipulateFromQml, 0, 0)

	widget.Layout().DestroyQObject()
	widget.SetLayout(layout)

	return widget
}

func newSeperator() *widgets.QFrame {
	var frame = widgets.NewQFrame(nil, 0)
	frame.SetFrameShape(widgets.QFrame__VLine)
	return frame
}

func newQmlWidget() *quick.QQuickWidget {
	quickWidget = quick.NewQQuickWidget(nil)
	quickWidget.SetResizeMode(quick.QQuickWidget__SizeRootObjectToView)

	qmlInitObject(quickWidget)
	qmlRegisterObject(quickWidget)
	qmlMessageObject(quickWidget)

	quickWidget.SetSource(core.NewQUrl3("qrc:///qml/bridge.qml", 0))

	qmlFindObjects(quickWidget)

	return quickWidget
}

func qmlInitObject(quickWidget *quick.QQuickWidget) {

	var m = map[string]map[string]string{
		"MainButton": map[string]string{
			"lightColor": "lightGray",
			"darkColor":  "darkGray",
			"Text":       "Call Go Function",
		},
	}

	var b, err = json.Marshal(m)
	if err != nil {
		log.Println("qmlInitObject", err)
	}
	quickWidget.RootContext().SetContextProperty2("qmlInitObject", core.NewQVariant14(string(b)))
}

func qmlRegisterObject(quickWidget *quick.QQuickWidget) {

	var (
		registerObject = widgets.NewQWidget(nil, 0)
		registerMutex  = new(sync.Mutex)
	)

	quickWidget.RootContext().SetContextProperty("qmlRegisterObject", registerObject)
	registerObject.ConnectWindowTitleChanged(func(objectName string) {
		if objectName != "" {
			registerMutex.Lock()
			registerObject.SetWindowTitle("")

			qmlObjects[objectName] = nil

			registerMutex.Unlock()
		}
	})
}

type message struct {
	Sender, Action, Data string
}

func qmlMessageObject(quickWidget *quick.QQuickWidget) {

	messageObject = widgets.NewQWidget(nil, 0)
	var messageMutex = new(sync.Mutex)

	quickWidget.RootContext().SetContextProperty("qmlMessageObject", messageObject)
	messageObject.ConnectWindowTitleChanged(func(jsonMessage string) {
		if jsonMessage != "" {
			messageMutex.Lock()
			messageObject.SetWindowTitle("")

			var msg = new(message)
			json.Unmarshal([]byte(jsonMessage), &msg)
			processQmlMessage(msg)

			messageMutex.Unlock()
		}
	})
}

func qmlFindObjects(quickWidget *quick.QQuickWidget) {
	for objectName, object := range qmlObjects {
		if object == nil {
			qmlObjects[objectName] = core.NewQObjectFromPointer(quickWidget.RootObject().FindChild(objectName, core.Qt__FindChildrenRecursively))
		}
	}
}

func processQmlMessage(msg *message) {
	switch msg.Sender {
	case "MainButton":
		{
			switch msg.Action {
			case "click":
				{
					rand.Seed(time.Now().UnixNano())
					manipulateFromQml.SetStyleSheet(fmt.Sprintf("background-color:%v;", colors[rand.Intn(len(colors))]))
				}

			default:
				{
					log.Println("Unknown Action:", msg)
				}
			}
		}

	default:
		{
			log.Println("Unknown Sender:", msg)
		}
	}
}
