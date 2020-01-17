import QtQuick 2.0

import Felgo 3.0

EntityBase {
  entityType: "player"

  // the key-pressed-signals get emitted from the scene when key presses are detected
  // key pressed cant be detected here, because this item has no size
  signal leftPressed(variant event)
  signal rightPressed(variant event)
  signal upPressed(variant event)
  signal downPressed(variant event)

  signal died

  // gets increased over time - it has the same value as the y value of the level
  property int score: 0
  // gets increased when a coin is collected
  property int bonusScore: 0
  // the total score is the one that gets displayed
  property int totalScore: score + (bonusScore * bonusScoreForCoin)
  // the total count of deaths within this gaming session
  property int deaths: 0

  // this gets added to bonusScore every time the player catches a coin
  property int bonusScoreForCoin: 300

  property alias controller: twoAxisController

  // these are the settings for balancing!
  property real upValue: 550
  property real downValue: 5
  property real rightValue: 250
  property real leftValue: -rightValue

  property bool __isJumping: true
  // cant be initialized with null
  property date lastJumpTime: new Date

  // this is needed internally to find out if the image should be inverted
  property bool __isLookingRight: true

  // this is called when debugging and modifying values at runtime
  onRightValueChanged: console.debug("rightValue changed to", rightValue)

  // by setting this to true, the removeAllEntities() does not affect this entity, which is called from Level.stopGame()
  preventFromRemovalFromEntityManager: true

  Image {
    id: sprite
    source: "../../assets/img/chicken-front.png"
    anchors.centerIn: parent

    width: 20
    height: 20
    visible: false // will be visible in default state
  }
  Image {
    id: spriteMovement
    source: "../../assets/img/chicken-left01.png"
    anchors.centerIn: parent
    mirror: __isLookingRight
    width: sprite.width
    height: sprite.height
    visible: false
  }
  Image {
    id: spriteFlying
    source: "../../assets/img/chicken-fly.png"
    anchors.centerIn: parent
    width: 25
    height: 20
    visible: false
  }

  // gets increased when a collision with a block occurs in onBeginContact, and decreased in onEndContact
  // when no blockCollisions happened, the fly-state is active!
  property int blockCollisions: 0

  BoxCollider {
    id: collider
    bodyType: Body.Dynamic
    fixedRotation: true

    // balancing settings:
    linearDamping: 5.0
    // set friction between 0 and 1
    friction: 0.6
    // restitution is bounciness - don't bounce, because then the state would be changed too often
    restitution: 0

    // this is needed, because otherwise when resetting the level (and the player position), and the body was sleeping before it wouldn't fall down immediately again, because it hasn't woken up from sleeping!
    sleepingAllowed: false

    anchors.fill: sprite

    fixture.onBeginContact: {
      var fixture = other;
      var body = fixture.getBody()
      var collidedEntity = body.target;
      var collidedEntityType = collidedEntity.entityType;
      if(collidedEntityType === "coin") {        
        // the coin is pooled for better performance
        collidedEntity.removeEntity();

        // increase the score, and also play a pling sound
        bonusScore++;

        coinSound.play()
      } else if(collidedEntityType === "roost") {
        blockCollisions++;
      }
    }

    fixture.onEndContact: {

      var fixture = other;
      var body = fixture.getBody();
      var collidedEntity = body.target;
      var collidedEntityType = collidedEntity.entityType;
      if(collidedEntityType === "roost") {
        blockCollisions--;
      }
    }
  }

  SoundEffect {
    id: coinSound
    source: "../../assets/snd/pling.wav"
  }

  TwoAxisController {
    id: twoAxisController

    onXAxisChanged: {
      console.debug("xAxis changed to", xAxis)
      if(xAxis>0)
        __isLookingRight = true;
      else if(xAxis<0)
        __isLookingRight = false;
    }
  }

  Timer {
    id: updateTimer
    interval: 60
    running: true
    repeat: true
    onTriggered: {
      // this must be done every frame, because the linearVelocity gets reduced because of the damping!
      var xAxis = controller.xAxis;
      if(xAxis) {
        collider.body.linearVelocity.x = xAxis*rightValue;
      }
    }
  }

  state: {
    if(blockCollisions==0)
      return "fly";
    else {
      if(controller.xAxis !== 0) {
        return "moveLeftRight";
      }
      return "";
    }
  }

  states: [
    State {
      name: ""
      PropertyChanges { target: sprite; visible: true }
    },
    State {
      name: "fly"
      PropertyChanges { target: spriteFlying; visible: true }
    },
    State {
      name: "moveLeftRight"
      PropertyChanges { target: spriteMovement; visible: true }
    }
  ]

}
