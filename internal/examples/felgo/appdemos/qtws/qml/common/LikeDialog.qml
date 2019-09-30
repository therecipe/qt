import QtQuick 2.0
import QtQuick.Controls 1.4
import QtQuick.Controls.Styles 1.4
import Felgo 3.0
import "../common"

// feedback window to contact Felgo
Dialog {
  id: likeDialog
  negativeAction: false
  positiveAction: false
  autoSize: true
  outsideTouchable: false

  Item {
    id: contentArea
    width: parent.width
    height: content.height

    Column {
      id: content
      width: parent.width
      anchors.horizontalCenter: parent.horizontalCenter
      anchors.top: parent.top
      anchors.margins: spacing
      spacing: dp(20)

      // like text
      Text {
        id: likeText
        horizontalAlignment: Text.AlignHCenter
        anchors.horizontalCenter: parent.horizontalCenter
        text: "Do you like this app?"
        color: Theme.textColor
        font.pixelSize: sp(18)
        width: parent.width * 0.8//- anchors.topMargin * 2
        wrapMode: Text.Wrap
      }

      // like / dislike buttons
      Row {
        anchors.horizontalCenter: parent.horizontalCenter
        spacing: parent.spacing * 2

        Rectangle {
          width: dp(72)
          height: width / redThumbImage.sourceSize.width * redThumbImage.sourceSize.height
          radius: dp(10)
          color: Theme.secondaryBackgroundColor
          anchors.verticalCenter: parent.verticalCenter

          AppImage {
            id: redThumbImage
            source: Qt.resolvedUrl("../../assets/RedThumb.png")
            anchors.fill: parent
            anchors.margins: dp(12)
            fillMode: AppImage.PreserveAspectFit
          }

          MouseArea {
            anchors.fill: parent
            onClicked: likeDialog.canceled()
          }
        }

        Rectangle {
          width: dp(72)
          height: width / greenThumbImage.sourceSize.width * greenThumbImage.sourceSize.height
          radius: dp(10)
          color: Theme.secondaryBackgroundColor
          anchors.verticalCenter: parent.verticalCenter

          AppImage {
            id: greenThumbImage
            source: Qt.resolvedUrl("../../assets/GreenThumb.png")
            anchors.fill: parent
            anchors.margins: dp(14)
            fillMode: AppImage.PreserveAspectFit
          }

          MouseArea {
            anchors.fill: parent
            onClicked: likeDialog.accepted()
          }
        }

      }
    }
  } // content area

}
