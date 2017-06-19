import QtQuick 2.7          //Image
import QtQuick.Controls 1.4 //TableView
import QtQuick.Layouts 1.3  //GridLayout

import Detail 1.0

DetailController {
  GroupBox {
    anchors.fill: parent

    title: "Details"
    GridLayout {
      anchors.fill: parent

      Image {
        id: imageLabel

        Layout.row: 0
        Layout.column: 0

        horizontalAlignment: Image.AlignHCenter
        verticalAlignment: Image.AlignVCenter

        source: "qrc:/images/image.png"
      }

      Image {
        id: iconLabel

        Layout.row: 0
        Layout.column: 1

        horizontalAlignment: Image.AlignBottom | Image.AlignRight
        verticalAlignment: Image.AlignBottom | Image.AlignRight

        source: "qrc:/images/icon.png"
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
        Layout.maximumWidth: parent.width * 0.5

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

   onShowImageLabel: {
    profileLabel.visible = false
    titleLabel.visible = false
    iconLabel.visible = false
    trackList.visible = false

    imageLabel.visible = true
  }

   onShowArtistProfile: {
    profileLabel.text = profileLabelText

    profileLabel.visible = true
    iconLabel.visible = true

    titleLabel.visible = false
    trackList.visible = false
    imageLabel.visible = false
  }

   onShowTitleAndAlbumDetails: {
    titleLabel.text = title
    titleLabel.visible = true
    trackList.visible = true

    trackList.model.clear()
    for (var i = 0; i < elements.length; i++) {
      trackList.model.append({"display":elements[i]})
    }
  }

  Component.onCompleted: showImageLabel()
}
