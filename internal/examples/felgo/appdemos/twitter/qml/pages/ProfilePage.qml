import QtQuick 2.3
import Felgo 3.0

import QtGraphicalEffects 1.0

import "../model"
import "../widgets"

ListPage {
  id: profilePage

  title: qsTr("Profile")
  backgroundColor: "white"

  navigationBarTranslucency: 1.0 // navigation bar is 100 percent translucent for this page
  useSafeArea: false // fill whole screen

  titleItem: Column {
    // Fade title item in depdending on scroll state
    opacity: 1 - profileImage.opacity

    AppText {
      anchors.horizontalCenter: Theme.isAndroid ? undefined : parent.horizontalCenter
      text: profilePage.profile.name
      color: "white"
      font.family: Theme.boldFont.name
      font.bold: true
      font.pixelSize: dp(14)
    }

    AppText {
      anchors.horizontalCenter: Theme.isAndroid ? undefined : parent.horizontalCenter
      text: qsTr("%1 Tweets").arg(profilePage.profile.statuses_count.toLocaleString(Qt.locale(), "f", 0))
      color: "white"
      font.pixelSize: dp(10)
    }
  }

  // The actual profile data to display
  property var profile

  readonly property real barHeight: dp(Theme.navigationBar.height) + (nativeUtils.safeAreaInsets.top > Theme.statusBarHeight ? nativeUtils.safeAreaInsets.top : Theme.statusBarHeight)

  Component { id: detailPage; DetailPage {} }

  model: dataModel.timeline
  delegate: TweetRow {
    id: row

    onSelected: {
      console.debug("Selected item at index:", index)
      navigationStack.push(detailPageComponent, { tweet: row.item })
    }

    onProfileSelected: {
      console.debug("Selected profile at index:", index)
      navigationStack.push(profilePageComponent, { profile: row.item.user })
    }
  }

  // Image header
  // Background color placeholder
  Rectangle {
    id: topImage

    z: 1

    y: listView.contentY + listView.headerItem.height > 0 ?
         -Math.min(listView.contentY + listView.headerItem.height, height - barHeight) : 0
    color: ("#" + (!!profile && profile.profile_background_color))
    width: parent.width
    height: dp(140)

    Image {
      id: imgHeader
      source: !!profile && !!profile.profile_banner_url ? profile.profile_banner_url : ""
      anchors.fill: parent
      fillMode: Image.PreserveAspectCrop
      visible: false
    }

    FastBlur {
      source: imgHeader
      anchors.fill: imgHeader
      radius: 64 * Math.max(0, Math.min(1, (listView.contentY * 0.5 + topImage.height) / listView.headerItem.contentHeight))
    }

    // Profile image
    RoundedImage {
      id: profileImage
      source: !!profile && !!profile.profile_image_url ? profile.profile_image_url.replace("_normal", "_bigger") : ""

      x: dp(10)
      y: -(height / 4) + parent.height
      width: dp(50)
      height: dp(50)

      opacity: Math.max(0, Math.min(1, -(listView.contentY + topImage.height) / topImage.height))
      scale: opacity

      border.color: "white"
      border.width: dp(2)
      backgroundColor: "white"
    }

    Rectangle {
      anchors.fill: parent
      color: "#35000000"
      opacity: 1-profileImage.opacity
    }
  }

  listView {
    emptyText.text: qsTr("No tweets")
    anchors.leftMargin: nativeUtils.safeAreaInsets.left
    anchors.rightMargin: nativeUtils.safeAreaInsets.right
    anchors.bottomMargin: nativeUtils.safeAreaInsets.bottom

    header: Column {
      id: contentColumn

      width: parent.width

      property real contentHeight: height - topImage.height

      // Spacer for image header
      Item {
        height: topImage.height
        width: parent.width
      }

      // Profile details
      Rectangle {
        width: parent.width
        height: profileContent.height + dp(10) + profileContent.y
        color: "white"

        Column {
          id: profileContent

          width: parent.width - dp(20)
          x: dp(10)
          y: profileImage.height * 0.75 + dp(10)
          spacing: dp(8)

          Column {
            spacing: dp(4)

            AppText {
              text: profile.name
              font.pixelSize: sp(16)
              font.bold: true
            }

            AppText {
              text: "@" + profile.screen_name
              font.pixelSize: sp(12)
              color: Theme.secondaryTextColor
            }
          }

          AppText {
            text: profile.description
            width: parent.width
            wrapMode: Text.WordWrap
            font.pixelSize: sp(12)
            lineHeight: 1.3
          }

          Flow {
            spacing: dp(6)
            width: parent.width

            Icon {
              icon: IconType.mapmarker
              size: dp(12)
              color: Theme.secondaryTextColor
              visible: profileLocation.visible
            }

            Text {
              id: profileLocation
              text: profile.location
              visible: !!profile.location
              font.family: Theme.normalFont.name
              font.pixelSize: sp(12)
              color: Theme.textColor
            }

            Icon {
              icon: IconType.link
              size: dp(12)
              color: Theme.secondaryTextColor
              visible: profileUrl.visible
            }

            Text {
              id: profileUrl
              text: profile.entities.url && profile.entities.url.urls[0].display_url || ""
              visible: !!profile.entities.url
              font.family: Theme.normalFont.name
              font.pixelSize: sp(12)
              color: Theme.textColor
            }
          }

          Flow {
            spacing: dp(6)
            width: parent.width

            Text {
              text: profile.friends_count
              font.family: Theme.normalFont.name
              font.bold: true
              font.pixelSize: sp(12)
              color: Theme.textColor
            }

            Text {
              text: qsTr("Following")
              font.family: Theme.normalFont.name
              font.pixelSize: sp(12)
              font.capitalization: Font.AllUppercase
              color: Theme.secondaryTextColor
            }

            Text {
              text: profile.followers_count
              font.family: Theme.normalFont.name
              font.bold: true
              font.pixelSize: sp(12)
              color: Theme.textColor
            }

            Text {
              text: qsTr("Followers")
              font.family: Theme.normalFont.name
              font.pixelSize: sp(12)
              font.capitalization: Font.AllUppercase
              color: Theme.secondaryTextColor
            }
          }
        }
      }

      Rectangle {
        y: parent.height - height
        width: parent.width
        height: 1
        color: Theme.dividerColor
      }
    }
  }
}
