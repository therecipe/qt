import Felgo 3.0
import QtQuick 2.0
import "../common"

FlickablePage {
  title: "Main"
  backgroundColor: app.secondaryTintColor

  // Timer for countdown until conference start
  Timer {
    running: parent.visible
    interval: 1000
    repeat: true

    property int conferenceStartTime: 1507622400 // 2017-10-10 10am GMT+2
    Component.onCompleted: triggered()
    onTriggered: {
      var remain = conferenceStartTime - (Date.now() / 1000)
      if(remain <= 0) {
        running = false
        remainDaysColumn.visible = false
        qwsWrapper.height = dp(150)
        return
      }

      remainDays.value = remain / 60 / 60 / 24
      remain -= remainDays.value * 60 * 60 * 24

      remainHours.value = remain / 60 / 60
      remain -= remainHours.value * 60 * 60

      remainMins.value = remain / 60
      remain -= remainMins.value * 60

      remainSecs.value = remain
    }
  }

  // set up navigation bar
  titleItem: Item {
    width: dp(150)
    implicitWidth: dp(150)
    height: dp(Theme.navigationBar.height)

    Image {
      id: img
      source: Theme.isIos && Theme.tabBar.backgroundColor != "#080808" ? "../../assets/QtWS2017_logo.png" : "../../assets/QtWS2017_logo_white.png"
      width: dp(150)
      height: parent.height
      fillMode: Image.PreserveAspectFit
      y: Theme.isAndroid ? dp(Theme.navigationBar.titleBottomMargin) : 0
    }
  }

  flickable.contentWidth: width
  flickable.contentHeight: Math.max(flickable.height, content.height + demosArea.height)

  Rectangle {
    width: parent.width
    height: content.height + 3000
    color: Theme.backgroundColor
  }

  // page content
  Column {
    id: content
    width: parent.width
    spacing: dp(10)

    property real descriptionTextMaxWidth: Math.min(parent.width - dp(Theme.navigationBar.defaultBarItemPadding) * 2, dp(600))

    Column {
      width: parent.width

      Rectangle {
        id: qwsWrapper
        height: dp(220)
        width: parent.width
        color: app.secondaryTintColor

        AppImage {
          anchors.fill: parent
          fillMode: AppImage.PreserveAspectCrop
          source: "../../assets/herobg_qtws17-1.jpg"
          opacity: 0.35
        }

        Column {
          width: parent.width
          anchors.verticalCenter: parent.verticalCenter
          spacing: dp(Theme.navigationBar.defaultBarItemPadding)

          AppText {
            width: parent.width
            horizontalAlignment: AppText.AlignHCenter
            color: "white"
            text: "The Future is\nWritten with Qt"
            font.pixelSize: sp(22)
          }

          AppText {
            width: parent.width
            horizontalAlignment: AppText.AlignHCenter
            color: "white"
            text: "Berlin, Germany / 10-12 October"
            font.pixelSize: sp(14)
          }

          // remaining days
          Column {
            id: remainDaysColumn
            anchors.horizontalCenter: parent.horizontalCenter
            spacing: parent.spacing * 0.5

            AppText {
              text: "\nstarting in"
              color: "white"
              font.pixelSize: sp(11)
              font.bold: true
              anchors.horizontalCenter: parent.horizontalCenter
            }

            Row {
              anchors.horizontalCenter: parent.horizontalCenter
              spacing: dp(5)
              CountDownBlock { id: remainDays; titleText: "Days" }
              CountDownBlock { id: remainHours; titleText: "Hours" }
              CountDownBlock { id: remainMins; titleText: "Min" }
              CountDownBlock { id: remainSecs; titleText: "Sec" }
            }
          }
        }
      } // colored section

      Rectangle {
        width: parent.width
        height: felgoBlockColumn.height
        color: Theme.secondaryBackgroundColor

        Column {
          id: felgoBlockColumn
          width: parent.width
          spacing: dp(Theme.navigationBar.defaultBarItemPadding)

          Item {
            width: parent.width
            height: px(1)
          }

          AppText {
            width: parent.width - dp(Theme.navigationBar.defaultBarItemPadding) * 2
            anchors.horizontalCenter: parent.horizontalCenter
            font.pixelSize: sp(13)
            horizontalAlignment: Text.AlignHCenter
            wrapMode: Text.WordWrap
            text: "This cross-platform mobile app<br/>was made with <b>Qt 5.9 & the Felgo SDK</b>!"
          }

          AppText {
            anchors.horizontalCenter: parent.horizontalCenter
            width: Math.min(parent.width - dp(Theme.navigationBar.defaultBarItemPadding), implicitWidth + dp(Theme.navigationBar.defaultBarItemPadding))
            font.pixelSize: sp(13)
            wrapMode: Text.WordWrap
            textFormat: AppText.RichText
            onLinkActivated: nativeUtils.openUrl(link)
            text: "<style>a:link { color: "+Theme.tintColor+";}</style>
<ul><li><b>Only ~2900 lines of code</b> - download the full source on GitHub below</li>

<li><b>Build cross-platform native apps</b> for" + (system.isPlatform(System.IOS) ? " all major mobile platforms, desktop and embedded devices</li>" : "</li>")
          }

          AppImage {
            width: parent.width * 0.8
            fillMode: AppImage.PreserveAspectFit
            source: "../../assets/platforms.png"
            anchors.horizontalCenter: parent.horizontalCenter
            visible: !system.isPlatform(System.IOS) // hide on iOS, instead display text above
          }

          AppButton {
            flat: false
            anchors.horizontalCenter: parent.horizontalCenter
            text: "More Info & Source Code"
            onClicked: {
              amplitude.logEvent("Click MainPage Button",{"label" : text})
              confirmOpenUrl()
            }
            verticalMargin: 0
          }

          Item {
            width: parent.width
            height: px(1)
          }

        } // felgo block column
      } // felgo block area
    } // main intro block


    AppText {
      width: content.descriptionTextMaxWidth
      anchors.horizontalCenter: parent.horizontalCenter
      font.pixelSize: sp(13)
      wrapMode: Text.WordWrap
      color: Theme.secondaryTextColor
      text: "With Qt you can build applications for tomorrow, while delivering value to your customers today. The Qt World Summit offers insight and inspiration to leading technology innovators, industry experts, startups, and the community."
    }

    AppButton {
      text: "Show Agenda"
      anchors.horizontalCenter: parent.horizontalCenter
      flat: true
      onClicked: navigationStack.push(Qt.resolvedUrl("RoomPage.qml"), { room: "Agenda" })
    }

    // Spacer
    Item {
      width: parent.width
      height: px(1)
    }

    // More Felgo Features
    AppText {
      width: content.descriptionTextMaxWidth
      anchors.horizontalCenter: parent.horizontalCenter
      font.pixelSize: sp(13)
      wrapMode: Text.WordWrap
      color: Theme.secondaryTextColor
      text: "Felgo <b>extends Qt</b> with APIs for:"
    }

    AppText {
      anchors.horizontalCenter: parent.horizontalCenter
      width: content.descriptionTextMaxWidth
      font.pixelSize: sp(13)
      wrapMode: Text.WordWrap
      color: Theme.secondaryTextColor
      textFormat: AppText.RichText
      text: "<p><b>Native UI & UX cross-platform:</b>
<ul>
<li>Native iOS Style</li>
<li>Native Navigation on iOS & Android</li>
<li>No platform-specific code</li>
<li>100% same source code</li>
</ul>"
    }

    AppText {
      anchors.horizontalCenter: parent.horizontalCenter
      width: content.descriptionTextMaxWidth
      font.pixelSize: sp(13)
      wrapMode: Text.WordWrap
      color: Theme.secondaryTextColor
      textFormat: AppText.RichText
      text: "
<b>Responsive Layouts</b> to support both smartphone and tablet with a <b>single code base</b>
<br/><br/>
<b>More native features like:</b>
<ul>
<li>Sharing to Facebook, Twitter, WhatsApp, etc.</li>
<li>Image picker from gallery or camera</li>
<li>Native dialogs</li>
<li>Address book access</li>
<li>And many more...</li>
</ul>"
    }

    AppText {
      anchors.horizontalCenter: parent.horizontalCenter
      width: content.descriptionTextMaxWidth
      font.pixelSize: sp(13)
      wrapMode: Text.WordWrap
      color: Theme.secondaryTextColor
      textFormat: AppText.RichText
      text: "<b>Native Plugins</b> for:"
    }

    AppImage {
      anchors.horizontalCenter: parent.horizontalCenter
      height: dp(250)
      fillMode: AppImage.PreserveAspectFit
      source: "../../assets/felgo-plugins.png"
    }

    // Spacer
    Item {
      width: parent.width
      height: dp(16)
    }
  } // Column


  // Link to Demos
  Rectangle {
    id: demosArea
    anchors.bottom: parent.bottom
    width: parent.width
    height: demosContent.height + 2 * dp(Theme.navigationBar.defaultBarItemPadding)
    color: "#eeeeee"

    Column {
      id: demosContent
      width: parent.width
      anchors.verticalCenter: parent.verticalCenter
      spacing: content.spacing

      AppText {
        id: demosText
        width: parent.width - dp(Theme.navigationBar.defaultBarItemPadding) * 2
        anchors.horizontalCenter: parent.horizontalCenter
        font.pixelSize: sp(13)
        horizontalAlignment: Text.AlignHCenter
        wrapMode: Text.WordWrap
        color: Theme.secondaryTextColor
        text: "Try more of the 60 other open-source Felgo Apps like:"
      }

      // Demos Grid
      Item {
        anchors.horizontalCenter: parent.horizontalCenter
        width: Math.min(parent.width - 2 * dp(Theme.navigationBar.defaultBarItemPadding), dp(300))
        height: Math.max(unoCol.height, showcaseCol.height)

        // One Card Demo
        Column {
          id: unoCol
          width: (parent.width - dp(Theme.navigationBar.defaultBarItemPadding)) * 0.5
          spacing: dp(Theme.navigationBar.defaultBarItemPadding) * 0.5

          AppText {
            text: "One Card! - UNO Game"
            font.pixelSize: sp(13)
            anchors.horizontalCenter: parent.horizontalCenter
          }

          AppImage {
            width: parent.width * 0.5
            fillMode: AppImage.PreserveAspectFit
            anchors.horizontalCenter: parent.horizontalCenter
            source: "../../assets/logo-onecard.png"
          }

          AppText {
            width: parent.width
            font.pixelSize: sp(12)
            color: Theme.secondaryTextColor
            horizontalAlignment: AppText.AlignHCenter
            text: "Play UNO with your friends in this multiplayer card game!"
          }
        } // One Card Column

        MouseArea {
          anchors.fill: unoCol
          onClicked: {
            amplitude.logEvent("Click MainPage Button",{"label" : "One Card! - UNO Game"})
            var url = Theme.isAndroid ? "https://play.google.com/store/apps/details?id=net.vplay.demos.ONECard" : "https://itunes.apple.com/at/app/id1112447141?mt=8"
            nativeUtils.openUrl(url)
          }
        }

        // Showcase Demo
        Column {
          id: showcaseCol
          width: unoCol.width
          anchors.left: unoCol.right
          anchors.leftMargin: spacing
          spacing: unoCol.spacing

          AppText {
            text: "Felgo Showcase App"
            font.pixelSize: sp(13)
            anchors.horizontalCenter: parent.horizontalCenter
          }

          AppImage {
            width: parent.width * 0.5
            fillMode: AppImage.PreserveAspectFit
            anchors.horizontalCenter: parent.horizontalCenter
            source: "../../assets/logo-showcase.png"
          }

          AppText {
            width: parent.width
            font.pixelSize: sp(12)
            color: Theme.secondaryTextColor
            horizontalAlignment: AppText.AlignHCenter
            text: "This app includes all open source demos & examples of Felgo Apps!"
          }
        } // Showcase Column

        MouseArea {
          anchors.fill: showcaseCol
          onClicked: {
            amplitude.logEvent("Click MainPage Button",{"label" : "Felgo Showcase App"})
            var url = Theme.isAndroid ? "https://play.google.com/store/apps/details?id=net.vplay.demos.apps.showcaseapp" : "https://itunes.apple.com/at/app/id1040477271?mt=8"
            nativeUtils.openUrl(url)
          }
        }
      } // Demos Grid
    } // Demos Content
  } // Demos Area

  // confirmOpenUrl - display confirm dialog before opening Felgo url
  function confirmOpenUrl() {
    NativeDialog.confirm("Leave the app?","This action opens your browser to visit https://felgo.com/qws-conference-in-app-2017.",function(ok) {
      if(ok)
        nativeUtils.openUrl("https://felgo.com/qws-conference-in-app-2017")
    })
  }


}
