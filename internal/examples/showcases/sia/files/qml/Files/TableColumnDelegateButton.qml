import QtQuick 2.7            //Image
import QtQuick.Controls 2.1   //Button
import QtGraphicalEffects 1.0 //ColorOverlay

import Theme 1.0              //Theme

Button {
  id: button

  property alias source: image.source
  property string accent

  background: null

  contentItem: Item {
    Image {
      id: image
      anchors.centerIn: parent
    }

    ColorOverlay {
      id: overlay
      anchors.fill: image
      source: image
      color: button.down || button.hovered ? button.accent : Theme.font
    }
  }
}