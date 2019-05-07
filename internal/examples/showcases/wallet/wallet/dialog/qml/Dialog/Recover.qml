import QtQuick 2.7              //TextInput
import QtQuick.Layouts 1.3      //Layout

import Theme 1.0                //Theme
import DialogTemplate 1.0       //RecoverTemplate

import "."                      //Dialog

Dialog {
  id: dialog
  property RecoverTemplate template: RecoverTemplate {  onShow: dialog.visible = (cident == "recover"); }

  customTitle: "The entire blockchain will be scanned for outputs belonging to the seed.\nThis takes a while.\nAfter the scan completes, these outputs will be sent to your wallet."
  customLabel: "Enter a seed to recover funds from."

  customContentItem: Input {
    id: textField

    Layout.alignment: Qt.AlignVCenter | Qt.AlignHCenter
    Layout.preferredWidth: parent.width * 0.8

    onAccepted: dialog.accept()
  }

  onAccepted: {
    var ret = template.recover(textField.text)
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
