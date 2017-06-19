import QtQuick.Controls 1.4 //TableView

import Album 1.0

AlbumController {
  GroupBox {
    anchors.fill: parent

    title: "AlbumNew"
    TableView {
      id: tableView

      anchors.fill: parent
      anchors.margins: 5

      sortIndicatorVisible: true
      onSortIndicatorColumnChanged: sortTableView(sortIndicatorColumn, sortIndicatorOrder)
      onSortIndicatorOrderChanged: sortTableView(sortIndicatorColumn, sortIndicatorOrder)

      onActivated: showAlbumDetails(model.index(row, 0))
      onClicked: showAlbumDetails(model.index(row, 0))

      model: viewModel

      TableViewColumn { role: "ID"; title: "ID"; width: 100; visible: false}
      TableViewColumn { role: "Title"; title: "Title"; width: 100 }
      TableViewColumn { role: "Artist"; title: "Artist"; width: 100 }
      TableViewColumn { role: "Year"; title: "Year"; width: 100 }
    }
  }

  onAlbumAdded: {
    tableView.selection.clear()
    tableView.selection.select(tableView.rowCount - 1)
    showAlbumDetails(tableView.model.index(tableView.rowCount - 1, 0))
  }

  onDeleteAlbumRequest: {
    var index = tableView.model.index(tableView.currentRow, 0)
    var title = tableView.model.data(index, Qt.UserRole + 2)
    var artist = tableView.model.data(index, Qt.UserRole + 3)
    deleteAlbumShowRequest(title, artist)
  }

  onDeleteAlbumCommand: {
    deleteAlbum(tableView.model.index(tableView.currentRow, 0))
    showImageLabel()
  }
}
