import Felgo 3.0
import QtQuick 2.0
import "../common"

ListPage {
  id: channelPage
  title: dataModel.channelInfo && dataModel.channelInfo.snippet.title || dataModel.userName
  backgroundColor: Theme.secondaryBackgroundColor

  rightBarItem: IconButtonBarItem {
    visible: !isOnline
    icon: IconType.exclamation
    onClicked: NativeDialog.confirm("Offline", "You are currently offline, please go online to use the app.")
  }

  emptyText.text: getOfflineStatus()

  property bool showDisclosure: true

  signal selected(var playlist)

  model: JsonListModel {
    id: listModel
    source: dataModel.channelPlaylists
    keyField: "id"
  }

  delegate: SimpleRow {
    property var item: listModel.get(index)
    text: item.snippet.title
    imageSource: item.snippet.thumbnails ? item.snippet.thumbnails.high.url : ""
    autoSizeImage: true
    style.showDisclosure: showDisclosure
    onSelected: {
      channelPage.selected(item)
    }
  }
}

