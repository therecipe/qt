import Felgo 3.0
import QtQuick 2.0
import QtQuick.Layouts 1.1
import "../common"

Item {
  id: scheduleListItem
  width: parent ? parent.width : 0
  height: gridWrapper.height + 2 * gridWrapper.y

  property var item: [] // talk data item

  readonly property bool loading: speakerImg.loading

  property StyleSimpleRow style: StyleSimpleRow { }

  property ListView parentView

  signal selected(int index)

  // default hover effect
  Rectangle {
    color: cellArea.pressed && !Theme.isAndroid ? scheduleListItem.style.selectedBackgroundColor : scheduleListItem.style.backgroundColor
    anchors.fill: parent

    Behavior on color {
      ColorAnimation { duration: 150; easing.type: Easing.InOutQuad }
    }
  }

  RippleMouseArea {
    id: cellArea
    enabled: scheduleListItem.enabled
    anchors.fill: parent
    circularBackground: false
    clip: true

    onReleased: selected(typeof index !== "undefined" ? index : 0)

    // use selected background color of the style for the ripple background
    backgroundColor: setAlpha(scheduleListItem.style.selectedBackgroundColor, 0.5)
    fillColor: setAlpha(scheduleListItem.style.selectedBackgroundColor, 0.5)

    function setAlpha(color, alpha) {
      return Qt.rgba(color.r,color.g,color.b, alpha)
    }
  }

  // track indicators
  Column {
    id: trackCol
    x: dp(2)
    width: dp(4)
    height: parent.height - 6 * x
    spacing: x
    anchors.verticalCenter: parent.verticalCenter

    Repeater {
      id: repeater
      model: item.tracks
      Rectangle {
        id: trackIndicator
        width: parent.width
        height: ((parent.height - (trackCol.spacing * (repeater.count - 1))) / repeater.count)
        color: loaderItem.getTrackColor(modelData)
      }
    }
  }

  // wrapper for grid to center it vertically
  Item {
    id: gridWrapper
    y: Math.max(0, (dp(scheduleListItem.style.minimumHeight) - innerGrid.height)) * 0.5
    width: parent.width
    height: innerGrid.height

    // Main cell content inside this item
    GridLayout {
      id: innerGrid
      property real trackIndent: trackCol.width + trackCol.x // additional indent for track rectangles

      // Auto-break after 3 columns, so we do not have to set row & column manually
      columns: 4
      rowSpacing: dp(5)
      columnSpacing: dp(scheduleListItem.style.indent)
      x: trackIndent + dp(scheduleListItem.style.indent) * 0.75
      width: parent.width - x - dp(scheduleListItem.style.indent)
      Layout.minimumWidth: parent.width - x - dp(scheduleListItem.style.indent)
      Layout.maximumWidth: parent.width - x - dp(scheduleListItem.style.indent)

      // Top spacer
      Item {
        id: topSpacer
        Layout.preferredWidth: parent.width
        Layout.preferredHeight: dp(scheduleListItem.style.spacing) - parent.rowSpacing
        Layout.minimumHeight: 0
        Layout.fillWidth: true
        Layout.columnSpan: 4
      }

      // speaker image (overlay on top of SimpleRow)
      SpeakerImage {
        id: speakerImg
        Layout.preferredWidth: height
        Layout.preferredHeight: dp(scheduleListItem.style.fontSizeText) * 2.5
        Layout.alignment: Qt.AlignCenter
        Layout.rowSpan: 2

        property var speakerId: item.persons && item.persons.length > 0 ? item.persons[0]["id"] : undefined
        source: speakerId !== undefined && dataModel.speakers && dataModel.speakers[speakerId] ? dataModel.speakers[speakerId].avatar : ""
      } // SpeakerImage

      // title text
      Item {
        // wrapper item for text handles available space in layout
        visible: textLabel.text.length > 0
        Layout.columnSpan: 2
        Layout.preferredWidth: parent.width - 2 * parent.columnSpacing - speakerImg.width - favoriteIcon.width
        Layout.fillWidth: true
        Layout.preferredHeight: textLabel.height
        Layout.alignment: Qt.AlignVCenter

        // actual text item
        Text {
          id: textLabel
          elide: Text.ElideRight
          maximumLineCount: 2
          wrapMode: Text.WordWrap
          color: scheduleListItem.style.textColor
          font.family: Theme.normalFont.name
          font.pixelSize: sp(scheduleListItem.style.fontSizeText)
          text: item && item.title || ""

          // calculate additional margin needed to align with title in navigation bar
          property real titleAlignmentMargin:
              dp(Theme.navigationBar.titleLeftMargin + Theme.navigationBar.defaultIconSize + Theme.navigationBar.defaultBarItemPadding)
              - parent.x + dp(5)

          // add x offset margin to align with navigation bar title text
          x: Theme.navigationBar.titleAlignLeft && titleAlignmentMargin > 0 ? titleAlignmentMargin : 0
          width: parent.width - x
          anchors.verticalCenter: parent.verticalCenter
        }
      }

      // favorite icon
      Icon {
        id: favoriteIcon
        Layout.preferredWidth: implicitWidth
        Layout.preferredHeight: implicitHeight
        size: dp(20)
        Layout.alignment: Qt.AlignCenter
        icon: isFavorite ? IconType.star : IconType.staro
        color: isFavorite || favoriteArea.pressed ? Theme.tintColor : Theme.secondaryTextColor

        property bool isFavorite: item && item.id ? dataModel.isFavorite(item.id) : false
        property bool updateWhenVisible: false

        // update UI when favorites change while visible
        Connections {
          target: dataModel
          onFavoritesChanged: {
            if(visible)
              favoriteIcon.isFavorite = item && item.id ? dataModel.isFavorite(item.id) : false
            else
              favoriteIcon.updateWhenVisible = true
          }
        }

        // update UI when favorites changed while invisible
        onVisibleChanged: {
          if(visible && updateWhenVisible) {
            favoriteIcon.isFavorite = item && item.id ? dataModel.isFavorite(item.id) : false
            updateWhenVisible = false
          }
        }

        // add/remove from favorites
        RippleMouseArea {
          id: favoriteArea
          width: parent.width * 3
          height: parent.height * 3
          fixedPosition: true
          touchPoint: Qt.point(width * 0.5, height * 0.5)
          anchors.centerIn: parent
          onClicked: scheduleListItem.toggleFavorite()
        }
      }

      // time
      Item {
        visible: detailTextLabel.text.length > 0
        Layout.preferredWidth: parent.width - 2 * parent.columnSpacing - speakerImg.width - roomTxt.width
        Layout.fillWidth: true
        Layout.columnSpan: 1
        Layout.preferredHeight: detailTextLabel.height

        AppText {
          id: detailTextLabel
          color: scheduleListItem.style.detailTextColor
          font.family: Theme.normalFont.name
          font.pixelSize: sp(scheduleListItem.style.fontSizeDetailText)
          maximumLineCount: 1
          text: item.start+" - "+item.end

          x: textLabel.x
          width: parent.width - x
          anchors.verticalCenter: parent.verticalCenter
        }
      }

      // room
      AppText {
        id: roomTxt
        Layout.minimumWidth: implicitWidth
        Layout.preferredHeight: implicitHeight
        Layout.columnSpan: 2
        text: item.room && item.room !== "Qt World Summit 2017" ? item.room : ""
        font.pixelSize: sp(scheduleListItem.style.fontSizeDetailText)
        font.italic: true
        wrapMode: Text.NoWrap
        color: scheduleListItem.style.detailTextColor
      }

      Item {
        id: bottomSpacer
        Layout.preferredWidth: parent.width
        Layout.preferredHeight: dp(scheduleListItem.style.spacing) - parent.rowSpacing + divider.height
        Layout.minimumHeight: divider.height
        Layout.columnSpan: 4
        Layout.fillWidth: true
      }
    } // grid layout

  }

  // Bottom cell divider
  Rectangle {
    id: divider

    property real leftSpacing: speakerImg.width + dp(scheduleListItem.style.dividerLeftSpacing)

    x: (scheduleListItem.style.dividerLeftSpacing > 0) ? leftSpacing + dp(scheduleListItem.style.dividerLeftSpacing) : 0
    width: parent.width - x
    color: cellArea.pressed && Theme.isIos ? scheduleListItem.style.selectedBackgroundColor : scheduleListItem.style.dividerColor
    height: px(scheduleListItem.style.dividerHeight)
    anchors.bottom: parent.bottom
  }

  // add or remove item from favorites
  function toggleFavorite() {
    logic.toggleFavorite(item)
  }
}
