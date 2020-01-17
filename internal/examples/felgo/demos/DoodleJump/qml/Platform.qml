import QtQuick 2.2
import Felgo 3.0

EntityBase {
  id: platform

  entityType: "Platform" // always name your entityTypes

  width: 64 // visual width of our platform
  height: 32 // visual height of our platform

  // leaf image for platform
  Image {
    id: platformImg
    source: "../assets/leaf.png"
    anchors.fill: platform
  }

  // BoxCollider responsible for collision detection
  BoxCollider {
    id: platformCollider
    width: parent.width // actual width is the same as the parent entity
    height: parent.height - 20 // actual height is slightly smaller so the collision works smoother
    bodyType: Body.Dynamic // only Dynamic bodies can collide with each other
    collisionTestingOnlyMode: true // collisions are detected, but no physics are applied to the colliding bodies

    fixture.onBeginContact: {
      var otherEntity = other.getBody().target
      var otherEntityType = otherEntity.entityType

      if(otherEntityType === "Border") {
        platform.x = utils.generateRandomValueBetween(32, gameScene.width-64) // generate random x
        platform.y = 0 // the top of the screen
      }
    }
  }

  // platform movement
  MovementAnimation {
    id: movement
    target: platform // which object will be affected
    property: "y" // which property will be affected
    velocity:  frog.impulse / 2 // impulse is y velocity of the frog
    running: frog.y < 210 // move only when the frog is jumping over the limit
  }

  // wobble animation
  ScaleAnimator {
    id: wobbleAnimation
    target: platform
    running: false // default is false and it gets activated on every collision
    from: 0.9
    to: 1
    duration: 1000
    easing.type: Easing.OutElastic // Easing used get an elastic wobbling instead of a linear scale change
  }

  // function to start wobble animation
  function playWobbleAnimation() {
    wobbleAnimation.start()
  }
}
