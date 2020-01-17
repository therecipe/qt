import Felgo 3.0
import QtQuick 2.0


App {
  id: app

  // You get free licenseKeys from https://felgo.com/licenseKey
  // With a licenseKey you can:
  //  * Publish your games & apps for the app stores
  //  * Remove the Felgo Splash Screen or set a custom one (available with the Pro Licenses)
  //  * Add plugins to monetize, analyze & improve your apps (available with the Pro Licenses)
  //licenseKey: "<generate one from https://felgo.com/licenseKey>"

  FontLoader {
    id: latoFont
    source: "../assets/fonts/Lato-Light.ttf"
  }

  onInitTheme: {
    Theme.colors.tintColor = Qt.rgba(0, 147/256.0, 201/256.0, 1)
    Theme.normalFont = latoFont
  }

  MapsMainPage { }
}
