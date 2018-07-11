import QtQuick 2.7            //Rectangle
//import QtGraphicalEffects 1.0 //FastBlur

import Theme 1.0              //Theme

import ViewTemplate 1.0       //ViewTemplate
import Left 1.0               //Left
import Top 1.0                //Top

ViewTemplate {
  Rectangle {
    id: root
    anchors.fill: parent
    color: Theme.background

    Left {
      id: leftControl
      anchors {
        top: parent.top
        left: parent.left
        bottom: parent.bottom
      }
      width: Theme.minWidth * 0.2
    }

    Top {
      id: topControl
      anchors {
        top: parent.top
        left: leftControl.right
        right: parent.right
      }
      height: Theme.minHeight * 0.09
    }

    Stack {
      anchors {
        top: topControl.bottom
        left: leftControl.right
        right: parent.right
        bottom: parent.bottom
      }
    }
  }

  /*
  FastBlur {
    anchors.fill: root
    source: root
    radius: 32
    visible: !root.visible
  }

  onBlur: root.visible = !b
  */

  onWidthChanged: Theme.currentWidth = width
  onHeightChanged: Theme.currentHeight = height
}
