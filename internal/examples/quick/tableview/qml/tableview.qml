import QtQuick 2.6
import QtQuick.Layouts 1.1
import QtQuick.Controls 1.4

Item {
  width: 400
  height: 300

  ListModel {
    id: libraryModel

    ListElement {
      title: "A Masterpiece"
      author: "Gabriel"
    }

    ListElement {
      title: "Brilliance"
      author: "Jens"
    }

    ListElement {
      title: "Outstanding"
      author: "Frederik"
    }
  }

  Connections {
    target: QmlBridge
    onAddItem: libraryModel.append({"title": title, "author": author})
    onRemoveItem: libraryModel.remove(libraryModel.count-1)
  }

  Connections {
    target: libraryModel
    onDataChanged: { //ListModel inherits core.QAbstractItemModel signals http://doc.qt.io/qt-5/qabstractitemmodel.html#dataChanged
      QmlBridge.editedModel(topLeft.row, roles, libraryModel.get(topLeft.row))
    }
  }

  TableView {
    id: tableView
    anchors.fill: parent

    selectionMode: SelectionMode.NoSelection

    TableViewColumn {
      role: "title"
      title: "Title"
      width: 100

      delegate: TextEdit {
        id: titleDelegate
        text: styleData.value

        Keys.onReturnPressed: titleDelegate.focus = false

        onEditingFinished: {
          libraryModel.setProperty(tableView.currentRow, "title", titleDelegate.text)
        }
      }
    }

    TableViewColumn {
      role: "author"
      title: "Author"
      width: 100

      delegate: TextEdit {
        id: authorDelegate
        text: styleData.value

        Keys.onReturnPressed: authorDelegate.focus = false

        onEditingFinished: {
          libraryModel.setProperty(tableView.currentRow, "author", authorDelegate.text)
        }
      }
    }

    model: libraryModel
  }
}
