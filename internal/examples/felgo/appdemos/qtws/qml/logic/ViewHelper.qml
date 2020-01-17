import QtQuick 2.0

Item {

  // prepareTracks - prepare track data for display in TracksPage
  function prepareTracks(tracks) {
    if(!dataModel.talks)
      return []

    var model = []
    for(var i in Object.keys(tracks)){
      var track = Object.keys(tracks)[i];
      var talks = []

      for(var j in Object.keys(dataModel.talks)) {
        var talkID = Object.keys(dataModel.talks)[j]
        var talk = dataModel.talks[parseInt(talkID)]

        if(talk !== undefined && talk.tracks.indexOf(track) > -1) {
          talks.push(talk)
        }
      }
      talks = prepareTrackTalks(talks)
      model.push({"title" : track, "talks" : talks})
    }
    model.sort(compareTitle)

    return model
  }

  // prepareTrackTalks - package talk data in array ready to be displayed by TimeTableDaySchedule item
  function prepareTrackTalks(trackTalks) {
    if(!trackTalks)
      return []

    var days = ['Sunday','Monday','Tuesday','Wednesday','Thursday','Friday','Saturday'];

    // get events and prepare data for sorting and sections
    for(var idx in trackTalks) {
      var data = trackTalks[idx]

      // prepare event date for sorting
      var date = new Date(data.day)
      data.dayTime = date.getTime()

      // prepare event section
      var weekday = isNaN(date.getTime()) ? "Unknown" : days[ date.getDay() ]
      data.section = weekday + ", " + (data.start.substring(0, 2) + ":00")

      trackTalks[idx] = data
    }

    // sort events
    trackTalks = trackTalks.sort(function(a, b) {
      if(a.dayTime == b.dayTime)
        return (a.start > b.start) - (a.start < b.start)
      else
        return (a.dayTime > b.dayTime) - (a.dayTime < b.dayTime)
    })

    return trackTalks
  }

  // sort tracks by title
  function compareTitle(a,b) {
    if (a.title < b.title)
      return -1;
    if (a.title > b.title)
      return 1;
    return 0;
  }

}
