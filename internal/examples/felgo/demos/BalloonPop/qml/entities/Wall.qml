import QtQuick 2.0
import Felgo 3.0

EntityBase {
  entityType: "wall"
  // default width and height
  width: 1
  height: 1

  // only collider since we want the wall to be invisible
  BoxCollider {
    anchors.fill: parent
    bodyType: Body.Static // the body shouldn't move
  }
}
