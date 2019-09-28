import Felgo 3.0
import "pages"


Page {
  // make navigation public, so app-demo launcher can track navigation changes and log screens with Google Analytics
  property alias childNavigationStack: globalNavStack
  property alias navigation: navigation
  useSafeArea: false // full screen

  NavigationStack {
    id: globalNavStack

    // Wrapper page
    Page {
      navigationBarHidden: Theme.isAndroid
      useSafeArea: false // full screen

      title: navigation.currentNavigationItem ? navigation.currentNavigationItem.title : ""

      Navigation {
        id: navigation

        navigationMode: navigationModeTabs

        NavigationItem {
          title: qsTr("Recent")
          icon: IconType.clocko

          MessageListPage {}
        }

        NavigationItem {
          id: groupsItem
          title: qsTr("Groups")
          icon: IconType.group

          ListPage {
            title: groupsItem.title
            emptyText.text: qsTr("No groups available.")
          }
        }

        NavigationItem {
          id: peopleItem
          title: qsTr("People")
          icon: IconType.list

          ListPage {
            title: peopleItem.title
            emptyText.text: qsTr("No contacts.")
          }
        }

        NavigationItem {
          id: settingsItem
          title: qsTr("Settings")
          icon: IconType.cog

          ListPage {
            title: settingsItem.title
            emptyText.text: qsTr("No settings.")
          }
        }

      }
    }
  }
}
