import QtQuick 2.0
import Felgo 3.0

EntityBase {
  entityType: "balloon"
  width: sprite.width
  height: sprite.height

  MultiResolutionImage {
    id: sprite
    source: "../../assets/img/balloon.png"
  }

  CircleCollider {
    radius: sprite.width/2
    // restitution is bounciness, balloons are quite bouncy
    fixture.restitution: 0.5
  }

  MouseArea {
    anchors.fill: sprite
    onPressed: {
      // if you touch a balloon and the game is running, it will pop
      if(balloonScene.gameRunning) {
        balloonScene.balloons--
        popSound.play()
        removeEntity()
      }
    }
  }

  // gives the balloon a random position when created
  Component.onCompleted: {
    x = utils.generateRandomValueBetween(0,balloonScene.width-sprite.width)
    y = balloonScene.height
  }
}
