import QtQuick 2.7            //Image
import QtQuick.Controls 2.1   //Label
import QtGraphicalEffects 1.0 //ColorOverlay

import Theme 1.0              //Theme

Item {
  property alias image : image.source
  property alias title : label.text
  default property alias contentItem : container.children

  anchors {
    top: parent.top
    left: parent.left
    right: parent.right
  }
  height: Theme.minHeight * 0.09

  Item {
    id: imageItem

    anchors {
      top: parent.top
      left: parent.left
      bottom: parent.bottom
    }
    width: height

    Image {
      id: image
      anchors.centerIn: parent

      sourceSize.width: 48
      sourceSize.height: 48
    }

    ColorOverlay {
      anchors.fill: image
      source: image
      color: Theme.fontHighlight
    }
  }

  Label {
    id: label

    anchors {
      top: parent.top
      left: imageItem.right
      bottom: parent.bottom
    }

    color: Theme.fontHighlight
    font.pointSize: 36
    font.bold: true

    horizontalAlignment: Text.AlignHCenter
    verticalAlignment: Text.AlignVCenter
  }

  Item {
    id: container

    anchors {
      top: parent.top
      left: label.right
      right: parent.right
      bottom: parent.bottom
    }
  }
}
