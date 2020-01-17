import QtQuick 2.4
import "pages"
import "model"

Item {
  anchors.fill: parent

  Component.onCompleted: logic.fetchTwitterData()

  DataModel {
    id: dataModel
    dispatcher: logic
  }

  Logic {
    id: logic
  }

  TwitterMainPage { }
}
