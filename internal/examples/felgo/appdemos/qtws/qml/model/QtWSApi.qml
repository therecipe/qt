import Felgo 3.0
import QtQuick 2.0

QtObject {
  id: _

  // qtws 2017 api urls
  property string qtwsApiScheduleUrl: Qt.resolvedUrl("https://www.qtworldsummit.com/api/schedule/all/")
  property string qtwsApiSpeakersUrl: Qt.resolvedUrl("https://www.qtworldsummit.com/api/speakers/all/")
  property string qtwsApiVersionUrl: Qt.resolvedUrl("https://www.qtworldsummit.com/api/version/show/")

  // fallback urls of locally stored version in assets
  property string fallbackScheduleUrl: Qt.resolvedUrl("../../assets/data/schedule.json")
  property string fallbackSpeakersUrl: Qt.resolvedUrl("../../assets/data/speakers.json")
  property string fallbackVersionUrl: Qt.resolvedUrl("../../assets/data/version.json")

  Component.onCompleted: {
    HttpNetworkActivityIndicator.activationDelay = 0
  }


  // sendGetRequest - load data from url with success handler
  function sendGetRequest(url, successHandler, errorHandler) {
    HttpRequest.get(url)
    .then(function(res) {
      var fixedResponse = res.text.replace(new RegExp("&amp;",'g'),"&")
      successHandler(JSON.parse(fixedResponse))
    })
    .catch(function(err) {
      console.error("Error: Failed to load data from "+url+", error = "+err.message)
      if(errorHandler !== undefined)
        errorHandler()
      else if(!loading)
        dataModel.loadingFailed()
    })
  }

  // checkAPIVersion - checks Qt WS API version and updates data if necessary
  function checkAPIVersion(useLocalData) {
    var versionUrl = useLocalData ? _.fallbackVersionUrl : _.qtwsApiVersionUrl

    _.sendGetRequest(versionUrl, function(data) {
      var currVersion = data.version

      // load new data when debug build, first call, or newer version available
      if(!system.publishBuild || dataModel.version === undefined || dataModel.version !== currVersion) {
        dataModel.version = currVersion
        _.loadSchedule(useLocalData) // also loads speakers
      }
    }, function() {
      // custom error handler
      if(dataModel.version === undefined && useLocalData === undefined)
        checkAPIVersion(true)
      else if(!loading)
        dataModel.loadingFailed()
    })
  }

  // loadSchedule - load Qt WS schedule from api
  function loadSchedule(useLocalData) {
    var scheduleUrl = useLocalData ? _.fallbackScheduleUrl : _.qtwsApiScheduleUrl

    _.sendGetRequest(scheduleUrl, function(data) {
      _.processScheduleData(data)
      // load speakers after schedule is processed
      _.loadSpeakers(useLocalData)
    })
  }

  // loadSpeakers - load Qt WS speakers from api
  function loadSpeakers(useLocalData) {
    var speakersUrl = useLocalData ? _.fallbackSpeakersUrl : _.qtwsApiSpeakersUrl

    _.sendGetRequest(speakersUrl, function(data) {
      _.processSpeakersData(data)

      // when schedule and speakers are loaded, all loading is done -> cache current API version
      storage.setValue("version", dataModel.version)
    })
  }

  // processScheduleData - process schedule data for usage in UI
  function processScheduleData(data) {
    // retrieve tracks and talks and build model for tracks, talks and schedule
    var tracks = {}
    var talks = {}
    for(var day in data.conference.days) {
      for(var room in data.conference.days[day]["rooms"])
        for (var eventIdx in data.conference.days[day]["rooms"][room]) {
          var event = data.conference.days[day]["rooms"][room][eventIdx]

          // calculate event end time
          var start = event.start.split(":")
          var duration = event.duration.split(":")
          var end = [parseInt(start[0])+parseInt(duration[0]),
                     parseInt(start[1])+parseInt(duration[1])]
          if(end[1] >= 60) {
            end[1] -= 60
            end[0] += 1
          }

          // format start and end time
          event.start = _.format2DigitTime(start[0]) + ":" + _.format2DigitTime(start[1])
          event.end = _.format2DigitTime(end[0]) + ":" + _.format2DigitTime(end[1])

          // clean-up false start time formatting (always 2 digits required)
          if(event.start.substring(1,2) == ':') {
            event.start = "0"+event.start
          }

          // add day of event (for favorites)
          event.day = day

          // build tracks model
          if(event["tracks"] !== undefined && Array.isArray(event["tracks"])) {
            for(var idx in event["tracks"])
              tracks[event["tracks"][idx]] = 0
          }

          // clean up incorrect room entries of API version 1.953
          if(event["room"] == "Berlin" || event["room"] == ":")
            event["room"] = ""

          // build talks model
          talks[event["id"]] = event

          // replace talks in schedule with talk-id
          data.conference.days[day]["rooms"][room][eventIdx] = event["id"]
        }
    }

    //  define track colors
    var hueDiff = 1 / Object.keys(tracks).length
    var i = 0
    for(var track in tracks) {
      tracks[track] = i * hueDiff
      i++
    }

    // store data
    dataModel.talks = talks
    dataModel.tracks = tracks
    dataModel.schedule = data
    storage.setValue("talks", talks)
    storage.setValue("tracks", tracks)
    storage.setValue("schedule", data)

    // force update of favorites as new data arrived
    var favorites = dataModel.favorites
    dataModel.favorites = undefined
    dataModel.favorites = favorites
  }

  // processSpeakersData - process schedule data for usage in UI
  function processSpeakersData(data) {
    // convert speaker data into model map with id as key
    var speakers = {}
    for(var i = 0; i < data.length; i++) {
      var speaker = data[i]
      speakers[speaker.id] = speaker

      var talks= []
      for (var j in Object.keys(dataModel.talks)) {
        var talkID = Object.keys(dataModel.talks)[j];
        var talk = dataModel.talks[parseInt(talkID)]
        var persons = talk.persons

        for(var k in persons) {
          if(persons[k].id === speaker.id) {
            talks.push(talkID.toString())
          }
        }
      }
      speakers[speaker.id]["talks"] = talks
    }
    // store data
    dataModel.speakers = speakers
    storage.setValue("speakers", speakers)
  }

  // format2DigitTime - adds leading zero to time (hour, minute) if required
  function format2DigitTime(time) {
    return (("" + time).length < 2) ? "0" + time : time
  }
}
