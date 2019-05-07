import QtQuick 2.7              //TextInput
import QtQuick.Layouts 1.3      //Layout
import QtQuick.Controls 2.1     //Label

import Theme 1.0                //Theme
import DialogTemplate 1.0       //SendTemplate

import "."                      //Dialog

Dialog{
  id: dialog
  property SendTemplate template: SendTemplate {  onShow: dialog.visible = (cident == "send"); }

  customTitle: "Send Amount <b>(C)</b>"

  customContentItem: [
    Input {
      id: amountField

      Layout.alignment: Qt.AlignVCenter | Qt.AlignHCenter
      Layout.preferredWidth: parent.width * 0.8
    },

    Label {
      Layout.alignment: Qt.AlignVCenter | Qt.AlignHCenter
      Layout.preferredWidth: parent.width * 0.8

      text: "To Address"
      color: Theme.fontHighlight
      font.pointSize: 14

      horizontalAlignment: Text.AlignHCenter
      verticalAlignment: Text.AlignVCenter
    },

    Input {
      id: addressField

      Layout.alignment: Qt.AlignVCenter | Qt.AlignHCenter
      Layout.preferredWidth: parent.width * 0.8
    }
  ]

  onAccepted: {
    var ret =  template.send(amountField.text, addressField.text)
    if (!ret[0]) {
      dialog.visible = true

      dialog.customError = ret[1]
    } else {
      dialog.reject()
    }
  }

  onRejected: {
    dialog.customError = ""
    amountField.clear()
    addressField.clear()
  }

  onVisibleChanged: template.blur(visible)
}
