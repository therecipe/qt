import QtQuick 2.7
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
      famous: true
    }

    ListElement {
      title: "Brilliance"
      author: "Jens"
      famous: true
    }

    ListElement {
      title: "Outstanding"
      author: "Frederik"
      famous: true
    }
  }

  Connections {
    target: QmlBridge
    onAddItem: libraryModel.append({"title": title, "author": author, "famous": famous})
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
        id: titleDelgate
        text: styleData.value

        Keys.onReturnPressed: titleDelgate.focus = false

        onEditingFinished: {
          libraryModel.setProperty(tableView.currentRow, "title", titleDelgate.text)
        }
      }
    }

    TableViewColumn {
      role: "author"
      title: "Author"
      width: 100

      delegate: TextEdit {
        id: titleDelgate
        text: styleData.value

        Keys.onReturnPressed: titleDelgate.focus = false

        onEditingFinished: {
          libraryModel.setProperty(tableView.currentRow, "author", titleDelgate.text)
        }
      }
    }

       TableViewColumn {
           title: "Famous"
           role: "famous"
           width: 100

           delegate: CheckBox {
               id: famousDelegate
               anchors.fill: parent
               checked: styleData.value
               enabled: true

               onClicked: {
                   libraryModel.setProperty(styleData.row, "famous", famousDelegate.checked)
               }
           }

       }

    model: libraryModel
  }
}
