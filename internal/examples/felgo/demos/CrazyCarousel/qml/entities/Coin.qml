import QtQuick 2.0
import Felgo 3.0

EntityBase {
    id: coin
    entityType: "coin"

    width: 32
    height: 33

    // properties for handling the movement towards a goal
    property double goalX: 0
    property double startX: 0
    property double startY: 0
    property double speed: 3000

    MultiResolutionImage {
        id: image
        width: parent.width
        height: parent.height
        x: -width / 2
        y: -height / 2
        source: "../../assets/coin.png"
    }

    // for collision detection
    BoxCollider {
        id: boxCollider
        width: 20
        height: 20
        anchors.centerIn: image
        collisionTestingOnlyMode: true

        fixture.onBeginContact: {
            // for access of the collided entity and the entityType and entityId:
            var collidedEntity = other.getBody().target
            var collidedEntityType = collidedEntity.entityType

             // only increase points if coin was collected
            if(collidedEntityType === "player") {
                boxCollider.active = false
                parent.visible = false
                coinSound.play()
                gameScene.logic.increasePoints()
            }
            else {
                entityManager.removeEntityById(parent.entityId)
            }
        }
    }

    scale: 0.2 // initial scale
    // parallel animation of x-position, y-position and scaling
    ParallelAnimation {
        running: true // all animations are started
        NumberAnimation {
            target: coin
            property: "x"
            from: startX
            to: goalX
            duration: speed
            easing.type: Easing.InCubic
        }
        NumberAnimation {
            target: coin
            property: "y"
            from: startY
            to: gameScene.gameWindowAnchorItem.height
            duration: speed
            easing.type: Easing.InCubic
        }
        NumberAnimation {
            target: coin
            property: "scale"
            to: 0.9
            duration: speed
            easing.type: Easing.InCubic
        }
    }

    // coin sound
    SoundEffect {
        id: coinSound
        source: "../../assets/sound/Coin.wav"
        onPlayingChanged: {
            if(coinSound.playing === false) {
                entityManager.removeEntityById(parent.entityId)
            }
        }
    }
}
