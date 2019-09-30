
import Felgo 3.0
import QtQuick 2.0

Rectangle {
  id: colorButton
  width: dp(40)
  height: width
  radius: width * 0.25

  property color referenceColor
  property bool selected: color === referenceColor

  property color _greyColor: "grey"

  border.color: selected ? Theme.tintColor :
                           (color === _greyColor ? "white" : "grey")
  border.width: selected ? dp(2) : Theme.backgroundColor == color ? dp(1) : 0


  signal clicked

  MouseArea {
    anchors.fill: parent
    onClicked: colorButton.clicked()
  }
}
