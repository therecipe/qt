import Felgo 3.0
import QtQuick 2.0
import "scenes"

GameWindow {
  id: gameWindow

  // You get free licenseKeys from https://felgo.com/licenseKey
  // With a licenseKey you can:
  //  * Publish your games & apps for the app stores
  //  * Remove the Felgo Splash Screen or set a custom one (available with the Pro Licenses)
  //  * Add plugins to monetize, analyze & improve your apps (available with the Pro Licenses)
  //licenseKey: "<generate one from https://felgo.com/licenseKey>"

  // window size
  screenWidth: 640
  screenHeight: 960

  activeScene: gameScene

  EntityManager {
    id: entityManager
    entityContainer: gameScene
  }

  // configure google analytics plugin
  GoogleAnalytics {
    id: ga

    // property tracking ID from Google Analytics dashboard
    propertyId: "UA-67377753-2"
  }

  // set up game network and achievements
  FelgoGameNetwork {
    id: gameNetwork
    gameId: 173 // put your gameId here
    secret: "doodlefrogsecret12345" // put your game secret here
    gameNetworkView: frogNetworkView

    achievements: [
      Achievement {
        key: "5opens"
        name: "Game Opener"
        target: 5
        points: 10
        description: "Open this game 5 times"
      },

      Achievement {
        key: "die100"
        name: "Y U DO DIS?"
        iconSource: "../assets/achievementImage.png"
        target: 100
        description: "Die 100 times"
      }
    ]
  }

  // scene for the actual game
  GameScene {
    id: gameScene
    onMenuScenePressed: {
      gameWindow.state = "menu"
      ga.logEvent("Click", "MenuScene")
    }
  }

  // the menu scene of the game
  MenuScene {
    id: menuScene
    onGameScenePressed: {
      gameWindow.state = "game"
      ga.logEvent("Click", "GameScene")
    }

    // the game network view shows the leaderboard and achievements
    GameNetworkView {
      id: frogNetworkView
      visible: false
      anchors.fill: parent.gameWindowAnchorItem

      onShowCalled: {
        frogNetworkView.visible = true
      }

      onBackClicked: {
        frogNetworkView.visible = false
      }
    }
  }

  // default state is menu -> default(starting) scene is menuScene
  state: "menu"

  // state machine, takes care reversing the PropertyChanges when changing the state like changing the opacity back to 0
  states: [
    State {
      name: "menu"
      PropertyChanges {target: menuScene; opacity: 1}
      PropertyChanges {target: gameWindow; activeScene: menuScene}
    },
    State {
      name: "game"
      PropertyChanges {target: gameScene; opacity: 1}
      PropertyChanges {target: gameWindow; activeScene: gameScene}
    }
  ]
}
