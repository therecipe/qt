import QtQuick 2.0
import Felgo 3.0

// player character entity
EntityBase {
    id: player
    entityType: "player"

    width: 65
    height: 96

    property bool isMoving: false // true when player is jumping
    property string nextMove: "" // for memorizing an additional player move
    property QtObject frontRide: null // current player ride for synchronizing the up/down-movement

    MultiResolutionImage {
        id: image
        width: parent.width
        height: parent.height
        x: -width / 2
        y: -height / 2
        source: "../../assets/player.png"
    }

    // for collision detection
    BoxCollider {
        id: boxCollider
        width: 25
        height: 50
        anchors.centerIn: image
        collisionTestingOnlyMode: true
    }

    // provides horicontal movement (switch ride)
    MovementAnimation {
        id: switchMovement
        target: player
        property: "x"
    }

    // provides vertical movement (jump)
    MovementAnimation {
        id: jumpMovement
        target: player
        property: "y"

        // limits the y property between 0 and the default player position
        minPropertyValue: 0
        maxPropertyValue: gameScene.playerY

        // trigger next move or stop movement, when player hits the ground
        onLimitReached: {
            isMoving = false
            image.source = "../../assets/player.png"
            if(nextMove !== "") {
                if(nextMove === "right")
                    moveRight()
                else
                    moveLeft()
            }
            else {
                jumpMovement.stop()
                startUpDownMovement()
            }
        }
    }

    // triggers player movement to the left
    function moveLeft() {
        // only if player is not on the first ride
        if(player.x - gameScene.laneWidth < 0)
            return

        if(!isMoving) {
            var distance = gameScene.laneWidth
            // use faster movement when game speed is higher
            var multiplier = 1 + (logic.speedUp / 8 * 3) // multiplier between 1 and 4

            // configure movement
            switchMovement.velocity = -distance * 4.5 * multiplier
            switchMovement.minPropertyValue = player.x - distance
            switchMovement.maxPropertyValue = gameScene.width

            // do animation
            isMoving = true
            nextMove=""
            upDownMovement.stop()
            switchMovement.start()
            jump()
        }
        else {
            nextMove = "left" //memorize attempted move
        }
    }

    // triggers player movement to the right
    function moveRight() {
        // only if player is not on the last ride
        if(player.x + gameScene.laneWidth > gameScene.width)
            return

        if(!isMoving) {
            var distance = gameScene.laneWidth

            // use faster movement when game speed is higher
            var multiplier = 1 + (logic.speedUp / 8 * 3) // multiplier between 1 and 4

            // configure movement
            switchMovement.velocity = distance * 4.5 * multiplier
            switchMovement.minPropertyValue = 0
            switchMovement.maxPropertyValue = player.x + distance

            // do animation
            isMoving = true
            nextMove = ""
            upDownMovement.stop()
            switchMovement.start()
            jump()
        }
        else {
            nextMove = "right" // memorize attempted move
        }
    }

    // start jump
    function jump() {
        // reset y position if player is under the limit (due to the moving rides)
        if(player.y >= gameScene.playerY)
            player.y = gameScene.playerY

        // use faster movement when game speed is higher
        var multiplier1 = 0.8 + (logic.speedUp / 8 * 1.0) // multiplier between 0.8 and 1.8
        var multiplier2 = 1 + (logic.speedUp / 8 * 4) // multiplier between 1 and 5

        // configure jump and start movement
        jumpMovement.velocity = -500 * multiplier1
        jumpMovement.acceleration = 2000 * multiplier2
        jumpMovement.start()
        jumpSound.stop()
        jumpSound.play()
        image.source = "../../assets/player_jump.png"
    }

    // stop all movement
    function stopMovement() {
        jumpMovement.stop()
        switchMovement.stop()
        upDownMovement.stop()
        isMoving = false
        nextMove = ""
    }

    // jump sound
    SoundEffect {
        id: jumpSound
        source: "../../assets/sound/Jump.wav"
    }

    // provides up/down movement along with rides
    MovementAnimation {
        id: upDownMovement
        target: player
        property: "y"

        // change direction after min/max is reached
        onLimitReached: {
            upDownMovement.velocity = frontRide.velocity
        }
    }

    // move player up and down along with ride he is currently standing on
    function startUpDownMovement() {
        // determine which ride the player is standing on
        var laneNr = 0
        for(var i = 1; i <= gameScene.lanes; i++) {
            var minX = (i - 1) * gameScene.laneWidth
            var maxX = i * gameScene.laneWidth
            var playerOnRide = player.x < maxX && player.x > minX

            if(playerOnRide) {
                laneNr = i
                break
            }
        }

        // retrieve entity of the current ride
        switch(laneNr) {
            case 1: player.frontRide = gameScene.ride1; break
            case 2: player.frontRide = gameScene.ride2; break
            case 3: player.frontRide = gameScene.ride3; break
            case 4: player.frontRide = gameScene.ride4; break
            default: player.frontRide = null
        }

        // move player with ride
        if(player.frontRide === null)
            player.y = gameScene.playerY
        else {
            // configure and start movement animation
            player.y = player.frontRide.y + 3
            upDownMovement.stop()
            upDownMovement.velocity = frontRide.velocity
            upDownMovement.maxPropertyValue = frontRide.maxY + 3
            upDownMovement.minPropertyValue = frontRide.minY + 3
            upDownMovement.start()
        }
    }
}
