import QtQuick 2.7           //Rectangle
import QtQuick.Controls 2.1  //ProgressBar

import Theme 1.0             //Theme
import LeftTemplate 1.0      //ProgressBarTemplate

ProgressBarTemplate {
  id: template

  ProgressBar {
    id: progressBar
    anchors.centerIn: parent
    height: parent.height * 0.25
    width: Theme.minWidth * 0.18

    from: 0
    to: 100
    value: template.value

    background: Rectangle {
      radius: 25
      color: Theme.darkBackground
    }

    contentItem: Item { Rectangle {
      width: progressBar.visualPosition * parent.width
      height: parent.height
      radius: 25
      color: Theme.accent
    } }

    Label {

      z: 1
      
      anchors.centerIn: parent

      font.pointSize: 12
      font.bold: true

      text: template.text
      color: Theme.walletTableHeader
    }
  }

  MouseArea {
    anchors.fill: progressBar
    onClicked: template.clicked()
  }
}
