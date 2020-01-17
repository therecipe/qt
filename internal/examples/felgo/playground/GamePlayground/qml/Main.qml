import Felgo 3.0
import QtQuick 2.0


GameWindow {

  // this rectangle fills the whole screen with grey
  Rectangle {
    color: "#f0f0f0"
    anchors.fill: parent
  }

  Scene {

    Text {
      text: "Welcome to Felgo"

      font.pixelSize: 16
      anchors.centerIn: parent
    }
  }
}
