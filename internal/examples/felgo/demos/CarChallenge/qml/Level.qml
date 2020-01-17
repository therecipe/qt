import QtQuick 2.0

import Felgo 3.0
import "entities"

Item {
    id: level
    width: parent.width
    height: parent.width

    // use this to insert the input action (which car should fire and which to steer) to the right car
    property alias player_red: car_red
    property alias player_blue: car_blue

    // background of the level
    Image {
        id: playgroundImage
        source: "../assets/img/playground.jpg"
        anchors.fill: parent
    }

    Car {
        id: car_red
        objectName: "car_red"
        variationType: "carRed"
        x: 90
        y: 200
        // rotation in degrees clockwise
        rotation: 0

    }

    Car {
        id: car_blue
        objectName: "car_blue"
        variationType: "carBlue"
        x: 390
        y: 400
        // rotation in degrees clockwise
        rotation: 225

        image.source: "../assets/img/car_blue.png"

        inputActionsToKeyCode: {
            "up": Qt.Key_W,
            "down": Qt.Key_S,
            "left": Qt.Key_A,
            "right": Qt.Key_D,
            "fire": Qt.Key_Space
        }
    }

    Wall {
        id: border_bottom

        height: 20
        anchors {
            left: parent.left
            right: parent.right
            bottom: parent.bottom
        }
    }

    Wall {
        id: border_top

        height: 20
        anchors {
            top: parent.top
            left: parent.left
            right: parent.right
        }
    }

    Wall {
        id: border_left
        width: 20
        anchors {
            top: parent.top
            bottom: parent.bottom
            left: parent.left
        }
    }

    Wall {
        id: border_right
        width: 20
        anchors {
            top: parent.top
            bottom: parent.bottom
            right: parent.right
        }
    }
}
