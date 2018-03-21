import QtQuick 2.7           //Rectangle
import QtQuick.Controls 2.1  //ProgressBar

import Theme 1.0             //Theme

ProgressBar {

  property alias progressBarText : progressBarLabel.text
  property bool error

  id: progressBar

  from: 0
  to: 100

  background: Rectangle {
    radius: 25
    color: styleData.row % 2 == 0 ? Theme.background : Theme.darkBackground
  }

  contentItem: Item { Rectangle {
    width: progressBar.visualPosition * parent.width
    height: parent.height
    radius: 25
    color: error ? Theme.walletTableHighlight : Theme.accent
  } }

  Label {
    id: progressBarLabel

    anchors.centerIn: parent

    font.pointSize: 12
    font.bold: true

    color: Theme.font
  }
}
