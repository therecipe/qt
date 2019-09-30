import Felgo 3.0
import QtQuick 2.0

Row {
  id: row

  property alias text: text.text

  Item {
    width: dp(16)
    height: dp(16)

    Rectangle {
      anchors.verticalCenter: parent.verticalCenter
      width: dp(5)
      height: dp(5)
      color: Theme.tintColor
    }
  }

  AppText {
    id: text
    width: parent.width - dp(20)
    font.pixelSize: sp(13)
    wrapMode: Text.WordWrap
    textFormat: AppText.RichText
    //color: Theme.secondaryTextColor
  }

}
