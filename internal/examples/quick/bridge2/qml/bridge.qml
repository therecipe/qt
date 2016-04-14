import QtQuick 2.0

Rectangle {
    id: page
    width: 320; height: 480
    color: "lightgray"
    MouseArea{
        anchors.fill: parent
        onClicked: {
            SignalHandler.callbackFromQml("hello world")
        }
    }
}