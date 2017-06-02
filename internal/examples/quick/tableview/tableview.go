package main

import (
	"fmt"
	"os"

	"github.com/therecipe/qt/core"
	"github.com/therecipe/qt/quick"
	"github.com/therecipe/qt/widgets"
)

type QmlBridge struct {
	core.QObject

	//from golang to qml's onAddItem
	_ func(title string, author string, famous bool) `signal:"addItem"`

	//from golang to qml's onRemoveItem
	_ func(index int) `signal:"removeItem"`

	//from qml to golang's ConnectEditedModel
	_ func(row int, role int, data *core.QObject) `slot:"editedModel"`
}

func main() {
	widgets.NewQApplication(len(os.Args), os.Args)

	var qmlBridge = NewQmlBridge(nil)
	qmlBridge.ConnectEditedModel(func(row int, role int, data *core.QObject) {
		fmt.Printf("row: %v\n", row)
		fmt.Printf("role: %v\n", role)

		sTitle := data.Property("title").ToString()
		fmt.Printf("title: %v\n", sTitle)

		sAuthor := data.Property("author").ToString()
		fmt.Printf("author: %v\n", sAuthor)

		bFamous := data.Property("famous").ToBool()
		fmt.Printf("famous: %v\n", bFamous)
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
		qmlBridge.AddItem("New Title", "New Author", false)
	})
	windowLayout.AddWidget(addButton, 0, 0)

	var removeButton = widgets.NewQPushButton2("removeItem", nil)
	removeButton.ConnectClicked(func(checked bool) {
		qmlBridge.RemoveItem(0)
	})
	windowLayout.AddWidget(removeButton, 0, 0)

	var view = quick.NewQQuickWidget(nil)
	view.RootContext().SetContextProperty("QmlBridge", qmlBridge)
	view.SetSource(core.NewQUrl3("qrc:/qml/tableview.qml", 0))
	view.SetResizeMode(quick.QQuickWidget__SizeRootObjectToView)
	windowLayout.AddWidget(view, 0, 0)

	window.Show()

	widgets.QApplication_Exec()
}
