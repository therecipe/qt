import QtQuick 2.0

Item {

  //let phone search for location, and use it instead of search
  signal useLocation()

  signal searchListings(string searchText, bool addToRecents)

  signal showRecentSearches()

  signal loadNextPage()

  signal toggleFavorite(var listingData)


}

