import QtQuick 2.0
import Felgo 3.0
 // needed for Body.Static

EntityBase {
  entityType: "henhouseWindow"

  poolingEnabled: true

  Image {
    id: sprite
    source: "../../assets/img/window3.png"

    // the size should not be bound to the grid as it is only a visual effect, but it gets set to the image size
    width: 64
    height: 60

  }

  property alias collider: collider
  BoxCollider {
    id: collider
    bodyType: Body.Static

    anchors.fill: sprite
    sensor: true

  }

}
