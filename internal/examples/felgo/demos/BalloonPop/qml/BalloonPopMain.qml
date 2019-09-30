import Felgo 3.0
import QtQuick 2.0

GameWindow {
  id: gameWindow

  // You get free licenseKeys from https://felgo.com/licenseKey
  // With a licenseKey you can:
  //  * Publish your games & apps for the app stores
  //  * Remove the Felgo Splash Screen or set a custom one (available with the Pro Licenses)
  //  * Add plugins to monetize, analyze & improve your apps (available with the Pro Licenses)
  //licenseKey: "<generate one from https://felgo.com/licenseKey>"

  // create a licenseKey to hide the splash screen
  onSplashScreenFinished: { ballonScene.start() }

  // our scene that contains the game
  BalloonScene {
    id: ballonScene
  }
}
