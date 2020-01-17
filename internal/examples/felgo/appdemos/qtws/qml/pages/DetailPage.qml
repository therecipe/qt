import Felgo 3.0
import QtQuick 2.0
import QtQuick.Layouts 1.1
import "../common"

FlickablePage {
  id: detailPage
  title: item.title
  rightBarItem: NavigationBarRow {
    showMoreButton: false
    ActivityIndicatorBarItem { visible: dataModel.loading || detailPage.loading ? true : false }
    IconButtonBarItem {
      icon: detailPage.isFavorite ? IconType.star : IconType.staro
      onClicked: detailPage.toggleFavorite()
      showItem: showItemAlways
    }
    IconButtonBarItem {
      icon: IconType.mapmarker
      onClicked: detailPage.navigationStack.push(Qt.resolvedUrl("RoomPage.qml"), { room: item.room })
      showItem: showItemAlways
    }
  }

  onPushed: amplitude.logEvent("Open Talk Detail", {"title" : title, "talkId" : item.id})

  property var item
  property bool isFavorite: item && item.id ? dataModel.isFavorite(item.id) : false
  readonly property bool loading: _.loadingCount > 0

  // update UI when favorites change
  Connections {
    target: dataModel
    onFavoritesChanged: detailPage.isFavorite = item && item.id ? dataModel.isFavorite(item.id) : false
  }

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

  flickable.contentHeight: contentCol.height + 2 * contentCol.y

  // content column
  Column {
    id: contentCol
    //x: 2 * spacing
    y: 2 * spacing
    width: parent.width// - 2 * x
    //spacing: _.rowSpacing

    Item {
      width: parent.width
      height: _.colSpacing
    }

    // title
    AppText {
      text: item.title
      width: parent.width - dp(Theme.navigationBar.defaultBarItemPadding) * 2
      anchors.horizontalCenter: parent.horizontalCenter
      wrapMode: Text.WordWrap
      font.bold: true
      font.weight: Font.Bold
      font.family: Theme.boldFont.name
    }

    // subtitle
    AppText {
      text: item.subtitle
      color: Theme.secondaryTextColor
      width: parent.width - dp(Theme.navigationBar.defaultBarItemPadding) * 2
      anchors.horizontalCenter: parent.horizontalCenter
      wrapMode: Text.WordWrap
      visible: text.length > 0
      font.pixelSize: sp(14)
    }

    // divider
    Item {
      width: parent.width
      height: _.colSpacing * 2
      Rectangle {
        color: _.dividerColor
        width: parent.width
        height: px(1)
        anchors.centerIn: parent
      }
    }

    Column {
      anchors.horizontalCenter: parent.horizontalCenter
      width: Math.min(parent.width, mainSpeakerRow + dp(50))
      Row {
        id: mainSpeakerRow
        //width: parent.width - dp(Theme.navigationBar.defaultBarItemPadding) * 2
        anchors.horizontalCenter: parent.horizontalCenter
        visible: item && item.persons ? true : false
        spacing: dp(Theme.navigationBar.defaultBarItemPadding)

        // speaker image
        SpeakerImage {
          id: speakerImg
          width:  _.speakerImgSize
          height: width
          anchors.verticalCenter: parent.verticalCenter
          source: item && item.persons && dataModel.speakers && dataModel.speakers[item.persons[0].id]
                  ? dataModel.speakers[item.persons[0].id].avatar : ""
          onLoadingChanged: {
            if(loading)
              _.loadingCount++
            else
              _.loadingCount--
          }
          activatePictureViewer: true
        }

        // speaker name
        AppText {
          id: speakerTxt
          width: _.speakerTxtWidth
          anchors.verticalCenter: parent.verticalCenter
          wrapMode: Text.WordWrap
          text: item && item.persons ? item.persons[0]["full_public_name"] : ""
          RippleMouseArea {
            width: mainSpeakerRow.width
            height: mainSpeakerRow.height
            anchors.right: parent.right
            anchors.verticalCenter: parent.verticalCenter
            fixedPosition: true
            centerAnimation: true
            touchPoint: Qt.point(speakerImg.width * 0.5, height * 0.5)
            onClicked: {
              detailPage.navigationStack.push(Qt.resolvedUrl("SpeakerDetailPage.qml"), { speakerID: item.persons[0].id })
            }
          }
        }
      }

      Item {
        width: parent.width
        height: _.colSpacing
      }

      TalkRow {
        talk: item
        onFavoriteClicked: toggleFavorite()
        onRoomClicked: detailPage.navigationStack.push(Qt.resolvedUrl("RoomPage.qml"), { room: item.room })
        onTrackClicked: {
          var obj = {}
          obj[track] = 0
          var model = viewHelper.prepareTracks(obj)
          console.debug(JSON.stringify(model))
          if(Theme.isAndroid)
            detailPage.navigationStack.push(Qt.resolvedUrl("TrackDetailPage.qml"), { track: model[0] })
          else
            detailPage.navigationStack.push(Qt.resolvedUrl("TrackDetailPage.qml"), { track: model[0] })
        }
      }
    }

    // divider
    Item {
      width: parent.width
      height: _.colSpacing
      Rectangle {
        color: _.dividerColor
        width: parent.width
        height: px(1)
        anchors.bottom: parent.bottom
      }
    }


    SimpleSection {
      title: "Rate This Talk"
    }

    // spacing
    Item {
      width: parent.width
      height: _.colSpacing
    }

    RatingRow {
      width: parent.width - dp(Theme.navigationBar.defaultBarItemPadding) * 2
      anchors.horizontalCenter: parent.horizontalCenter
      talk: item
    }

    // divider
    Item {
      width: parent.width
      height: _.colSpacing
      visible: abstractTxt.visible
      Rectangle {
        color: _.dividerColor
        width: parent.width
        height: px(1)
        anchors.bottom: parent.bottom
      }
    }

    SimpleSection {
      title: "Abstract"
      visible: abstractTxt.visible
    }

    // spacing
    Item {
      width: parent.width
      height: _.rowSpacing
      visible: abstractTxt.visible
    }

    // abstract
    AppText {
      id: abstractTxt
      text: '"' + item.abstract + '"'
      width: parent.width - dp(Theme.navigationBar.defaultBarItemPadding) * 2
      anchors.horizontalCenter: parent.horizontalCenter
      wrapMode: Text.WordWrap
      horizontalAlignment: Text.AlignHCenter
      font.italic: true
      visible: item.abstract.length > 0
    }

    // divider
    Item {
      width: parent.width
      height: _.colSpacing
      Rectangle {
        color: _.dividerColor
        width: parent.width
        height: px(1)
        anchors.bottom: parent.bottom
      }
    }

    SimpleSection {
      title: "Description"
    }

    // spacing
    Item {
      width: parent.width
      height: _.rowSpacing
    }

    // description
    AppText {
      id: descriptionTxt
      text: item.description
      width: parent.width - dp(Theme.navigationBar.defaultBarItemPadding) * 2
      anchors.horizontalCenter: parent.horizontalCenter
      wrapMode: Text.WordWrap
      font.pixelSize: sp(15)
      visible: item.description.length > 0
    }

    // spacing
    Item {
      width: parent.width
      height: 1
    }

    // id note
    Rectangle {
      width: parent.width
      height: idNote.height
      color: Theme.isIos ? Theme.secondaryBackgroundColor : Theme.backgroundColor

      Column {
        id: idNote
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

    Column {
      width: parent.width

      Item {
        width: parent.width
        height: dp(Theme.navigationBar.defaultBarItemPadding)
        visible: Theme.isAndroid
      }

      SimpleSection {
        title: "Speakers"
      }

      Repeater {
        id: speakerRepeater
        model: item && item.persons ? item.persons : []
        delegate: SpeakerRow {
          speaker: dataModel.speakers && dataModel.speakers[modelData.id]
          onClicked: {
            detailPage.navigationStack.push(Qt.resolvedUrl("SpeakerDetailPage.qml"), { speakerID: modelData.id })
          }
        }
      }
    }
  } // content column

  // add or remove item from favorites
  function toggleFavorite() {
    logic.toggleFavorite(item)
  }
} // page
