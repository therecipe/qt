import Felgo 3.0
import QtQuick 2.0
import QtQuick.Layouts 1.1
import "../common"

Rectangle {
  id: container
  width: parent.width
  height: isListItem ? talkRow.height + dp(Theme.navigationBar.defaultBarItemPadding) : talkRow.height
  color: isListItem ? Theme.listItem.backgroundColor : "transparent"

  signal clicked
  signal favoriteClicked
  signal roomClicked
  signal trackClicked(var track)
  property var talk
  property bool isFavorite: talk && talk.id ? dataModel.isFavorite(talk.id) : false
  property bool small: false
  property bool isListItem: false

  // private members
  Item {
    id: _
    readonly property color dividerColor: Theme.dividerColor
    readonly property color iconColor:  Theme.secondaryTextColor
    readonly property var rowSpacing: dp(10)
    readonly property var colSpacing: dp(20)
    readonly property real speakerImgSize: dp(Theme.navigationBar.defaultIconSize) * 4
    readonly property real speakerTxtWidth: sp(150)
    readonly property real favoriteTxtWidth: sp(150)
    property int loadingCount: 0
  }

  RippleMouseArea {
    anchors.fill: parent
    circularBackground: false
    enabled: isListItem
    onClicked: {
      container.clicked()
    }
  }

  Column {
    id: talkRow
    width: parent.width
    spacing: dp(Theme.navigationBar.defaultBarItemPadding)
    anchors.verticalCenter: parent.verticalCenter

    Row {
      visible: small
      width: parent.width - dp(Theme.navigationBar.defaultBarItemPadding) * 2
      spacing: dp(Theme.navigationBar.defaultBarItemPadding)
      anchors.horizontalCenter: parent.horizontalCenter

      AppText {
        text: talk && talk.title || ""
        width: parent.width - parent.spacing - favoriteIconSmall.width
        anchors.verticalCenter: parent.verticalCenter
        wrapMode: Text.WordWrap
        font.bold: true
        font.weight: Font.Bold
        font.family: Theme.boldFont.name
      }
      // favorite icon
      Icon {
        id: favoriteIconSmall
        icon: isFavorite ? IconType.star : IconType.staro
        color: isFavorite || favoriteArea.pressed ? Theme.tintColor : _.iconColor
        anchors.verticalCenter: parent.verticalCenter
        size: dp(20)

        // add/remove from favorites
        RippleMouseArea {
          enabled: small
          width: parent.width * 3
          height: parent.height * 3
          fixedPosition: true
          touchPoint: Qt.point(width * 0.5, height * 0.5)
          anchors.centerIn: parent
          onClicked: container.favoriteClicked()
        }
      }
    }

    Column {
      width: parent.width
      spacing: small ? talkRow.spacing / 2 : talkRow.spacing

      Column {
        width: parent.width
        spacing: small ? talkRow.spacing / 2 : talkRow.spacing
        // tracks
        Repeater {
          model: talk && talk.tracks ? talk.tracks : []
          delegate: Row {
            x: grid.x
            width: talkRow.width - x
            spacing: talkRow.spacing

            // track
            Icon {
              id: trackIcon
              anchors.verticalCenter: parent.verticalCenter
              icon: IconType.tag
              color: loaderItem.getTrackColor(modelData)
            }

            AppText {
              id: trackText
              anchors.verticalCenter: parent.verticalCenter
              width: parent.width - trackIcon.width - _.colSpacing
              wrapMode: Text.WordWrap
              text: modelData
              color: trackIcon.color

              RippleMouseArea {
                height: parent.height * 2
                width: parent.width + trackIcon.width + spacing
                anchors.right: parent.right
                anchors.verticalCenter: parent.verticalCenter
                fixedPosition: true
                centerAnimation: true
                touchPoint: Qt.point(trackIcon.width * 0.5, trackIcon.height)
                enabled: !isListItem
                onClicked: container.trackClicked(trackText.text)
              }
            }
          }
        }
      }

      // icon grid
      GridLayout {
        id: grid
        width: Math.min(implicitWidth, parent.width - dp(Theme.navigationBar.defaultBarItemPadding) * 2)
        x: talkRow.spacing
        columns: 2
        rowSpacing: small ? talkRow.spacing / 2 : talkRow.spacing
        columnSpacing: talkRow.spacing

        // time
        Icon {
          Layout.preferredWidth: implicitWidth
          Layout.alignment: Qt.AlignVCenter
          icon: IconType.clocko
          color: _.iconColor
        }

        AppText {
          Layout.fillWidth: true
          Layout.alignment: Qt.AlignVCenter
          wrapMode: Text.WordWrap
          color: _.iconColor
          text: getWeekDay() + "  " + talk.start+" - "+talk.end
          function getWeekDay() {
            var days = ['Sun','Mon','Tue','Wed','Thu','Fri','Sat'];
            var date = new Date(talk.day+"T00:00.000Z")
            return days[ date.getUTCDay() ]
          }
        }

        // date
        /*Icon {
          Layout.preferredWidth: implicitWidth
          Layout.alignment: Qt.AlignVCenter
          icon: IconType.calendaro
          color: _.iconColor
        }

        AppText {
          Layout.fillWidth: true
          Layout.alignment: Qt.AlignVCenter
          wrapMode: Text.WordWrap
          color: _.iconColor
          text: talk.day
        }*/

        // language
        Icon {
          Layout.preferredWidth: implicitWidth
          Layout.alignment: Qt.AlignVCenter
          icon: IconType.flag
          color: _.iconColor
          visible: langTxt.visible
        }

        AppText {
          id: langTxt
          Layout.fillWidth: true
          Layout.alignment: Qt.AlignVCenter
          wrapMode: Text.WordWrap
          text: talk.language
          color: _.iconColor
          visible: talk.language.length > 0
        }

        // type
        Icon {
          Layout.preferredWidth: implicitWidth
          Layout.alignment: Qt.AlignVCenter
          icon: IconType.dotcircleo
          color: _.iconColor
          visible: typeTxt.visible
        }

        AppText {
          id: typeTxt
          Layout.fillWidth: true
          Layout.alignment: Qt.AlignVCenter
          wrapMode: Text.WordWrap
          text: talk.type
          color: _.iconColor
          visible: talk.type.length > 0
        }

        // room
        Icon {
          id: roomIcon
          Layout.preferredWidth: implicitWidth
          Layout.alignment: Qt.AlignVCenter
          icon: IconType.mapmarker
          color: roomTxt.color
          visible: roomTxt.visible
        }

        AppText {
          id: roomTxt
          Layout.fillWidth: true
          Layout.alignment: Qt.AlignVCenter
          wrapMode: Text.WordWrap
          text: isListItem ? talk.room : "Show " + talk.room
          color: isListItem ? _.iconColor : Theme.tintColor
          visible: talk.room.length > 0

          RippleMouseArea {
            height: parent.height * 2
            width: parent.width + roomIcon.width + grid.rowSpacing
            anchors.right: parent.right
            anchors.verticalCenter: parent.verticalCenter
            fixedPosition: true
            centerAnimation: true
            touchPoint: Qt.point(0, roomIcon.height)
            enabled: !isListItem
            onClicked: container.roomClicked()
          }
        }

        // favorite
        Item {
          Layout.columnSpan: 2
          Layout.minimumWidth: favoriteIcon.implicitWidth + talkRow.spacing + _.favoriteTxtWidth
          Layout.fillWidth: true
          Layout.preferredHeight: Math.max(favoriteIcon.height, favoriteTxt.height)
          visible: !small

          Icon {
            id: favoriteIcon
            icon: isFavorite ? IconType.star : IconType.staro
            color: Theme.tintColor//isFavorite || favoriteArea.pressed ? Theme.tintColor : _.iconColor
            anchors.verticalCenter: parent.verticalCenter
          }

          AppText {
            id: favoriteTxt
            width: _.favoriteTxtWidth
            anchors.left: favoriteIcon.right
            anchors.leftMargin: talkRow.spacing
            anchors.right: parent.right
            wrapMode: Text.WordWrap
            text: isFavorite ? "Favorited" : "Add to Favorites"
            color: Theme.tintColor
            anchors.verticalCenter: parent.verticalCenter
          }

          RippleMouseArea {
            id: favoriteArea
            enabled: !small
            width: parent.width
            height: parent.height * 2
            fixedPosition: true
            centerAnimation: true
            touchPoint: Qt.point(favoriteIcon.width * 0.5, favoriteIcon.height)
            anchors.centerIn: parent
            onClicked: container.favoriteClicked()
          }
        }

        // website
        Item {
          Layout.columnSpan: 2
          Layout.minimumWidth: webIcon.implicitWidth + talkRow.spacing + _.favoriteTxtWidth
          Layout.fillWidth: true
          Layout.preferredHeight: Math.max(webIcon.height, webTxt.height)
          visible: !small

          Icon {
            id: webIcon
            icon: IconType.globe
            color: Theme.tintColor
            anchors.verticalCenter: parent.verticalCenter
          }

          AppText {
            id: webTxt
            width: _.favoriteTxtWidth
            anchors.left: webIcon.right
            anchors.leftMargin: talkRow.spacing
            anchors.right: parent.right
            wrapMode: Text.WordWrap
            text: "Open in Browser"
            color: Theme.tintColor
            anchors.verticalCenter: parent.verticalCenter
          }

          RippleMouseArea {
            enabled: !small
            width: parent.width
            height: parent.height * 2
            fixedPosition: true
            centerAnimation: true
            touchPoint: Qt.point(webIcon.width * 0.5, webIcon.height)
            anchors.centerIn: parent
            onClicked: {
              amplitude.logEvent("Open Talk In Browser", {"title" : talk.title, "talkId" : talk.id, "url" : "http://www.qtworldsummit.com/sessions/"+talk.slug})
              nativeUtils.openUrl("http://www.qtworldsummit.com/sessions/"+talk.slug)
            }
          }
        }

        // share
        Item {
          Layout.columnSpan: 2
          Layout.minimumWidth: shareIcon.implicitWidth + talkRow.spacing + _.favoriteTxtWidth
          Layout.fillWidth: true
          Layout.preferredHeight: Math.max(shareIcon.height, shareTxt.height)
          visible: !small

          Icon {
            id: shareIcon
            icon: IconType.share
            color: Theme.tintColor
            anchors.verticalCenter: parent.verticalCenter
          }

          AppText {
            id: shareTxt
            width: _.favoriteTxtWidth
            anchors.left: shareIcon.right
            anchors.leftMargin: talkRow.spacing
            anchors.right: parent.right
            wrapMode: Text.WordWrap
            text: "Share this Session"
            color: Theme.tintColor
            anchors.verticalCenter: parent.verticalCenter
          }

          RippleMouseArea {
            enabled: !small
            width: parent.width
            height: parent.height * 2
            fixedPosition: true
            centerAnimation: true
            touchPoint: Qt.point(shareIcon.width * 0.5, shareIcon.height)
            anchors.centerIn: parent
            onClicked: {
              amplitude.logEvent("Share Talk", {"title" : talk.title, "talkId" : talk.id})
              nativeUtils.share("I am attending \"" + talk.title + "\" at Qt World Summit 2017!","http://www.qtworldsummit.com/sessions/"+talk.slug)
            }
          }
        }
      } // icon grid
    }
  }

  Rectangle {
    anchors.right: parent.right
    anchors.bottom: parent.bottom
    width: Theme.isIos ? parent.width - dp(Theme.navigationBar.defaultBarItemPadding) : parent.width
    color: Theme.listItem.dividerColor
    height: px(1)
    visible: isListItem
  }

  Icon {
    icon: IconType.angleright
    anchors.right: parent.right
    anchors.rightMargin: Theme.navigationBar.defaultBarItemPadding
    anchors.verticalCenter: parent.verticalCenter
    color: Theme.dividerColor
    visible: Theme.isIos && isListItem
    size: dp(24)
  }
}
