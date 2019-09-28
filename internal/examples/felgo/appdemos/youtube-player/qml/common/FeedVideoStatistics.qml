import Felgo 3.0
import QtQuick 2.0

Item {
  width: parent.width
  height: dp(24)

  property var statistics: undefined
  property real spacing: dp(Theme.navigationBar.defaultBarItemPadding)

  VideoStatisticRow {
    anchors.verticalCenter: parent.verticalCenter
    spacing : parent.spacing

    iconItem.icon: IconType.eye
    textItem.text: statistics.viewCount
  }

  VideoStatisticRow {
    anchors.centerIn: parent
    spacing : parent.spacing

    iconItem.icon: IconType.comment
    textItem.text: statistics.commentCount
  }

  VideoStatisticRow {
    anchors.right: parent.right
    anchors.verticalCenter: parent.verticalCenter
    spacing: parent.spacing

    iconItem.icon: IconType.thumbsup
    textItem.text: statistics.likeCount
  }
}
