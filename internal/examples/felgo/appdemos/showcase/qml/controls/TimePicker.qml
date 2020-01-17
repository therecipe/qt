import QtQuick 2.5

Rectangle {
  id: timePicker
  color: "#fff"
  clip: true
  property int numberOfItems: 7
  property string fontFamily
  property var time: ({hour: 0, minute: 0})
  Rectangle {
    width: parent.width
    height: 1
    y: (parent.height - listView.delegateHeight) / 2
    color: "#D0D0D0"
  }
  Rectangle {
    width: parent.width
    height: 1
    y: (parent.height - listView.delegateHeight) / 2 + listView.delegateHeight
    color: "#D0D0D0"
  }
  ListView {
    id: listView
    x: 0
    width: parent.width
    y: -parent.height * 0.5
    height: parent.height * 2
    property int delegateHeight: height / numberOfItems
    model: 96000
    spacing: 1
    highlightRangeMode: ListView.StrictlyEnforceRange
    preferredHighlightBegin: (height - delegateHeight) / 2
    preferredHighlightEnd: (height + delegateHeight) / 2
    delegate: Item {
      id: contentItem
      width: listView.width
      height: listView.delegateHeight
      Rectangle {
        anchors.fill: parent
        color: "#3333ee33"
        visible: false
      }
      Text {
        id: innerText
        property int hour: Math.floor(index % 96 / 4)
        property int minute: 15 * (index % 4)
        text: leadingZero(hour) + ":" + leadingZero(minute)
        anchors.centerIn: parent
        font.pixelSize: listView.delegateHeight * 0.8
        font.family: fontFamily
        color: contentItem.ListView.isCurrentItem ? "black" : "#999"
        transform: [
          Rotation {
            origin.x: innerText.width / 2
            origin.y: innerText.height / 2
            axis { x: 1; y: 0; z: 0 }
            angle: {
              var middle = contentItem.ListView.view.contentY - contentItem.y + contentItem.ListView.view.height / 2
              var calculated = (middle - contentItem.height / 2) / contentItem.height * 40
              if (calculated < -90)
                return -90
              else if (calculated > 90)
                return 90
              else
                return calculated
            }
          },
          Scale {
            origin.x: innerText.width / 2
            origin.y: innerText.height / 2
            xScale: {
              // scaled 1 in middle position -> 0 when reaching edges
              var scaled = (contentItem.y - contentItem.ListView.view.contentY + contentItem.height * 0.5) / contentItem.ListView.view.height * 2
              if (scaled > 1) scaled = 2 - scaled
              return Math.max(0, scaled)
            }
            yScale: xScale
          },
          Translate {
            y: {
              var scaled = (contentItem.y - contentItem.ListView.view.contentY + contentItem.height * 0.5) / contentItem.ListView.view.height * 2
              scaled = Math.max(0, scaled)
              scaled = 1 - scaled
              return scaled * scaled * scaled * contentItem.height * 3
            }
          }
        ]
      }
    }
    Component.onCompleted: {
      // Scrolls to middle of list
      positionViewAtIndex(model * 0.5 - (listView.numberOfItems > 2 ? 1 : 0), ListView.SnapPosition)
    }
    onMovementEnded: {
      var item = currentIndex % 96
      var hour = Math.floor(item / 4)
      var minute = (item - (hour * 4)) * 15
      console.debug("TIME IS:", hour + ":" + minute)
      timePicker.time = {hour: hour, minute: minute}
    }
  }
  function setTime(newTime) {
    listView.positionViewAtIndex(newTime && newTime.hour * 4 + newTime.minute / 15 || 0, ListView.Center)
    timePicker.time = newTime
  }
  function leadingZero(number) {
    return ('00' + number).slice(-2)
  }
}
