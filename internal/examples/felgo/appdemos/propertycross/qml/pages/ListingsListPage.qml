import QtQuick 2.0
import Felgo 3.0

ListPage {
  id: listPageWrapper

  property var scrollPos: null
  property bool favorites

  Component {
    id: detailPageComponent
    ListingDetailPage { }
  }

  rightBarItem: ActivityIndicatorBarItem {
    visible: dataModel.loading
  }

  model: JsonListModel {
    id: listModel
    source: favorites ? dataModel.favoriteListings : dataModel.listings
    fields: [ "text", "detailText", "image", "model" ]
  }

  title: favorites
         ? qsTr("Favorites")
         : qsTr("%1 of %2 matches").arg(
             dataModel.numListings).arg(
             dataModel.numTotalListings)

  emptyText.text: favorites
                  ? qsTr("You have not added any properties to your favourites.")
                  : qsTr("No listings available.")

  listView.footer: VisibilityRefreshHandler {
    // visible if NOT favorites shown and more listings are available
    visible: !favorites && dataModel.numListings < dataModel.numTotalListings

    onRefresh: {
      scrollPos = listView.getScrollPosition()
      logic.loadNextPage()
    }
  }

  delegate: SimpleRow {
    item: listModel.get(index)
    autoSizeImage: true
    imageMaxSize: dp(40)
    image.fillMode: Image.PreserveAspectCrop

    onSelected: navigationStack.popAllExceptFirstAndPush(detailPageComponent, {model: item.model})
  }

  listView.onModelChanged: if(scrollPos) listView.restoreScrollPosition(scrollPos)
}


