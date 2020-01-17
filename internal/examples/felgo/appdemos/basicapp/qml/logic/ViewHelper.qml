import QtQuick 2.0

Item {

  // function to format the todo title for display
  // appends draft number for drafts
  function formatTitle(todo) {
    if(!todo)
      return ""

    return qsTr("Todo") + " " + todo.id
  }


}
