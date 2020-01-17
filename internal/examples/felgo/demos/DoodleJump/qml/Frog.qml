import QtQuick 2.0
import Felgo 3.0

EntityBase {
  id:frogEntity // the id we use as a reference inside this class

  entityType: "Frog" // always name your entityTypes

  state: frogCollider.linearVelocity.y < 0 ?  "jumping" : "falling" // change state according to the frog's y velocity

  property int impulse: y-frogCollider.linearVelocity.y // to move platforms

  property alias controller: twoAxisController // we make the frog's twoAxisController visible and accessible for the scene

  // frogCollider uses TwoAxisController to move the frog left or right.
  TwoAxisController {
    id: twoAxisController
  }

  // sprite for the frog, he can either be sitting or jumping
  SpriteSequence {
    id: frogAnimation

    defaultSource: "../assets/spritesheet.png"
    scale: 0.35 // our image is too big so we reduce the size of the original image to 35%
    anchors.centerIn: frogCollider
    // when frog jumps it turns to the direction he moves
    rotation: frogEntity.state == "jumping" ?
                 (system.desktopPlatform ?
                    twoAxisController.xAxis * 15
                    : (accelerometer.reading !== null ? -accelerometer.reading.x * 10 : 0))
                : 0

    Sprite {
      name: "sitting"
      frameWidth: 128
      frameHeight: 128
      startFrameColumn: 2
    }

    Sprite {
      name: "jumping"
      frameCount: 4
      frameRate: 8

      frameWidth: 128
      frameHeight: 256
      frameX: 0
      frameY: 128
    }
  }

  // collider to check for collisions and apply gravity
  BoxCollider {
    id: frogCollider

    width: 25 // width of the frog collider
    height: 5 // height of the frog collider

    bodyType: gameScene.state == "playing" ?  Body.Dynamic : Body.Static // do not apply gravity when the frog is dead or the game is not started

    // move the frog left and right
    linearVelocity.x: system.desktopPlatform ?
                        twoAxisController.xAxis * 200 :  //  for desktop
                        (accelerometer.reading !== null ? -accelerometer.reading.x * 100 : 0)   // for mobile

    fixture.onContactChanged: {
      var otherEntity = other.getBody().target
      var otherEntityType = otherEntity.entityType

      if(otherEntityType === "Border") {
        frogEntity.die()
      }
      else if(otherEntityType === "Platform" && frogEntity.state == "falling") {
        frogCollider.linearVelocity.y = -400

        otherEntity.playWobbleAnimation()
      }
    }
  }

  // animations handling
  onStateChanged: {
    if(frogEntity.state == "jumping") {
      frogAnimation.jumpTo("jumping") // change the current animation to jumping
    }
    if(frogEntity.state == "falling") {
      frogAnimation.jumpTo("sitting") // change the current animation to sitting
    }
  }

  onYChanged: {
    if(y < 200) {
      y = 200 // limit the frog's y value

      score += 1 // increase score
    }
  }

  // die and restart game
  function die() {
    // reset position
    frogEntity.x = gameScene.width / 2
    frogEntity.y = 220

    // reset velocity
    frogCollider.linearVelocity.y = 0

    // reset animation
    frogAnimation.jumpTo("sitting")

    gameNetwork.reportScore(score) // report the current score to the gameNetwork
    score = 0

    gameScene.state = "gameOver"

    gameNetwork.incrementAchievement("die100")
  }
}
