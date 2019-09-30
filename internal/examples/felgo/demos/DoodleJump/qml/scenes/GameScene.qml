import QtQuick 2.0
import Felgo 3.0
import QtSensors 5.5
import "../"

SceneBase {
  id: gameScene

  // actual scene size
  width: 320
  height: 480

  state: "start"

  property int score: 0

  signal menuScenePressed

  // increment achievement whenever game scene is created (once per app start)
  Component.onCompleted: gameNetwork.incrementAchievement("5opens")

  // background image
  Image {
    anchors.fill: parent.gameWindowAnchorItem
    source: "../../assets/background.png"
  }

  // key input will be handled by the controller in our frog-entity
  Keys.forwardTo: frog.controller

  // accelerometer can be used to react to tilting the phone
  Accelerometer {
    id: accelerometer
    active: true
  }

  // add physics world to use gravity and collision detection
  PhysicsWorld {
    debugDrawVisible: false // turn it on for debugging
    updatesPerSecondForPhysics: 60
    gravity.y: 20 // how much gravity do you want?
  }

  // the repeater adds ten platforms to the scene
  Repeater {
    model: 10 // every platorm gets recycled so we define only ten of them
    Platform {
      x: utils.generateRandomValueBetween(0, gameScene.width) // random value
      y: gameScene.height / 10 * index // distribute the platforms across the screen
    }
  }

  // this platform is placed directly under the frog for the start
  Platform {
    x: 120
    y: 300
  }

  // the frog entity (player character)
  Frog {
    id: frog
    x: gameScene.width / 2 // place the frog in the horizontal center
    y: 220
  }

  // border at the bottom of the screen, used to check game-over
  Border {
    id: border
    x: -gameScene.width*2
    y: gameScene.height-10 // subtract a small value to make the border just visible in our scene
  }

  // show current player score
  Image {
    id: scoreCounter
    source: "../../assets/scoreCounter.png"
    height: 80
    x: -15
    y: -15
    // text component to show the score
    Text {
      id: scoreText
      anchors.centerIn: parent
      color: "white"
      font.pixelSize: 32
      text: score
    }
  }

  // start the game when user touches the screen
  MouseArea {
    id: mouseArea
    anchors.fill: gameScene.gameWindowAnchorItem
    onClicked: {
      if(gameScene.state === "start") { // if the game is ready and the screen is clicked we start the game
        gameScene.state = "playing"
      }
      if(gameScene.state === "gameOver") // if the frog is dead and the screen is clicked we restart the game
      {
        gameScene.state = "start"
      }
    }
  }

  // show info text depending on state (gameover, start)
  Image {
    id: infoText
    anchors.centerIn: parent
    source: gameScene.state == "gameOver" ? "../../assets/gameOverText.png" : "../../assets/clickToPlayText.png"
    visible: gameScene.state !== "playing"
  }

  // options button to jump back to menu
  Image {
    id: menuButton
    source: "../../assets/optionsButton.png"
    x: gameScene.width - 96
    y: -40
    scale: 0.5

    MouseArea {
      id: menuButtonMouseArea
      anchors.fill: parent
      onClicked: {
        menuScenePressed()

        // reset the gameScene
        frog.die()
        gameScene.state = "start"
      }
    }
  }
}
