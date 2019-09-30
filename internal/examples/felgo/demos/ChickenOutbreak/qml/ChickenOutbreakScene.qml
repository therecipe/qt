import QtQuick 2.0
import Felgo 3.0
import "entities"

SceneBase {
  id: scene

  property alias level: level
  property alias player: level.player
  property alias entityContainer: level

  // make 1 grid (so 1 block and the player size) 48 logical px - a roost has this size, so 320/48= 6.6 can be displayed in one scene
  gridSize: 48

  // place it on bottom, because otherwise it would be unfair compared to different devices because the player would see more to the bottom playfield!
  sceneAlignmentY: "bottom"

  onBackButtonPressed: {
    // it is important to call stopGame() here, because otherwise the entities would not be deleted!
    level.stopGame();
    window.state = "main"
  }

  // allows collision detection between the entities, so if the player collides with roosts or the corn, and for removal of the out-of-scene entities
  PhysicsWorld {
    // this id gets used in Level, to test if a new roost or window can be built on the new position or if there is an entity already
    id: physicsWorld
    // this puts it on top of all other items with lower z for the physics debug renderer
    z: 10

    // for physics-based games, this should be set to 60!
    updatesPerSecondForPhysics: 60
    // make objects fall faster by increasing gravity
    gravity.y: 60

    // this should be increased so it looks good, by default it is set to 1 to save performance
    velocityIterations: 5
    positionIterations: 5

    // by default this would be enabled in the debug version, but disable it as most of the time we are not interested in it, only if we debug physics issues
    debugDrawVisible: false
  }

  // handles the repositioning of the background, if they are getting out of the scene
  // internally, 4 images are created below each other so it appears to the user as being one continuous background
  ParallaxScrollingBackground {
    id: levelBackground
//    property int initialOffsetx: -(levelBackground.image.width-scene.width)/2
//    property int initialOffsety: -(levelBackground.image.height-scene.height)/2
//    x: initialOffsetx
//    y: initialOffsety-level.y
    anchors.horizontalCenter: parent.horizontalCenter
    y: parent.height-parent.gameWindowAnchorItem.height
    sourceImage: "../assets/img/background-wood2.png"
    // do not mirror it vertically, because the image is prepared to match on the top and the bottom
    mirrorSecondImage: false
    movementVelocity: Qt.point(0, level.levelMovementAnimation.velocity)
    running: level.levelMovementAnimation.running // start non-running, gets set to true in startGame
  }

  Level {
    id: level

    onGameLost: {
      console.debug("ChickenOutbreakScene: gameLost()");
      level.stopGame();
      // lastScore is used in GameOverScreen to check if a new highscore is reached
      lastScore = player.totalScore;
      // Increment the player's death count
      player.deaths++;
      // change to state gameOver showing the GameOverScene - in the SceneBase a crossfade effect is implemented
      window.state = "gameOver"
    }
  }

  // this allows usage of the left and right keys on desktop systems or mobiles with physical keyboards
  Keys.forwardTo: player.controller

  // input graphics - the handling of the input is done below in the MouseArea
  Image {
    source: "../assets/img/arrow-left.png"
    opacity: 0.5
    width: 48
    height: 48
    anchors {
      left: scene.gameWindowAnchorItem.left
      bottom: scene.gameWindowAnchorItem.bottom
      leftMargin: 10
      bottomMargin: 10
    }
  }
  Image {
    source: "../assets/img/arrow-left.png"
    opacity: 0.5
    width: 48
    height: 48
    mirror: true
    anchors {
      right: scene.gameWindowAnchorItem.right
      bottom: scene.gameWindowAnchorItem.bottom
      rightMargin: 10
      bottomMargin: 10
    }
  }
  MouseArea {
    // use the full window as control item, press anywhere on the left half for steering left, on the right half for steering right
    anchors.fill: scene.gameWindowAnchorItem
    onPressed: {
      console.debug("onPressed, mouseX", mouseX)
      if(mouseX > scene.gameWindowAnchorItem.width/2)
        player.controller.xAxis = 1;
      else
        player.controller.xAxis = -1;
    }
    onPositionChanged: {
      if(mouseX > scene.gameWindowAnchorItem.width/2)
        player.controller.xAxis = 1;
      else
        player.controller.xAxis = -1;
    }
    onReleased: player.controller.xAxis = 0
  }

  Text {
    x: 5
    // place it on top of the window, not on top of the logical scene
    anchors.top: scene.gameWindowAnchorItem.top
    anchors.topMargin: 5

    text: qsTr("Score: ") + player.totalScore
    font.family: fontHUD.name
    font.pixelSize: 22
    color: "white"
  }

  // gets called by ChickenOutbreakMain when this scene gets active
  // starts a game again - stopGame() was called before so it is save to call that here
  function enterScene() {
    level.startGame();
  }


  // ------------------- debugging-only code ------------------- //

  // this button is only for testing and only visible in debug mode
  // it shows the ingameMenu to quickly restart the game
  SimpleButton {
    id: hud
    width: 64
    height: 64
    // place it on top of the window, not on top of the logical scene
    anchors.top: scene.gameWindowAnchorItem.top
    anchors.right: scene.gameWindowAnchorItem.right
    visible: system.debugBuild // only display in debug mode - the menu button for ingame testing should not be visible in retail builds for the store (and also not in release builds)
    text: "Menu"
    onClicked: {
      console.debug("Menu button clicked")
      // this activates the ingameMenu state, which will show the IngameMenu item
      scene.state = "ingameMenu"
    }
  }
  // this gets only displayed when the menu button is pressed, which is only allowed in debug builds
  IngameMenu {
    id: ingameMenu
    // in the default state, this should be invisible!
    visible: false
    anchors.centerIn: parent
  }
  // the ingameMenue state is only enabled when in debug mode and when the menu button is pressed
  onStateChanged: console.debug("Scene.state changed to", state)
  states: [
    State {
      name: ""
      StateChangeScript {
        script: {
          console.debug("scene: entered state ''")
          level.resumeGame();
        }
      }
    },
    State {
      name: "ingameMenu"
      PropertyChanges { target: ingameMenu; visible: true}
      StateChangeScript {
        script: {
          console.debug("scene: entered state 'ingameMenu'")
          level.pauseGame();
        }
      }
    }

  ]
}
