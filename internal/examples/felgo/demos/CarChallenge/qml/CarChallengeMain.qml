import QtQuick 2.0
import Felgo 3.0

GameWindow {
    id: window
    screenWidth: 960
    screenHeight: 640

    activeScene: scene

    // You get free licenseKeys from https://felgo.com/licenseKey
    // With a licenseKey you can:
    //  * Publish your games & apps for the app stores
    //  * Remove the Felgo Splash Screen or set a custom one (available with the Pro Licenses)
    //  * Add plugins to monetize, analyze & improve your apps (available with the Pro Licenses)
    //licenseKey: "<generate one from https://felgo.com/licenseKey>"

    // for creating and destroying entities at runtime (rockets)
    EntityManager {
        id: entityManager
        entityContainer: level
    }

    Scene {
        id: scene
        width: 640
        height: 640

        // background image
        Image {
            id: background
            source: "../assets/img/asphalt_background.png"
            // since we are using a simple Image element to define our background, we stretch it to avoid any kind of black borders on any device
            width: parent.width*1.8
            height: parent.height*1.8
            anchors.centerIn: parent
        }

        // use a physics world because we need collision detection
        PhysicsWorld {
            id: world
            updatesPerSecondForPhysics: 60
        }

        Level {
            // this gets accessed by its id from JoystickControllerHUD below
            id: level

        }

        focus: true
        // forward the input keys to both players
        Keys.forwardTo: [level.player_red.controller, level.player_blue.controller]
    }


    JoystickControllerHUD {
        anchors.bottom: parent.bottom

        // the joystick width is the space from the left to the start of the logical scene, so the radius is its half
        joystickRadius: scene.x/2


        // this allows setting custom images for the JoystickControllerHUD component
        source: "../assets/img/joystick_background.png"
        thumbSource: "../assets/img/joystick_thumb.png"


        // this is the mapping between the output of the JoystickControllerHUD to the input of the player's TwoAxisController
        // this is a performance improvement, to not have to bind every time the position changes
        property variant playerTwoxisController: level.player_red.getComponent("TwoAxisController")
        onControllerXPositionChanged: playerTwoxisController.xAxis = controllerXPosition;
        onControllerYPositionChanged: playerTwoxisController.yAxis = controllerYPosition;
    }
}

