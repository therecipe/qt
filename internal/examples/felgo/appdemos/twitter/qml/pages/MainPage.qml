import QtQuick 2.4
import QtQuick.Controls 1.2

import Felgo 3.0

import "../widgets"
import "../model"

ListPage {
  id: page

  title: qsTr("Home")

  rightBarItem: IconButtonBarItem {
    icon: IconType.plus
    onClicked: InputDialog.inputTextMultiLine(app,
                                              qsTr("New tweet"),
                                              qsTr("Enter text..."),
                                              function(ok, text) {
                                                if(ok) logic.addTweet(text)
                                              })
  }

  titleItem: Icon {
    icon: IconType.twitter
    color: "white"
    size: dp(24)
  }

  property int numItems: 10
  model: JsonListModel {
    id: listModel
    source: dataModel.timeline && dataModel.timeline.slice(0, numItems)
    keyField: "id"
  }
  delegate: TweetRow {
    id: row
    item: listModel.get(index)

    onSelected: {
      console.debug("Selected item at index:", index)
      navigationStack.push(detailPageComponent, { tweet: row.item })
    }

    onProfileSelected: {
      console.debug("Selected profile at index:", index)
      navigationStack.push(profilePageComponent, { profile: row.item.user })
    }
  }

  //load older tweets with visibility handler
  listView.footer: VisibilityRefreshHandler {
    canRefresh: dataModel.timeline ? numItems < dataModel.timeline.length : false
    onRefresh: loadOldTimer.start()
  }

  //load new tweets with pull handler
  pullToRefreshHandler {
    pullToRefreshEnabled: true
    onRefresh: loadNewTimer.start()
    refreshing: loadNewTimer.running
  }

  Timer {
    // Fake loading of new tweets in background
    id: loadNewTimer
    interval: 5000
    onTriggered: {
      logic.addTweet("Lorem Ipsum.")
    }
  }

  Timer {
    // Fake loading of older tweets in background
    id: loadOldTimer
    interval: 2000
    onTriggered: {
      var pos = listView.getScrollPosition()
      numItems += 10
      listView.restoreScrollPosition(pos)
    }
  }
}
