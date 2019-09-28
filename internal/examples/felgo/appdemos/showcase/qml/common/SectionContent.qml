import QtQuick 2.0
import Felgo 3.0

Rectangle {
  id: sectionContent
  color: Theme.backgroundColor
  width: parent.width
  height: contentContainer.height + 2 * verticalMargin

  property real horizontalMargin: dp(Theme.navigationBar.defaultBarItemPadding)
  property real verticalMargin: dp(Theme.navigationBar.defaultBarItemPadding)

  property var contentItem
  onContentItemChanged: contentItem.parent = contentContainer

  Item {
    id: contentContainer
    anchors.centerIn: parent
    width: parent.width - 2 * horizontalMargin
    height: contentItem ? contentItem.height : 0
  }
}
