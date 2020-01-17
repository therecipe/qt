import QtQuick 2.4
import Felgo 3.0

ListPage {
  title: qsTr("Recent")

  emptyText.text: qsTr("No recent messages")

  // Use a predefined delegate but change some of its layout parameters
  delegate: SimpleRow {
    image.radius: image.height
    image.fillMode: Image.PreserveAspectCrop
    autoSizeImage: true
    style.showDisclosure: false
    imageMaxSize: dp(48)
    detailTextItem.maximumLineCount: 1
    detailTextItem.elide: Text.ElideRight

    onSelected: {
      globalNavStack.popAllExceptFirstAndPush(detailPageComponent, {
                                     person: item.text,
                                     newMsgs: [{me: true, text: item.detailText}]
                                   })
    }
  }

  Component { id: detailPageComponent;  ConversationPage { } }


  // Model should be loaded from your messaging backend
  model: [
    { text: "Tom McEloy", detailText: "Sorry for the late reply ...", image: Qt.resolvedUrl("../../assets/portrait0.jpg") },
    { text: "Leah Douglas", detailText: "Hahaha :D", image: Qt.resolvedUrl("../../assets/portrait1.jpg") }
  ]
}
