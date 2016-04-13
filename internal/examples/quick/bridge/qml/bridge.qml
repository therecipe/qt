import QtQuick 2.5

Rectangle
{
	Rectangle
	{
		id: mainButton
		objectName: "MainButton"

		anchors
		{
			top: parent.top
			left: parent.left
			right: parent.right
		}

		height: parent.height * 0.15

		color: JSON.parse(qmlInitObject).MainButton.lightColor

		Text
		{
			anchors.fill: parent

			text: JSON.parse(qmlInitObject).MainButton.Text

			horizontalAlignment: Text.AlignHCenter
			verticalAlignment: Text.AlignVCenter
		}

		MouseArea
		{
			id: mainButtonMouseArea
			anchors.fill: parent

			onClicked: clickMainButton()
		}

		states: State
		{
			name: "down"; when: mainButtonMouseArea.pressed
			PropertyChanges { target: mainButton; color: JSON.parse(qmlInitObject).MainButton.darkColor }
		}

		Component.onCompleted: qmlRegisterObject.setWindowTitle(objectName)
		Component.onDestruction: qmlRegisterObject.setWindowTitle(objectName)
	}

	function clickMainButton()
	{
		qmlMessageObject.setWindowTitle(JSON.stringify({"Sender":"MainButton","Action":"click","Data":""}))
	}

	Rectangle
	{
		id: manipulateFromGo
		objectName: "ManipulateFromGo"
		property string messageFromGo

		anchors
		{
			top: mainButton.bottom
			left: parent.left
			right: parent.right
			bottom: parent.bottom
		}

		onMessageFromGoChanged: processGoMessage(messageFromGo)

		Component.onCompleted: qmlRegisterObject.setWindowTitle(objectName)
		Component.onDestruction: qmlRegisterObject.setWindowTitle(objectName)
	}

	function processGoMessage(message)
	{
		if (message != "")
		{
			var msg = JSON.parse(message)

			switch (msg.Sender)
			{
				case "MainButton":
				{
					switch (msg.Action)
					{
						case "click":
						{
							manipulateFromGo.color = msg.Data
							break;
						}

						default:
						{
							console.log("Unknown Action:", message)
							break;
						}
					}

					break;
				}

				default:
				{
					console.log("Unknown Sender:", message)
					break;
				}
			}
		}
	}
}
