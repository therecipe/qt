import QtQuick 2.0
import Felgo 3.0

Item {
  id: api

  // configure Google Project API Key for YouTube Data API
  property string apiKey: "AIzaSyD3FPmFGDTyO_wHiUVMYSb7MQ4eFeveJw4"

  property bool requestsPending: HttpNetworkActivityIndicator.enabled

  Component.onCompleted: {
    HttpNetworkActivityIndicator.setActivationDelay(0)
  }

  // getChannelIdForUsername
  function getChannelForUsername(name, success, error) {
    var reqUrl = _.getChannelForUsernameRequest.arg(name).arg(apiKey);
    _.sendGetRequest(reqUrl, function(data) {
      if(data && data.items[0])
        success(data.items[0])
      else
        success(undefined)
    }, error)
  }

  // getChannelPlaylists
  function getChannelPlaylists(channelId, success, error) {
    var reqUrl = _.getChannelPlaylistsRequest.arg(channelId).arg(apiKey);
    _.sendGetRequest(reqUrl, function(data) {
      if(data && data.items)
        success(data.items)
      else
        success([])
    }, error)
  }

  function getPlaylistById(playlistId, success, error) {
    var reqUrl = _.getPlaylistByIdRequest.arg(playlistId).arg(apiKey);
    _.sendGetRequest(reqUrl, function(data) {
      if(data && data.items[0])
        success(data.items[0])
      else
        success(undefined)
    }, error)
  }

  // getPlaylistItems
  function getPlaylistItems(playlistId, success, error) {
    var reqUrl = _.getPlaylistItemsRequest.arg(playlistId).arg(apiKey);
    _.sendGetRequest(reqUrl, function(data) {
      if(data && data.items)
        success(data.items)
      else
        success([])
    }, error)
  }

  // getVideos
  function getVideos(videoIds, success, error) {
    var idParam = ""
    if(videoIds.length)  // it's an array
      idParam = videoIds.join()
    else
      idParam = videoIds

    var reqUrl = _.getVideosRequest.arg(idParam).arg(apiKey);
    _.sendGetRequest(reqUrl, function(data) {
      if(data && data.items)
        success(data.items)
      else
        success([])
    }, error)
  }

  // item for private members
  QtObject {
    id: _

    // YouTube API URLs
    property string getChannelForUsernameRequest: "https://www.googleapis.com/youtube/v3/channels?part=id,snippet,brandingSettings,contentDetails&forUsername=%1&key=%2"
    property string getChannelPlaylistsRequest: "https://www.googleapis.com/youtube/v3/playlists?part=snippet,contentDetails&channelId=%1&maxResults=50&key=%2"
    property string getPlaylistByIdRequest: "https://www.googleapis.com/youtube/v3/playlists?part=snippet,contentDetails&id=%1&key=%2"
    property string getPlaylistItemsRequest: "https://www.googleapis.com/youtube/v3/playlistItems?part=snippet,contentDetails,status&playlistId=%1&maxResults=50&key=%2"
    property string getVideosRequest: "https://www.googleapis.com/youtube/v3/videos?part=snippet,contentDetails,status,statistics&id=%1&maxResults=50&key=%2"

    // sendGetRequest - load data from url with success handler
    function sendGetRequest(url, successHandler, errorHandler) {
      HttpRequest.get(url)
      .then(function(response) {
        successHandler(response.body)
      })
      .catch(function(error) {
        console.error("Error: Failed to load data from "+url+", status = "+xmlHttpReq.status+", response = "+XMLHttpRequest.responseText)
        if(errorHandler !== undefined)
          errorHandler(error.response.status, error.message)
      })
    }
  }
}
