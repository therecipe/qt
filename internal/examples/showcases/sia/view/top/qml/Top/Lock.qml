import QtQuick 2.7            //Image
import QtGraphicalEffects 1.0 //ColorOverlay

import Theme 1.0              //Theme
import TopTemplate 1.0        //LockTemplate

LockTemplate {
  id: template

  Rectangle {
    anchors.fill: parent

    color: Theme.background

    Image {
      id: color
      anchors.centerIn: parent

      source: template.locked ? "qrc:/qml/assets/ic_lock_outline_black_24px.svg" : "qrc:/qml/assets/ic_lock_open_black_24px.svg"
    }

    ColorOverlay {
      id: overlay
      anchors.fill: color
      source: color
      color: Theme.accent
    }

    MouseArea {
      id: mouseArea
      anchors.fill: parent
      onClicked: template.change()
    }
  }
}
