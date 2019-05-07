import QtQuick 2.7            //Image
import QtGraphicalEffects 1.0 //ColorOverlay

import Theme 1.0              //Theme
import TopTemplate 1.0        //ColorTemplate

ColorTemplate {
  id: template

  Rectangle {
    anchors.fill: parent

    color: Theme.background

    Image {
      id: color
      anchors.centerIn: parent

      source: "qrc:/qml/assets/ic_invert_colors_black_24px.svg"
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

      hoverEnabled: true
      onEntered: overlay.color = Theme.nextAccent
      onExited: overlay.color = Theme.accent

      onClicked: template.change()
    }
  }
}
