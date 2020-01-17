import QtQuick 2.0
import Felgo 3.0

// is shown at game start and shows the maximum highscore and a button for starting the game
SceneBase {
  id: mainScene

  // only if this is set to true, the exit dialog should quit the app
  property bool exitDialogShown: false
  property bool vplayLinkShown: false

  onBackButtonPressed: {
    exitDialogShown = true
    // instead of immediately shutting down the app, ask the user if he really wants to exit the app with a native dialog
    nativeUtils.displayMessageBox(qsTr("Really quit the game?"), "", 2);
  }

  Connections {
    // nativeUtils should only be connected, when this is the active scene
      target: activeScene === mainScene ? nativeUtils : null
      onMessageBoxFinished: {
        console.debug("the user confirmed the Ok/Cancel dialog with:", accepted)
        if(accepted && exitDialogShown) {
          Qt.quit()
        } else if(accepted && vplayLinkShown) {
          flurry.logEvent("MainScene.Show.VPlayWeb")
          nativeUtils.openUrl("https://felgo.com/showcases/?utm_medium=game&utm_source=chickenoutbreak&utm_campaign=chickenoutbreak#chicken_outbreak");
        }


        // set it to false again
        exitDialogShown = false
        vplayLinkShown = false
      }
  }

  MultiResolutionImage {
    source: "../assets/img/mainMenuBackground.png"
    anchors.centerIn: parent
  }

  Column {
    spacing: 15
    anchors.horizontalCenter: parent.horizontalCenter
    y: 5

    MenuText {
      text: qsTr("Chicken Outbreak")
      font.pixelSize: 35
    }

    MenuText {
      text: qsTr("Highscore: ") + gameNetwork.userHighscoreForCurrentActiveLeaderboard
    }

    Item {
      width: 1
      height: 0
    }

    MenuButton {
      text: qsTr("Play")
      onClicked: window.state = "game"
      textSize: 30
    }

    Item {
      width: 1
      height: 25
    }    

    MenuButton {
      text: qsTr("Highscores")

      width: 170 * 0.8
      height: 60 * 0.8

      onClicked: {
        window.state = "gameNetwork"
        // this would not need to be called, then no request is sent when showing it every time
        gameNetwork.showLeaderboard()
      }
    }

//    MenuButton {
//      text: qsTr("Game Center")

//      width: 170 * 0.8
//      height: 60 * 0.8

//      // this button opens the gamecenter leaderboards - only show it if the gamecenter is available (so iOS only)
//      // still show the Game Center button, as on iOS it is also very popular!
//      // on iOS VPGN and GameCenter both can be used together
//      visible: gameCenter.authenticated

//      onClicked: {
//          flurry.logEvent("GameCenter.Show")
//          gameCenter.showLeaderboard();
//      }
//    }

    MenuButton {
      text: qsTr("Credits")

      width: 170 * 0.8
      height: 60 * 0.8

      onClicked: window.state = "credits";
    }

    MenuButton {
      text: settings.soundEnabled ? qsTr("Sound on") : qsTr("Sound off")

      width: 170 * 0.8
      height: 60 * 0.8

      onClicked: {
        settings.soundEnabled = !settings.soundEnabled
        // also set the musicEnabled to the same value as soundEnabled
        settings.musicEnabled = settings.soundEnabled
      }


      // this button should only be displayed on Symbian & Meego, because on the other platforms the volume hardware keys work; but on Sym & Meego the volume cant be adjusted as the hardware volume keys are not working
      // also, display it when in debug build for quick toggling the sound
      // always display it, because it is easier to turn sound on and off with a button than pressing the hardware keys often
      //visible: system.debugBuild
    }
  }

  Image {
   id: logo
    anchors.right: mainScene.gameWindowAnchorItem.right
    anchors.rightMargin: 10
    anchors.bottom: mainScene.gameWindowAnchorItem.bottom
    anchors.bottomMargin: 10
    source: "../assets/img/felgo-logo.png"
    // the image size is bigger (for hd2 image), so only a single image no multiresimage can be used
    // this scene is not performance sensitive anyway!
    fillMode: Image.PreserveAspectFit
    height: 55

    MouseArea {
      anchors.fill: parent
      onClicked: {
        vplayLinkShown = true
        flurry.logEvent("MainScene.ShowDialog.VPlayWeb")
        nativeUtils.displayMessageBox(qsTr("Felgo"), qsTr("This game is built with Felgo. The source code is available in the free Felgo SDK - so you can build your own Chicken Outbreak in minutes! Visit Felgo.net now?"), 2)
      }
    }

    SequentialAnimation {
      running: true
      loops: -1
      NumberAnimation { target: logo; property: "opacity"; to: 0.1; duration: 1200 }
      NumberAnimation { target: logo; property: "opacity"; to: 1; duration: 1200 }
    }
  }

  // this allows navigation through key presses
  Keys.onReturnPressed: {
    window.state = "game"
  }
}
