import QtQuick 2.0
import QtPositioning 5.3
import Qt.labs.settings 1.0

Item {

  // property to configure target dispatcher / logic
  property alias dispatcher: logicConnection.target

  //called when a page of listings has been received from the server
  signal listingsReceived

  //called when the location of the device has been received
  signal locationReceived

  //location type, a coordinate type with latitude, longitude and isValid property
  readonly property var location: positionSource.position.coordinate

  //true when there is data being loaded in the background
  //or when the location is being retrieved
  readonly property bool loading: client.loading || positionSource.active

  //total number of found listings, listings property only contains the first page
  readonly property alias numTotalListings: _.numTotalListings

  //number of listings currently loaded
  readonly property int numListings: _.listings.length

  //list model for listings list page
  readonly property var listings: _.createListingsModel(_.listings)

  //list model for listings list page
  readonly property var favoriteListings: _.createListingsModel(_.favoriteListingsValues, true)

  //list model for search page location list
  readonly property var locations: _.createLocationsModel()

  //list model for search page recent searches list
  readonly property var recentSearches: _.createRecentSearchesModel()

  readonly property bool isSuggested: _.locationSource === _.locationSourceSuggested
  readonly property bool isRecent: _.locationSource === _.locationSourceRecent
  readonly property bool isError: _.locationSource === _.locationSourceError

  readonly property bool positioningSupported: positionSource.supportedPositioningMethods !==
                                               PositionSource.NoPositioningMethods &&
                                               positionSource.valid

  // handle logic actions
  Connections {
    id: logicConnection

    // action 1 - useLocation
    onUseLocation: {
      if(positionSource.position.coordinate.isValid) {
        //already have a location
        _.searchByLocation()
      } else {
        _.locationSearchPending = true
        positionSource.update()
      }
    }

    // action 2 - searchListings
    onSearchListings: {
      _.lastSearchText = addToRecents ? searchText : ""
      _.listings = []
      client.search(searchText, _.responseCallback)
    }

    // action 3 - showRecentSearches
    onShowRecentSearches: {
      _.locationSource = _.locationSourceRecent
    }

    // action 4 - loadNextPage
    onLoadNextPage: {
      client.repeatForPage(_.currentPage + 1, _.responseCallback)
    }

    // action 5 - toggleFavorite
    onToggleFavorite: {
      var listingDataStr = JSON.stringify(listingData)
      var index =  _.favoriteListingsValues.indexOf(listingDataStr)

      if(index < 0) {
        _.favoriteListingsValues.push(listingDataStr) // add entry to favorites
      } else {
        _.favoriteListingsValues.splice(index, 1)  // remove entry in favorites
      }

      _.favoriteListingsValuesChanged()
    }
  }

  function isFavorite(listingData) {
    return _.favoriteListingsValues.indexOf(JSON.stringify(listingData)) >= 0
  }

  Settings {
    property string recentSearches: JSON.stringify(_.recentSearches)
    property string favoriteListingsValues: JSON.stringify(_.favoriteListingsValues)

    Component.onCompleted: {
      _.recentSearches = recentSearches && JSON.parse(recentSearches) || {}
      _.favoriteListingsValues = favoriteListingsValues && JSON.parse(favoriteListingsValues) || []
    }
  }

  Client {
    id: client
  }

  PositionSource {
    id: positionSource
    active: false

    onActiveChanged: {
      var coord = position.coordinate
      console.log("Coordinates:", coord.latitude, coord.longitude,
                  "valid:", coord.isValid,
                  "source active:", active)

      if(!active) {
        if(coord.isValid && _.locationSearchPending) {
          _.searchByLocation()
          _.locationSearchPending = false
        }
      }
    }

    onUpdateTimeout: console.log("location timed out")
  }

  Item {
    id: _ //private members

    property int locationSource: locationSourceRecent

    property var favoriteListingsValues: []

    property var recentSearches: ({})

    property var locations: []
    property var listings: []
    property int numTotalListings

    property int currentPage: 1

    property string lastSearchText: ""

    readonly property int locationSourceSuggested: 1
    readonly property int locationSourceRecent: 2
    readonly property int locationSourceError: 3

    readonly property var successCodes: ["100", "101", "110"]
    readonly property var ambiguousCodes: ["200", "202"]

    property bool locationSearchPending: false

    function searchByLocation() {
      locationReceived()
      var coord = positionSource.position.coordinate
      client.searchByLocation(coord.latitude, coord.longitude, _.responseCallback)
    }

    function responseCallback(obj) {
      var response = obj.response
      var code = response.application_response_code
      console.debug("Server returned application code:", code)
      if(successCodes.indexOf(code) >= 0) {
        //found a location -> display listings
        currentPage = parseInt(response.page)
        listings = listings.concat(response.listings)
        numTotalListings = response.total_results || 0
        console.debug("Server returned", response.listings.length, "listings")
        addRecentSearch(qsTr("%1 (%2 listings)").arg(lastSearchText).arg(numTotalListings))
        listingsReceived()
      } else if(ambiguousCodes.indexOf(code) >= 0) {
        //found ambiguous locations -> display locations
        locations = response.locations
        console.debug("Server returned", response.locations.length, "locations")
        addRecentSearch(qsTr("%1 (%2 locations)").arg(lastSearchText).arg(response.locations.length))
        locationSource = locationSourceSuggested
      }
      else if(code === "210") {
        // no result found
        locations = []
        locationSource = locationSourceSuggested
      }
      else {
        //found nothing -> display error
        locations = []
        locationSource = locationSourceError
      }
    }

    function createLocationsModel() {
      return locations.map(function(data) {
        return {
          heading: "Please select a location below",
          text: data.title,
          detailText: data.long_title,
          model: data,
          searchText: data.place_name
        }
      })
    }

    function createRecentSearchesModel() {
      return Object.keys(recentSearches).map(function(text) {
        return {
          heading: "Recent Searches",
          text: recentSearches[text].displayText,
          searchText: text
        }
      })
    }

    function createListingsModel(source, parseValues) {
      return source.map(function(data) {
        if(parseValues)
          data = JSON.parse(data)

        return {
          text: data.price_formatted,
          detailText: data.title,
          image: data.thumb_url,
          model: data
        }
      })
    }

    function addRecentSearch(displayText) {
      if(lastSearchText) {
        recentSearches[lastSearchText] = {
          displayText: displayText
        }
        _.recentSearchesChanged()
      }
    }
  }
}

