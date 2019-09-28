import Felgo 3.0
import QtQuick 2.0

App {

  NavigationStack {

    Page {
      title: "My First App"

      AppText {
        text: "Welcome to Felgo"
        anchors.centerIn: parent
      }
    }
  }
}
