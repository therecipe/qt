import QtQuick 2.0
import Felgo 3.0

NotificationBar {
  z: 1
  y: -height
  width: parent.width // parent should be game window, bar fills window then
  height: Theme.statusBarHeight + dp(60)
  enabled: visible
  visible: false // notification bar will be set visible when display(...) is called

  // duration that notification will show, in ms
  property int duration: 3000

  // colors of notification bar
  property color tintColor: "#f05352" // default value, will be set to MultiplayerView::tintColor when used by FelgoMultiplayer
  property color backgroundColor: Qt.rgba(tintColor.r, tintColor.g, tintColor.b, 0.75)
  property color textColor: "white"

  // content of last received notification
  property string lastNotificationMessage: ""
  property var lastNotificationData: []
  property bool lastNotificationIsActive: false

  // is emitted when user manually dismisses the notification by clicking
  signal dismiss(string message, var additionalData, bool isActive)

  // display function is called when FelgoMultiplayer receives a notification
  onDisplay: {
    console.debug("Display notification "+message+"; data: "+JSON.stringify(additionalData)+"; isActive: "+isActive)

    // memorize notification
    lastNotificationMessage = message
    lastNotificationData = additionalData
    lastNotificationIsActive = isActive

    // set notification text
    notificationMessage.text = message
    notificationTitle.text = additionalData.title

    // show notification bar
    y = 0
    visible = true
    hideAfterDelay.start()
  }

  // background
  Rectangle {
    anchors.fill: parent
    color: parent.backgroundColor
  }

  // text
  Text {
    // set font
    id: notificationTitle
    font.pixelSize: sp(18)
    font.bold: true
    color: parent.textColor
    text: ""
    y: dp(8) + Theme.statusBarHeight
    x: dp(8)
  }

  Text {
    // set font
    id: notificationMessage
    font.pixelSize: sp(15)
    x: dp(8)
    color: parent.textColor
    text: ""
    anchors.top: notificationTitle.bottom
  }

  // handle clicks for dismiss
  MouseArea {
    anchors.fill: parent
    onClicked: {
      hide()
      dismiss(lastNotificationMessage, lastNotificationData, lastNotificationIsActive)
    }
  }

  // hide automatically after interval passes
  Timer {
    id: hideAfterDelay
    interval: duration
    onTriggered: {
      hide()
    }
  }

  // animate the y property to let the notification bar slide in and out
  Behavior on y {
    SequentialAnimation {
      // first do animation
      NumberAnimation {
        duration: 600
        easing.type: Easing.InOutCubic
      }
      // if animation completed and it was a slide-out, reset notification bar
      ScriptAction {
        script: {
          if(y <= -height) {
            notificationMessage.text = ""
            notificationTitle.text = ""
            visible = false
          }
        }
      }
    }
  }

  // hide function slides out the notification bar
  function hide() {
    hideAfterDelay.stop()
    y =- height
  }
}
