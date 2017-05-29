package main

func layouts() {

	//Vertical Layout
	createWidget("Vertical Layout", `
	ColumnLayout {
		spacing: 0

		Rectangle {
			Layout.fillWidth: true
			Layout.preferredHeight: parent.height * 0.3
			color: "blue"
		}

		Rectangle {
			Layout.fillWidth: true
			Layout.fillHeight: true
			color: "red"
		}
	}`)

	//Horizontal Layout
	createWidget("Horizontal Layout", `
	RowLayout {
		spacing: 0

		Rectangle {
			Layout.preferredWidth: parent.width * 0.3
			Layout.fillHeight: true
			color: "blue"
		}

		Rectangle {
			Layout.fillWidth: true
			Layout.fillHeight: true
			color: "red"
		}
	}`)

	//Grid Layout
	createWidget("Grid Layout", `
	GridLayout {
		rowSpacing: 0
		columnSpacing: 0
		columns: 2

		Rectangle {
			Layout.fillWidth: true
			Layout.fillHeight: true
			color: "blue"
		}

		Rectangle {
			Layout.fillWidth: true
			Layout.fillHeight: true
			color: "red"
		}

		Rectangle {
			Layout.fillWidth: true
			Layout.fillHeight: true
			color: "green"
		}

		Rectangle {
			Layout.fillWidth: true
			Layout.fillHeight: true
			color: "yellow"
		}
	}`)

	//Stack Layout
	createWidget("Stack Layout", `
	MouseArea {

		Text {
			z: 1
			anchors.centerIn: parent
			text: "Click Me!"
		}

		StackLayout {
			id: stackLayout
			anchors.fill: parent

			Rectangle {
				Layout.fillWidth: true
				Layout.fillHeight: true
				color: "blue"
			}

			Rectangle {
				Layout.fillWidth: true
				Layout.fillHeight: true
				color: "red"
			}

			Rectangle {
				Layout.fillWidth: true
				Layout.fillHeight: true
				color: "green"
			}

			Rectangle {
				Layout.fillWidth: true
				Layout.fillHeight: true
				color: "yellow"
			}

			onCurrentIndexChanged: if (currentIndex >= count) { currentIndex = 0 }
		}

		onClicked: stackLayout.currentIndex++
	}`)

}
