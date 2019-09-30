import QtQuick 2.0

import Felgo 3.0


// the 2 BorderRegion entities (one on top and one on bottom of the screen) are not visible because they are offscreen
// if the topRegion collides with a roost, a coin or a window, they get removed and used for pooling
// if the topRegion collides with the player, that means the game is lost as the player got out of the scene on the top
// the bottomRegion is only for detecting collision with the player, so if he falls through and cant stand on a roost, the game is lost
EntityBase {
  id: entity
  entityType: "borderRegion"
  // either is topRegion or BorderRegion - topRegion removes the collided block or coin entity at collision, whereas BorderRegion doesn't!
  variationType: "topRegion"

  // will be handled in Level when the region collides with the player, which means the game is lost
  signal playerCollision

  // these only exist once and should not be pooled
  preventFromRemovalFromEntityManager: true

  BoxCollider {
    id: boxCollider
    // the body must be of type dynamic, otherwise no collisions between the other static bodies (roosts & coins) would be possible!
    bodyType: Body.Dynamic
    // this is required, because the position should not be modified from the physics system, but from the QML positioning!
    // if only sensor would be set to true it would not be enough, because then the body would fall down based on gravity!
    collisionTestingOnlyMode: true

    fixture.onBeginContact: {

      var fixture = other;
      var body = other.getBody();
      var collidedEntity = body.target;
      var collidedEntityType = collidedEntity.entityType;

      console.debug("BorderRegion: collided with entity type:", collidedEntityType);

      if(collidedEntityType === "player")
        // game is lost, player reached top or bottom of scene
        entity.playerCollision();
      else {
        // don't remove the collided entity if this is the bottomRegion, because entities are created outside the scene and thus they would constantly collide with the bottomRegion
        if(variationType === "bottomRegion")
          return;

        // do not remove this entity, as it is the one where the player should stand when it gets deleted!
        if(collidedEntity.entityId === "playerInitialBlock")
          return;

        // collision with another entity, so either a block gotten out of range, or a coin, which should be removed now (it uses pooling though!)
        collidedEntity.removeEntity();
      }
    }
  }
}
