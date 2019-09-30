import QtQuick 2.0
import Felgo 3.0
import QtGraphicalEffects 1.0
import "pages"
import "common"
import "common/social"

Item {
  anchors.fill: parent

  // make navigation public
  property alias navigation: navigation

  // social view (only once per app)
  property alias socialViewItem: socialView //publicly accessible

  Component.onCompleted: {
    if(system.publishBuild) {
      // give 1 point for opening the app
      if(gameNetwork.userScoresInitiallySynced)
        gameNetwork.reportRelativeScore(1)
      else
        gameNetwork.addScoreWhenSynced += 1
    }
    notificationTimer.start() // schedule notification at app-start
    checkFeedbackDialog() // check app starts and show feedback dialog
  }

  // handle data loading failed
  Connections {
    target: dataModel
    onLoadingFailed: NativeDialog.confirm("Failed to update conference data, please try again later.")
    onFavoriteAdded: {
      console.debug("favorite added")
      if(gameNetwork.userScoresInitiallySynced)
        gameNetwork.reportRelativeScore(1)
      else
        gameNetwork.addScoreWhenSynced += 1


      // only schedule a notification for the changed talk, not for all again
      scheduleNotificationForTalk(talk.id)

      amplitude.logEvent("Favor Talk",{"title" : talk.title, "talkId" : talk.id})
    }
    onFavoriteRemoved: {
      console.debug("favorite removed")
      if(gameNetwork.userScoresInitiallySynced && gameNetwork.userHighscoreForCurrentActiveLeaderboard > 0)
        gameNetwork.reportRelativeScore(-1)
      else if(!gameNetwork.userScoresInitiallySynced)
        gameNetwork.addScoreWhenSynced -= 1

      cancelNotificationForTalk(talk.id)

      amplitude.logEvent("Unfavor Talk",{"title" : talk.title, "talkId" : talk.id})
    }
    onFavoritesChanged: {
      // the schedule call is stopped below anyways, if the notificationTimer did not complete
      // with this check, we would not allow push not syncing if the user is offline, thus do not add this
      //if(!gameNetwork.userInitiallyInSync) {
      //  console.debug("favorties changed, but user is not synced with server yet thus wait")
      //  return
      //}

      console.debug("onFavoritesChanged")
      // dont reschedule all when favorites changed - they are scheduled indidividually instead as a complete reschedule is a lenghty operation
      //scheduleNotificationsForFavorites()
    }
    onNotificationsEnabledChanged: {
       console.debug("onNotificationsEnabledChanged, reschedule notifications")
      scheduleNotificationsForFavorites()
    }
  }

  // timer to schedule notifications several seconds after app startup
  Timer {
    id: notificationTimer
    interval: 8000 // we can delay this, is not time-critical to happen initially
    //running: true // start the timer when the compoent was loaded - it is started from onCompleted after the navigation was setup
    onTriggered: {
      console.debug("notificationTimer.triggered", running)
      scheduleNotificationsForFavorites()
    }
  }

  Connections {
    target: notificationManager
    // display alert for upcoming sessions
    onNotificationFired: {
      if(notificationId >= 0) {
        // session reminder
        if(dataModel.loaded && dataModel.talks && dataModel.talks[notificationId]) {
          var talk = dataModel.talks[notificationId]
          var text = talk["title"]+" starts "+talk.start+" at "+talk["room"]+"."
          var title = "Session Reminder"
          NativeDialog.confirm(title, text, function(){}, false)
        }
      }
      else {
        // default notification
        NativeDialog.confirm("The conference starts soon!", "Thanks for using our app, we wish you a great Qt World Summit 2017!", function(){}, false)
      }
    }
  }


  // scheduleNotificationsForFavorites
  function scheduleNotificationsForFavorites() {
    console.debug("attempting scheduleNotificationsForFavorites()")

    if(notificationTimer.running) {
      console.debug("notificationTimer at initialization is currently running, dont update yet")
      return
    }

    // if used within Demo Launcher app project, we do not use notifications
    if(typeof notificationManager === "undefined")
      return

    console.debug("scheduling notifications now")

    // TODO: only re-schedule, if the current nofications changed. this may be a lengthy process

    notificationManager.cancelAllNotifications()
    if(!dataModel.notificationsEnabled || !dataModel.favorites || !dataModel.talks)
      return

    for(var idx in dataModel.favorites) {
      var talkId = dataModel.favorites[idx]
      scheduleNotificationForTalk(talkId)
    }

    // add notification before world summit starts!
    var nowTime = new Date().getTime()
    var eveningBeforeConferenceTime = new Date("2017-10-09T21:00.000"+dataModel.timeZone).getTime()
    if(nowTime < eveningBeforeConferenceTime) {
      var text = "Felgo wishes all the best for Qt World Summit 2017!"
      var notification = {
        notificationId: -1,
        message: text,
        timestamp: Math.round(eveningBeforeConferenceTime / 1000) // utc seconds
      }
      notificationManager.schedule(notification)
    }
  }

  // scheduleNotificationForTalk
  function scheduleNotificationForTalk(talkId) {
    if(dataModel.loaded && dataModel.talks && dataModel.talks[talkId]) {
      var talk = dataModel.talks[talkId]
      var text = talk["title"]+" starts "+talk.start+" at "+talk["room"]+"."

      var nowTime = new Date().getTime()
      var utcDateStr = talk.day+"T"+talk.start+".000"+dataModel.timeZone
      var notificationTime = new Date(utcDateStr).getTime()
      notificationTime = notificationTime - 10 * 60 * 1000 // 10 minutes before

      if(nowTime < notificationTime) {
        var notification = {
          notificationId: talkId,
          message: text,
          timestamp: Math.round(notificationTime / 1000) // utc seconds
        }
        notificationManager.schedule(notification)
      }
    }
  }

  function cancelNotificationForTalk(talkId) {
    notificationManager.cancelNotification(talkId)
  }

  // app navigation
  Navigation {
    id: navigation
    property string previousPageTitle: ""
    property string previousPagePlatform: ""
    property var currentPage: {
      if(!currentNavigationItem)
        return null

      if(currentNavigationItem.navigationStack)
        return currentNavigationItem.navigationStack.currentPage
      else
        return currentNavigationItem.page
    }
    // track previous page & platform changes to restore previous page on platform-switch when navigation changes
    onCurrentPageChanged: {
      if(previousPagePlatform !== Theme.platform && previousPageTitle !== "") {
        previousPagePlatform = Theme.platform

        // current page changed due to platform switch -> restore previous page
        restorePageTimer.previousPageTitle = previousPageTitle
        restorePageTimer.restart()
      }
      else if(!!currentPage) {
        previousPagePlatform = Theme.platform
        previousPageTitle = currentPage.title
      }
    }

    // automatically load data if not loaded and schedule/favorites page is opened
    onCurrentIndexChanged: {
      if(currentIndex > 0 && currentIndex < 3) {
        if(!dataModel.loaded && isOnline)
          logic.loadData()
      }
    }
    onCurrentNavigationItemChanged: {
      amplitude.logEvent("Open Page",{"title" : currentNavigationItem.title})
    }

    // Android drawer header item
    headerView: Item {
      width: parent.width
      height: dp(75) + Theme.statusBarHeight
      clip: true

      Rectangle {
        anchors.fill: parent
        color: Theme.tintColor
      }

      AppImage {
        width: parent.width
        fillMode: AppImage.PreserveAspectFit
        source: "../assets/venue_photo.jpg"
        anchors.verticalCenter: parent.verticalCenter
      }

      AppImage {
        width: parent.width
        fillMode: AppImage.PreserveAspectFit
        source: "../assets/venue_photo.jpg"
        anchors.verticalCenter: parent.verticalCenter
        opacity: 0.5
        layer.enabled: true
        layer.effect: Colorize {
          id: titleImgColorize
          lightness: 0.1
          saturation: 0.5

          // we set the hue for the colorize effect based on the Theme.tintColor
          // this could be done with a simple property binding, but that strangely causes issues on Linux Qt 5.8
          // which is why this workaround with manual signal handling is used:
          property color baseColor
          Component.onCompleted: updateHue()
          Connections {
            target: app
            onSecondaryTintColorChanged: titleImgColorize.updateHue()
          }
          function updateHue() {
            titleImgColorize.baseColor = app.secondaryTintColor
            var hslColor = loaderItem.colorToHsl(titleImgColorize.baseColor)
            titleImgColorize.hue = hslColor[0]
            titleImgColorize.saturation = hslColor[1]
            titleImgColorize.lightness = hslColor[2]
          }
        }
      }

      AppImage {
        width: parent.width * 0.75
        source: "../assets/QtWS2017_logo_white.png"
        fillMode: AppImage.PreserveAspectFit
        anchors.horizontalCenter: parent.horizontalCenter
        y: Theme.statusBarHeight + ((parent.height - Theme.statusBarHeight) - height) * 0.5
        layer.enabled: true
        layer.effect: DropShadow {
          color: Qt.rgba(0,0,0,0.5)
          radius: 16
          samples: 16
        }
      }
    }

    NavigationItem {
      title: "About"
      iconComponent: Component {
        Item {
          height: !!parent ? parent.height : 0
          width: height

          property bool selected: parent && parent.selected

          Icon {
            anchors.centerIn: parent
            width: height
            height: parent.height
            icon: IconType.home
            color: !parent.selected ? Theme.textColor  : Theme.tintColor
            visible: !felgoIcon.visible
          }

          Image {
            id: felgoIcon
            height: parent.height
            anchors.horizontalCenter: parent ? parent.horizontalCenter : undefined
            fillMode: Image.PreserveAspectFit
            source: !parent.selected ? (Theme.isAndroid ? "../assets/Qt_logo_Android_off.png" : "../assets/Qt_logo_iOS_off.png") : "../assets/Qt_logo.png"
            visible: true // Theme.isIos || Theme.backgroundColor.r == 1 && Theme.backgroundColor.g == 1 && Theme.backgroundColor.b == 1
          }
        }
      }

      NavigationStack {
        navigationBarShadow: false
        MainPage {}
      }
    } // main

    NavigationItem {
      title: "Timetable"
      icon: IconType.calendaro

      NavigationStack {
        splitView: tablet && landscape
        // if first page, reset leftColumnIndex (may change when searching)
        onTransitionFinished: {
          if(depth === 1)
            leftColumnIndex = 0
        }

        TimetablePage {
        }
      }
    } // timetable

    NavigationItem {
      title: "Favorites"
      icon: IconType.star

      asynchronous: true
      sourceComponent: favoritesComponent
    } // favorites

    NavigationItem {
      title: "Speakers"
      icon: IconType.microphone

      asynchronous: true
      sourceComponent: speakersComponent
    } // speakers

    NavigationItem {
      title: "More"
      icon: IconType.ellipsish
      showItem: !Theme.isAndroid

      sourceComponent: moreComponent
      asynchronous: true
    }

    // social
    NavigationItem {
      title: "Business Meet"
      icon: IconType.group
      showItem: Theme.isAndroid

      asynchronous: true
      sourceComponent: businessMeetComponent
    } // business meet

    NavigationItem {
      title: "Profile"
      icon: IconType.user
      showItem: Theme.isAndroid

      asynchronous: true
      sourceComponent: profileComponent
    } // profile

    NavigationItem {
      title: "Chat"
      icon: IconType.comment
      showItem: Theme.isAndroid

      asynchronous: true
      sourceComponent: chatComponent
    } // chat

    NavigationItem {
      title: "Leaderboard"
      icon: IconType.flagcheckered
      showItem: Theme.isAndroid

      asynchronous: true
      sourceComponent: leaderboardComponent
    } // leaderboard

    NavigationItem {
      title: "Tracks"
      icon: IconType.tag
      showItem: Theme.isAndroid

      asynchronous: true
      sourceComponent: tracksComponent
    } // tracks


    NavigationItem {
      title: "Venue"
      icon: IconType.building
      showItem: Theme.isAndroid

      sourceComponent: venueComponent
      asynchronous: true
    } // venue

    NavigationItem {
      title: "QR Contacts"
      icon: IconType.qrcode
      showItem: Theme.isAndroid

      asynchronous: true
      sourceComponent: contactsComponent
    } // qr contacts

    NavigationItem {
      title: "Settings"
      icon: IconType.gears
      showItem: Theme.isAndroid

      sourceComponent: settingsComponent
      asynchronous: true
    } // settings

    NavigationItem {
      title: "About Felgo"
      showItem: Theme.isAndroid
      iconComponent: Item {
        height: parent.height
        width: height

        property bool selected: parent && parent.selected

        Icon {
          anchors.centerIn: parent
          width: height
          height: parent.height
          icon: IconType.home
          color: !parent.selected ? Theme.textColor  : Theme.tintColor
          visible: !felgoIcon.visible
        }

        Image {
          id: felgoIcon
          height: parent.height
          anchors.horizontalCenter: parent ? parent.horizontalCenter : undefined
          fillMode: Image.PreserveAspectFit
          source: !parent.selected ? "../assets/Felgo_icon_nav_off.png" : "../assets/Felgo_icon_nav.png"
          visible: Theme.isIos || Theme.backgroundColor.r == 1 && Theme.backgroundColor.g == 1 && Theme.backgroundColor.b == 1
        }
      }

      sourceComponent: aboutFelgoComponent
      asynchronous: true
    } // About Felgo
  } // nav

  Component {
    id: favoritesComponent
    NavigationStack {
      splitView: tablet && landscape
      Component.onCompleted: push(Qt.resolvedUrl("pages/FavoritesPage.qml"))
    }
  }

  Component {
    id: speakersComponent
    NavigationStack {
      splitView: landscape && tablet
      Component.onCompleted: push(Qt.resolvedUrl("pages/SpeakersPage.qml"))
    }
  }

  Component {
    id: moreComponent
    NavigationStack {
      splitView: tablet && landscape
      Component.onCompleted: push(Qt.resolvedUrl("pages/MorePage.qml"))
    }
  }

  Component {
    id: businessMeetComponent
    NavigationStack {
      Component.onCompleted: push(socialView.businessMeetPage)
    }
  }

  Component {
    id: profileComponent
    NavigationStack {
      Component.onCompleted: push(socialView.profilePage)
    }
  }

  Component {
    id: chatComponent
    NavigationStack {
      Component.onCompleted: push(socialView.inboxPage)
    }
  }

  Component {
    id: leaderboardComponent
    NavigationStack {
      Component.onCompleted: push(socialView.leaderboardPage)
    }
  }

  Component {
    id: tracksComponent
    NavigationStack {
      splitView: landscape && tablet
      Component.onCompleted: push(Qt.resolvedUrl("pages/TracksPage.qml"))
    }
  }

  Component {
    id: venueComponent
    NavigationStack {
      Component.onCompleted: push(Qt.resolvedUrl("pages/VenuePage.qml"))
    }
  }

  Component {
    id: contactsComponent
    NavigationStack {
      // Note: QZXing library for barcode scanning is not available with QML Live Reloading
      Component.onCompleted: push(Qt.resolvedUrl("pages/ContactsPage.qml"))
    }
  }

  Component {
    id: settingsComponent
    NavigationStack {
      Component.onCompleted: push(Qt.resolvedUrl("pages/SettingsPage.qml"))
    }
  }

  Component {
    id: aboutFelgoComponent
    NavigationStack {
      Component.onCompleted: push(Qt.resolvedUrl("pages/AboutFelgoPage.qml"))
    }
  }

  SocialView {
    id: socialView
    visible: false
    anchors.fill: parent
    tintColor: Theme.tintColor

    profileUserDelegate: SocialProfileDelegate { }
    leaderboardUserDelegate: SocialLeaderboardDelegate { }

    property Component businessMeetPage: SocialUserSearchPage {
      id: businessMeetPage
      title: "Business Meet"
      filterToUsersWithCustomData: true // only search users with custom data

      // replace default user delegate to also show custom data
      userSearchUserDelegate: SocialSearchDelegate { }

      // open profile when user is selected
      onUserSelected: {
        socialViewItem.pushProfilePage(gameNetworkUser, businessMeetPage.navigationStack,
                               { friendStatus: modelData.friendStatus})
      }
    }
  }

  Timer {
    id: restorePageTimer
    interval: 5
    property string previousPageTitle: ""
    onTriggered: {
      var activeTitle = restorePageTimer.previousPageTitle

      // set active Android NavigationItem in drawer
      if(Theme.isAndroid) {
        if(activeTitle === "About Felgo")
          navigation.currentIndex = 12
        else if(activeTitle === "Settings")
          navigation.currentIndex = 11
        else if(activeTitle === "Contacts")
          navigation.currentIndex = 10
        else if(activeTitle === "Venue")
          navigation.currentIndex = 9
        else if(activeTitle === "Tracks")
          navigation.currentIndex = 8
        else if(activeTitle === "Highscores")
          navigation.currentIndex = 7
        else if(activeTitle === "Inbox")
          navigation.currentIndex = 6
        else if(activeTitle === "User Profile")
          navigation.currentIndex = 5
        else if(activeTitle === "Business Meet")
          navigation.currentIndex = 4
      }
      else {
        // push active Page to More-Page on iOS
        var target = ""
        if(activeTitle === "Highscores")
          target = socialView.leaderboardPage
        else if(activeTitle === "Inbox")
          target = socialView.inboxPage
        else if(activeTitle === "User Profile")
          target = socialView.profilePage
        else if(activeTitle === "Business Meet")
          target = socialView.businessMeetPage
        else if(activeTitle === "About Felgo")
          target = Qt.resolvedUrl("pages/AboutFelgoPage.qml")
        else if(activeTitle === "Settings")
          target = Qt.resolvedUrl("pages/SettingsPage.qml")
        else if(activeTitle === "Contacts")
          target = Qt.resolvedUrl("pages/ContactsPage.qml")
        else if(activeTitle === "Venue")
          target = Qt.resolvedUrl("pages/VenuePage.qml")
        else if(activeTitle === "Tracks")
          target = Qt.resolvedUrl("pages/TracksPage.qml")

        if (target != "") {
          navigation.currentIndex = navigation.countVisible - 1 // open more page
          navigation.currentPage.navigationStack.push(target)
        }
      }
    }
  }

  // openInbox activates inbox navigation item on Android or more navigation item with inbox page on iOS
  function openInbox() {
    if(Theme.isAndroid)
      navigation.currentIndex = 6 // go to inbox navigation item
    else {
      navigation.currentIndex = navigation.count - 1 // open more page
      if(!navigation.currentPage || navigation.currentPage.title === "Inbox")
        return

      navigation.currentPage.navigationStack.push(socialView.inboxPage) // push inbox page
    }
  }

  // check app starts and show feedback dialog if required
  function checkFeedbackDialog() {
    if(dataModel.localAppStarts > 5 && !dataModel.feedBackSent) {
      likeDialog.open()
    }
  }

  FeedbackDialog {
    id: feedbackDialog
  }

  RatingDialog {
    id: ratingDialog
  }

  LikeDialog {
    id: likeDialog
    onCanceled: {
      amplitude.logEvent("Dislike App")

      // open the feedback dialog instead
      likeDialog.close()
      feedbackDialog.open()
    }
    onAccepted: {
      amplitude.logEvent("Like App")

      // open the rating dialog instead
      likeDialog.close()
      ratingDialog.open()
    }
  }
}
