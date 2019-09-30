import QtQuick 2.0

import Felgo 3.0

EntityBase {
  id: entity
  entityType: "rocket"

  Component.onCompleted: {
      console.debug("Rocket.onCompleted, width:", width);
      applyForwardImpulse();
  }

  property real angleDeg

  rotation: angleDeg

  BoxCollider {
    id: boxCollider

    // the image and the physics will use this size; this is important as it specifies the mass of the body! it is in respect to the world size
    width: 50
    height: 20

    anchors.centerIn: parent

    density: 0.003
    friction: 0.4
    restitution: 0.5
    body.bullet: true
    // we prevent the physics engine from applying rotation to the rocket, because we will do it ourselves
    body.fixedRotation: true

    property var lastWall: null

    fixture.onBeginContact: {
      var fixture = other;
      var body = other.getBody();
      var otherEntity = body.target

      // get the entityType of the colliding entity
      var collidingType = otherEntity.entityType

      if(collidingType === "car" ||
              collidingType === "rocket") {
          entity.removeEntity()
          return
      }

      //can't hit the same wall twice, but onBeginContact called again after rotation has changed
      if(otherEntity === lastWall) {
        return;
      }
      lastWall = otherEntity

      //apply law of reflection, all calculations in degrees
      var normalAngle = 180 / Math.PI * Math.atan2(contactNormal.y, contactNormal.x)
      var angleDiff = normalAngle - entity.rotation
      var newAngle = entity.rotation + 2 * angleDiff + 180

      // manually set the entity rotation, because it is the target and its rotation will be used for the physics body
      entity.rotation = newAngle

      // it's important to clear the old velocity before applying the impulse, otherwise the rocket would get faster every time it collides with a wall!
      boxCollider.body.linearVelocity = Qt.point(0,0)

      applyForwardImpulse();
    }
  }

  Image {
    id: image
    source: "../../assets/img/rocket_green.png"
    anchors.centerIn: parent
    width: boxCollider.width
    height: boxCollider.height
  }

  function applyForwardImpulse() {
    var power = 1500
    var rad = entity.rotation / 180 * Math.PI

    //can't use body.toWorldVector() because the rotation is not instantly
    var localForward = Qt.point(power * Math.cos(rad), power * Math.sin(rad))
    boxCollider.body.applyLinearImpulse(localForward, boxCollider.body.getWorldCenter())
  }
}
