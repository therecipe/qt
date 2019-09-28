import Felgo 3.0
import QtQuick 2.5

// rectangle that holds AppTextEdit
Rectangle {
  implicitWidth: dp(200)
  implicitHeight: textEdit.contentHeight + 2 * textEdit.textMargin + textEdit.y * 2
  border.width: dp(1)
  border.color: textEdit.activeFocus ? Theme.tintColor : Theme.controlBackgroundColor
  clip: true

  AppTextEdit {
    id: textEdit
    y: dp(Theme.navigationBar.defaultBarItemPadding) * 0.5
    height: contentHeight + 2 * textEdit.textMargin
    width: parent.width - dp(Theme.navigationBar.defaultBarItemPadding)
    placeholderText: "Enter text ...";
    font.pixelSize: dp(14)
    anchors.horizontalCenter: parent.horizontalCenter
    verticalAlignment: Text.AlignVCenter
    wrapMode: AppTextEdit.WordWrap
  }
}
