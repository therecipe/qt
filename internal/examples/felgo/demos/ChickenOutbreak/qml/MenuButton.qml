import QtQuick 2.0
import Felgo 3.0

// gets used for the Play button in MainScene and for the Continue button in GameOverScene
Image {
  id: button

  // width & height must get set from outside, these are the default values!
  width: 170
  height: 60
  source: "../assets/img/button.png"

  anchors.horizontalCenter: parent.horizontalCenter

  property alias text: buttonText.text
  property alias textColor: buttonText.color
  property alias textSize: buttonText.font.pixelSize
  property alias textItem: buttonText
  property alias font: buttonText.font

  signal clicked

  Text {
    id: buttonText
    anchors.centerIn: parent
    font.pixelSize: 22
    color: "#ca840a"

    font.family: fontHUD.name
  }

  MouseArea {
    id: mouseArea
    anchors.fill: parent
    hoverEnabled: true
    onClicked: button.clicked()
  }
}
