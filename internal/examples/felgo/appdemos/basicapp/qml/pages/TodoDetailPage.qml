import Felgo 3.0
import QtQuick 2.0

Page {
  id: page
  title: viewHelper.formatTitle(todoData)

  // network activity indicator
  rightBarItem: ActivityIndicatorBarItem {
    id: activityBarItem
    visible: dataModel.isBusy && !dataModel.isStoringTodos
    showItem: showItemAlways // do not collapse into sub-menu on Android
  }

  // target id
  property int todoId: 0

  // data property for page
  property var todoData: dataModel.todoDetails[todoId]

  // load data initially or when id changes
  onTodoIdChanged: {
    logic.fetchTodoDetails(todoId)
  }

  // column to show all todo object properties, if data is available
  Column {
    y: spacing
    width: parent.width - 2 * spacing
    anchors.horizontalCenter: parent.horizontalCenter
    spacing: dp(Theme.navigationBar.defaultBarItemPadding)
    visible: !noDataMessage.visible

    // Repeater creates copies of given item based on configured model data
    Repeater {
      enabled: parent.visible
      model: !!todoData ? Object.keys(todoData) : undefined

      // Text Item to show each property - value pair
      AppText {
        property string propName: modelData
        property string value: todoData[propName]

        width: parent.width
        anchors.horizontalCenter: parent.horizontalCenter
        height: implicitHeight

        text: "<strong>"+propName+":</strong> "+value
        wrapMode: AppText.WrapAtWordBoundaryOrAnywhere
      }
    }
  }

  // show message if data not available
  AppText {
    id: noDataMessage
    anchors.verticalCenter: parent.verticalCenter
    text: qsTr("Todo data not available. Please check your internet connection.")
    width: parent.width
    horizontalAlignment: Qt.AlignHCenter
    visible: !todoData && !dataModel.isBusy
  }
}
