import QtQuick.Controls 1.4 //ComboBox

import Artist 1.0

ArtistController {
  GroupBox {
    anchors.fill: parent
    title: "Artist"

    ComboBox {
      id: artistView
      anchors.fill: parent
      anchors.margins: 5

      model: listModel
      textRole: "display"

      onActivated: changeArtist(index)
    }
  }
}
