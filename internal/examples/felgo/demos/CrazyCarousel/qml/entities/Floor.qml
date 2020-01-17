import QtQuick 2.0
import Felgo 3.0

// non-visible floor for collision testing
EntityBase {
    id: floor
    entityType: "floor"

    width: parent.width // width of the scene
    height: 5

    // for collision testing
    BoxCollider {
        anchors.fill: parent
        bodyType: Body.Static // the body shouldn't move
    }
}
