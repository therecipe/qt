import QtQuick 2.4
import QtQuick.Controls 1.2

import Felgo 3.0
import "."
import "../widgets"


Page {
  id: twitterMainPage

  // These can be used from anywhere in the app - this way the QML files are parsed only once
  Component { id: mainPageComponent; MainPage { } }
  Component { id: detailPageComponent; DetailPage { } }
  Component { id: profilePageComponent; ProfilePage { } }
  Component { id: listsPageComponent; ListsPage { } }

  // make page navigation public, so app-demo launcher can track navigation changes and log screens with Google Analytics
  property alias navigation: navigation

  useSafeArea: false // fill full screen

  Navigation {
    id: navigation
    drawer.drawerPosition: drawer.drawerPositionLeft
    headerView: NavHeader {}
    footerView: NavFooter {}

    //this overrides the default mode of drawer on android and tabs elsewhere
    //navigationMode: navigationModeTabsAndDrawer

    NavigationItem {
      title: "Home"
      icon: IconType.home

      NavigationStack {
        MainPage { }
      }
    }

    NavigationItem {
      title: "Lists"
      icon: IconType.bars

      NavigationStack {
        ListsPage { }
      }
    }

    NavigationItem {
      title: "Messages"
      icon: IconType.envelope

      NavigationStack {
        MessagesPage { }
      }
    }

    NavigationItem {
      title: "Me"
      icon: IconType.user

      NavigationStack {

        // manually push profilePage to fix initial scroll position of profile page
        // (due to bug when using ListView::headerItem)
        Component.onCompleted: {
          push(profilePageComponent, { profile: dataModel.currentProfile })
        }
      }
    }
  }
}
