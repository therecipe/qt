import Felgo 3.0
import QtQuick 2.0

Row {
  property alias iconItem: icon
  property alias textItem: text

  Icon {
    id: icon
    icon: IconType.eye
    color: Theme.secondaryTextColor
    size: dp(12)
  }

  AppText {
    id: text
    text: "12345"
    color: Theme.secondaryTextColor
    font.pixelSize: dp(12)
  }
}
