import QtQuick 2.7            //Item
import WalletTemplate 1.0     //WalletTemplate
import Dialog 1.0             //Unlock

WalletTemplate {
  id: template

  Unlock {}
  Receive {}
  Send {}
  Recover {}

  Header {
    id: header
  }

  Item {
    anchors {
      top: header.bottom
      left: parent.left
      right: parent.right
      bottom: parent.bottom
      leftMargin: 5
      rightMargin: 5
    }

    Table {
      anchors.fill: parent
      template: template
    }
  }
}
