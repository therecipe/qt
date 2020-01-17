import Felgo 3.0
import QtQuick 2.0
import "../common"

FlickablePage {
  title: "About Felgo"

  flickable.contentWidth: width
  flickable.contentHeight: Math.max(flickable.height, content.height + demosArea.height)

  Rectangle {
    width: parent.width
    height: content.height + 1000
    color: app.secondaryTintColor
    anchors.bottom: content.top
  }

  // page content
  Column {
    id: content
    width: parent.width
    spacing: dp(10)

    Column {
      width: parent.width

      // Banner
      Rectangle {
        width: parent.width
        height: dp(160)
        color: app.secondaryTintColor

        Column {
          anchors.centerIn: parent
          spacing: dp(15)

          Image {
            anchors.horizontalCenter: parent.horizontalCenter
            source: "../../assets/felgo.png"
            width: dp(150)
            fillMode: Image.PreserveAspectFit
          }

          AppText {
            anchors.horizontalCenter: parent.horizontalCenter
            color: "#fff"
            font.pixelSize: sp(15)
            text: "We made this App for you!"
            font.bold: true
            font.capitalization: Font.AllUppercase
          }

          AppText {
            anchors.horizontalCenter: parent.horizontalCenter
            color: "#fff"
            font.pixelSize: sp(13)
            text: "And there is <b>much more</b> we can do."
          }

        }
      }

      Rectangle {
        width: parent.width
        height: partnerCol.height + dp(Theme.navigationBar.defaultBarItemPadding)*2
        color: Theme.secondaryBackgroundColor

        Column {
          id: partnerCol
          width: parent.width
          anchors.centerIn: parent
          spacing: dp(Theme.navigationBar.defaultBarItemPadding)

          // Felgo as partner
          AppText {
            width: parent.width - dp(Theme.navigationBar.defaultBarItemPadding) * 2
            anchors.horizontalCenter: parent.horizontalCenter
            font.pixelSize: sp(13)
            wrapMode: Text.WordWrap
            text: "Choose Felgo as your <b>partner</b> and get:"
          }

          AppTextWithBullet {
            anchors.horizontalCenter: parent.horizontalCenter
            width: Math.min(parent.width - dp(Theme.navigationBar.defaultBarItemPadding) * 2, dp(600))
            text: "Technical consulting for your <b>mobile app & game development</b> for Qt C++ or QML projects"
          }
          AppTextWithBullet {
            anchors.horizontalCenter: parent.horizontalCenter
            width: Math.min(parent.width - dp(Theme.navigationBar.defaultBarItemPadding) * 2, dp(600))
            text: "Outsourcing partner to flexibly develop apps & games together with your development team"
          }
          AppTextWithBullet {
            anchors.horizontalCenter: parent.horizontalCenter
            width: Math.min(parent.width - dp(Theme.navigationBar.defaultBarItemPadding) * 2, dp(600))
            text: "Qt Trainings: both on-site & remote"
          }
          AppTextWithBullet {
            anchors.horizontalCenter: parent.horizontalCenter
            width: Math.min(parent.width - dp(Theme.navigationBar.defaultBarItemPadding) * 2, dp(600))
            text: "Access to experts in native iOS & Android development"
          }
          AppTextWithBullet {
            anchors.horizontalCenter: parent.horizontalCenter
            width: Math.min(parent.width - dp(Theme.navigationBar.defaultBarItemPadding) * 2, dp(600))
            text: "10 years experience in mobile development with Qt"
          }

          AppButton {
            flat: false
            anchors.horizontalCenter: parent.horizontalCenter
            text: "Request a Free Intro Call"
            onClicked: {
              amplitude.logEvent("Click MainPage Button",{"label" : text})
              confirmOpenUrl()
            }
            verticalMargin: 0
          }

        }

      }

      // Companies
      Rectangle {
        width: parent.width
        height: dp(100)
        color: Theme.backgroundColor

        Rectangle {
          width: parent.width
          height: px(1)
          color: Theme.dividerColor
        }

        Column {
          anchors.centerIn: parent
          anchors.verticalCenterOffset: -dp(3)
          spacing: dp(10)

          AppText {
            anchors.horizontalCenter: parent.horizontalCenter
            color: Theme.textColor
            font.pixelSize: sp(13)
            text: "Join Felgo customers like:"
            font.bold: true
          }

          Image {
            anchors.horizontalCenter: parent.horizontalCenter
            source: "../../assets/companies1.jpg"
            width: dp(240)
            fillMode: Image.PreserveAspectFit
          }

        }
      }

      // Tech Partner
      Rectangle {
        width: parent.width
        height: dp(100)
        color: Theme.secondaryBackgroundColor

        Rectangle {
          width: parent.width
          height: px(1)
          color: Theme.dividerColor
        }

        Rectangle {
          width: parent.width
          height: px(1)
          color: Theme.dividerColor
          anchors.bottom: parent.bottom
        }

        Column {
          anchors.centerIn: parent
          anchors.verticalCenterOffset: -dp(3)
          spacing: dp(10)

          AppText {
            anchors.horizontalCenter: parent.horizontalCenter
            color: Theme.textColor
            font.pixelSize: sp(13)
            text: "Felgo is proud to be:"
            font.bold: true
          }

          Image {
            anchors.horizontalCenter: parent.horizontalCenter
            source: "../../assets/qt_tech_partner.png"
            width: dp(180)
            fillMode: Image.PreserveAspectFit
          }
        }
      }

    }

    // More Felgo Features
    AppText {
      width: Math.min(parent.width - dp(Theme.navigationBar.defaultBarItemPadding) * 2, dp(600))
      anchors.horizontalCenter: parent.horizontalCenter
      font.pixelSize: sp(13)
      wrapMode: Text.WordWrap
      color: Theme.secondaryTextColor
      text: "Do you develop a <b>mobile Qt-based app in your company</b>?
Save time and money by using the Felgo SDK."
    }

    AppText {
      width: Math.min(parent.width - dp(Theme.navigationBar.defaultBarItemPadding) * 2, dp(600))
      anchors.horizontalCenter: parent.horizontalCenter
      font.pixelSize: sp(13)
      wrapMode: Text.WordWrap
      color: Theme.secondaryTextColor
      text: "Felgo <b>extends Qt</b> with APIs for:"
    }

    AppText {
      anchors.horizontalCenter: parent.horizontalCenter
      width: Math.min(parent.width - dp(Theme.navigationBar.defaultBarItemPadding) * 2, dp(600))
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
      width: Math.min(parent.width - dp(Theme.navigationBar.defaultBarItemPadding) * 2, dp(600))
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
      width: Math.min(parent.width - dp(Theme.navigationBar.defaultBarItemPadding) * 2, dp(600))
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
    NativeDialog.confirm("Leave the app?","This action opens your browser to visit https://felgo.com/qtws17-free-intro/.",function(ok) {
      if(ok)
        nativeUtils.openUrl("https://felgo.com/qtws17-free-intro/")
    })
  }
}
