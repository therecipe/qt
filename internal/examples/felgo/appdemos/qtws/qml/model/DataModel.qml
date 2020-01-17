import Felgo 3.0
import QtQuick 2.0

Item {
  id: dataModel

  // property to configure target dispatcher / logic
  property alias dispatcher: logicConnection.target

  // stored data
  property var version: undefined
  property var schedule: undefined
  property var speakers: undefined
  property var tracks: undefined
  property var favorites: undefined
  property var talks: undefined
  property var contacts: undefined
  property var ratings: undefined

  property int localAppStarts: 0
  property bool feedBackSent: false

  property string timeZone: "+0200"

  property var webStorage: undefined // reference to WebStorage for favorites
  readonly property bool loading: api.loadingCount > 0
  readonly property bool loaded: !!schedule && !!speakers

  property bool initialized: false

  property bool notificationsEnabled: true
  onNotificationsEnabledChanged: storage.setValue("notificationsEnabled", notificationsEnabled)

  signal loadingFailed()
  signal favoriteAdded(var talk)
  signal favoriteRemoved(var talk)

  Connections {
    id: logicConnection

    // action 1 - initializeDataModel from storage - get data from local storage
    onInitializeDataModel: {
      dataModel.version = storage.getValue("version")
      dataModel.schedule = storage.getValue("schedule")
      dataModel.speakers = storage.getValue("speakers")
      dataModel.tracks = storage.getValue("tracks")
      dataModel.talks = storage.getValue("talks")
      dataModel.ratings = storage.getValue("ratings")
      dataModel.notificationsEnabled = storage.getValue("notificationsEnabled") !== undefined ? storage.getValue("notificationsEnabled") : true

      dataModel.localAppStarts = storage.getValue("localAppStarts") || 0
      dataModel.feedBackSent = storage.getValue("feedBackSent") || false

      // get favorites and contacts from web storage
      dataModel.webStorage = webStorageItem
      dataModel.favorites = webStorage.getValue("favorites")
      dataModel.contacts = webStorage.getValue("contacts")

      dataModel.initialized = true
    }

    onResetOrInitializeDataModel: {
      if(!userInitiallyInSync) {
        // initially in sync changed to false -> user switched, clear favorites
        dataModel.favorites = undefined
        dataModel.initialized = false
      }
      else if(!dataModel.initialized) {
        // initially in sync changed to true again -> initialize data
        logic.initializeDataModel(webStorage)
      }
    }

    // action 3 - increaseLocalAppStarts - increase local app start counter
    onIncreaseLocalAppStarts: {
      if(!initialized)
        return

      localAppStarts++
      storage.setValue("localAppStarts",localAppStarts)
    }

    // action 4 - loadData - fetches all data from Qt WS api
    onLoadData: {
      if(initialized && !loading) {
        api.checkAPIVersion() // checks version and loads data if necessary
      }
    }

    // action 5 - setFeedBackSent - store whether feedback was sent
    onSetFeedBackSent: {
      if(!initialized)
        return

      feedBackSent = !!value  // ensures boolean type
      storage.setValue("feedBackSent", feedBackSent)
    }

    // action 6 - storeRating
    onStoreRating: {
      if(!ratings) ratings = {}
      ratings[id] = rating
      storage.setValue("ratings", ratings)
    }

    // action 7 - toggleFavorite - add or remove item from favorites
    onToggleFavorite: {
      if(dataModel.favorites === undefined)
        dataModel.favorites = { }

      if(dataModel.favorites[item.id]) {
        delete dataModel.favorites[item.id]
        dataModel.favoriteRemoved(item)
      }
      else {
        dataModel.favorites[item.id] = item.id
        dataModel.favoriteAdded(item)
      }

      // store favorites
      webStorage.setValue("favorites", dataModel.favorites, function(data) {
        // in case setValue merges favorites with data from server, we update the local value
        // the merged server data comes as stringified string, thus call json.parse here
        //dataModel.favorites = webStorage.getValue("favorites") // no need to call getValue() again, we get the value from the callback
        console.debug("server favorites data:", JSON.stringify(data))
        // only update the local favorites, if the server has new merged data, in this case conflict is true
        if(data["conflict"]) {
          dataModel.favorites = data.mergedData
        }

      })
      // call the changed here, to schedule the notifications with the new local data
      favoritesChanged()
    }

    // action 8 - removeContact - remove contact from list
    onRemoveContact: {
      if(dataModel.contacts !== undefined && dataModel.contacts[id]) {
        delete dataModel.contacts[id]

        // store contacts
        webStorage.setValue("contacts", dataModel.contacts, function(data) {
          // in case setValue merges contacts with data from server, we update the local value
          // only update the local favorites, if the server has new merged data, in this case conflict is true
          if(data["conflict"]) {
            dataModel.contacts = data.mergedData
          }
        })

        // signal that contacts changed
        dataModel.contactsChanged()
      }
    }

    // action 9 - loadContact - loads contact from Eventbrite
    onLoadContact: {
      api.sendGetRequest(Qt.resolvedUrl("https://www.qtworldsummit.com/api/attendee/show/?id="+id), successHandler, errorHandler)
    }

    // action 10 - addContact - store contact from barcode scanner
    onAddContact: {
      if(dataModel.contacts === undefined)
        dataModel.contacts = { }

      // store contact
      dataModel.contacts[id] = contact
      webStorage.setValue("contacts", dataModel.contacts, function(data) {
        // in case setValue merges contacts with data from server, we update the local value
        // only update the local favorites, if the server has new merged data, in this case conflict is true
        if(data["conflict"]) {
          dataModel.contacts = data.mergedData
        }
      })

      // signal that contacts changed
      dataModel.contactsChanged()
    }

    // action 10 - clearContacts
    onClearContacts: {
      webStorage.setValue("contacts",{})
    }

    // action 11 - clears locally stored data
    onClearCache: {
      // reset dataModel, but keep favorites and contacts
      var favorites = dataModel.favorites
      var contacts = dataModel.contacts
      _.reset()
      dataModel.favorites = favorites
      dataModel.contacts = contacts

      // clear local storage, favorites and contacts are still in webStorage
      storage.clearAll()
      initialized = true // also reloads api data after reset
    }

    // action 12 - setNotificationsEnabled
    onSetNotificationsEnabled: {
      notificationsEnabled = value
    }
  }

  // api for data access
  QtWSApi {
    id: api
  }

  // storage for caching data
  Storage {
    id: storage
    databaseName: "localStorage"
  }

  // private
  Item {
    id: _

    // reset data model
    function reset() {
      dataModel.initialized = false
      dataModel.version = undefined
      dataModel.schedule = undefined
      dataModel.speakers = undefined
      dataModel.tracks = undefined
      dataModel.talks = undefined
      dataModel.ratings = undefined
      dataModel.favorites = undefined
      dataModel.contacts = undefined
      dataModel.notificationsEnabled = true

      localAppStarts = 0
      feedBackSent = false
    }
  }

  // isFavorite - check if item is favorited
  function isFavorite(id) {
    return dataModel.favorites !== undefined && dataModel.favorites[id] !== undefined
  }

  function getRating(id) {
    if(!ratings || !(id in ratings)) return -1
    else return ratings[id]
  }
}
