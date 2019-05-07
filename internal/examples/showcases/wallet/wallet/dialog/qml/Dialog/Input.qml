import QtQuick 2.7            //Rectangle
import QtQuick.Controls 2.1   //TextField

import Theme 1.0

TextField {
  font.pointSize: 12

  horizontalAlignment: Text.AlignHCenter
  verticalAlignment: Text.AlignVCenter

  background: Rectangle {
    color: Theme.inputFieldBackground
  }

  color: Theme.accent
  selectionColor: Theme.font
}
