
import Felgo 3.0
import QtQuick 2.0
import "../common"

Item {
  id: dayScheduleItem

  property var scheduleData: []

  signal itemClicked(var item);

  readonly property bool loading: _.loadingCount > 0

  property alias emptyText: listView.emptyText

  property bool searchAllowed: true

  property alias listView: listView

  signal searchAccepted(string text)

  Component.onCompleted: {
    // helps to correctly update firstItemIndex and lastItemIndex of listView
    listView.positionViewAtEnd()
    listView.positionViewAtBeginning()
    updater.updateJumpToNow()
    jumpToNow() // scroll to talk of current time initially
  }

  // private members
  Item {
    id: _
    property int loadingCount: 0
  }

  // jump to now button (top)
  JumpToNowButton {
    id: jumpToNowTop
    z: 1
    y: show ? dp(5) : -height
    opacity: show ? 1 : 0
    text: "now"
    visible: !Theme.isIos

    // whether jump to now button should be visible, will be set while scrolling through view
    property bool show: false

    // onClicked - jump to list view position
    onClicked: {
      if(text === "now")
        jumpToNow()
      else
        listView.positionViewAtBeginning()
      show = false
    }
  }

  // jump to now button (bottom)
  JumpToNowButton {
    id: jumpToNowBottom
    z: 1
    y: show ? parent.height - height - dp(5) : parent.height
    opacity: show ? 1 : 0
    text: "now"

    // whether jump to now button should be visible, will be set while scrolling through view
    property bool show: false

    // onClicked - jump to list view position
    onClicked: { jumpToNow(); show = false; }
  }

  // show/hide jump to now buttons
  Connections {
    id: updater
    target: listView
    onFirstItemIndexChanged: updateJumpToNow()
    onLastItemIndexChanged: updateJumpToNow()

    // update visibility settings of jump to now button
    function updateJumpToNow() {
      if(!scheduleData || scheduleData.length == 0) {
        jumpToNowTop.show = false
        jumpToNowBottom.show = false
        return
      }

      var now = getCurrentDateTime()
      var nowItemIdx = itemIndexForTime(now.getTime())

      // check if nowItem (nearest) is not today -> always show scroll to top then
      var nowItemDate = getEventDateTime(scheduleData[nowItemIdx])
      if(nowItemDate.getFullYear() !== now.getFullYear() || nowItemDate.getMonth() !== now.getMonth() || nowItemDate.getDate() !== now.getDate())
        nowItemIdx = 0

      // use scroll to top if first
      if(nowItemIdx === 0)
        jumpToNowTop.text = "to top"
      else
        jumpToNowTop.text = "now"

      // use scroll to bottom if last item
      if(nowItemIdx === (scheduleData.length - 1))
        jumpToNowBottom.text = "to bottom"
      else
        jumpToNowBottom.text = "now"

      // show/hide correct buttons
      if(nowItemIdx >= listView.firstItemIndex && nowItemIdx <= listView.lastItemIndex) {
        // currently ongoing talk is visible
        jumpToNowTop.show = false
        jumpToNowBottom.show = false
      }
      else if(nowItemIdx > listView.firstItemIndex && listView.firstItemIndex >= 0) {
        jumpToNowTop.show = false
        jumpToNowBottom.show = true
      }
      else if(nowItemIdx < listView.lastItemIndex) {
        jumpToNowTop.show = true
        jumpToNowBottom.show = false
      }
      else {
        jumpToNowTop.show = false
        jumpToNowBottom.show = false
      }
    }
  }

  // app list view with daily schedule
  AppListView {
    id: listView
    width: parent.width
    height: parent.height - y
    model: JsonListModel {
      id: listModel
      source: scheduleData
      keyField: "id"
    }
    cacheBuffer: 10000 // allows to cache delegate items to fill up whole listview for invisible area at top and bottom

    // search header for listview
    header: SearchBar {
      visible: searchAllowed
      height: searchAllowed ? implicitHeight : 0
      onAccepted: searchAccepted(text)
      iosAlternateStyle: true
    }

    // provide index of first and last visible item
    property int firstItemIndex: nearestIndexAt(listView.width * 0.5, contentY, true)
    property int lastItemIndex: nearestIndexAt(listView.width * 0.5, contentY + listView.height, false)

    // nearestIndexAt - returns nearest item index for x,y position
    property real relaxationStep: dp(10)
    function nearestIndexAt(xPos, yPos, nextFirst) {
      var index = listView.indexAt(xPos, yPos)
      var relaxationStep = dp(10)

      // if nothing found (e.g. section is at chosen point) -> retry with relaxation factor
      var relaxation = relaxationStep
      while(index === -1 && relaxation <= listView.height * 0.5) {
        var next = listView.indexAt(listView.width * 0.5, yPos + relaxation)
        var prev = listView.indexAt(listView.width * 0.5, yPos - relaxation)
        if(nextFirst)
          index = next > -1 ? next : prev
        else
          index = prev > -1 ? prev : next
        relaxation += relaxationStep
      }

      return index
    }

    // scrollToIndex - scrolls listview to item with given index
    function scrollToIndex(idx) {
      listView.positionViewAtIndex(idx, ListView.Beginning)
    }

    // row item config
    delegate: ScheduleListItem {
      item: listModel.get(index)
      onSelected: dayScheduleItem.itemClicked(item)
      parentView: listView
      onLoadingChanged: {
        if(loading)
          _.loadingCount++
        else
          _.loadingCount--
      }
      Component.onDestruction: if(loading) _.loadingCount--
    }

    // section config
    section.property: "section"
    section.delegate: SimpleSection {
      style.compactStyle: true
      textItem.horizontalAlignment: Text.AlignHCenter
    }

    emptyText.text: "No events."
  }

  // add or remove item from favorites
  function toggleFavorite(item) {
    logic.toggleFavorite(item)
  }

  // itemIndexForTime - get item index for time in ms
  function itemIndexForTime(time) {
    var i = -1
    for(var idx in scheduleData) {
      i++
      var data = scheduleData[idx]
      var eventTime = getEventDateTime(data).getTime()
      if(eventTime >= time) {
        break
      }
    }
    return i
  }

  // getEventDateTime - build JS date object for event
  function getEventDateTime(event) {
    var date = new Date(event.day+"T"+event.start+".000"+dataModel.timeZone)
    return date
  }

  // getCurrentDateTime - returns JS date object for current time
  function getCurrentDateTime() {
    return new Date()
  }

  // jumpToNow - jumps to list entry matching current time
  function jumpToNow() {
    var now = getCurrentDateTime().getTime()
    var nowItemIndex = itemIndexForTime(now)
    if(nowItemIndex >= 0)
      listView.scrollToIndex(nowItemIndex)
  }
}
