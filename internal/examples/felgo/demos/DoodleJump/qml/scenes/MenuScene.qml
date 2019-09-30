import Felgo 3.0
import QtQuick 2.0
import "../"

SceneBase {
  id:menuScene

  // signal indicating that the gameScene should be displayed
  signal gameScenePressed

  // background image
  Image {
    anchors.fill: menuScene.gameWindowAnchorItem
    source: "../../assets/background.png"
  }

  // column aligns its child components within a column
  Column {
    anchors.centerIn: parent
    spacing: 20

    // play button to start game
    Rectangle {
      width: 150
      height: 50
      color: "orange"
      Image {
        id: gameSceneButton
        source: "../../assets/playButton.png"
        anchors.centerIn: parent
      }

      MouseArea {
        id: gameSceneMouseArea
        anchors.fill: parent
        onClicked: gameScenePressed()
      }
    }

    // score button to open leaderboard
    Rectangle {
      width: 150
      height: 50
      color: "orange"
      Image {
        id: scoreSceneButton
        source: "../../assets/scoreButton.png"
        anchors.centerIn: parent
      }
      MouseArea {
        id: scoreSceneMouseArea
        anchors.fill: parent
        onClicked: frogNetworkView.visible = true
      }
    }

    // row component aligns its children within a row
    Row {
      // button to open chartboost interstitial
      Image {
        id: chartboostButton
        source: "../../assets/chartboostButton.png"

        MouseArea {
          anchors.fill: parent
          onClicked: {
            if(system.desktopPlatform)
              nativeUtils.displayMessageBox("Ads are only available on iOS and Android. To test the Chartboost plugin, open the Doodle Jump game project and deploy the game to your iOS or Android device.")
            else
              chartboost.showInterstitial()
          }
        }
      }

      // button to open admob interstitial
      Image {
        id: admobButton
        source: "../../assets/admobButton.png"

        MouseArea {
          anchors.fill: parent
          onClicked: {
            if(system.desktopPlatform)
              nativeUtils.displayMessageBox("Ads are only available on iOS and Android. To test the AdMob plugin, open the Doodle Jump game project and deploy the game to your iOS or Android device.")
            else
              interstitial.loadInterstitial()
          }
        }
      }
    }
  }

  // configure the admob banner
  AdMobBanner {
    //adUnitId: "ca-app-pub-5893836292676877/1041507143"
    adUnitId: "ca-app-pub-5893836292676877/7730182343"

    banner: system.desktopPlatform ? 0 : AdMobBanner.Smart // theAdMobBanner enum is only available on iOS & Android; to prevent a warning, only set it on these platforms to a Smart Banner

    anchors.horizontalCenter: parent.Center
    visible: !frogNetworkView.visible

    testDeviceIds: ["7AA2ECE6469E41E0C8E5ABAFCC7A0BB9", "3385843ff1e43633521e3750a6d57fed", "28CA0A7F16015163A1C70EA42709318A"]
  }

  // configure the admob interstitial
  AdMobInterstitial {
    id: interstitial

    //adUnitId: "ca-app-pub-5893836292676877/2518240347"
    adUnitId: "ca-app-pub-5893836292676877/9206915541"

    onInterstitialReceived: {
      showInterstitialIfLoaded()
    }

    onInterstitialFailedToReceive: {
      console.debug("Interstitial not loaded")
    }

    testDeviceIds: ["7AA2ECE6469E41E0C8E5ABAFCC7A0BB9", "3385843ff1e43633521e3750a6d57fed", "28CA0A7F16015163A1C70EA42709318A"]
  }

  // configure the chartboost interstitial
  Chartboost {
    id: chartboost

    appId: Qt.platform.os === System.IOS ? "55f2a8145b145373586f4b16" : "55f2a8155b145373586f4b18"
    appSignature: Qt.platform.os === System.IOS ? "8b2d667bac3e06e60c08c06e4e63898fcad52cfd" : "f8b21bf14f5b85cc62612765300ce4c205664b0a"

    // Do not use reward videos in this example
    shouldDisplayRewardedVideo: false

    // allows to show interstitial also at first app startup.
    // see http://www.felgo.com/doc/felgo-chartboost/#shouldRequestInterstitialsInFirstSession-prop
    shouldRequestInterstitialsInFirstSession: true

    onInterstitialCached: {
      console.debug("InterstitialCached at location:", location)
    }

    onInterstitialFailedToLoad: {
      console.debug("InterstitialFailedToLoad at location:", location, "error:", error)
    }
  }
}
