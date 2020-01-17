import Felgo 3.0
import QtQuick 2.0
import "../common"

FlickablePage {
  id: speakerDetailPage
  title: dataModel.speakers && dataModel.speakers[speakerID] ? dataModel.speakers[speakerID].first_name + " " + dataModel.speakers[speakerID].last_name : ""
  property string speakerID

  onPushed: amplitude.logEvent("Open Speaker Detail",{"title" : title, "speakerId" : speakerID})

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


  flickable.contentWidth: parent.width
  flickable.contentHeight: contentCol.height

  Column {
    id: contentCol
    width: parent.width

    Item {
      width: parent.width
      height: dp(Theme.navigationBar.defaultBarItemPadding)
    }

    Row {
      id: row1
      spacing: dp(Theme.navigationBar.defaultBarItemPadding)
      width: parent.width - dp(Theme.navigationBar.defaultBarItemPadding) * 2
      anchors.horizontalCenter: parent.horizontalCenter

      SpeakerImage {
        id: avatar
        source: dataModel.speakers && dataModel.speakers[speakerID] ? dataModel.speakers[speakerID].avatar : ""
        width: dp(40)
        height: width
        activatePictureViewer: true
      }
      AppText {
        width: parent.width - avatar.width - row1.spacing
        text: dataModel.speakers && dataModel.speakers[speakerID] ? dataModel.speakers[speakerID].first_name + " " + dataModel.speakers[speakerID].last_name : ""
        font.bold: true
        font.weight: Font.Bold
        font.family: Theme.boldFont.name
        anchors.verticalCenter: parent.verticalCenter
      }
    }

    Item {
      width: parent.width
      height: dp(Theme.navigationBar.defaultBarItemPadding) / 2
    }

    AppText {
      width: parent.width - dp(Theme.navigationBar.defaultBarItemPadding) * 2
      anchors.horizontalCenter: parent.horizontalCenter
      text: dataModel.speakers && dataModel.speakers[speakerID] ? dataModel.speakers[speakerID].abstract : ""
      wrapMode: AppText.WordWrap
    }

    // spacer / divider
    Rectangle {
      width: parent.width
      height: spacerDividerCol.height
      color: Theme.isIos ? Theme.secondaryBackgroundColor : Theme.backgroundColor

      Column {
        id: spacerDividerCol
        width: parent.width

        Rectangle {
          width: parent.width
          height: _.colSpacing
          color: Theme.backgroundColor
        }

        Rectangle {
          color: _.dividerColor
          width: parent.width
          height: px(1)
        }
      }
    }

    Item {
      width: parent.width
      height: dp(Theme.navigationBar.defaultBarItemPadding)
      visible: Theme.isAndroid
    }

    SimpleSection {
      title: "Talks"
    }

    Repeater {
      id: speakerRepeater
      model: dataModel.speakers && dataModel.speakers[speakerID] ? dataModel.speakers[speakerID].talks : []
      delegate: TalkRow {
        talk: dataModel.talks && dataModel.talks[modelData]
        isListItem: true
        small: true
        onClicked: {
          speakerDetailPage.navigationStack.push(Qt.resolvedUrl("DetailPage.qml"), { item: talk })
        }
        onFavoriteClicked: {
          toggleFavorite(talk)
        }
      }
    }
  }

  function toggleFavorite(item) {
    logic.toggleFavorite(item)
  }
}
