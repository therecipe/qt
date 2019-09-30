import Felgo 3.0
import QtQuick 2.0

Rectangle {
  id: item
  width: dp(40)
  height: dp(45)
  color: "white"

  property string titleText: "Days"
  property int value: 21

  Column {
    anchors.centerIn: parent
    AppText {
      text: item.value < 10 ? "0" + item.value : item.value
      anchors.horizontalCenter: parent.horizontalCenter
      font.pixelSize: sp(16)
    }
    AppText {
      text: item.titleText
      anchors.horizontalCenter: parent.horizontalCenter
      font.pixelSize: sp(8)
      font.capitalization: Font.AllUppercase
    }
  }

}
