package main

func displayWidgets() {

	//Label
	createWidget("Label", `
  Item {
    Label {
			anchors.centerIn: parent
      text: "This Is A Label"
    }
  }`)

	//Text
	createWidget("Text", `
  Item {
    Text {
			anchors.centerIn: parent
      text: "This Is A Text Item"
    }
  }`)

	//Calendar Widget
	createWidget("Calendar Widget", `
	Calendar {
		onClicked: console.log("Calender Widget Selected Date Changed To:", date)
	}`)

	//Progress Bar
	createWidget("Progress Bar", `
	Item {
		ProgressBar {
			id: progressBar
			anchors.centerIn: parent

			minimumValue: 0
			maximumValue: 1000
			value: 500

			NumberAnimation on value {
				from: 0
				to: 1000
				duration: 4000
				loops: Animation.Infinite
			}
		}
	}`)

	//Image
	createWidget("Image", `
	Image {
		fillMode: Image.PreserveAspectFit
		source: "qrc:/qml/earth.png"
	}`)

	//BusyIndicator
	createWidget("Busy Indicator", `
	Item {
		BusyIndicator {
			anchors.centerIn: parent
		}
	}`)

}
