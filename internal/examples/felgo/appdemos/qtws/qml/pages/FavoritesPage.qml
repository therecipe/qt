import Felgo 3.0
import QtQuick 2.0
import "../common"
import QtQuick.Controls 2.0 as QtQuick2

Page {
  id: page
  title: "Favorites"
  rightBarItem: ActivityIndicatorBarItem { opacity: dataModel.loading || scheduleItem.loading ? 1 : 0 }

  property var favoritesModel: dataModel.favorites ? prepareFavoritesModel(dataModel.favorites) : []
  readonly property bool dataAvailable: favoritesModel.length > 0

  signal floatingButtonClicked
  onFloatingButtonClicked: navigation.currentIndex = 1

  // update UI when favourites change
  Connections {
    target: dataModel
    onFavoritesChanged: favoritesModel = dataModel.favorites ? prepareFavoritesModel(dataModel.favorites) : []
  }

  AppText {
    text: "No talks added to favorites yet."
    visible: !dataAvailable
    anchors.centerIn: parent
  }

  TimetableDaySchedule {
    id: scheduleItem
    anchors.fill: parent
    scheduleData: favoritesModel
    searchAllowed: false
    onItemClicked: {
      page.navigationStack.popAllExceptFirstAndPush(Qt.resolvedUrl("DetailPage.qml"), { item: item })
    }
    visible: dataAvailable
  }

  FloatingActionButton {
    icon: IconType.calendaro
    onClicked: floatingButtonClicked()

    property int currentVisibleItemCount: !!favoritesModel ? favoritesModel.length : 0

    opacity: (scheduleItem.listView.lastItemIndex === currentVisibleItemCount - 1) ? 0 : 1
    enabled: opacity > 0
    Behavior on opacity {
      NumberAnimation { duration: 150 }
    }
  }

  // prepareFavoritesModel - package favorites data in array ready to be displayed by TimeTableDaySchedule item
  function prepareFavoritesModel(favorites) {
    if(!(favorites && dataModel.talks))
      return []

    var days = ['Sunday','Monday','Tuesday','Wednesday','Thursday','Friday','Saturday'];

    // get events and prepare data for sorting and sections
    var events = []
    for(var id in favorites) {
      var data = dataModel.talks[favorites[id]]
      if(data !== undefined) {
        // prepare event date for sorting
        var date = new Date(data.day+"T00:00.000Z")
        data.dayTime = date.getTime()

        // prepare event section
        var weekday = isNaN(date.getUTCDay()) ? "Unknown" : days[ date.getUTCDay() ]
        data.section = weekday + ", " + (data.start.substring(0, 2) + ":00")

        events.push(data)
      }
    }

    // sort events
    events = events.sort(function(a, b) {
      if(a.dayTime == b.dayTime)
        return (a.start > b.start) - (a.start < b.start)
      else
        return (a.dayTime > b.dayTime) - (a.dayTime < b.dayTime)
    })

    return events
  }
}
