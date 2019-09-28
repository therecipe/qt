import Felgo 3.0
import QtQuick 2.0
import "../common"

Page {
  id: mypage
  backgroundColor: Theme.secondaryBackgroundColor

  property var playlist: null

  title: playlist && playlist.snippet.title || ""

  onPlaylistChanged: {
    if(!playlist)
      return

    logic.fetchPlaylistVideos(playlist)
  }

  rightBarItem: IconButtonBarItem {
    visible: !isOnline
    icon: IconType.exclamation
    onClicked: NativeDialog.confirm("Offline", "You are currently offline, please go online to use the app.")
  }

  AppListView {
    anchors.fill: parent
    model: JsonListModel {
      id: listModel
      source: playlist && dataModel.playlistVideos[playlist.id] || []
      keyField: "id"
    }
    delegate: FeedVideo {
      property var item: listModel.get(index)
      title: item.snippet.title
      description: item.snippet.description
      videoId: item.id
      videoStatistics: item.statistics
      previewImageSource: {
        // use best available quality as preview image
        var qualities = Object.keys(item.snippet.thumbnails)
        return item.snippet.thumbnails[qualities.pop()].url
      }
    }

    emptyText.text: getOfflineStatus()

    // when scrolling list, hide top-level video
    onContentYChanged: {
      // hide when scrolling is set to false for 1 second after video becomes active
      // this avoids the video to close in case touch interaction  affects contentY
      if(Theme.isIos && youTubePlayer.visible && youTubePlayer.hideWhenScrolling) {
        youTubePlayer.stop()
        youTubePlayer.visible = false
      }
    }

  }
}



