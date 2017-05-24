package main

func buttons() {

	//Push Button
	createWidget("Push Button", `
	Item {
		Button {
			anchors.centerIn: parent
			text: "Push Button"
			onClicked: console.log("You Clicked The Push Button")
		}
	}
`)

	//Tool Button (+ ToolBar)
	createWidget("Tool Button", `
	Item {
		ToolBar {
			RowLayout {
				anchors.fill: parent

				ToolButton {
					Layout.fillHeight: true
					text: "Tool Button"
					onClicked: console.log("You Clicked The Tool Button")
				}
			}
		}
	}
`)

	//Radio Button
	createWidget("Radio Button", `
	Item {
		ColumnLayout {
			anchors.fill: parent

			ExclusiveGroup { id: radioGroup }

			RadioButton {
				text: "Radio Button 0"
				onClicked: console.log("You Clicked The " + text)
				exclusiveGroup: radioGroup
			}

			RadioButton {
				text: "Radio Button 1"
				onClicked: console.log("You Clicked The " + text)
				exclusiveGroup: radioGroup
			}

			RadioButton {
				text: "Radio Button 2"
				onClicked: console.log("You Clicked The " + text)
				exclusiveGroup: radioGroup
			}
		}
	}`)

	//Check Box
	createWidget("Check Box", `
	Item {
		CheckBox {
			anchors.centerIn: parent
			text: "Check Box"
			onClicked: console.log("You Clicked The Check Box :", checked)
		}
	}`)

	//Dialog Button Box
	createWidget("Dialog Button Box", `
	Item {
		RowLayout {
			anchors.fill: parent

			Button {
				text: "Cancel"
				onClicked: console.log("You Clicked The Accept Button")
			}

				Button {
				text: "OK"
				isDefault: true
				onClicked: console.log("You Clicked The Reject Button")
			}
		}
	}`)

}
