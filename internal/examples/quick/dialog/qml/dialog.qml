import QtQuick 2.5
import QtQuick.Controls 1.4
import QtQuick.Dialogs 1.2

Rectangle {
  id: myRect

  Text {
    anchors.centerIn: parent
    text: "click me!"
  }

  MouseArea {
      id: mouseArea
      anchors.fill: parent
      onClicked: myDialog.open()
  }

  Dialog {
      id: myDialog

      visible: false
      title: "Blue sky dialog"
      modality: Qt.ApplicationModal

      contentItem: Rectangle {
          color: "lightskyblue"
          implicitWidth: myRect.width / 2
          implicitHeight: myRect.width / 2
          Text {
              text: "Hello blue sky!"
              color: "navy"
              anchors.centerIn: parent
          }
      }
  }
}
