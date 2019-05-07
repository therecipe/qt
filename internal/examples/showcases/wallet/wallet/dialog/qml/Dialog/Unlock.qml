import QtQuick 2.7              //TextInput
import QtQuick.Layouts 1.3      //Layout

import Theme 1.0                //Theme
import DialogTemplate 1.0       //UnlockTemplate

import "."                      //Dialog

Dialog {
  id: dialog
  property UnlockTemplate template: UnlockTemplate {  onShow: dialog.visible = (cident == "unlock"); }

  customTitle: "Your Wallet Is Locked"
  customLabel: "Please Enter Your Password To Unlock Your Wallet"

  customContentItem: Input {
    id: textField

    Layout.alignment: Qt.AlignVCenter | Qt.AlignHCenter
    Layout.preferredWidth: parent.width * 0.8

    echoMode: TextInput.Password
    onAccepted: dialog.accept()
  }

  onAccepted: {
    var ret = template.unlock(textField.text)
    if (!ret[0]) {
      dialog.visible = true

      dialog.customError = ret[1]
      textField.clear()
    } else {
      dialog.reject()
    }
  }

  onRejected: {
    dialog.customError = ""
    textField.clear()
  }

  onVisibleChanged: template.blur(visible)
}
