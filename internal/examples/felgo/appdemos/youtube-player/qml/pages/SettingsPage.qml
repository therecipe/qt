import Felgo 3.0
import QtQuick 2.0
import QtQuick.Layouts 1.3
import "../common"

Page {
  id: page
  title: "Settings"

  rightBarItem: IconButtonBarItem {
    visible: !isOnline
    icon: IconType.exclamation
    onClicked: NativeDialog.confirm("Offline", "You are currently offline, please go online to use the app.")
  }

  Column {
    id: settingsCol
    y: spacing
    width: parent.width - 2 * spacing
    spacing: dp(Theme.navigationBar.defaultBarItemPadding)
    anchors.horizontalCenter: parent.horizontalCenter

    Row {
      spacing: parent.spacing

      AppText {
        text: "User Name:"
        anchors.verticalCenter: parent.verticalCenter
      }

      AppTextField {
        id: userNameField
        text: dataModel.userName
        placeholderText: "Enter User Name"
        anchors.verticalCenter: parent.verticalCenter
        backgroundColor: Theme.secondaryBackgroundColor
        onAccepted: {
          logic.clearData()
          logic.setUserName(userNameField.text)
          page.forceActiveFocus()
        }
      }
    }

    // channel info
    AppText {
      id: infoText
      text: "Channel: "+(dataModel.channelInfo ? dataModel.channelInfo.snippet.title : "")
      visible: !!dataModel.channelInfo
      color: Theme.secondaryTextColor
    }

    // divider
    Rectangle {
      width: parent.width
      height: px(1)
      color: Theme.dividerColor
    }

    // spotlight
    Item {
      id: spotlightItem
      width: parent.width
      height: spotlightRow.height
      visible: dataModel.channelPlaylists.length > 0

      RowLayout {
        id: spotlightRow
        width: parent.width
        spacing: settingsCol.spacing

        AppText {
          Layout.minimumWidth: implicitWidth
          text: "Spotlight Playlist:"
        }

        AppText {
          text: (dataModel.mainPlaylist && dataModel.mainPlaylist.snippet && dataModel.mainPlaylist.snippet.title || "none")
          color: Theme.secondaryTextColor
          Layout.fillWidth: true
          wrapMode: AppText.WrapAtWordBoundaryOrAnywhere
        }

        Icon {
          Layout.minimumWidth: implicitWidth
          icon: IconType.edit
          color: Theme.tintColor
        }
      }

      MouseArea {
        anchors.fill: parent
        onClicked: page.navigationStack.push(spotlightChooserComponent)
      }
    }

    // divider
    Rectangle {
      width: parent.width
      height: px(1)
      color: Theme.dividerColor
      visible: spotlightItem.visible
    }
  }

  Component {
    id: spotlightChooserComponent

    ChannelPage {
      id: chooser
      title: "Choose Playlist"
      model: dataModel.channelPlaylists.concat(noPlaylistOption)
      showDisclosure: false
      onSelected: {
        logic.setSpotlightPlaylist(playlist)
        chooser.navigationStack.pop()
      }

      property var noPlaylistOption: [{
          id: "",
          snippet: { title: "No Spotlight Playlist" }
        }]
    }
  }
}


