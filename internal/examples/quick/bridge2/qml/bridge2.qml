import QtQuick 2.5

Rectangle
{
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

	MouseArea
	{
		anchors.fill: parent

		onClicked: console.log(QmlBridge.sendToGo("hello from qml"))
	}
}
