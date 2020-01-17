import QtQuick 2.0
import QtQuick.Controls 1.2
import Felgo 3.0
import "../common"

FlickablePage {

  flickable.contentHeight: content.height

  Column {
    id: content
    width: parent.width

    // storage
    SectionHeader { icon: IconType.database; text: "Local Storage" }
    SectionDescription { text: "Read and write local SQLite databases or use a simple key-value based storage." }
    SectionContent { contentItem: Row {
        anchors.horizontalCenter: parent.horizontalCenter
        spacing: dp(10)
        AppText { text: "Counter: "+localStorage.counter; anchors.verticalCenter: parent.verticalCenter }
        AppButton { flat: false; text: "Increase"; onClicked: localStorage.increaseCounter(); anchors.verticalCenter: parent.verticalCenter }
      }
      Storage {
        id: localStorage
        property int counter: 0
        Component.onCompleted: counter = localStorage.getValue("counter") || 0
        function increaseCounter() {
          counter++
          localStorage.setValue("counter", counter)
        }
      }
    }

    // sensors
    SectionSpacer { }
    SectionHeader {
      icon: IconType.eye
      text: "Sensors"
    }

    SectionDescription {
      text: "The Sensors API provides access to sensor hardware via QML and C++ interfaces. It also provides a motion gesture recognition API for devices."
    }

    SectionContent {
      contentItem: Column {
        anchors.centerIn: parent
        width: parent.width
        AppText {
          text: "features"
          font.pixelSize: sp(10)
          anchors.horizontalCenter: parent.horizontalCenter
          horizontalAlignment: Text.AlignHCenter
        }

        AppText {
          text: "Accelerometer
Altimeter
Compass
Gyroscope
Light
Orientation
Rotation
Tilt"
          anchors.horizontalCenter: parent.horizontalCenter
          horizontalAlignment: Text.AlignHCenter
        }
      }
    }

    // spacer
    SectionSpacer { }

    // communication
    SectionHeader {
      icon: IconType.wifi
      text: "Connectivity"
    }

    SectionDescription {
      text: "Felgo supports many different forms of network communication, for example a simple integration of services using XmlHttpRequests."
    }

    SectionContent {
      contentItem: Column {
        anchors.centerIn: parent
        width: parent.width
        AppText {
          text: "features"
          font.pixelSize: sp(10)
          anchors.horizontalCenter: parent.horizontalCenter
          horizontalAlignment: Text.AlignHCenter
        }

        AppText {
          text: "TCP/UDP
WebSockets
XmlHttpRequests
Bluetooth
Bluetooth LE
NFC"
          anchors.horizontalCenter: parent.horizontalCenter
          horizontalAlignment: Text.AlignHCenter
        }
      }
    }

  }
}
