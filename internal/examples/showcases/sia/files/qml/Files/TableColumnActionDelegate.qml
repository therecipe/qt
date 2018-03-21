import QtQuick 2.7          //Component
import QtQuick.Controls 1.4 //TableViewColumn
import QtQuick.Controls 2.1 //Label

import Theme 1.0            //Theme
import FilesTemplate 1.0    //ActionButtonTemplate

import "."                  //ProgressBar

TableViewColumn {

  property QtObject tableView
  property ActionButtonTemplate template: ActionButtonTemplate{}

  title: role
  resizable: false
  movable: false
  width: parent.width * 0.25

  delegate: Component { Item {

    TableColumnDelegateButton {
      id: downloadButton

      anchors {
        top: parent.top
        left: parent.left
        bottom: parent.bottom
        }
      width: parent.width * 0.75

      source: "qrc:/qml/assets/ic_cloud_download_black_24px.svg"
      accent: Theme.accent
      onClicked: template.showDownload(tableView.model.data(tableView.model.index(styleData.row, 0), Qt.UserRole + 1))

      visible: styleData.value.available != null ? styleData.value.available : false
    }

    ProgressBar {
      id: progressBar

      anchors {
        verticalCenter: parent.verticalCenter
        left: parent.left
      }
      width: parent.width * 0.75
      height: parent.height * 0.6

      progressBarText: styleData.value.text != null ? styleData.value.text : ""
      value: styleData.value.value != null ? styleData.value.value : 0
      error: styleData.value.error != null ? styleData.value.error : false

      visible: !downloadButton.visible
    }

    TableColumnDelegateButton {
      anchors {
        top: parent.top
        left: downloadButton.visible ? downloadButton.right : progressBar.right
        right: parent.right
        bottom: parent.bottom
      }

      source: "qrc:/qml/assets/ic_delete_forever_black_24px.svg"
      accent: Theme.walletTableHighlight
      onClicked: template.deleteRequest(tableView.model.data(tableView.model.index(styleData.row, 0), Qt.UserRole + 1))
    }
  } }
}
