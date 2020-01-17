import QtQuick 2.4
import QtQuick.Controls 1.2

import Felgo 3.0
import "../widgets"
import "../model"

ListPage {
  id: page

  title: qsTr("Messages")

  listView.emptyText.text: qsTr("No messages")

  model: JsonListModel {
    id: listModel
    source: dataModel.messages
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
}
