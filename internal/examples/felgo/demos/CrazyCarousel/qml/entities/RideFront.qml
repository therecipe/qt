import QtQuick 2.0
import Felgo 3.0

EntityBase {
    id: rideFront
    entityType: "rideFront"

    z: 9

    property alias image: image
    property alias minY: upDownMovement.minPropertyValue
    property alias maxY: upDownMovement.maxPropertyValue
    property alias velocity: upDownMovement.velocity
    property double speed: 10 * gameScene.logic.speedUp

    MultiResolutionImage {
        id: image
        width: parent.width
        height: parent.height
    }

    // provides up/down movement
    MovementAnimation {
        id: upDownMovement
        target: rideFront
        property: "y"
        velocity: -rideFront.speed // start with up movement
        running: true

        // change direction after min/max is reached
        onLimitReached: {
            if(rideFront.velocity > 0)
                rideFront.velocity = -rideFront.speed
            else
                rideFront.velocity = rideFront.speed
        }
    }
}
