import QtQuick 2.0
import Felgo 3.0
 // needed for Body.Static

EntityBase {
  entityType: "coin"

  poolingEnabled: true

  // put them before the windows
  z:1

  Image {
    id: sprite
    source: "../../assets/img/corn.png"

    width: 7
    height: 10

    anchors.centerIn: parent
  }

  property alias collider: collider
  BoxCollider {
    id: collider
    bodyType: Body.Static

    anchors.fill: sprite
    // it should not affect the movement of the player
    sensor: true

  }

}
