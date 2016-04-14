import QtQuick 2.0

Rectangle {
    id: page
    width: 640; height: 480
    color: "black"
    MouseArea{
        anchors.fill: parent
        onClicked: {
            SignalHandler.callbackFromQml("hello world")
        }
    }
	
	Text {
        id: title
        color: "white"
        anchors.centerIn: parent
        font.pointSize: 16
        text: ""

        Connections{
            target: SignalHandler
            onSendToQml: {
                title.text = value
            }
        }
    }
}