import QtQuick 2.7              //TextInput
import QtQuick.Layouts 1.3      //Layout

import Theme 1.0                //Theme
import DialogTemplate 1.0       //ReceiveTemplate

import "."                      //Dialog

Dialog {
  id: dialog
  property ReceiveTemplate template: ReceiveTemplate {  onShow: dialog.visible = (cident == "receive"); }

  customTitle: "You can receive Coins using the following address:"

  isCancelable: false

  customContentItem: Input {
    id: textField

    Layout.alignment: Qt.AlignVCenter | Qt.AlignHCenter
    Layout.preferredWidth: parent.width * 0.8

    text: template.address

    readOnly: true
    selectionColor: Theme.font
    selectByMouse: true

    onAccepted: dialog.accept()
  }

  onVisibleChanged: template.blur(visible)
}
