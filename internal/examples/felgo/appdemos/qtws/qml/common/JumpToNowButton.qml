import Felgo 3.0
import QtQuick 2.0

AppButton {
  id: jumpToNowButton
  anchors.horizontalCenter: parent.horizontalCenter
  text: "jump to now"

  // styling
  flat: false
  radius: width * 0.12
  backgroundColor: Theme.tintColor
  backgroundColorPressed: Theme.tintLightColor
  textColor: Theme.backgroundColor
  textSize: sp(12)
  minimumHeight: 0
  minimumWidth: 0

  // animate opacity and y position
  Behavior on opacity {
    PropertyAnimation { duration: 250 }
  }
  Behavior on y {
    PropertyAnimation { duration: 250 }
  }
}
