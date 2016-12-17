package main

import (
	"fmt"
	"os"

	"github.com/therecipe/qt/core"
	"github.com/therecipe/qt/quick"
	"github.com/therecipe/qt/widgets"
)

//go:generate qtmoc
type QmlBridge struct {
	core.QObject

	_ func(title, author string) `signal:"addItem"`
	_ func(index int)            `signal:"removeItem"`

	_ func(row, role int, data *core.QObject) `slot:"editedModel"`
}

func main() {
	widgets.NewQApplication(len(os.Args), os.Args)

	var qmlBridge = NewQmlBridge(nil)
	qmlBridge.ConnectEditedModel(func(row int, role int, data *core.QObject) {
		fmt.Printf("row: %v role: %v title: %v author: %v\n", row, role, data.Property("title").ToString(), data.Property("author").ToString())
	})

	var (
		window       = widgets.NewQMainWindow(nil, 0)
		windowLayout = widgets.NewQVBoxLayout()
	)

	var centralWidget = widgets.NewQWidget(window, 0)
	centralWidget.SetLayout(windowLayout)
	window.SetCentralWidget(centralWidget)

	var addButton = widgets.NewQPushButton2("addItem", nil)
	addButton.ConnectClicked(func(checked bool) {
		qmlBridge.AddItem("New Title", "New Author")
	})
	windowLayout.AddWidget(addButton, 0, 0)

	var removeButton = widgets.NewQPushButton2("removeItem", nil)
	removeButton.ConnectClicked(func(checked bool) {
		qmlBridge.RemoveItem(0)
	})
	windowLayout.AddWidget(removeButton, 0, 0)

	var view = quick.NewQQuickWidget(nil)
	view.RootContext().SetContextProperty("QmlBridge", qmlBridge)
	view.SetSource(core.NewQUrl3("qrc:///qml/tableview.qml", 0))
	view.SetResizeMode(quick.QQuickWidget__SizeRootObjectToView)
	windowLayout.AddWidget(view, 0, 0)

	window.Show()

	widgets.QApplication_Exec()
}
