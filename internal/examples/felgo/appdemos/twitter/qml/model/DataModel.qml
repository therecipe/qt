import QtQuick 2.0
import Felgo 3.0

Item {
  id: dataModel

  // property to configure target dispatcher / logic
  property alias dispatcher: logicConnection.target

  readonly property alias currentProfile: _.currentProfile
  readonly property alias timeline: _.timeline
  readonly property alias messages: _.messages

  readonly property alias firstTweetData: _.firstTweetData

  Connections {
    id: logicConnection

    onFetchTwitterData:  {
      console.debug("Loading datamodel...")

      // Profile
      HttpRequest.get(Qt.resolvedUrl("../data/user.json"))
      .then(function(res) {
        _.currentProfile = JSON.parse(res.body)
      })

      // Feed
      HttpRequest.get(Qt.resolvedUrl("../data/feed.json"))
      .then(function(res) {
        var json = JSON.parse(res.body)
        var model = []

        _.firstTweetData = json[0]
        for (var i = 0; i < json.length; i++) {
          model.push(_.tweetModel(json[i]))
        }
        _.timeline = model
      })

      // Messages
      HttpRequest.get(Qt.resolvedUrl("../data/messages.json"))
      .then(function(res) {
        var raw = JSON.parse(res.body)
        var model = []

        for (var i = 0; i < raw.length; i++) {
          var names = model.map(function(val) { return val.handle } )
          var message = JSON.parse(JSON.stringify(raw[i]))
          message.user = message.sender
          message = _.tweetModel(message)

          if (names.indexOf("@" + message.user.screen_name) === -1) {
            message.actionsHidden = true
            delete message.sender
            delete message.image
            delete message.retweeted
            delete message.favorited
            delete message.retweet_count
            delete message.favorite_count
            model.push(message)
          }
        }
        _.messages = model
      })
    }

    // action 2 - addTweet
    onAddTweet: {
      //create fake tweet as copy of first tweet with new text
      var newTweet = JSON.parse(JSON.stringify(_.firstTweetData))
      newTweet.user = _.currentProfile
      newTweet.text = text
      _.timeline.splice(0, 0, _.tweetModel(newTweet)) //insert at position 0
      _.timelineChanged()
    }
  }

  // private
  Item {
    id: _

    property var currentProfile
    property var timeline
    property var messages

    property var firstTweetData


    function tweetModel(data) {
      console.log("TWEET MODEL: "+JSON.stringify(data))
      // Twitter uses custom format not recognized by new Date(string)
      var date = new Date(data.created_at.replace("+0000 ", "") + " UTC")

      // Check for URLs
      var text = data.text
      if (!!data.entities && !!data.entities.urls) {
        for (var j = 0; j < data.entities.urls.length; j++) {
          var url = data.entities.urls[j]
          text = text.replace(url.url, "<a href=\"" + url.url + "\">" + url.display_url + "</a>")
        }
      }

      // Check for image media
      var image = null
      if (!!data.entities && !!data.entities.media) {
        for (var j = 0; j < data.entities.media.length; j++) {
          var med = data.entities.media[j]
          if (med.type === "photo") {
            // Save image reference
            image = med.media_url

            // Remove possible url from text
            text = text.replace(med.url, "")

            // We just use the first found photo for now
            break
          }
        }
      }

      return {
        "id": data.id,
        "name": data.user.name,
        "handle": "@" + data.user.screen_name,
        "icon": data.user.profile_image_url.replace("_normal", "_bigger"),
        "text": text,
        "image": image,
        "time": DateFormatter.prettyDateTime(date),
        "retweet_count": data.retweet_count,
        "favorite_count": data.favorite_count,
        "retweeted": data.retweeted,
        "favorited": data.favorited,
        "user": data.user,
        "actionsHidden": undefined // workaround because dynamic add/remove of properties has troubles on iOS with Qt. 5.6.0
      }
    }
  }
}
