import Felgo 3.0
import QtQuick 2.0
import "."

Rectangle {
  id: trackRow
  height: col.height + dp(10) * 2
  width: parent ? parent.width : 0
  color: Theme.listItem.backgroundColor
  property var track
  signal clicked

  property StyleSimpleRow style: StyleSimpleRow { }

  RippleMouseArea {
    anchors.fill: parent
    circularBackground: false
    onClicked: {
      trackRow.clicked()
    }
  }

  Column {
    id: col
    width: parent.width - dp(Theme.navigationBar.defaultBarItemPadding) * 2
    x: dp(Theme.navigationBar.defaultBarItemPadding) * 2
    anchors.verticalCenter: parent.verticalCenter

    AppText {
      text: track.title
      width: parent.width
    }
    AppText {
      width: parent.width
      elide: AppText.ElideRight
      maximumLineCount: 1
      text: track.talks.length + " Sessions"
      color: Theme.secondaryTextColor
      font.pixelSize: sp(12)
    }
  }

  // track indicators
  Rectangle {
    id: trackCol
    x: dp(2)
    width: dp(4)
    height: parent.height - x * 6
    anchors.verticalCenter: parent.verticalCenter
    color: loaderItem.getTrackColor(track.title)
  }

  Rectangle {
    anchors.right: parent.right
    anchors.bottom: parent.bottom
    width: Theme.isIos ? parent.width - dp(Theme.navigationBar.defaultBarItemPadding) : parent.width
    color: Theme.listItem.dividerColor
    height: px(1)
  }

  Icon {
    icon: IconType.angleright
    anchors.right: parent.right
    anchors.rightMargin: Theme.navigationBar.defaultBarItemPadding
    anchors.verticalCenter: parent.verticalCenter
    color: Theme.dividerColor
    visible: Theme.isIos
    size: dp(24)
  }
}
