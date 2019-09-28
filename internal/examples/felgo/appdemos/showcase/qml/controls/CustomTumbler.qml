import QtQuick 2.0
import QtQuick.Controls 2.0
import Felgo 3.0

Tumbler {
  id: customTumbler
  implicitWidth: textItem.implicitWidth + leftPadding + rightPadding
  implicitHeight: dp(100)
  topPadding: 0
  bottomPadding: topPadding
  leftPadding: dp(16)
  rightPadding: leftPadding

  // allows to define function to process the delegate text before displaying
  property var processDelegateText: undefined

  // hidden text for width calculation
  AppText {
    id: textItem
    visible: false
    text: processDelegateText ? processDelegateText(0) : 0
  }

  // styling
  delegate: AppText {
    text: processDelegateText ? processDelegateText(modelData) : modelData
    horizontalAlignment: Text.AlignHCenter
    verticalAlignment: Text.AlignVCenter
    color: blendColor(Theme.tintColor, Theme.textColor, 1 - Math.abs(Tumbler.displacement) / (visibleItemCount / 3))
    opacity: 1.0 - Math.abs(Tumbler.displacement) / (visibleItemCount / 2)

    function blendColor(col1, col2, mix) {
      var r = col1.r * (mix) + col2.r * (1 - mix)
      var g = col1.g * (mix) + col2.g * (1 - mix)
      var b = col1.b * (mix) + col2.b * (1 - mix)
      return Qt.rgba(r, g, b, 1)
    }
  }

  background: Rectangle {
    border.color: Theme.controlBackgroundColor
  }
}
