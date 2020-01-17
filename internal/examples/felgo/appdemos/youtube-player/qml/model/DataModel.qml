import QtQuick 2.0
import Felgo 3.0

Item {
  id: dataModel

  // property to configure target dispatcher / logic
  property alias dispatcher: logicConnection.target

  // configuration properties
  readonly property alias userName: _.userName
  readonly property alias channelId: _.channelId
  readonly property alias mainPlaylistId: _.mainPlaylistId
  readonly property alias uploadsPlaylistId: _.uploadsPlaylistId

  // data properties
  readonly property alias channelInfo: _.channelInfo
  readonly property alias channelPlaylists: _.channelPlaylists
  readonly property alias playlistVideos: _.playlistVideos

  // spotlight playlist
  readonly property alias mainPlaylist: _.mainPlaylist
  readonly property alias mainPlaylistItems: _.mainPlaylistItems
  readonly property alias mainPlaylistVideos: _.mainPlaylistVideos

  readonly property alias requestsPending: api.requestsPending

  Connections {
    id: logicConnection

    // action 1 - loadYouTubeData
    onLoadYouTubeData: {
      if(!_.userName)
        return

      _.playlistVideos = { }

      api.getChannelForUsername(_.userName, function(channelData) {
        if(!channelData)
          return

        // store channel data
        _.channelInfo = channelData
        _.channelId = channelData.id
        _.uploadsPlaylistId = channelData.contentDetails.relatedPlaylists.uploads

        // get channel playlists
        api.getChannelPlaylists(_.channelId, function(playlists) {
          _.channelPlaylists = playlists

          // get uploads playlist
          api.getPlaylistById(_.uploadsPlaylistId, function(playlist) {
            _.channelPlaylists.unshift(playlist)
            _.channelPlaylistsChanged()
          })
        }) // get playlists
      }) // get channel
    }

    // action 2 - loadMainPlaylistVideos
    onLoadMainPlaylistVideos: {
      if(_.mainPlaylistId) {
        // get items of main playlist
        api.getPlaylistItems(_.mainPlaylistId, function(items) {
          _.mainPlaylistItems = items

          // get video data for playlist items
          var videoIds = _.mainPlaylistItems.map(function(item) { return item.contentDetails.videoId })
          api.getVideos(videoIds, function(videos) {
            _.mainPlaylistVideos = videos
          }) // get spotlight videos
        }) // get spotlight items
      }
    }

    // action 3 - getPlaylistVideos
    onFetchPlaylistVideos: {
      api.getPlaylistItems(playlist.id, function(items) {
        var videoIds = items.map(function(item) { return item.contentDetails.videoId })
        api.getVideos(videoIds, function(videos) {
          _.playlistVideos[playlist.id] = videos
          _.playlistVideosChanged()
        })
      })
    }

    // action 4 - clearData
    onClearData: {
      _.channelId = ""
      _.mainPlaylistId = ""

      _.channelInfo = undefined
      _.channelPlaylists = []
      _.mainPlaylist = undefined
      _.mainPlaylistItems = []
      _.mainPlaylistVideos = []
      _.playlistVideos = { }
    }

    // action 5 - setSpotlightPlaylist
    onSetSpotlightPlaylist: {
      if(playlist.id === "") {
        _.mainPlaylistId = ""
        _.mainPlaylist = undefined
        _.mainPlaylistItems = undefined
        _.mainPlaylistVideos = undefined
      }
      else {
        _.mainPlaylistId = playlist.id
        _.mainPlaylist = playlist
      }
    }

    // action 6 - setUserName
    onSetUserName: {
      _.userName = name
    }
  }

  Storage {
    id: storage
    Component.onCompleted: {
      _.userName = storage.getValue("userName") || _.userName
      _.mainPlaylistId = storage.getValue("mainPlaylistId") || ""
    }
  }

  YouTubeAPI {
    id: api
  }

  // private
  Item {
    id: _

    // configuration properties
    property string userName: "VPlayEngine"
    property string channelId: ""
    property string mainPlaylistId: ""
    property string uploadsPlaylistId: ""

    // data properties
    property var channelInfo: undefined
    property var channelPlaylists: []
    property var playlistVideos: null

    // spotlight playlist
    property var mainPlaylist: undefined
    property var mainPlaylistItems: undefined
    property var mainPlaylistVideos: undefined

    // try to retrieve main playlist if playlists change
    onChannelPlaylistsChanged: {
      var newMainPlaylist = undefined
      for(var i = 0; i < _.channelPlaylists.length; i++) {
        if(_.channelPlaylists[i].id === _.mainPlaylistId)
          newMainPlaylist = _.channelPlaylists[i]
      }
      _.mainPlaylist = newMainPlaylist
    }

    // data storage
    onUserNameChanged: storage.setValue("userName", _.userName)
    onMainPlaylistIdChanged: storage.setValue("mainPlaylistId", _.mainPlaylistId)
  }
}
