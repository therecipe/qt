import QtQuick 2.0
import Felgo 3.0
import QtQuick.Layouts 1.2


GridLayout {
  id: _

  property string stationName
  property int stationBikes
  property int stationBoxes
  property alias stationFavorited: favButton.selected

  property bool backVisible

  signal backPressed
  signal favoritedPressed

  columns: 2
  columnSpacing: 0
  rowSpacing: 0

  Item {
    Layout.columnSpan: 2
    Layout.fillWidth: true
    Layout.preferredHeight: parent.height * 0.4

    AppText {
      id: innerText
      anchors.top: parent.top
      anchors.horizontalCenter: parent.horizontalCenter
      font.pixelSize: dp(24)
      color: blueDarkColor
      text: _.stationName
    }

    IconButton {
      anchors.left: parent.left
      anchors.leftMargin: dp(10)
      anchors.verticalCenter: innerText.verticalCenter
      icon: IconType.angleleft
      color: blueDarkColor
      selectedColor: blueLightColor
      size: dp(22)
      visible: backVisible

      onClicked: {
        backPressed()
      }
    }

    IconButton {
      id: favButton
      anchors.right: parent.right
      anchors.rightMargin: dp(10)
      anchors.verticalCenter: innerText.verticalCenter
      icon: IconType.hearto
      selectedIcon: IconType.heart
      color: blueDarkColor
      selectedColor: blueLightColor
      size: dp(22)

      toggle: true
      onToggled: favoritedPressed()
    }
  }

  Item {
    Layout.preferredWidth: parent.width / 2
    Layout.fillHeight: true

    Column {
      width: parent.width
      anchors.verticalCenter: parent.verticalCenter
      spacing: 0

      AppText {
        text: qsTr("Free Bikes")
        font.pixelSize: sp(16)
        color: blueDarkColor
        anchors.horizontalCenter: parent.horizontalCenter
      }

      AppText {
        font.pixelSize: sp(56)
        color: blueLightColor
        text: "" + _.stationBikes
        anchors.horizontalCenter: parent.horizontalCenter
      }
    }
  }

  Item {
    Layout.preferredWidth: parent.width / 2
    Layout.fillHeight: true

    Rectangle {
      width: px(1)
      height: parent.height
      color: greyLineColor
    }

    // Duplicate code
    Column {
      width: parent.width
      anchors.verticalCenter: parent.verticalCenter
      spacing: 0

      AppText {
        text: qsTr("Free Boxes")
        font.pixelSize: sp(16)
        color: blueDarkColor
        anchors.horizontalCenter: parent.horizontalCenter
      }

      AppText {
        font.pixelSize: sp(56)
        color: blueLightColor
        anchors.horizontalCenter: parent.horizontalCenter
        text: "" + _.stationBoxes
      }
    }
  }
}
