import QtQuick 2.0

Rectangle
{
	id: root
	width: 320
	height: 240

	color: "black"

	Text
	{
		id: title
		anchors.centerIn: parent

		color: "white"
		font.pointSize: 16

		Connections
		{
			target: QmlBridge
			onSendToQml: title.text = data
		}
	}

	Text
	{
		id: titleShadow
		anchors
		{
			horizontalCenter: parent.horizontalCenter
			verticalCenter: parent.verticalCenter
		}

		color: "white"
		font.pointSize: 16
		opacity: 0.75

		Connections
		{
			target: QmlBridge
			onSendToQml:
			{
				titleShadow.text = data
				titleShadow.state = "up"
			}
		}

		states: State
		{
			name: "up";
			AnchorChanges
			{
				target: titleShadow
				anchors.horizontalCenter: undefined
				anchors.verticalCenter: undefined
				anchors.bottom: root.top
			}
		}

		transitions: Transition
		{
			AnchorAnimation { duration: 900 }

			onRunningChanged:
			{
				if ((titleShadow.state == "up") && (!running)) //after the transition is completed
				{
					enabled = false
					titleShadow.state = ""
					enabled = true
				}
			}
		}
	}

	MouseArea
	{
		anchors.fill: parent

		onClicked: console.log(QmlBridge.sendToGo("hello from qml"))
	}
}
