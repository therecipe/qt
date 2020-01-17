import Felgo 3.0
import QtQuick 2.0
import "../common"
import "../entities"

SceneBase {
  id:gameScene

  // the "logical size" - the scene content is auto-scaled to match the GameWindow size
  width: 320
  height: 480

  // background image is based on game window (bigger as logical scene)
  BackgroundImage {
      anchors.centerIn: gameScene.gameWindowAnchorItem
      source: "../../assets/CarouselBackground.jpg"
  }

  // background music plays automatically
  BackgroundMusic {
    id: backgroundMusic
    source: "../../assets/sound/Carousel.mp3"
  }

  // setup scene layout
  property int lanes: 4 // number of lanes (rides)
  property double laneWidth: gameScene.width / gameScene.lanes // pixelWidth of a lane
  property int startLane: 0 // start lane of player (will be randomly set on game start)
  property double playerY: 400 // y position of player in scene

  // make important objects accessible for other entities
  property alias logic: logic
  property alias ride1: horse
  property alias ride2: boat
  property alias ride3: bike
  property alias ride4: car

  // player score
  property int points: 0

  // add game logic element
  Logic {
      id: logic
  }

  // initialize a new game
  function startNewGame() {
    gameScene.points = 0

    // initial speed multiplier
    gameScene.logic.speedUp = 1.0

    // reset bullet and coin timeouts
    gameScene.logic.bulletTimeout = 4000
    gameScene.logic.coinTimeout = 2000

    // set new random starting lane
    player.stopMovement()
    gameScene.startLane = Math.random() * gameScene.lanes + 1
    player.x = (gameScene.startLane * gameScene.laneWidth) - (gameScene.laneWidth / 2)
    player.startUpDownMovement()
  }


  // basic physics
  PhysicsWorld {
      gravity.y: 0 // we use fixed animations instead of gravity
      z: 5 // draw the debugDraw on top of the entities

      // these are performance settings to avoid boxes colliding too far together
      // set them as low as possible so it still looks good
      updatesPerSecondForPhysics: 60
      velocityIterations: 5
      positionIterations: 5
      // set this to true to see the debug draw of the physics system
      // this displays all bodies, joints and forces which is great for debugging
      debugDrawVisible: false
  }

  // display player score
  Text {
      text: gameScene.points
      color: "#e52222"
      y: 36
      z: 1
      font.pixelSize: 12
      font.family: "Arial"
      font.weight: Font.Bold
      width: gameScene.width
      horizontalAlignment: Text.AlignHCenter
  }

  // player character
  Player {
      id: player
      z: 10
      x: (gameScene.startLane * gameScene.laneWidth) - (gameScene.laneWidth / 2)
      y: gameScene.playerY
  }

  // player rides
  // ride1 - horse
  RideFront {
     id: horse
     x: -20; y: 397
     minY: 397 - (Math.random() * 10 + 5); maxY: 397 + Math.random() * 10 + 5
     width: 92.5; height: 128
     image.source: "../../assets/ride1_f.png"
  }
  RideBack {
     frontRide: horse
     image.source: "../../assets/ride1_b.png"
  }
  // ride2 - boat
  RideFront {
     id: boat
     x: 73.5; y: 397
     minY: 397 - (Math.random() * 10 + 5); maxY: 397 + Math.random() * 10 + 5
     width: 94.5; height: 128
     image.source: "../../assets/ride2_f.png"
  }
  RideBack {
     frontRide: boat
     image.source: "../../assets/ride2_b.png"
  }
  // ride3 - bike
  RideFront {
     id: bike
     x: 167; y: 397
     minY: 397 - (Math.random() * 10 + 5); maxY: 397 + Math.random() * 10 + 5
     width: 67.5; height: 128
     image.source: "../../assets/ride3_f.png"
  }
  RideBack {
     frontRide: bike
     image.source: "../../assets/ride3_b.png"
  }
  // ride4 - car
  RideFront {
     id: car
     x: 234.5; y: 397
     minY: 397 - (Math.random() * 10 + 5); maxY: 397 + Math.random() * 10 + 5
     width: 105.5; height: 128
     image.source: "../../assets/ride4_f.png"
  }
  RideBack {
     frontRide: car
     image.source: "../../assets/ride4_b.png"
  }

  // enemies are grouped as a single enemies entity
  Enemies { id: enemies; y: 61 }

  // invisible floor
  Floor {
      z: 100
      anchors.bottom: gameScene.gameWindowAnchorItem.bottom
  }

  // handle user input (check for swipe to left/right)
  MouseArea {
    anchors.fill: gameScene.gameWindowAnchorItem // check full game window
    property bool touch: false // true when user touches screen
    property int firstX: 0 // x position of swipe start point

    // recognize start of swipe on press
    onPressed: {
      if(touch == false)
          firstX = mouseX
      touch = true
    }

    // recognize end of swipe on release
    onReleased: {
      if(touch == true)
          checkSwipe(15)
      touch = false
    }

    // also recognize swipe if mouse moved a long distance
    onPositionChanged: {
        if(touch)
          checkSwipe(50)
    }

    // move player based on swipe (left or right)
    function checkSwipe(minDistance) {
        var distance = mouseX - firstX
        if(Math.abs(distance) > minDistance) {
            if(distance > 0)
                player.moveRight()
            else
                player.moveLeft()
          touch = false
        }
    }
  }

}
