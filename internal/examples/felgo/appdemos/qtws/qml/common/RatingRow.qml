import Felgo 3.0
import QtQuick 2.0
import "."

Item {
  id: ratingRow

  property int stars: 0
  property bool past: false
  property bool isRated: false
  property var talk

  height: col.height

  Component.onCompleted: {
    checkTimer.start()
    // the model will return -1 if nothing is stored, or 1-5 if a stored rating is successfully loaded
    var rating = dataModel.getRating(talk.id)
    isRated = rating > -1
    stars = rating
  }

  Column {
    id: col
    width: parent.width
    Item {
      width: parent.width
      height: dp(40)

      Row {
        anchors.centerIn: parent

        Repeater {
          model: 5

          Item {
            width: dp(50)
            height: dp(40)

            Image {
              source: stars > index ? "../../assets/Qt_logo.png" : "../../assets/Qt_logo_iOS_off.png"
              opacity: past ? stars > index ? 1 : 0.2 : 0
              width: parent.width - dp(20)
              fillMode: Image.PreserveAspectFit
              anchors.centerIn: parent
            }

            // the stars can only be clicked if the talk has ended and it is not yet rated
            MouseArea {
              anchors.fill: parent
              enabled: past && !isRated
              onClicked: starClicked(index+1)
            }
          }
        }
      }

      AppText {
        anchors.centerIn: parent
        visible: !past
        text: "Rating opens after talk ends at " + talk.end
        wrapMode: Text.WordWrap
        font.pixelSize: sp(14)
        width: parent.width
        horizontalAlignment: AppText.AlignHCenter
        font.italic: true
      }
    }

    Item {
      id: okButtonWrapper
      visible: false
      width: parent.width
      clip: true
      Behavior on height {NumberAnimation{duration: 250}}

      Item {
        width: parent.width
        height: dp(40)
        y: dp(20)

        AppButton {
          anchors.centerIn: parent
          flat: false
          enabled: past && !isRated
          text: "Rate"
          onClicked: {
            saveRating()
          }
        }
      }
    }
  }

  // just checking every 10 seconds if the talk end time is reached is surely enough
  Timer {
    id: checkTimer
    interval: 10000
    repeat: true
    triggeredOnStart: true
    onTriggered: checkIfPast()
  }

  function starClicked(rating) {
    stars = rating
    // if this was the first time the stars are clicked, we can show the rate button
    if(!okButtonWrapper.visible) {
      okButtonWrapper.visible = true
      okButtonWrapper.height = dp(60)
    }
  }

  function saveRating() {
    okButtonWrapper.height = 0
    isRated = true
    amplitude.logEvent("Rate Talk", {"title" : talk.title, "talkId" : talk.id, "rating" : stars})
    logic.storeRating(talk.id, stars)
    if(gameNetwork.userScoresInitiallySynced)
      gameNetwork.reportRelativeScore(1)
    else
      gameNetwork.addScoreWhenSynced += 1
  }

  function checkIfPast() {
    var nowTime = new Date().getTime()
    var utcDateStr = talk.day+"T"+talk.end+".000"+dataModel.timeZone
    var talkEndTime = new Date(utcDateStr).getTime() - (15 * 60 * 1000)

    if(nowTime > talkEndTime) {
      checkTimer.stop()
      past = true
    }
  }
}
