import QtQuick 2.7
import QtQuick.Controls 2.0
import QtQuick.VirtualKeyboard 2.0

ApplicationWindow {
    visible: true
    width: 500
    height: 400

    TextField {
        anchors.centerIn: parent
        placeholderText: "Write something ..."
    }

    InputPanel {
        y: active ? parent.height - height : parent.height
        anchors.left: parent.left
        anchors.right: parent.right
    }
}