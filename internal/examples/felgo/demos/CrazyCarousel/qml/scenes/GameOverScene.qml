import Felgo 3.0
import QtQuick 2.0

TitleScene {
    id: gameOverScene

    // switch background image
    backgroundImage.source: "../../assets/CarouselGameOver.jpg"

    // display the points
    Text {
        text: gameScene.points
        color: "white"
        x: 188
        y: 300
        z: 1
        font.pixelSize: 34
        font.family: "Arial"
        width: 108
        horizontalAlignment: Text.AlignHCenter
    }

}
