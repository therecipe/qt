import QtQuick 2.5
import Felgo 3.0

Page {
  id: searchPage
  title: qsTr("Property Cross")

  rightBarItem: NavigationBarRow {
    ActivityIndicatorBarItem {
      visible: dataModel.loading
      showItem: showItemAlways
    }

    IconButtonBarItem {
      icon: IconType.heart
      onClicked: showListings(true)
      title: qsTr("Favorites")
    }
  }

  Column {
    id: contentCol
    anchors.left: parent.left
    anchors.top: parent.top
    anchors.right: parent.right
    anchors.margins: contentPadding
    spacing: contentPadding

    AppText {
      width: parent.width
      wrapMode: Text.WrapAtWordBoundaryOrAnywhere
      font.pixelSize: sp(12)
      text: qsTr("Use the form below to search for houses to buy. You can search by place-name, postcode, or click 'My location', to search in your current location.")
    }

    AppText {
      width: parent.width
      font.pixelSize: sp(12)
      color: Theme.secondaryTextColor
      font.italic: true
      text: "Hint: You can quickly find and view results by searching 'London'."
      wrapMode: Text.WrapAtWordBoundaryOrAnywhere
    }

    AppTextField {
      id: searchInput
      width: parent.width

      showClearButton: true
      placeholderText: qsTr("Search")
      inputMethodHints: Qt.ImhNoPredictiveText

      onTextChanged: showRecentSearches()
      onEditingFinished: if(navigationStack.currentPage === searchPage) search()
    }

    Row {
      spacing: contentPadding

      AppButton {
        text: qsTr("Go")
        onClicked: search()
      }

      AppButton {
        text: qsTr("My location")
        enabled: dataModel.positioningSupported

        onClicked: {
          searchInput.text = ""
          searchInput.placeholderText = qsTr("Looking for location...")
          logic.useLocation()
        }
      }
    }

    AppText {
      visible: dataModel.isError
      text: qsTr("There was a problem with your search")
    }
  }

  AppListView {
    id: listView

    width: parent.width
    anchors.top: contentCol.bottom
    anchors.bottom: parent.bottom

    visible: !dataModel.isError

    // Show either the recents searches or the currently found locations depending on search mode
    model: JsonListModel {
      source: dataModel.isRecent ? dataModel.recentSearches : dataModel.locations
      keyField: "searchText"
      fields: [ "searchText", "heading", "text", "model", "detailText" ]
    }

    // Show a section header for listings
    section.property: "heading"
    section.delegate: SimpleSection { }

    delegate: SimpleRow {
      item: listView.model.get(index)
      // do not add suggestions to recent searches
      onSelected: logic.searchListings(item.searchText, false)
    }

    emptyText.text: dataModel.isRecent
                    ? qsTr("No recent searches.")
                    : qsTr("No suggested locations.")
  }

  Connections {
    target: dataModel
    onListingsReceived: showListings(false)
    onLocationReceived: if(searchInput.placeholderText === "Looking for location...") searchInput.placeholderText = "Search"
  }

  Component {
    id: listPageComponent
    ListingsListPage {}
  }

  function showListings(favorites) {
    if(navigationStack.depth === 1) {
      navigationStack.popAllExceptFirstAndPush(listPageComponent, { favorites: favorites })
    }
  }

  function search() {
    logic.searchListings(searchInput.text, true)
  }

  function showRecentSearches() {
    logic.showRecentSearches()
  }
}
