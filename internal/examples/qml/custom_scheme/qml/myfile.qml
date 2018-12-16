import QtQuick 2.0
import "custom:///qml/"

Rectangle {
  id: page
  width: 320; height: 480
  color: "lightblue"

  Text {
    id: helloText
    text: "hello from qrc:///"
    y: 30
    anchors.horizontalCenter: page.horizontalCenter
    font.pointSize: 24; font.bold: true
  }

  SomeClass {
    height: parent.height/2

    anchors {
      left: parent.left
      right: parent.right
      bottom: parent.bottom
    }
  }
}
