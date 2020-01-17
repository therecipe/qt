import QtQuick 2.7
import Felgo 3.0
import "pages"
import "common"
import "model"
import "logic"

Rectangle {
  id: loaderItem
  anchors.fill: parent
  color: Theme.backgroundColor

  // make navigation public (required for Felgo app demo launcher)
  property Navigation navigation: mainLoader.item && mainLoader.item.navigation

   // if notification was opened, go to inbox after loading
  property bool openInboxAfterLoading: false
  onNavigationChanged: {
    if(navigation && openInboxAfterLoading)
      mainLoader.item.openInbox()
  }

  property SocialView socialViewItem: mainLoader.item && mainLoader.item.socialViewItem || null

  // check for available app updates
  Loader {
    id: versionChecker
    anchors.fill: parent
    visible: false
    asynchronous: true
    property string updateCheckUrl: system.publishBuild ? "https://felgo.com/qml-sources/qtws2017/QtWSVersionCheck.qml" : "https://felgo.com/qml-sources/qtws2017/QtWSVersionCheck-test.qml"
    source: !system.desktopPlatform ? updateCheckUrl : ""
    onLoaded: versionChecker.visible = true // show result on successful load
    z: 1
  }

  DataModel {
    id: dataModel
    dispatcher: logic
    onInitializedChanged: logic.loadData() // load/update data after initialization
  }

  Logic {
    id: logic
  }

  ViewHelper {
    id: viewHelper
  }

  // storage for caching data
  WebStorage {
    id: webStorage
    property var gameWindow: app
    gameNetworkItem: gameNetwork
    clearAllAtStartup: gameNetwork.clearAllUserDataAtStartup // allows to simulate with a clean app, without favored talks and scanned contacts

    // initialize data model with stored data at startup
    Component.onCompleted: {
      logic.initializeDataModel(webStorage)
      logic.increaseLocalAppStarts() // increase local app starts after first initialization
    }

    // update data model with fresh synchronized data
    onInitiallyInServerSyncChanged: {
      if(initiallyInServerSync)
        logic.initializeDataModel(webStorage)
    }
  }

  // facebook
  Facebook {
    id: facebook
    appId: AppSettings.facebookAppId
    readPermissions: [ "public_profile", "email", "user_friends" ]
  }

  // amplitude
  Amplitude {
    id: amplitude
    // From Amplitude Settings
    apiKey: AppSettings.amplitudeApiKey

    onPluginLoaded: {
      amplitude.logEvent("Start App", {"platform" : (system.isPlatform(System.IOS) ? "iOS" : "Android")})
    }
  }

  // game network
  FelgoGameNetwork {
    id: gameNetwork
    gameId: AppSettings.gameId
    secret: AppSettings.gameSecret
    facebookItem: facebook
    defaultUserName: "User %1"
    defaultPerPageCount: 100 // increase to show more users in leaderboard, default would be 30

    // this saves the get_user_scores request at app startup if the user already logged in before. it can be synced manually in the profile view
    autoLoadUserScoresAndAchievemenstWhenAuthenticated: false

    //clearAllUserDataAtStartup: system.desktopPlatform // this can be enabled during development to simulate a first-time app start
    clearOfflineSendingQueueAtStartup: true // clear any stored requests in the offline queue at app start, to avoid starting errors
    user.deviceId: system.UDID

    // update highscore with changes after scores are in sync initially
    property int addScoreWhenSynced: 0
    onUserScoresInitiallySyncedChanged: {
      if(userScoresInitiallySynced && !system.publishBuild) {
        console.log("Debug Build - reset current score of "+gameNetwork.userHighscoreForCurrentActiveLeaderboard+" to 0")
        var targetScore = 0
        if(dataModel.favorites) {
          for(var id in dataModel.favorites)
            targetScore++
        }
        gameNetwork.reportRelativeScore(targetScore - gameNetwork.userHighscoreForCurrentActiveLeaderboard)
      }
      else if(userScoresInitiallySynced && addScoreWhenSynced != 0)
        gameNetwork.reportRelativeScore(addScoreWhenSynced)
    }

    // reset / initialize data model when when GameNetwork user switches
    onUserInitiallyInSyncChanged: logic.resetOrInitializeDataModel(userInitiallyInSync)
  }

  // multiplayer
  FelgoMultiplayer {
    id: multiplayer
    gameNetworkItem: gameNetwork
    appKey: AppSettings.appKey
    pushKey: AppSettings.pushKey
    notificationBar: appNotificationBar // notification bar that also takes statusBarHeight into account
  }

  AppNotificationBar {
    id: appNotificationBar
    tintColor: Theme.tintColor
    onDismiss: {
      if(navigation)
        mainLoader.item.openInbox()
      else
        loaderItem.openInboxAfterLoading = true
    }
  }

  Component.onCompleted: {
    loaderTimer.start()
  }

  // load main item dynamically
  Loader {
    id: mainLoader
    // setting asynchronous to true causes issues at loading on Desktop. the item sometimes doesnt get fully loaded
    asynchronous: !system.desktopPlatform
    // the visible setting is irrelevant, as we move the item to the parent anyways
    visible: false
    //visible: status === Loader.Ready // this would be the ideal setting if we would display the item within loader, and not move it to the parent
    //onVisibleChanged: console.debug("mainLoader.visible changed to", visible)
    //onStatusChanged: console.debug("mainLoader.status changed to", status, " 0=null, 1=ready, 2=loading, 3=error")
    //source: Qt.resolvedUrl("QtWSMainItem.qml") // this would initially start loading, avoid this as we first wan to show the loading page
    onLoaded:{
      console.debug("xxx-QtWSLoaderItem: finished loading main qml file, showing it now")
      // this is required. nothing is displayed without moving it to the parent
      item.parent = loaderItem
      loadingFadeOut.start() // fade out loading screen (will reveal loaded item)
    }
  }

  Timer {
    id: loaderTimer
    interval: 100 // start loading asap after the items were completed. was set to 500ms before, but rather start faster, the loading screen animation is shown anyhow
    onTriggered: mainLoader.source = Qt.resolvedUrl("QtWSMainItem.qml")
  }

  // loading screen
  Rectangle {
    id: loading
    anchors.fill: parent
    color: "#f0f1f2"
    z: 2

    // Qt image
    AppImage {
      id: loadImage
      fillMode: AppImage.PreserveAspectCrop
      anchors.fill: parent
      source: "../assets/loader.png"
    }

    Column {
      anchors.centerIn: parent
      anchors.verticalCenterOffset: -dp(15)
      spacing: dp(15)

      // Spinner
      Item {
        id: loadSpinner
        width: dp(30)
        height: dp(30)
        anchors.horizontalCenter: parent.horizontalCenter

        Rectangle {
          width: dp(10)
          height: dp(10)
          radius: width/2
          color: Theme.tintColor
          anchors.horizontalCenter: parent.horizontalCenter
        }
        Rectangle {
          width: dp(10)
          height: dp(10)
          radius: width/2
          color: Theme.tintColor
          anchors.horizontalCenter: parent.horizontalCenter
          anchors.bottom: parent.bottom
        }

        RotationAnimator {
          target: loadSpinner
          running: true
          loops: Animation.Infinite
          from: 0
          to: 360
          duration: 2000
        }
      }

      // fade out
      NumberAnimation {
        id: loadingFadeOut
        target: loading
        property: "opacity"
        to: 0
        duration: 300
      }
    }
  }

  Connections {
    target: getApplication()

    // load data if not available and device goes online
    onIsOnlineChanged: {
      if(!dataModel.loaded && isOnline)
        loadDataTimer.start() // use timer to delay load as immediate calls might not get through (network not ready yet)
    }
  }

  // timer to load data after 1 second delay when going online
  Timer {
    id: loadDataTimer
    interval: 1000
    onTriggered: logic.loadData()
  }

  // we set the lightness of the used track colors based on the Theme.backgroundColor
  // this could be done with a simple property binding, but that strangely causes issues on Linux Qt 5.8
  // which is why this workaround with manual signal handling is used:
  Item {
    id: trackColorBindingFix
    property color baseTrackColor
    property var baseTrackLightness

    Component.onCompleted: updateBaseTrackLightness()
    Connections {
      target: Theme
      onBackgroundColorChanged: trackColorBindingFix.updateBaseTrackLightness()
    }

    function updateBaseTrackLightness() {
      trackColorBindingFix.baseTrackColor = Theme.backgroundColor
      trackColorBindingFix.baseTrackLightness = colorToHsl(trackColorBindingFix.baseTrackColor)[2]
    }
  }

  // getTrackColor - determines track color
  function getTrackColor(track) {
    if(!dataModel.tracks || dataModel.tracks[track] === undefined)
      return Theme.secondaryTextColor

    var light = 0.45 + 0.25 * (0.5 - trackColorBindingFix.baseTrackLightness)
    return Qt.hsla(dataModel.tracks[track], 1, light, 1)
  }

  // color to HSL conversion
  function colorToHsl(color) {
    var r = color.r /= 255
    var g = color.g /= 255
    var b = color.b /= 255
    var max = Math.max(r, g, b), min = Math.min(r, g, b);
    var h, s, l = (max + min) / 2;

    if(max == min) {
      h = s = 0; // achromatic
    }
    else {
      var d = max - min;
      s = l > 0.5 ? d / (2 - max - min) : d / (max + min);
      switch(max) {
      case r: h = (g - b) / d + (g < b ? 6 : 0); break;
      case g: h = (b - r) / d + 2; break;
      case b: h = (r - g) / d + 4; break;
      }
      h /= 6;
    }
    return [h, s, l];
  }

}
