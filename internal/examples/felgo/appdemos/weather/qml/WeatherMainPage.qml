import QtQuick 2.4
import QtQuick.Controls 1.2

import Felgo 3.0

Page {
  id: page
  property var today: alldata[currentIndex]
  property int currentIndex: 0

  property var alldata: [
    { temperature: 12, description: qsTr("Foggy") },
    { temperature: 32, description: qsTr("Sunny") }
  ]

  // Background
  Rectangle {
    x: -page.safeArea.x
    y: -page.safeArea.y
    width: page.width
    height: page.height

    gradient: Gradient {
      GradientStop {
        position: 0
        color: today.temperature < 20 ? "#1AD6FD" : "#FF5E3A"

        Behavior on color { ColorAnimation { duration: 1500 } }
      }
      GradientStop {
        position: 1
        color: today.temperature < 20 ? "#1D62F0" : "#FF2A68"

        Behavior on color { ColorAnimation { duration: 1000 } }
      }
    }
  }

  // Top content
  Column {
    anchors.horizontalCenter: parent.horizontalCenter
    y: dp(10)
    spacing: dp(10)

    // Current time
    AppText {
      id: timeLabel

      font.pixelSize: sp(14)
      anchors.horizontalCenter: parent.horizontalCenter

      Timer {
        running: true
        interval: 1000 * 30
        triggeredOnStart: true
        repeat: true
        onTriggered: {
          timeLabel.text = new Date().toLocaleTimeString(Qt.locale(), Locale.ShortFormat)
        }
      }
    }

    // City
    AppText {
      text: "Vienna, AT"
      font.pixelSize: sp(22)
      anchors.horizontalCenter: parent.horizontalCenter
    }

  }

  // Centered content
  Column {
    id: col
    anchors.centerIn: parent

    // Temperature
    AppText {
      id: tempText

      property int temperature: today.temperature

      Component.onCompleted: text = temperature

      font.pixelSize: sp(140)
      anchors.horizontalCenter: parent.horizontalCenter

      onTemperatureChanged: {
        var currentTemp = parseInt(tempText.text)
        textTimer.restart()
      }

      Timer {
        id: textTimer
        interval: 40

        onTriggered: {
          // Check if we have to change the text
          var currentTemp = parseInt(tempText.text)

          if (tempText.temperature > currentTemp) {
            tempText.text = ++currentTemp
            restart()
          }
          else if (tempText.temperature < currentTemp) {
            tempText.text = --currentTemp
            restart()
          }
        }
      }
    }

    // Description
    AppText {
      id: descText
      text: today.description
      font.pixelSize: sp(28)
      anchors.horizontalCenter: parent.horizontalCenter

      Behavior on text {
        SequentialAnimation {
          NumberAnimation { target: descText; property: "opacity"; to: 0 }
          PropertyAction { }
          NumberAnimation { target: descText; property: "opacity"; to: 1 }
        }
      }
    }
  }

  // Bottom content
  Grid {
    id: bottomGrid

    width: Math.min(parent.width - dp(20), dp(450))
    anchors.horizontalCenter: parent.horizontalCenter
    y: parent.height - height - dp(10)
    columns: 5

    Repeater {
      model: [
        { day: Qt.locale().dayName(0, Locale.ShortFormat), high: 34, low: 24, icon: IconType.suno },
        { day: Qt.locale().dayName(1, Locale.ShortFormat), high: 32, low: 20, icon: IconType.suno },
        { day: Qt.locale().dayName(2, Locale.ShortFormat), high: 33, low: 20, icon: IconType.suno },
        { day: Qt.locale().dayName(3, Locale.ShortFormat), high: 28, low: 21, icon: IconType.suno },
        { day: Qt.locale().dayName(4, Locale.ShortFormat), high: 24, low: 16, icon: IconType.suno }
      ]

      Column {
        width: bottomGrid.width / 5
        spacing: dp(5)

        Icon {
          icon: modelData.icon
          anchors.horizontalCenter: parent.horizontalCenter
          size: dp(20)
        }

        Item {
          width: 1
          height: dp(5)
        }

        AppText {
          text: modelData.high
          anchors.horizontalCenter: parent.horizontalCenter
          font.pixelSize: sp(14)
        }

        AppText {
          text: modelData.low
          color: "#aaffffff"
          anchors.horizontalCenter: parent.horizontalCenter
          font.pixelSize: sp(14)
        }

        Item {
          width: 1
          height: dp(5)
        }

        AppText {
          text: modelData.day
          anchors.horizontalCenter: parent.horizontalCenter
          font.pixelSize: sp(14)
        }
      }
    }
  }

  Timer {
    interval: 1000
    repeat: true
    running: true
    onTriggered: {
      interval = 3500
      if (currentIndex < alldata.length - 1)
        currentIndex++
      else
        currentIndex = 0
    }
  }
}
