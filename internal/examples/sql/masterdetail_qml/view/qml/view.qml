import QtQuick.Controls 1.4 //Menu
import QtQuick.Dialogs 1.2  //MessageDialog
import QtQuick.Layouts 1.3  //GridLayout

import View 1.0
import Artist 1.0
import Album 1.0
import Detail 1.0
import Dialog 1.0

ApplicationWindow {
  id: app
  visible: true
  title: "Music Archive"
  minimumWidth: 850; minimumHeight: 400
  width: minimumWidth; height: minimumHeight
  property ViewController core: ViewController{}

  AddDialog {}
  DeleteDialog {}

  MessageDialog {
    id: aboutDialog
    title: "About Music Archive"
    text: "<p>The <b>Music Archive</b> example shows how to present data from different data sources in the same application. The album titles, and the corresponding artists and release dates, are kept in a database, while each album's tracks are stored in an XML file. </p><p>The example also shows how to add as well as remove data from both the database and the associated XML file using the API provided by the Qt SQL and Qt XML modules, respectively.</p>"
  }

  menuBar: MenuBar {
    Menu {
      title: "&File"
      MenuItem { text: "&Add album..."; shortcut: "Ctrl+A"; onTriggered: core.addAlbumShowRequest() }

      MenuItem {
        text: "&Delete album..."
        shortcut: "Ctrl+D"
        onTriggered: core.deleteAlbumRequest()
      }

      MenuSeparator {}

      MenuItem { text: "&Quit"; shortcut: "Ctrl+Q"; onTriggered: Qt.quit() }
    }

    Menu {
      title: "&Help"
      MenuItem { text: "&About"; onTriggered: aboutDialog.visible = true }
      MenuItem { text: "About &Qt"; onTriggered: core.aboutQt() }
    }
  }

  GridLayout {
    anchors.fill: parent
    Artist {
      Layout.row: 0
      Layout.column: 0

      Layout.minimumWidth: 500
      Layout.preferredHeight: 75

      Layout.leftMargin: 12
      Layout.rightMargin: 5
    }

    Album {
      Layout.row: 1
      Layout.column: 0

      Layout.preferredWidth: 500
      Layout.fillHeight: true

      Layout.leftMargin: 12
      Layout.rightMargin: 5
      Layout.bottomMargin: 12
    }

    Detail {
      Layout.row: 0
      Layout.column: 1
      Layout.rowSpan: 2
      Layout.columnSpan: 1

      Layout.fillWidth: true
      Layout.fillHeight: true

      Layout.leftMargin: 6
      Layout.rightMargin: 15
      Layout.bottomMargin: 12
    }
  }
}
