import Felgo 3.0
import QtQuick 2.0
import "../common"
import QtQuick.Controls 2.0 as QtQuick2

Page {
  id: page
  title: "Timetable"
  rightBarItem: ActivityIndicatorBarItem { opacity: dataModel.loading || currentContentLoading ? 1 : 0 }

  readonly property var daysModel: dataModel.schedule ? prepareDaysModel(dataModel.schedule) : []
  readonly property bool dataAvailable: daysModel.length > 0

  readonly property var currentContentItem: swipeView.itemAt(swipeView.currentIndex)
  readonly property bool currentContentLoading: !currentContentItem ? false : currentContentItem.loading

  signal floatingButtonClicked
  onFloatingButtonClicked: navigation.currentIndex = 2


  AppText {
    text: "No data available."
    visible: !dataAvailable
    anchors.centerIn: parent
  }

  // tab bar
  AppTabBar {
    id: appTabBar
    showIcon: false
    visible: dataAvailable
    contentContainer: swipeView
    onCurrentItemChanged: {
      if(dataAvailable) amplitude.logEvent("View Day",{"day" : currentItem.text})
    }

    // auto select tab of current day after tab-bar was built
    property int autoSelectTab: -1
    onCountChanged: {
      if(autoSelectTab > 0 && dataAvailable && appTabBar.count === daysModel.length)
        appTabBar.currentIndex = autoSelectTab
    }

    Repeater {
      // dummyData to avoid tabcontrol issue when no children
      property var dummyData: [{ "day" : 0, "weekday" : "", "schedule": undefined }]
      model: dataAvailable ? daysModel : dummyData
      delegate: AppTabButton {
          text: Theme.isAndroid ? modelData.weekday.substring(0, 3) : modelData.weekday
      } // AppTabButton

    } // Repeater
  } // AppTabBar

  // tab contents
  QtQuick2.SwipeView {
    id: swipeView
    y: appTabBar.height
    height: parent.height - y
    width: parent.width
    clip: true

    visible: dataAvailable
    currentIndex: appTabBar.currentIndex

    Repeater {
      property var dummyData: [{ "day" : 0, "weekday" : "", "schedule": undefined }]
      model: dataAvailable ? daysModel : dummyData
      delegate: TimetableDaySchedule {
        // swipe view has troubles readjusting its width, so we change it explicitly here
        width: page.navigationStack.splitViewActive ? page.navigationStack.leftColumnWidth : page.width
        height: parent.height
        scheduleData: modelData.schedule
        onSearchAccepted: {
          if(text !== "") {
            amplitude.logEvent("Search Talk",{"term" : text})
            var result = logic.search(text, dataModel.talks)
            page.navigationStack.leftColumnIndex = 1
            page.navigationStack.popAllExceptFirstAndPush(Qt.resolvedUrl("SearchPage.qml"), { searchModel: result })
          }
          else
            page.navigationStack.popAllExceptFirst()
        }
        onItemClicked: {
          page.navigationStack.leftColumnIndex = 0
          page.navigationStack.popAllExceptFirstAndPush(Qt.resolvedUrl("DetailPage.qml"), { item: item })
        }
      }
    }
  }

  FloatingActionButton {
    icon: IconType.star
    onClicked: floatingButtonClicked()

    property int currentVisibleItemCount: !!daysModel && !!daysModel[swipeView.currentIndex] ? daysModel[swipeView.currentIndex].schedule.length : 0

    opacity: (swipeView.currentItem.listView.lastItemIndex === currentVisibleItemCount - 1) ? 0 : 1
    enabled: opacity > 0
    Behavior on opacity {
      NumberAnimation { duration: 150 }
    }
  }

  // prepareDaysModel - package schedule data in array with conference days (for tabs)
  function prepareDaysModel(data) {
    if(!(data.conference && data.conference.days))
      return []


    var currentDate = new Date()
    var days = ['Sunday','Monday','Tuesday','Wednesday','Thursday','Friday','Saturday'];

    var model = []
    for(var day in data.conference.days) {
      var date = new Date(day+"T00:00.000Z")
      var weekday = isNaN(date.getUTCDay()) ? "Unknown" : days[ date.getUTCDay() ]
      var schedule = data.conference.days[day]
      var dayModel = {
        day: day,
        weekday: weekday,
        schedule: prepareScheduleModel(schedule, day, weekday)
      }

      model.push(dayModel)

      // activate tab if its current day
      if(currentDate.getDate() === date.getDate() &&
          currentDate.getMonth() === date.getMonth() &&
          currentDate.getFullYear() === date.getFullYear()) {
        appTabBar.autoSelectTab = model.length - 1
      }
    }

    return model
  }

  // prepareScheduleModel - creates the events list from schedule data (per day, tab content)
  function prepareScheduleModel(data, day, weekday) {
    if(!(data && data.rooms && dataModel.talks))
      return []

    // get events for all rooms and prepare sections
    var events = []
    for(var room in data.rooms) {
      for(var eventIdx in data.rooms[room]) {
        var talk = dataModel.talks[data.rooms[room][eventIdx]]
        if(talk !== undefined) {
          // prepare section and add talk to events
          talk.section = weekday+", "+talk.start.substring(0, 2) + ":00"
          events.push(talk)
        }
      }
    }

    // sort events
    events = events.sort(function(a, b) {
      return (parseInt(a.start) > parseInt(b.start)) - (a.start < b.start)
    })

    return events
  }
}
