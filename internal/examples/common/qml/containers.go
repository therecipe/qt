package main

func containers() {

	//Group Box
	createWidget("Group Box", `
	Item {
		GroupBox {
			anchors.centerIn: parent
			title: "Group Box"

			ColumnLayout {
				Button {
					text: "Push Button: 0"
				}

				Button {
					text: "Push Button: 1"
				}

				Button {
					text: "Push Button: 2"
				}
			}
		}
	}`)

	//Scroll Area
	createWidget("Scroll Area", `
	ScrollView {
		GridLayout {
			width: 250
			height: 250

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
		}
	}`)

	//Tab Widget
	createWidget("Tab Widget", `
	TabView {
		Tab {
				title: "red"
				Rectangle { color: "red" }
		}
		Tab {
				title: "blue"
				Rectangle { color: "blue" }
		}
		Tab {
				title: "green"
				Rectangle { color: "green" }
		}
		Tab {
				title: "yellow"
				Rectangle { color: "yellow" }
		}
	}`)

	//Stacked Widget
	createWidget("Stacked Widget", `
		MouseArea {

			Text {
				z: 1
				anchors.centerIn: parent
				text: "Click Me!"
			}

			StackView {
				id: stack
				anchors.fill: parent

				initialItem: blue

				Rectangle {
					id: blue
					color: "blue"
				}

				Rectangle {
					id: red
					color: "red"
				}

				Rectangle {
					id: green
					color: "green"
				}

				Rectangle {
					id: yellow
					color: "yellow"
				}
			}

			onClicked: {
				switch (stack.depth) {
					case 1:
						stack.push(red)
						break;

					case 2:
						stack.push(green)
						break;

					case 3:
						stack.push(yellow)
						break;

					case 4:
						stack.pop(blue)
						break;
				}
			}
		}
	`)

}
