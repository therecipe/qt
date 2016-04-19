import QtQuick 2.5

Rectangle
{
	width: 160
	height: 240

	Rectangle
	{
		id: qmlButton
		objectName: "QmlButton"

		anchors
		{
			top: parent.top
			left: parent.left
			right: parent.right
		}

		height: parent.height * 0.15

		color: JSON.parse(qmlInitContext)[objectName].color

		Text
		{
			anchors.centerIn: parent

			text: JSON.parse(qmlInitContext)[parent.objectName].text
		}

		MouseArea
		{
			id: qmlButtonMouseArea
			anchors.fill: parent

			onClicked: qmlBridge.sendToGo(parent.objectName, "click", "")
		}

		states: State
		{
			name: "down"; when: qmlButtonMouseArea.pressed
			PropertyChanges { target: qmlButton; color: JSON.parse(qmlInitContext)[objectName].pressedColor }
		}

		Component.onCompleted: qmlBridge.registerToGo(this)
		Component.onDestruction: qmlBridge.deregisterToGo(objectName)
	}

	Rectangle
	{
		id: manipulatedFromGo
		objectName: "ManipulatedFromGo"

		anchors
		{
			top: qmlButton.bottom
			left: parent.left
			right: parent.right
			bottom: parent.bottom
		}

		Connections
		{
			target: qmlBridge
			onSendToQml:
			{
				if (source == "GoButton" && action == "click")
					manipulatedFromGo.color = data
			}
		}

		Component.onCompleted: qmlBridge.registerToGo(this)
		Component.onDestruction: qmlBridge.deregisterToGo(objectName)
	}
}
