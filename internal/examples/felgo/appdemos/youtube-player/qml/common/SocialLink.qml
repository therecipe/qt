import Felgo 3.0
import QtQuick 2.0

Item {
  property string targetUrl: ""
  property alias source: image.source

  width: dp(36)
  height: width
  visible: targetUrl !== ""

  Image {
    id: image
    width: parent.width
    height: width / sourceSize.width * sourceSize.height
    anchors.verticalCenter: parent.verticalCenter
    source: "../../assets/spotify.png"
  }

  MouseArea {
    anchors.fill: parent
    onClicked: nativeUtils.openUrl(targetUrl)
  }
}
