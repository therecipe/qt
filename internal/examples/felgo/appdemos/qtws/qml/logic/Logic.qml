import QtQuick 2.0

Item {

  signal initializeDataModel(var webStorageItem)

  signal resetOrInitializeDataModel(bool userInitiallyInSync)

  signal increaseLocalAppStarts()

  signal loadData()

  signal setFeedBackSent(bool value)

  signal storeRating(int id, var rating)

  signal toggleFavorite(var item)

  signal removeContact(string id)

  signal loadContact(string id, var successHandler, var errorHandler)

  signal addContact(string id, var contact)

  signal clearContacts()

  signal clearCache()

  signal setNotificationsEnabled(bool value)

  // search - get talks with certain keyword in title or description
  function search(query, talks) {
    if(!talks)
      return []

    query = query.toLowerCase().split(" ")
    var result = []

    // check talks
    for(var id in talks) {
      var talk = talks[id]
      var contains = 0

      // check query
      for (var key in query) {
        var term = query[key].trim()
        if(talk.title.toLowerCase().indexOf(term) >= 0 ||
            talk.description.toLowerCase().indexOf(term) >= 0) {
          contains++
        }
        for(var key2 in talk.persons) {
          var speaker = talk.persons[key2]
          if(speaker.full_public_name.toLowerCase().indexOf(term) >= 0) {
            contains++
          }
        }
      }

      if(contains == query.length)
        result.push(talk)
    } // check talks

    return result
  }
}
