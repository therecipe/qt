import QtQuick 2.7            //Image
import QtGraphicalEffects 1.0 //ColorOverlay

import Theme 1.0              //Theme
import LeftTemplate 1.0       //LogoTemplate

LogoTemplate {
  id: template

  Image {
    id: logo
    source: "qrc:/qml/assets/logo.png"

    fillMode: Image.PreserveAspectFit
    width: parent.width
    height: parent.height
    mipmap: true

    //or with svg: sourceSize.width: parent.width
    //or with svg: sourceSize.height: parent.height
  }

  ColorOverlay {
    anchors.fill: logo
    source: logo
    color: Theme.accent
  }

  MouseArea {
    anchors.fill: template
    onClicked: template.clicked()
  }
}
