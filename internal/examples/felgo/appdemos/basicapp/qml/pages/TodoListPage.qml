import Felgo 3.0
import QtQuick 2.0

Page {
  id: page
  title: qsTr("Todo List") // use qsTr for strings you want to translate

  rightBarItem: NavigationBarRow {
    // network activity indicator
    ActivityIndicatorBarItem {
      enabled: dataModel.isBusy
      visible: enabled
      showItem: showItemAlways // do not collapse into sub-menu on Android
    }

    // add new todo
    IconButtonBarItem {
      icon: IconType.plus
      showItem: showItemAlways
      onClicked: {
        // use qsTr for strings you want to translate
        var title = qsTr("New Todo")

        // this logic helper function creates a todo
        logic.addTodo(title)
      }
    }
  }

  // when a todo is added, we open the detail page for it
  Connections {
    target: dataModel
    onTodoStored: {
      page.navigationStack.popAllExceptFirstAndPush(detailPageComponent, { todoId: todo.id })
    }
  }

  // JsonListModel
  // A ViewModel for JSON data that offers best integration and performance with list views
  JsonListModel {
    id: listModel
    source: dataModel.todos // show todos from data model
    keyField: "id"
    fields: ["id", "title", "completed"]
  }

  // show sorted/filterd todos of data model
  AppListView {
    id: listView
    anchors.fill: parent

    // the model specifies the data for the list view
    model: listModel

    // the delegate is the template item for each entry of the list
    delegate: SimpleRow {
      text: viewHelper.formatTitle(model)

      // push detail page when selected, pass chosen todo id
      onSelected: page.navigationStack.popAllExceptFirstAndPush(detailPageComponent, { todoId: model.id })
    }
  }

  // component for creating detail pages
  Component {
    id: detailPageComponent
    TodoDetailPage { }
  }
}
