import QtQuick 2.0
import Felgo 3.0

Item {
  id: row
  x: dp(16)

  //height: col.height + dp(10) // not good, because then the size is different (ie higher) if 2 text items are shown
  width: parent.width - 2 * x // fill parent width with padding
  height: Math.max(col.height, dp(40))
  anchors.verticalCenter: parent.verticalCenter

  property alias text: name.text
  property alias subtext: subtext.text
  property alias detailText: detailText.text
  property alias userImageSource: userImage.source
  property alias locale: userImage.locale

  property alias textItem: name
  property alias subtextItem: subtext

  property real spacing: dp(15) // spacing between the UserImage and the text column

  property var details

  Item {
    id: userImageContainer
    height: dp(40)
    width: userImage.width
    anchors.verticalCenter: parent.verticalCenter

    SocialUserImage {
      id: userImage
      width: height
      height: parent.height * 0.66 // the locale flag needs a bit of extra space
      anchors.verticalCenter: parent.verticalCenter
    }
  }

  Column {
    id: col
    anchors.verticalCenter: parent.verticalCenter
    anchors.left: userImageContainer.right
    anchors.right: parent.right
    anchors.leftMargin: row.spacing
    spacing: dp(3)

    Text {
      id: name
      font.pixelSize: sp(14)
      width: col.width
      elide: Text.ElideRight
    }

    Item {width: 1;height: 1; visible: detailText.visible } // spacer

    Text {
      id: detailText
      font.pixelSize: sp(12)
      width: col.width
      elide: Text.ElideRight
      visible: detailText.text !== undefined && detailText.text !== ""
    }

    Flow {
      width: col.width
      spacing: dp(8)

      Repeater {
        id: detailRepeater
        model: details

        Row {
          spacing: dp(4)

          Icon {
            id: detailIcon
            icon: modelData.icon
            color: "#999"
            size: sp(8)
            textItem.width: width
            anchors.verticalCenter: parent.verticalCenter
          }

          Text {
            font.pixelSize: sp(10)
            font.family: socialViewItem.bodyFontName
            text: modelData.text
            color: "#000"
            anchors.verticalCenter: parent.verticalCenter
          }
        }
      }
    }

    Item {width: 1;height: 1; visible: detailRepeater.count > 0 } // spacer

    Text {
      id: subtext
      font.family: socialViewItem.bodyFontName
      font.pixelSize: name.font.pixelSize * 0.8 // make a bit smaller than the main text
      visible: subtext.text !== undefined && subtext.text !== "" // if it is invisible, the only text shown is the above one, which will then be vertically centered
      width: name.width
      elide: Text.ElideRight
      color: socialViewItem.tintColor
    }
  }
}// Row
