import Felgo 3.0
import QtQuick 2.0

FlickablePage {
  title: "Venue"

  flickable.contentWidth: parent.width
  flickable.contentHeight: contentCol.height

  Column {
    id: contentCol
    width: parent.width

    Item {
      width: parent.width
      height: landscape ? dp(300) : img.height * 0.75
      clip: true

      AppImage {
        id: img
        width: parent.width
        height: width / sourceSize.width * sourceSize.height
        fillMode: AppImage.PreserveAspectFit
        source: "../../assets/venue_photo.jpg"
        anchors.bottom: parent.bottom
      }
    }

    Item {
      width: parent.width
      height: addressCol.height + dp(Theme.navigationBar.defaultBarItemPadding) * 2
      Column {
        id: addressCol
        width: parent.width - dp(Theme.navigationBar.defaultBarItemPadding) * 2
        anchors.centerIn: parent
        spacing: dp(Theme.navigationBar.defaultBarItemPadding)
        Item {
          width: parent.width
          height: 1
          visible: Theme.isIos
        }
        Column {
          width: parent.width
          AppText {
            width: parent.width
            wrapMode: AppText.WordWrap
            text: "bcc Berlin Congress Center"
          }
          AppText {
            width: parent.width
            wrapMode: AppText.WordWrap
            color: Theme.secondaryTextColor
            font.pixelSize: sp(13)
            text: "Alexanderstraße 11"
          }
          AppText {
            width: parent.width
            wrapMode: AppText.WordWrap
            color: Theme.secondaryTextColor
            font.pixelSize: sp(13)
            text: "10178 Berlin, Germany"
          }
        }

        AppButton {
          text: "Plan Route"
          minimumWidth: parent.width - dp(Theme.navigationBar.defaultBarItemPadding) * 2
          anchors.horizontalCenter: parent.horizontalCenter
          onClicked: {
            amplitude.logEvent("Plan Route")
            if (Theme.isIos){
              Qt.openUrlExternally("http://maps.apple.com/?q=bcc, Alexanderstraße 11, 10178 Berlin, Germany")
            } else {
              Qt.openUrlExternally("geo:0,0?q=52.520427,13.416309")
            }
          }
        }
      }
      Rectangle {
        anchors.top: parent.top
        width: parent.width
        color: Theme.listItem.dividerColor
        height: px(1)
      }
      Rectangle {
        anchors.bottom: parent.bottom
        width: parent.width
        color: Theme.listItem.dividerColor
        height: px(1)
      }
    }

    AppImage {
      width: parent.width
      height: dp(200)
      fillMode: AppImage.PreserveAspectCrop
      visible: status === Image.Ready || status === Image.Loading
      property int imgWidth: Math.min(1000,Math.max(320,width))
      property int imgHeight: Math.min(1000,Math.max(240,height))
      source: "https://api.mapbox.com/v4/mapbox.streets/pin-l-building+008000(13.416309,52.520427)/13.416309,52.520427,15/"+imgWidth+"x"+imgHeight+".png?access_token=pk.eyJ1IjoiaDAwYnMiLCJhIjoiY2lvdmNsbDI3MDA2OXc5bHdxNWE5NGdtOCJ9.me4r_KETvbmcohxomTuhvQ"

      // use map image from assets as a fallback in case MapBox image doesn't work
      property string fallbackImageSource: "../../assets/venue.png"
      onStatusChanged: {
        if(status === Image.Error && source !== fallbackImageSource)
          source = fallbackImageSource
      }

      RippleMouseArea {
        anchors.fill: parent
        circularBackground: false
        onClicked: {
          if (Theme.isIos){
            Qt.openUrlExternally("http://maps.apple.com/?q=bcc, Alexanderstraße 11, 10178 Berlin, Germany")
          } else {
            Qt.openUrlExternally("geo:0,0?q=52.520427,13.416309")
          }
        }
      }
    }

    Column {
      width: parent.width - dp(Theme.navigationBar.defaultBarItemPadding) * 2
      anchors.horizontalCenter: parent.horizontalCenter
      spacing: dp(Theme.navigationBar.defaultBarItemPadding)

      Item {
        width: parent.width
        height: 1
      }

      AppText {
        width: parent.width
        wrapMode: AppText.WordWrap
        onLinkActivated: nativeUtils.openUrl(link)
        textFormat: AppText.StyledText
        font.pixelSize: sp(14)
        linkColor: Theme.tintColor
        text: "<b>Berlin, one of the biggest tech hubs in Europe will host the Qt World Summit again this year.</b>
<br/><br/>
<a href='http://bcc-berlin.de/' target='_blank'>bcc Berlin</a> is right in the heart of Berlin, next to Alexanderplatz.\n
Whether travelling from Germany or overseas, you’ll benefit from easy access to the venue."
      }

      Item {
        width: parent.width
        height: 1
      }
    }
  }
}
