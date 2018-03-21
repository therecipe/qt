import QtQuick 2.7            //Rectangle
import QtQuick.Controls 1.4   //ScrollView
import QtQuick.Controls 2.1   //TextArea

import Theme 1.0              //Theme

Rectangle {
  property alias text : output.text

  ScrollView {
    anchors.fill: parent

    TextArea {
      id: output

      readOnly: true
      wrapMode: TextEdit.Wrap
      selectByMouse: true

      selectionColor: Theme.accent
    }
  }
}
