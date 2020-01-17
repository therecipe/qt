import Felgo 3.0
import QtQuick 2.0
import "../common"

ListPage {
  id: speakersPage

  property var speakersModel: dataModel.speakers !== undefined ? dataModel.speakers : {}

  title: "Speakers"

  model: JsonListModel {
    id: listModel
    source: prepareSpeakers(speakersModel)
    keyField: "id"
    fields: [ "id","first_name","last_name","abstract","avatar","talks","firstLetter" ]
  }
  section.property: "firstLetter"
  section.delegate: SimpleSection {
    style.compactStyle: Theme.isIos
  }

  delegate: SpeakerRow {
    speaker: listModel.get(index)
    small: true
    onClicked: {
      if(Theme.isAndroid)
        speakersPage.navigationStack.popAllExceptFirstAndPush(Qt.resolvedUrl("SpeakerDetailPage.qml"), { speakerID: speaker.id })
      else
        speakersPage.navigationStack.push(Qt.resolvedUrl("SpeakerDetailPage.qml"), { speakerID: speaker.id })
    }
  }

  listView.scrollIndicatorVisible: false

  SectionSelect {
    id: sectionSelect
    anchors.right: parent.right
    target: speakersPage.listView
  }

  // prepareSpeakers - build speaker model for display
  function prepareSpeakers(speakers) {
    var model = []
    for(var i in Object.keys(speakers)){
      var speakerID = Object.keys(speakers)[i];
      var speaker = speakers[parseInt(speakerID)]
      speaker["firstLetter"] = speaker["last_name"].charAt(0).toUpperCase()
      console.log("\nSPEAKERS: "+JSON.stringify(speaker))
      model.push(speaker)
    }
    model.sort(compareLastName);
    return model
  }

  function compareLastName(a,b) {
    if (a.last_name < b.last_name)
      return -1;
    if (a.last_name > b.last_name)
      return 1;
    return 0;
  }
}
