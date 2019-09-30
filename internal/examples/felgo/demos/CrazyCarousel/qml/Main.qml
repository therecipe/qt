import Felgo 3.0
import QtQuick 2.0
import "scenes"

GameWindow {
    id: gameWindow

    // start size of the window
    // usually set to resolution of main target device - this resolution is for iPhone 4 & iPhone 4S
    screenWidth: 320
    screenHeight: 480


    // the entity manager allows dynamic creation of entities within the game scene (e.g. bullets and coins)
    EntityManager {
        id: entityManager
        entityContainer: gameScene
    }

    // title scene
    TitleScene {
      id: titleScene

      // the title scene is our start scene, so if back is pressed there (on android) we ask the user if he wants to quit the application
      onBackButtonPressed: {
          nativeUtils.displayMessageBox("Really quit the game?", "", 2)
      }
      // listen to the return value of the MessageBox
      Connections {
          target: nativeUtils
          onMessageBoxFinished: {
              // only quit, if the activeScene is titleScene - the messageBox might also get opened from other scenes in your code
              if(accepted && gameWindow.activeScene === titleScene)
                  Qt.quit()
          }
      }
    }

    // game scene to play
    GameScene {
      id: gameScene

      // the game scene is our main scene of the game, if back is pressed there we return to the title scene
      onBackButtonPressed: {
        gameWindow.state = "title"
      }
    }

    // game over scene
    GameOverScene {
        id: gameOverScene

        // the game over scene is very similar to the title scene, if back is pressed there we ask the user if he wants to quit the application
        onBackButtonPressed: {
            nativeUtils.displayMessageBox("Really quit the game?", "", 2)
        }
        // listen to the return value of the MessageBox
        Connections {
            target: nativeUtils
            onMessageBoxFinished: {
                // only quit, if the activeScene is gameOverScene - the messageBox might also get opened from other scenes in your code
                if(accepted && gameWindow.activeScene === gameOverScene)
                    Qt.quit()
            }
        }
    }

    // the default state is title -> now our default scene is the titleScene
    state: "title"

    // state machine, takes care of reversing the PropertyChanges when switching the state, e.g. changing the opacity back to 0
    states: [
        State {
            name: "title"
            PropertyChanges {target: titleScene; opacity: 1}
            PropertyChanges {target: gameWindow; activeScene: titleScene}
        },
        State {
            name: "game"
            PropertyChanges {target: gameScene; opacity: 1}
            PropertyChanges {target: gameWindow; activeScene: gameScene}
        },
        State {
            name: "gameover"
            PropertyChanges {target: gameOverScene; opacity: 1}
            PropertyChanges {target: gameWindow; activeScene: gameOverScene}
        }

    ]
}

