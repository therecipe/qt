import QtQuick 2.7
import QtQuick.Controls 2.0
import QtQuick.VirtualKeyboard 2.0
import CustomModule 1.0

ApplicationWindow {
    visible: true
    width: 500
    height: 400

    Embedded {
        anchors.centerIn: parent
        focus: true
    }

    InputPanel {
        y: active ? parent.height - height : parent.height
        anchors.left: parent.left
        anchors.right: parent.right
    }
}