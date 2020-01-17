import QtQuick 2.3
import QtQuick.Layouts 1.1

import Felgo 3.0

Item {
  id: cell

  property var item: modelData

  property string icon: item && item.icon || ""
  property string name: item && item.name || ""
  property string handle: item && item.handle || ""
  property string time: item && item.time || ""
  property string text: item && item.text || ""
  property string image: item && item.image || ""

  property bool actionsHidden: item && item.actionsHidden || false

  property int favs: item && item.favorite_count || 0
  property int retweets: item && item.retweet_count || 0

  property bool isFaved: item && item.favorited || false
  property bool isRetweeted: item && item.retweeted || false
  // Temp, should be in theme!
  property color favColor: "#ffac33"
  property color retweetColor: "#77b255"
  property color inactiveColor: "#ccc"

  width: parent ? parent.width : 0
  // minimum of 48dp for interactable items recommended
  height: Math.max(dp(48), innerGrid.height)


  signal selected(int index)
  signal profileSelected(int index)


  Rectangle {
    color: cellArea.pressed ? "#eee" : "#fff"
    anchors.fill: parent

    Behavior on color {
      ColorAnimation { duration: 150; easing.type: Easing.InOutQuad }
    }
  }

  MouseArea {
    id: cellArea

    enabled: cell.enabled
    anchors.fill: parent

    onClicked: selected(typeof index !== "undefined" ? index : 0)
  }

  // Main cell content inside this item
  GridLayout {
    id: innerGrid

    // Auto-break after 4 columns, so we do not have to set row & column manually
    columns: 4
    rowSpacing: dp(4)
    columnSpacing: dp(8)

    x: dp(10)
    width: parent.width - 2 * x

    // Top spacer
    Item {
      id: topSpacer
      width: parent.width
      height: dp(6)

      Layout.columnSpan: parent.columns
      Layout.fillWidth: true
    }

    RoundedImage {
      id: avatarImage

      img.source: cell.icon
      img.defaultSource: "../../assets/user_default.png"

      Layout.preferredWidth: dp(50)
      Layout.preferredHeight: dp(50)
      Layout.rowSpan: image.visible ? 4 : 3
      Layout.alignment: Qt.AlignTop

      MouseArea {
        anchors.fill: parent
        onClicked: profileSelected(index || 0)
      }
    }

    Text {
      id: nameText
      elide: Text.ElideRight
      maximumLineCount: 1
      color: Theme.textColor
      font.family: Theme.normalFont.name
      font.bold: true
      font.pixelSize: dp(14)
      lineHeight: dp(16)
      lineHeightMode: Text.FixedHeight
      text: cell.name
    }

    Text {
      id: handleText
      elide: Text.ElideRight
      maximumLineCount: 1
      color: Theme.secondaryTextColor
      font.family: Theme.normalFont.name
      font.pixelSize: dp(12)
      lineHeight: dp(16)
      lineHeightMode: Text.FixedHeight
      text: cell.handle
      Layout.fillWidth: true
      verticalAlignment: Text.AlignBottom
      Layout.preferredWidth: parent.width
    }

    Text {
      id: timeText
      elide: Text.ElideRight
      maximumLineCount: 1
      color: Theme.secondaryTextColor
      font.family: Theme.normalFont.name
      font.pixelSize: dp(12)
      lineHeight: dp(16)
      lineHeightMode: Text.FixedHeight
      verticalAlignment: Text.AlignBottom
      text: cell.time
      Layout.alignment: Qt.AlignRight
    }

    Text {
      id: contentText
      color: Theme.textColor
      linkColor: Theme.tintColor
      font.family: Theme.normalFont.name
      font.pixelSize: dp(14)
      lineHeight: 1.15
      text: cell.text
      wrapMode: Text.WordWrap
      Layout.columnSpan: 3
      Layout.fillWidth: true
      Layout.fillHeight: true
      Layout.alignment: Qt.AlignTop


      onLinkActivated: {
        Qt.openUrlExternally(link)
      }
    }

    RoundedImage {
      id: image

      fillMode: Image.PreserveAspectCrop
      visible: cell.image.length > 0 && !img.error

      source: cell.image

      Layout.columnSpan: 3
      Layout.fillWidth: true
      Layout.preferredWidth: contentText.width
      Layout.preferredHeight: dp(320)
      Layout.maximumWidth: contentText.width
      Layout.maximumHeight: dp(160)

      MouseArea {
        anchors.fill: parent

        onClicked: {
          PictureViewer.show(app, image.source)
        }
      }
    }

    // actions
    Row {
      visible: !actionsHidden

      spacing: dp(4)

      Layout.columnSpan: 3
      Layout.fillWidth: true
      Layout.preferredWidth: parent.width
      Layout.preferredHeight: replyIcon.height + dp(4)
      Layout.alignment: Qt.AlignBottom

      Icon {
        id: replyIcon
        icon: IconType.reply
        color: inactiveColor
      }

      Item {
        width: dp(28)
        height: 1
      }

      Icon {
        icon: IconType.retweet
        color: isRetweeted ? retweetColor : inactiveColor
      }

      Text {
        text: retweets
        visible: retweets > 0
        color: isRetweeted ? retweetColor : inactiveColor
        font.pixelSize: sp(13)
      }

      Item {
        width: dp(28)
        height: 1
      }

      Icon {
        icon: IconType.star
        color: isFaved ? favColor : inactiveColor
      }

      Text {
        text: favs
        visible: favs > 0
        color: isFaved ? favColor : inactiveColor
        font.pixelSize: sp(13)
      }
    }

    Item {
      id: bottomSpacer

      width: parent.width
      height: dp(6)

      Layout.columnSpan: parent.columns
      Layout.fillWidth: true
    }
  }

  // Bottom cell divider
  Rectangle {
    id: divider
    width: parent.width
    color: Theme.dividerColor
    height: px(1)
    anchors.bottom: parent.bottom
  }
}
