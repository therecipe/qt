package main

import (
	"fmt"
	"os"
	"plugin"

	"github.com/therecipe/qt/core"
	"github.com/therecipe/qt/qml"
	"github.com/therecipe/qt/quick"
	"github.com/therecipe/qt/widgets"
)

//go:generate cp main.go ./plugin
//go:generate qtrcc desktop ./plugin
//go:generate qtmoc desktop ./plugin
//go:generate qtminimal desktop ./plugin
//go:generate rm ./plugin/main.go
//go:generate go build -tags=minimal -buildmode=plugin -o ./plugin/plugin.so ./plugin

//go:generate go build -tags=minimal

func init() {
	if _, err := plugin.Open("plugin/plugin.so"); err != nil {
		panic(err)
	}
}

func main() {
	widgets.NewQApplication(len(os.Args), os.Args)
	view := quick.NewQQuickView(nil)
	view.Engine().AddImportPath("qrc:/")
	view.SetResizeMode(quick.QQuickView__SizeRootObjectToView)

	mainComponent := qml.NewQQmlComponent2(view.Engine(), nil)
	mainComponent.ConnectStatusChanged(func(status qml.QQmlComponent__Status) {
		if status == qml.QQmlComponent__Ready {

			item := quick.NewQQuickItemFromPointer(mainComponent.Create(view.RootContext()).Pointer()) //create item and "cast" it to QQuickItem
			item.SetParent(view)                                                                       //add invisible item to invisible parent (for auto-deletion ...)
			item.SetParentItem(view.ContentItem())                                                     //add visible item to visible parent

		} else {
			fmt.Println("failed with status:", status)
			for _, e := range mainComponent.Errors() {
				fmt.Println("error:", e.ToString())
			}
		}
	})

	qmlString := `
import QtQuick 2.0
import CustomComponents 1.0

Item {
	anchors.fill: parent

	SomeComponent {
		id: rootItem
		anchors.centerIn: parent
		width: 100; height: 100
		color: "blue"

		Component.onCompleted: {
			var subComponent = Qt.createQmlObject('import QtQuick 2.0; import CustomComponents 1.0; SomeComponent { anchors.centerIn: rootItem; width: 50; height: 50; color: "red" }', rootItem);
			var subSubComponent = Qt.createComponent("qrc:///CustomComponents/SomeComponent.qml").createObject(subComponent, { width: 25, height: 25, color: "green" });
			subSubComponent.anchors.centerIn = subComponent;
		}
	}
}
`

	mainComponent.SetData(core.NewQByteArray2(qmlString, -1), core.NewQUrl())

	view.Show()
	widgets.QApplication_Exec()
}
