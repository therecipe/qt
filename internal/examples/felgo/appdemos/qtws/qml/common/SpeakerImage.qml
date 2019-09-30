import Felgo 3.0
import QtQuick 2.0

Rectangle {
  id: speakerImg
  width: webImg.implicitWidth
  height: webImg.implicitHeight

  color: Theme.backgroundColor
  border.color: Theme.secondaryTextColor
  border.width: dp(1)
  radius: width * 0.5

  readonly property bool loaded: webImg.img.status === Image.Ready
  readonly property bool loading: webImg.loading

  property bool activatePictureViewer: false

  property alias img: webImg
  property alias defaultImg: defaultImg
  property alias source: webImg.actualSource
  property real padding: dp(6)

  RoundedImage {
    id: defaultImg
    anchors.fill: webImg
    radius: webImg.radius
    img.defaultSource: Qt.resolvedUrl("../../assets/icon_engineer.png")
    visible: !webImg.visible
  }

  RoundedImage {
    id: webImg
    width: parent.width - padding
    height: parent.width - padding
    anchors.centerIn: parent
    radius: parent.radius / parent.width * width
    source: actualSource
    visible: img.status === Image.Ready

    property string actualSource: ""
    property bool loading: false

    // handle loading status for image
    Connections {
      target: webImg.img
      onStatusChanged: {
        if(webImg.img.status === Image.Loading && !webImg.loading)
          webImg.loading = true
        else if(webImg.img.status !== Image.Loading && webImg.loading)
          webImg.loading = false
      }
    }

    // reload image be resetting source property
    function reload() {
      webImg.source = ""
      webImg.source = Qt.binding(function() { return actualSource })
    }
  }

  // allow image reload on clicking if failed before, or call picture viewer
  MouseArea {
    enabled: (webImg.img.status !== Image.Ready && webImg.img.status !== Image.Loading) || activatePictureViewer
    anchors.fill: parent
    onClicked: {
      if(webImg.img.status !== Image.Ready && webImg.img.status !== Image.Loading)
        webImg.reload()
      else
        PictureViewer.show(getApplication(), webImg.source, false)
    }
  }

  // auto-reload image when device is online again
  Connections {
    target: webImg.img.status !== Image.Ready ? getApplication() : null
    onIsOnlineChanged: if(isOnline) webImg.reload()
  }
}
