import QtQuick.Controls 1.4 //TextField
import QtQuick.Dialogs 1.2  //MessageDialog
import QtQuick.Layouts 1.3  //GridLayout

import Dialog 1.0

AddDialogController {

  MessageDialog {
    id: addDialogIssue
    title: "Add Album"
    text: "Please provide both the name of the artist and the title of the album."
    icon: StandardIcon.Information

    onAccepted: addDialog.visible = true
  }

  Dialog {
    id: addDialog
    width: 350
    height: 250
    property bool failedAccept

    GroupBox {
      anchors.fill: parent
      title: "Add Album"

      GridLayout {
        anchors.fill: parent

        Label {
          Layout.row: 0
          Layout.column: 0
          Layout.topMargin: 6
          Layout.leftMargin: 6
          text: "Artist:"
        }

        TextField {
          id: artistEditor

          Layout.row: 0
          Layout.column: 1
          Layout.topMargin: 6
          Layout.rightMargin: 6
          Layout.fillWidth: true
        }

        Label {
          Layout.row: 1
          Layout.column: 0
          Layout.leftMargin: 6
          text: "Title:"
        }

        TextField {
          id: titleEditor

          Layout.row: 1
          Layout.column: 1
          Layout.rightMargin: 6
          Layout.fillWidth: true
        }

        Label {
          Layout.row: 2
          Layout.column: 0
          Layout.leftMargin: 6
          text: "Year:"
        }

        SpinBox {
          id: yearEditor

          Layout.row: 2
          Layout.column: 1
          Layout.rightMargin: 6
          Layout.fillWidth: true

          minimumValue: 1900
          maximumValue: new Date().getFullYear()
          value: maximumValue
        }

        Label {
          Layout.row: 3
          Layout.column: 0
          Layout.leftMargin: 6
          Layout.rightMargin: 6
          Layout.fillWidth: true
          Layout.rowSpan: 1
          Layout.columnSpan: 2
          text: "Tracks (seperated by comma):"
        }

        TextField {
          id: tracksEditor

          Layout.row: 4
          Layout.column: 0
          Layout.leftMargin: 6
          Layout.rightMargin: 6
          Layout.bottomMargin: 6
          Layout.fillWidth: true
          Layout.rowSpan: 1
          Layout.columnSpan: 2
        }
      }
    }

    standardButtons: StandardButton.Close | StandardButton.Reset | StandardButton.Save

    onRejected: {
      addDialog.reset()
    }

    onAccepted: {
      if (artistEditor.text == "" || titleEditor.text == "") {
        failedAccept = true
      } else {
        addAlbum(artistEditor.text, titleEditor.text, yearEditor.value, tracksEditor.text)
        addDialog.reset()
      }
    }

    onReset: {
      artistEditor.text = ""
      titleEditor.text = ""
      yearEditor.value = yearEditor.maximumValue
      tracksEditor.text = ""
    }

    onFailedAcceptChanged: {
      if (failedAccept) {
        failedAccept = false
        Qt.callLater(function() { addDialogIssue.visible = true })
      }
    }
  }

  onAddAlbumShowRequest: addDialog.visible = true
}
