import Felgo 3.0
import QtQuick 2.0
import "../common"

SceneBase {
    id:titleScene

    // the "logical size" - the scene content is auto-scaled to match the GameWindow size
    width: 320
    height: 480

    // property alias to allow switching the background later
    property alias backgroundImage: background

    // background image is based on game window (bigger as logical scene)
    BackgroundImage {
        id: background
        anchors.centerIn: titleScene.gameWindowAnchorItem
        source: "../../assets/CarouselTitle.jpg"
    }

    // play button to start game
    Text {
        id: text
        text: "play!"
        color: "black"
        x: 60
        y: 380
        z: 1
        font.pixelSize:  20
        font.family: "Arial"

        // start new game when start is clicked
        MouseArea {
            width: parent.width
            height: parent.height

            onClicked: {
                clickSound.play()
                gameScene.startNewGame()
                gameWindow.state = "game"
            }
        }
    }

    // alternate text color of play button over time
    Timer {
        id: colorTimer
        interval: 1000 // 1 sec
        repeat: true
        running: true
        onTriggered: {
            if(text.color == "#000000")
                text.color = "yellow"
            else
                text.color = "black"
        }
    }

    // sound effect when clicking the play button
    SoundEffect {
        id: clickSound
        source: "../../assets/sound/Click.wav"
    }
}
