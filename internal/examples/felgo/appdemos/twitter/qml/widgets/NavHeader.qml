import QtQuick 2.0
import Felgo 3.0

import "../model"

Item {
  width: parent.width
  height: dp(120)

  property var profile: dataModel.currentProfile || null

  RoundedImage {
    id: bgImage
    backgroundColor: profile ? ("#" + profile.profile_background_color) : Theme.tintColor
    radius: 0
    anchors.fill: parent
    fillMode: Image.PreserveAspectCrop
    border.width: 0
    source: profile ? profile.profile_banner_url : ""

    Column {
      anchors.centerIn: parent
      width: parent.width

      // Profile image
      RoundedImage {
        id: profileImage
        source: profile && profile.profile_image_url ? profile.profile_image_url.replace("_normal", "_bigger") : ""

        width: dp(50)
        height: dp(50)

        border.color: "white"
        border.width: dp(2)
        backgroundColor: "white"
        anchors.horizontalCenter: parent.horizontalCenter
      }

      Item { height: dp(5); width: parent.width } // spacer

      AppText {
        width: parent.width
        horizontalAlignment: Text.Center
        color: "white"
        text: profile ? profile.name : ""
        font.pixelSize: sp(14)
        font.bold: true
        font.family: Theme.boldFont.name
      }

      AppText {
        width: parent.width
        text: profile ? ("@" + profile.screen_name) : ""
        horizontalAlignment: Text.Center
        color: "white"
        font.pixelSize: sp(13)
      }
    }
  }
}
