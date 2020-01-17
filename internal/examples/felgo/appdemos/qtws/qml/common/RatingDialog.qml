import QtQuick 2.0
import QtQuick.Controls 1.4
import QtQuick.Controls.Styles 1.4
import Felgo 3.0
import "../common"

// feedback window to contact Felgo
Dialog {
  id: ratingDialog
  negativeAction: true
  negativeActionLabel: "Close"
  positiveActionLabel: "Rate"
  autoSize: true
  outsideTouchable: false

  property string ratingUrl: system.isPlatform(System.IOS) ? "itms-apps://itunes.apple.com/at/app/id1286758361?mt=8" :
                             system.isPlatform(System.Android) ? "http://play.google.com/store/apps/details?id=net.vplay.demos.qtws2017" :
                             "https://felgo.com/qws-conference-in-app-2017"

  onCanceled: ratingDialog.close()
  onAccepted: {
    amplitude.logEvent("RateInStore")
    logic.setFeedBackSent(true)

    // open the store site to rate the game instead
    nativeUtils.openUrl(ratingUrl)
    ratingDialog.close()
  }

  Item {
    id: contentArea
    width: parent.width
    height: content.height

    Column {
      id: content
      width: parent.width
      anchors.horizontalCenter: parent.horizontalCenter
      anchors.top: parent.top
      anchors.margins: dp(20)
      spacing: dp(10)

      // rating header
      Text {
        id: ratingText
        horizontalAlignment: Text.AlignHCenter
        anchors.horizontalCenter: parent.horizontalCenter
        text: "Rate This App"
        color: Theme.textColor
        font.pixelSize: sp(18)
        width: parent.width * 0.8//- anchors.topMargin * 2
        wrapMode: Text.Wrap
      }

      // rating note
      Text {
        id: ratingNote
        horizontalAlignment: Text.AlignHCenter
        anchors.horizontalCenter: parent.horizontalCenter
        text: "Support us by rating the app in the store!"
        color: Theme.textColor
        font.pixelSize: sp(10)
        width: parent.width * 0.8//- 20
        wrapMode: Text.Wrap
      }

      AppButton {
        text: "I already rated this app"
        flat: true
        anchors.horizontalCenter: parent.horizontalCenter
        textSize: sp(12)
        onClicked: {
          logic.setFeedBackSent(true)

          // close the window
          ratingDialog.close()
        }
      }

      // spacer
      Item {
        width: parent.width
        height: parent.spacing
      }
    }

  } // content area
}
