import Felgo 3.0
import QtQuick 2.0
import "../common"

Page {
  id: page
  title: track.title

  onPushed: amplitude.logEvent("View Track",{"track" : title})

  property var track

  TimetableDaySchedule {
    width: parent.width
    height: parent.height
    scheduleData: track.talks
    searchAllowed: false
    onItemClicked: {
      page.navigationStack.push(Qt.resolvedUrl("DetailPage.qml"), { item: item })
    }
  }

}
