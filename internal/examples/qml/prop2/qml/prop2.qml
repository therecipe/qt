import QtQuick 2.5
import QtQuick.Controls 2.0

ApplicationWindow {
  visible: true

  id: myRect

  Text {
    anchors.centerIn: parent
    text: "click me!"
  }

  MouseArea {
    id: mouseArea
    anchors.fill: parent
    onClicked: {
      console.log(bridge.boolProp)
      console.log(bridge.intProp)
      console.log(bridge.stringProp)
      console.log(bridge.stringListProp)
      console.log(bridge.objectProp.objectName)

      bridge.boolProp = false
      bridge.intProp = 321
      bridge.stringProp = "hello from qml"
      bridge.stringListProp = ["hello","from","qml"]
      bridge.objectProp.objectName = "newQmlObjectName"
    }
  }

  Connections {
    target: bridge

    onBoolPropChanged: console.log("changed bool ->", bridge.boolProp)
    onIntPropChanged: console.log("changed int ->", bridge.intProp)
    onStringPropChanged: console.log("changed string ->", bridge.stringProp)
    onStringListPropChanged: console.log("changed stringlist ->", bridge.stringListProp)
    onObjectPropChanged: console.log("changed object ->", bridge.objectProp.objectName)
  }
}
