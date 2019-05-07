import QtQuick 2.7  //Item

import Theme 1.0    //Theme
import View 1.0     //StackItemHeader

StackItemHeader {
  image: "qrc:/qml/assets/ic_account_balance_wallet_black_24px.svg"
  title: "Wallet"

  contentItem: Item {
    anchors.fill: parent

    HeaderButton {
      anchors {
        top: parent.top
        right: send.left
        bottom: parent.bottom
      }
      width: Theme.minimumWidth * 0.1

      text: "Receive Coin"
      code: "receive"
      image: "qrc:/qml/assets/ic_file_download_black_24px.svg"
    }

    HeaderButton {
      id: send

      anchors {
        top: parent.top
        right: recover.left
        bottom: parent.bottom
      }
      width: Theme.minimumWidth * 0.1

      text: "Send Coin"
      code: "send"
      image: "qrc:/qml/assets/ic_file_upload_black_24px.svg"
    }

    HeaderButton {
      id: recover

      anchors {
        top: parent.top
        right: parent.right
        bottom: parent.bottom
      }
      width: Theme.minimumWidth * 0.1

      text: "Recover Lost Funds"
      code: "recover"
      image: "qrc:/qml/assets/ic_vpn_key_black_24px.svg"
    }
  }
}
