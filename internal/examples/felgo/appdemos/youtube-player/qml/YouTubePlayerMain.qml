import Felgo 3.0
import QtQuick 2.0
import "pages"
import "common"
import "model"

App {
  id: app
  // You get free licenseKeys from https://felgo.com/licenseKey
  // With a licenseKey you can:
  //  * Publish your games & apps for the app stores
  //  * Remove the Felgo Splash Screen or set a custom one (available with the Pro Licenses)
  //  * Add plugins to monetize, analyze & improve your apps (available with the Pro Licenses)
  //licenseKey: "<generate one from https://felgo.com/licenseKey>"

  DataModel {
    id: dataModel
    dispatcher: logic
  }

  Logic {
    id: logic
  }

  // data handling
  onIsOnlineChanged: Qt.callLater(logic.loadYouTubeData)

  Connections {
    target: dataModel
    onUserNameChanged: Qt.callLater(logic.loadYouTubeData)
    onMainPlaylistChanged: Qt.callLater(logic.loadMainPlaylistVideos)
  }

  Component.onCompleted: Qt.callLater(logic.loadYouTubeData)

  // theming
  onInitTheme: {
    var darkBgColor = "#161614"
    var darkTextColor = "white"

    Theme.navigationBar.backgroundColor = darkBgColor
    Theme.navigationBar.titleColor = darkTextColor
    Theme.navigationBar.itemColor = darkTextColor

    Theme.tabBar.backgroundColor = darkBgColor
    Theme.tabBar.titleColor = darkTextColor

    Theme.colors.secondaryBackgroundColor = "#EFEFF4"

    Theme.colors.statusBarStyle = Theme.colors.statusBarStyleWhite
  }

  // app navigaiton
  Navigation {
    id: navigation

    NavigationItem {
      icon: IconType.star
      title: "Spotlight"

      NavigationStack {
        PlaylistPage {
          title: "Spotlight"
          playlist: dataModel.mainPlaylist
          //model: dataModel.mainPlaylistVideos || []
        }
      }
      showItem: !!dataModel.mainPlaylistId
      onShowItemChanged: navigation.currentIndex = 0 // open spotlight
    }

    NavigationItem {
      icon: IconType.compass
      title: "Browse"

      NavigationStack {
        ChannelPage {
          id: channelPage

          // push selected playlist when clicked
          onSelected: channelPage.navigationStack.push(Qt.resolvedUrl("pages/PlaylistPage.qml"), { playlist: playlist })
        }
      }
    }

    NavigationItem {
      icon: IconType.gear
      title: "Settings"

      NavigationStack {
        SettingsPage { }
      }
    }
  }

  // YouTube Player, will be re-parented to video items of feed as required
  property alias youTubePlayer: player

  YouTubeWebPlayer {
    id: player
    visible: false
    onVisibleChanged: {
      if(visible) {
        hideWhenScrolling = false
        videoStartedTimer.start()
      }
    }

    // if users scrolls one second after video was started, hide the player
    property bool hideWhenScrolling: false
    Timer {
      id: videoStartedTimer
      interval: 1000
      onTriggered: player.hideWhenScrolling = true
    }
  }

  // offline status to show for empty list views
  function getOfflineStatus() {
    if(dataModel.requestsPending)
      return "loading ..."
    else if(!isOnline)
      return "You are offline. Please go online to use the app."
    else if (dataModel.channelId === "")
      return "Please log-in with your user name."
    else
      return "No videos."
  }
}


