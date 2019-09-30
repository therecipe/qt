import QtQuick 2.0
import Felgo 3.0

EntityBase {
    id: enemy
    entityType: "enemy"

    property alias image: image
    property double speed: 10 * gameScene.logic.speedUp

    MultiResolutionImage {
        id: image
        width: parent.width
        height: parent.height
    }

    // provides up/down movement
    MovementAnimation {
        id: upDownMovement
        target: enemy // set animation target to enemy object
        property: "y" // we animate the y position
        velocity: -enemy.speed // start with up movement
        running: true // animation starts automatically

        // limits the y property (defines possible movement area)
        // we set a random area for each enemy between
        minPropertyValue: -(Math.random() * 10 + 5) // min y is random between -5 and -15
        maxPropertyValue: Math.random() * 10 + 5 // max y is random between +5 to +15

        // change direction after min/max is reached (e.g. move up after down movement is finished)
        onLimitReached: {
            if(upDownMovement.velocity > 0)
                upDownMovement.velocity = -enemy.speed;
            else
                upDownMovement.velocity = enemy.speed;
        }
    }
}
