import QtQuick 2.4
import QtQuick.Controls 2.0 as Quick2
import QtQuick.Layouts 1.3
import Felgo 3.0

import "pages"

Page {
  // make page navigation public, so app-demo launcher can track navigation changes and log screens with Google Analytics
  property alias navigation: navigation

  // for windows platform check
  property bool isWindows: system.isPlatform(4) || system.isPlatform(9) || system.isPlatform(10) // windows, winPhone or winRT

  useSafeArea: false // main page fills whole window

  Navigation {
    id: navigation

    NavigationItem {
      id: componentsItem
      title: "Components"
      icon: IconType.calculator

      NavigationStack {
        // we use initial page and component is here instead of setting the Page as a child
        // this is only required to avoid initialization issues when running this app within Felgo Sample Launcher
        initialPage: appComponentsPage

        Component {
          id: appComponentsPage
          Page {
            id: page
            title: componentsItem.title
            rightBarItem: stackLayout.currentIndex >= 0 ? stackLayout.children[stackLayout.currentIndex].rightBarItem : null

            AppTabBar {
              id: tabBar
              contentContainer: stackLayout
              AppTabButton { text: "Overview" }
              AppTabButton { text: "Example Page" }
            }

            StackLayout {
              id: stackLayout
              width: parent.width
              anchors.top: tabBar.bottom
              anchors.bottom: parent.bottom
              clip: true

              // we need a Wrapper-Item because the Page always fills its parent,
              // which would be the bigger container within SwipeView
              AppComponentsPage { }
              ExamplePage { }
            }
          }
        } // initial page component

      }
    }

    NavigationItem {
      id: effectsItem
      title: "Effects"
      icon: IconType.bolt

      NavigationStack {
        EffectsPage { title: effectsItem.title }
      }
    }

    NavigationItem {
      id: controlsItem
      title: "Controls"
      icon: IconType.toggleoff

      NavigationStack {
        navigationBarShadow: false
        ControlsPage { title: controlsItem.title }
      }
    }

    NavigationItem {
      id: featuresItem
      title: "Features"
      icon: IconType.listalt

      NavigationStack {
        FeaturesPage { title: featuresItem.title }
      }
    }

    NavigationItem {
      id: listViewItem
      title: "List View"
      icon: IconType.list

      NavigationStack {
        splitView: tablet
        SimpleListPage { title: listViewItem.title }
      }
    }
  }
}
