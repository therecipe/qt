import QtQuick 2.0
import Felgo 3.0

EntityBase {
    id: bullet
    entityType: "bullet"

    width: 32
    height: 37

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
        source: "../../assets/bullet.png"
    }

    // for collision detection
    BoxCollider {
        id: boxCollider
        width: 15
        height: 15
        anchors.centerIn: image
        collisionTestingOnlyMode: true

        fixture.onBeginContact: {
            // for access of the collided entity and the entityType and entityId:
            var collidedEntity = other.getBody().target
            var collidedEntityType = collidedEntity.entityType

            // check if bullet was avoided
            if(collidedEntityType === "floor") {
                bulletSound.play()
                parent.visible = false
            }
            else {
                gameScene.logic.gameOver()
                entityManager.removeEntityById(parent.entityId)
            }
        }
    }

    scale: 0.2 // initial scale
    // parallel animation of x-position, y-position and scaling
    ParallelAnimation {
        running: true // all animations are started
        NumberAnimation {
            target: bullet
            property: "x"
            from: startX
            to: goalX
            duration: speed
            easing.type: Easing.InCubic
        }
        NumberAnimation {
            target: bullet
            property: "y"
            from: startY
            to: gameScene.gameWindowAnchorItem.height
            duration: speed
            easing.type: Easing.InCubic
        }
        NumberAnimation {
            target: bullet
            property: "scale"
            to: 1.0
            duration: speed
            easing.type: Easing.InCubic
        }
    }

    // bullet sound
    SoundEffect {
        id: bulletSound
        source: "../../assets/sound/Bullet.wav"
        onPlayingChanged: {
            if(bulletSound.playing === false) {
                entityManager.removeEntityById(parent.entityId)
            }
        }
    }
}
