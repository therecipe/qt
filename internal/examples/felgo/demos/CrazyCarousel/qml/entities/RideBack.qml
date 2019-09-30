import QtQuick 2.0
import Felgo 3.0

EntityBase {
    id: rideBack
    entityType: "rideBack"

    z: 11
    x: frontRide.x
    y: frontRide.y
    width: frontRide.width
    height: frontRide.height

    property QtObject frontRide: null // for synchronizing movement animation
    property alias image: image

    MultiResolutionImage {
        id: image
        width: parent.width
        height: parent.height
    }

    // provides up/down movement
    MovementAnimation {
        id: upDownMovement
        target: rideBack
        property: "y"
        velocity: frontRide.velocity // start with up movement
        running: true

        // limits the y property (defines possible movement area)
        minPropertyValue: frontRide.minY
        maxPropertyValue: frontRide.maxY

        // change direction after min/max is reached
        onLimitReached: {
            upDownMovement.velocity = frontRide.velocity
        }
    }
}
