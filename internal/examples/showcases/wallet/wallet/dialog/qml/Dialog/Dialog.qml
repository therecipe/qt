import QtQuick 2.7                //Rectangle
import QtQuick.Layouts 1.3        //ColumnLayout
import QtQuick.Controls 2.1       //Label

import Theme 1.0                  //Theme

Dialog {
  modal: true
  closePolicy: Popup.CloseOnEscape

  property alias customTitle : titleLabel.text
  property alias customLabel : labelLabel.text
  property alias customError : errorField.text
  property alias customContentItem : column.children
  property alias isCancelable : cancel.visible

  width: Theme.minWidth * 0.7
  height: Theme.minHeight * 0.4

  x: ( ( Theme.currentWidth - width ) * 0.5 ) - Theme.minWidth * 0.2
  y: ( ( Theme.currentHeight - height ) * 0.5 ) - Theme.minHeight * 0.09

  background: null

  Rectangle {
    id: root
    anchors.fill: parent

    radius: 5
    color: Theme.walletTableHeader

    ColumnLayout {
      id: column

      anchors.centerIn: parent

      spacing: 10

      Label {
        id: titleLabel

        Layout.alignment: Qt.AlignVCenter | Qt.AlignHCenter
        Layout.preferredWidth: root.width * 0.8

        color: Theme.fontHighlight
        font.pointSize: 16
        font.bold: true

        horizontalAlignment: Text.AlignHCenter
        verticalAlignment: Text.AlignVCenter
      }

      Label {
        id: labelLabel

        Layout.alignment: Qt.AlignVCenter | Qt.AlignHCenter
        Layout.preferredWidth: root.width * 0.8

        color: Theme.fontHighlight
        font.pointSize: 14

        horizontalAlignment: Text.AlignHCenter
        verticalAlignment: Text.AlignVCenter

        visible: customLabel != ""
      }
    }
  }

  Label {
    id: errorField

    Layout.alignment: Qt.AlignVCenter | Qt.AlignHCenter
    Layout.preferredWidth: root.width * 0.8

    wrapMode: TextEdit.Wrap

    color: "red"
    font.pointSize: 16
    font.bold: true

    horizontalAlignment: Text.AlignHCenter
    verticalAlignment: Text.AlignVCenter

    visible: customError != ""
  }

  RowLayout {
    id: buttons

    Layout.alignment: Qt.AlignVCenter | Qt.AlignHCenter
    Layout.fillWidth: true

    Button {
      id: control

      background: null

      contentItem: Label {
        text: "OK"

        color: control.down || control.hovered ? Theme.fontHighlight : Theme.font
        font.pointSize: 16
        font.bold: true

        horizontalAlignment: Text.AlignHCenter
        verticalAlignment: Text.AlignVCenter
      }

      onClicked: dialog.accept()
    }

    Button {
      id: cancel

      background: null

      contentItem: Label {
        text: "CANCEL"

        color: cancel.down || cancel.hovered ? Theme.fontHighlight : Theme.font
        font.pointSize: 16
        font.bold: true

        horizontalAlignment: Text.AlignHCenter
        verticalAlignment: Text.AlignVCenter
      }

      onClicked: dialog.reject()
    }
  }

  Component.onCompleted: {
    //relocate errorField and buttons
    //to the end of the column layout

    errorField.parent = column
    buttons.parent = column
  }
}
