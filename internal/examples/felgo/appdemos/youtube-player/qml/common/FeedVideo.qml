import Felgo 3.0
import QtQuick 2.0
import QtQuick.Layouts 1.3
import QtMultimedia 5.0
import QtGraphicalEffects 1.0

Item {
  id: feedVideo
  width: parent.width - 2 * spacing
  height: topSpacer.height + content.height + bottomSpacer.height
  anchors.horizontalCenter: parent.horizontalCenter

  // properties
  property alias previewImageSource: previewImage.source
  property alias title: titleItem.text
  property string description: ""
  property string richDescription: ""
  property string richDescriptionPreview: richDescription.substring(0, 150)
  property var videoStatistics: []

  property real spacing: dp(Theme.navigationBar.defaultBarItemPadding)
  property string videoId: ""
  property bool showDetails: false

  // cross-linking
  onDescriptionChanged: updateCrossLinks()
  //Component.onCompleted: updateCrossLinks()

  property string otherUrl :""
  property string spotifyUrl: ""
  property string itunesUrl: ""
  property string amazonUrl: ""
  property string soundcloudUrl: ""
  property string facebookUrl: ""
  property string instagramUrl: ""
  property string twitterUrl: ""

  function updateCrossLinks() {
    var regex = /(?:(?:https?):\/\/)?([a-zA-Z/\-=%]+\.)+[\w/\-?=%]{2,}/ig;
    if(description) {
      var results = description.match(regex)
      for(var key in results) {
        var link = results[key]

        // add http to link if not specified
        if(!link.startsWith("http://"))
          link = "http://" + link

        if(results[key].includes("spoti.fi") && !spotifyUrl)
          spotifyUrl = link
        else if(results[key].includes("apple.co") && !itunesUrl)
          itunesUrl = link
        else if(results[key].includes("amzn.to") && !amazonUrl)
          amazonUrl = link
        else if(results[key].includes("soundcloud") && !soundcloudUrl)
          soundcloudUrl = link
        else if(results[key].includes("facebook") && !facebookUrl)
          facebookUrl = link
        else if(results[key].includes("instagram") && !instagramUrl)
          instagramUrl = link
        else if(results[key].includes("twitter") && !twitterUrl)
          twitterUrl = link
        else if(!otherUrl)
          otherUrl = link
      }

      // create rich text with links
      richDescription = description.replace(regex, function(link) { return "<a href='"+link+"'>"+link+"</a>" })

      // add line breaks
      richDescription = "<p>"+richDescription.replace("\n", "</p>\n<p>")+"</p>"
    }
    else
      richDescription = ""
  }

  // spacer
  Item {
    id: topSpacer
    height: spacing / 2
  }

  // mouse area / playback
  MouseArea {
    anchors.fill: parent
    onClicked: {
      youTubePlayer.visible = false
      youTubePlayer.parent = youtubeVideoContainer
      youTubePlayer.loadVideo(videoId, true) // load and play video
      youTubePlayer.visible = true
    }
  }

  // content
  Item {
    id: content
    width: parent.width
    height: previewImageContainer.height + textPanel.anchors.topMargin + textPanel.height
    anchors.top: topSpacer.bottom

    layer.enabled: true
    layer.effect: DropShadow {
      radius: dp(8)
      samples: 16
      horizontalOffset: dp(2)
      verticalOffset: horizontalOffset
      color: "#16161400"
    }

    // preview image
    Item {
      id: previewImageContainer
      anchors.fill: youtubeVideoContainer
      clip: true

      AppImage {
        id: previewImage
        source: Qt.resolvedUrl("../../assets/felgo-logo.png")
        width: parent.width
        height: parent.height
        anchors.centerIn: parent
        fillMode: AppImage.PreserveAspectCrop
      }
    }

    // embedded video
    Item {
      id: youtubeVideoContainer
      width: parent.width
      height: width / 16 * 9
    }

    // details
    Rectangle {
      id: textPanel
      width: parent.width
      height: detailsCol.height + 2 * detailsCol.spacing
      color: "white"
      anchors.top: previewImageContainer.bottom

      Column {
        id: detailsCol
        width: parent.width - 4 * spacing
        anchors.centerIn: parent
        spacing: feedVideo.spacing / 2

        // spacer
        Item { width: parent.width; height: px(1) }

        // title
        AppText {
          id: titleItem
          color: Theme.secondaryTextColor
          width: parent.width
          wrapMode: AppText.WrapAtWordBoundaryOrAnywhere
          font.pixelSize: sp(14)
        }

        // spacer
        Item { width: parent.width; height: px(1); visible: socialLinks.visible }

        // social links
        GridLayout {
          id: socialLinks
          anchors.horizontalCenter: parent.horizontalCenter
          visible: height > 0

          property real spacing: feedVideo.spacing
          columnSpacing: spacing
          rowSpacing: spacing
          columns: parent.width / (dp(36) + spacing)

          property int maxWidth: parent.width - 2 * spacing
          property int itemWidth: children.length * children[0].width + children.legnth - 1 * spacing

          SocialLink {
            targetUrl: otherUrl
            source: "../../assets/web.png"
          }

          SocialLink {
            targetUrl: spotifyUrl
            source: "../../assets/spotify.png"
          }

          SocialLink {
            targetUrl: itunesUrl
            source: "../../assets/itunes.png"
          }

          SocialLink {
            targetUrl: amazonUrl
            source: "../../assets/amazon.png"
          }

          SocialLink {
            targetUrl: soundcloudUrl
            source: "../../assets/soundcloud.png"
          }

          SocialLink {
            targetUrl: instagramUrl
            source: "../../assets/instagram.png"
          }

          SocialLink {
            targetUrl: facebookUrl
            source: "../../assets/facebook.png"
          }

          SocialLink {
            targetUrl: twitterUrl
            source: "../../assets/twitter.png"
          }
        }

        // spacer
        Item { width: parent.width; height: px(1); visible: socialLinks.visible }

        // description
        Item {
          width: parent.width
          height: !showDetails ? descPreviewItem.height : descItem.height

          Behavior on height {
            PropertyAnimation { duration: 150 }
          }

          // preview details
          AppText {
            id: descPreviewItem
            text: richDescriptionPreview + "..."
            width: parent.width
            wrapMode: AppText.WrapAtWordBoundaryOrAnywhere
            color: Theme.secondaryTextColor
            font.pixelSize: sp(12)
            textFormat: AppText.RichText
            visible: !showDetails && showMoreButton.visible
            onLinkActivated: nativeUtils.openUrl(link)
          }

          // full details
          AppText {
            id: descItem
            text: richDescription
            width: parent.width
            wrapMode: AppText.WrapAtWordBoundaryOrAnywhere
            color: Theme.secondaryTextColor
            font.pixelSize: sp(12)
            textFormat: AppText.RichText
            visible: !descPreviewItem.visible
            onLinkActivated: nativeUtils.openUrl(link)
          }
        }

        // show more
        AppButton {
          id: showMoreButton
          text: !showDetails ? "show more" : "show less"
          flat: true
          textSize: dp(10)
          minimumHeight: dp(24)
          anchors.horizontalCenter: parent.horizontalCenter
          onClicked: showDetails = !showDetails
          visible: richDescriptionPreview !== richDescription
        }

        FeedVideoStatistics { statistics: videoStatistics }
      }
    }
  }

  // spacer
  Item {
    id: bottomSpacer
    anchors.bottom: parent.bottom
    height: spacing / 2
  }
}
