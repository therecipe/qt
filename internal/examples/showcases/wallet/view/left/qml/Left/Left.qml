import QtQuick 2.7            //Image
import QtQuick.Controls 2.1   //ButtonGroup
import QtQuick.Layouts 1.3    //ColumnLayout

import Theme 1.0              //Theme

import LeftTemplate 1.0       //LeftTemplate
import "." as T               //needed for name clash with std Controls.Button

LeftTemplate {
  ColumnLayout {
    anchors.fill: parent

    Item {
      Layout.preferredHeight: Theme.minHeight * 0.005
    }

    Logo {
      Layout.fillWidth: true
      Layout.preferredHeight: Theme.minHeight * 0.15
    }

    Item {
      Layout.fillHeight: true
    }

    ButtonGroup {
      buttons: column.children
    }

    ColumnLayout {
      id: column
      Layout.fillWidth: true
      Layout.preferredHeight: Theme.minHeight * 0.4

      T.Button {
        id: dashboard
        Layout.fillWidth: true

        text: "Dashboard"
        code: "dashboard"
        image: "qrc:/qml/assets/ic_dashboard_black_24px.svg"
      }

      T.Button {
        id: files
        Layout.fillWidth: true

        text: "Files"
        code: "files"
        image: "qrc:/qml/assets/ic_insert_drive_file_black_24px.svg"
      }

      T.Button {
        id: hosting
        Layout.fillWidth: true

        text: "Hosting"
        code: "hosting"
        image: "qrc:/qml/assets/ic_folder_black_24px.svg"
      }

      T.Button {
        id: wallet
        Layout.fillWidth: true

        text: "Wallet"
        code: "wallet"
        image: "qrc:/qml/assets/ic_account_balance_wallet_black_24px.svg"

        checked: true
      }

      T.Button {
        id: terminal
        Layout.fillWidth: true

        text: "Terminal"
        code: "terminal"
        image: "qrc:/qml/assets/ic_code_black_24px.svg"
      }
    }

    Item {
      Layout.fillHeight: true
    }

    T.ProgressBar {
      Layout.fillWidth: true
      Layout.preferredHeight: Theme.minHeight * 0.125
    }
  }

  onClicked: {
    switch (cident) {
      case "dashboard":
        dashboard.clicked();
        break;

      case "files":
        files.clicked();
        break;

      case "hosting":
        hosting.clicked();
        break;

      case "wallet":
        wallet.clicked();
        break;

      case "terminal":
        terminal.clicked();
        break;
    }
  }
}
