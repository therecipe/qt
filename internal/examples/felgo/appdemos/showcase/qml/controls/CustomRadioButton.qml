import Felgo 3.0
import QtQuick 2.5
import QtQuick.Controls 2.0 as Quick2

Quick2.RadioButton {
  id: radioButton
  implicitWidth: leftPadding + indicator.implicitWidth + spacing + contentItem.implicitWidth + rightPadding
  implicitHeight: topPadding + indicator.implicitHeight + bottomPadding
  topPadding: 0
  bottomPadding: 0
  leftPadding: 0
  rightPadding: 0

  // overwrite style for density-independent sizes
  contentItem: AppText {
    text: parent.text
    font.pixelSize: sp(12)
    anchors.left: parent.left
    anchors.leftMargin: parent.indicator.width + parent.indicator.x + parent.spacing
  }

  indicator: Item {
    implicitWidth: dp(36)
    implicitHeight: implicitWidth
    x: parent.leftPadding
    y: parent.height / 2 - height / 2
    Rectangle {
      anchors.centerIn: parent
      implicitWidth: dp(20)
      implicitHeight: implicitWidth
      radius: width * 0.5
      border.color: radioButton.checked ? Theme.tintColor : Theme.secondaryTextColor
      border.width: 2

      Rectangle {
        width: parent.width * 0.5
        height: width
        anchors.centerIn: parent
        radius: width * 0.5
        color: Theme.tintColor
        visible: radioButton.checked
      }
    }
  }
}
