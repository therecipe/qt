package main

import (
	"github.com/therecipe/qt/core"
)

func inputWidgets() {

	//Abstract List Model Used By The Combo Box
	listDB := []string{"Some", "Combo", "Box", "Items"}
	listModel := core.NewQAbstractListModel(nil)
	listModel.ConnectRowCount(func(parent *core.QModelIndex) int {
		return len(listDB)
	})
	listModel.ConnectData(func(index *core.QModelIndex, role int) *core.QVariant {
		if role != int(core.Qt__DisplayRole) {
			return core.NewQVariant()
		}
		return core.NewQVariant14(listDB[index.Row()])
	})

	//Combo Box
	createWidget("Combo Box", `
	Item {
		ComboBox {
			anchors.centerIn: parent
			model: ["Some", "Combo", "Box", "Items"]
			onActivated: console.log("Combo Box Index Changed To:", index, model[index])
		}
	}`)

	//Combo Box + List Model
	createWidget("Combo Box + List Model", `
	Item {
		ComboBox {
			anchors.centerIn: parent
			model: ListModel {
				ListElement { text: "Some" }
				ListElement { text: "Combo" }
				ListElement { text: "Box" }
				ListElement { text: "Items" }
			}

			onActivated: console.log("Combo Box Index Changed To:", index, model.get(index).text)
		}
	}`)

	//Combo Box + Go Model
	createWidgetWithContext("Combo Box + Go Model", `
	Item {
		ComboBox {
			anchors.centerIn: parent
			model: goListModel

			textRole: "display"
			onActivated: console.log("Combo Box Index Changed To:", index, model.data(model.index(index, 0)))
		}
	}`, "goListModel", listModel)

	//Combo Box + Go String List Model
	createWidgetWithContext("Combo Box + Go String List Model", `
	Item {
		ComboBox {
			anchors.centerIn: parent
			model: goListModel

			textRole: "display"
			onActivated: console.log("Combo Box Index Changed To:", index, model.data(model.index(index, 0)))
		}
	}`, "goListModel", core.NewQStringListModel2(listDB, nil))

	//Text Field
	createWidget("Text Field", `
	Item {
		TextField {
			anchors.centerIn: parent
			onTextChanged: console.log("Text Field Text Changed To:", text)
		}
	}
	`)

	//Text Input
	createWidget("Text Input", `
	Item {
		TextInput {
			anchors.centerIn: parent
			text: "Some Editable Text"
			onTextChanged: console.log("Text Input Text Changed To:", text)
		}
	}
	`)

	//Text Area
	createWidget("Text Area", `
	TextArea {
		onTextChanged: console.log("Text Area Text Changed To:", text)
	}
	`)

	//Text Edit
	createWidget("Text Edit", `
	TextEdit {
		onTextChanged: console.log("Text Edit Text Changed To:", text)
	}
	`)

	//Spin Box
	createWidget("Spin Box", `
	Item {
		SpinBox {
			anchors.centerIn: parent
			minimumValue: 0
			maximumValue: 1000
			value: 500
			onValueChanged: console.log("Spin Box Value Changed To:", value)
		}
	}`)

	//Double Spin Box
	createWidget("Double Spin Box", `
	Item {
		SpinBox {
			anchors.centerIn: parent
			minimumValue: 0
			maximumValue: 1000
			value: 500
			decimals: 2
			onValueChanged: console.log("Spin Box Value Changed To:", value)
		}
	}`)

	//Switch
	createWidget("Switch", `
	Item {
		Switch {
			anchors.centerIn: parent
			onClicked: console.log("Switch Button Clicked:", checked)
		}
	}`)

	//Horizontal Slider
	createWidget("Horizontal Slider", `
		Item {
			Slider {
				anchors.centerIn: parent
				minimumValue: 0
				maximumValue: 1000
				value: 500
				orientation: Qt.Horizontal
				stepSize: 1
				onValueChanged: console.log("Horizontal Slider Value Changed To:", value)
			}
		}
	`)

	//Vertical Slider
	createWidget("Vertical Slider", `
	Item {
		Slider {
			anchors.centerIn: parent
			minimumValue: 0
			maximumValue: 1000
			value: 500
			orientation: Qt.Vertical
			stepSize: 1
			onValueChanged: console.log("Vertical Slider Value Changed To:", value)
		}
	}
`)
}
