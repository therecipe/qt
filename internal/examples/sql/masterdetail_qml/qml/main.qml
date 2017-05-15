import QtQuick 2.0
import QtQuick.Controls 1.4
import QtQuick.Dialogs 1.2
import QtQuick.Layouts 1.3

ApplicationWindow {
  id: app
  visible: true
  title: "Music Archive"
  minimumWidth: 850; minimumHeight: 400

  MessageDialog {
    id: aboutDialog
    title: "About Music Archive"
    text: "<p>The <b>Music Archive</b> example shows how to present data from different data sources in the same application. The album titles, and the corresponding artists and release dates, are kept in a database, while each album's tracks are stored in an XML file. </p><p>The example also shows how to add as well as remove data from both the database and the associated XML file using the API provided by the Qt SQL and Qt XML modules, respectively.</p>"
  }

  MessageDialog {
    id: deleteDialogDismissed
    title: "Delete Album"
    text: "Select the album you want to delete."
    icon: StandardIcon.Information
  }

  MessageDialog {
    id: deleteDialog
    title: "Delete Album"
    text: ""
    icon: StandardIcon.Question
    standardButtons: StandardButton.Yes | StandardButton.No

    onYes: {
      controller.deleteAlbum(tableView.model.index(tableView.currentRow, 0))
      showImageLabel()
    }
    onNo: deleteDialogDismissed.visible = true
  }

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
        controller.addAlbum(artistEditor.text, titleEditor.text, yearEditor.value, tracksEditor.text)
        tableView.selection.clear()
        tableView.selection.select(tableView.rowCount - 1)
        controller.showAlbumDetails(tableView.model.index(tableView.rowCount - 1, 0))
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

  menuBar: MenuBar {
    Menu {
      title: "&File"
      MenuItem { text: "&Add album..."; shortcut: "Ctrl+A"; onTriggered: addDialog.visible = true }

      MenuItem {
        text: "&Delete album..."
        shortcut: "Ctrl+D"
        onTriggered: {
          var index = tableView.model.index(tableView.currentRow, 0)
          var title = tableView.model.data(index, Qt.UserRole + 2)
          var artist = tableView.model.data(index, Qt.UserRole + 3)

          deleteDialog.text = "Are you sure you want to delete '"+title+"' by '"+artist+"'?"
          deleteDialog.visible = true
        }
      }

      MenuSeparator {}

      MenuItem { text: "&Quit"; shortcut: "Ctrl+Q"; onTriggered: Qt.quit() }
    }

    Menu {
      title: "&Help"
      MenuItem { text: "&About"; onTriggered: aboutDialog.visible = true }
      MenuItem { text: "About &Qt"; onTriggered: controller.aboutQt() }
    }
  }

  GridLayout {
    anchors.fill: parent

    GroupBox {
      Layout.row: 0
      Layout.column: 0
      Layout.minimumWidth: 500
      Layout.leftMargin: 12
      Layout.rightMargin: 5

      title: "Artist"
      ComboBox {
        anchors.fill: parent
        anchors.margins: 5

        id: artistView
        model: listModel
        textRole: "display"

        onActivated: controller.changeArtist(index)
      }
    }

    GroupBox {
      Layout.row: 1
      Layout.column: 0
      Layout.minimumWidth: 500
      Layout.fillHeight: true
      Layout.leftMargin: 12
      Layout.rightMargin: 5
      Layout.bottomMargin: 12

      title: "Album"
      TableView {
        id: tableView

        anchors.fill: parent
        anchors.margins: 5

        sortIndicatorVisible: true
        onSortIndicatorColumnChanged: controller.sortTableView(sortIndicatorColumn, sortIndicatorOrder)
        onSortIndicatorOrderChanged: controller.sortTableView(sortIndicatorColumn, sortIndicatorOrder)

        onActivated: controller.showAlbumDetails(model.index(row, 0))
        onClicked: controller.showAlbumDetails(model.index(row, 0))

        model: viewModel

        TableViewColumn { role: "ID"; title: "ID"; width: 100; visible: false}
        TableViewColumn { role: "Title"; title: "Title"; width: 100 }
        TableViewColumn { role: "Artist"; title: "Artist"; width: 100 }
        TableViewColumn { role: "Year"; title: "Year"; width: 100 }
      }
    }

    GroupBox {
      Layout.row: 0
      Layout.column: 1
      Layout.rowSpan: 2
      Layout.columnSpan: 1
      Layout.fillWidth: true
      Layout.fillHeight: true
      Layout.leftMargin: 6
      Layout.rightMargin: 15
      Layout.bottomMargin: 12

      title: "Details"
      GridLayout {
        anchors.fill: parent

        Image {
          id: imageLabel

          Layout.row: 0
          Layout.column: 0

          horizontalAlignment: Image.AlignHCenter
          verticalAlignment: Image.AlignVCenter

          source: "../images/image.png"
        }

        Image {
          id: iconLabel

          Layout.row: 0
          Layout.column: 1

          horizontalAlignment: Image.AlignBottom | Image.AlignRight
          verticalAlignment: Image.AlignBottom | Image.AlignRight

          source: "../images/icon.png"
        }

        Label {
          id: profileLabel

          Layout.row: 1
          Layout.column: 0
          Layout.columnSpan: 2

          wrapMode: Text.WordWrap
          horizontalAlignment: Text.AlignBottom
          verticalAlignment: Text.AlignBottom
        }

        Label {
          id: titleLabel

          Layout.row: 2
          Layout.column: 0
          Layout.columnSpan: 2

          wrapMode: Text.WordWrap
          horizontalAlignment: Text.AlignBottom
          verticalAlignment: Text.AlignBottom
        }

        TableView {
          id: trackList

          Layout.row: 3
          Layout.column: 0
          Layout.columnSpan: 2

          alternatingRowColors: false
          headerVisible: false

          model: ListModel {}

          TableViewColumn { role: "display" }
        }
      }
    }
  }

  function showImageLabel() {
    profileLabel.visible = false
    titleLabel.visible = false
    iconLabel.visible = false
    trackList.visible = false

    imageLabel.visible = true
  }

  function showArtistProfile(profileLabelText) {
    profileLabel.text = profileLabelText

    profileLabel.visible = true
    iconLabel.visible = true

    titleLabel.visible = false
    trackList.visible = false
    imageLabel.visible = false
  }

  function showTitleAndAlbumDetails(title, elements) {
    titleLabel.text = title
    titleLabel.visible = true

    trackList.model.clear()
    for (var i = 0; i < elements.length; i++) {
      trackList.model.append({"display":elements[i]})
    }

    if (trackList.model.count != 0) {
      trackList.visible = true
    }
  }

  function resetComboBoxModel() {
    artistView.model = null
    artistView.model = listModel
  }

  Connections {
    target: controller
    onShowImageLabel: showImageLabel()
    onShowArtistProfile: showArtistProfile(profileLabelText)
    onShowTitleAndAlbumDetails: showTitleAndAlbumDetails(title, elements)
    onResetComboBoxModel: resetComboBoxModel()
  }

  Component.onCompleted: showImageLabel()
}
