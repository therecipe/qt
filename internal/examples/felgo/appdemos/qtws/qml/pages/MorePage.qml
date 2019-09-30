import Felgo 3.0
import QtQuick 2.0
import "../common"

ListPage {
  id: morePage
  title: "More"

  // Note: no need to e.g. use JsonListModel here, this model menu items won't change while using the app
  model: [
    { text: "Business Meet", section: "Social", page: socialViewItem.businessMeetPage },
    { text: "Your Profile", section: "Social", page: socialViewItem.profilePage },
    { text: "Chat", section: "Social", page: socialViewItem.inboxPage },
    { text: "Leaderboard", section: "Social", page: socialViewItem.leaderboardPage },
    { text: "Tracks", section: "General", page: Qt.resolvedUrl("TracksPage.qml") },
    { text: "Venue", section: "General", page: Qt.resolvedUrl("VenuePage.qml") },
    { text: "QR Contacts", section: "General", page: Qt.resolvedUrl("ContactsPage.qml")},
    { text: "Settings", section: "General", page: Qt.resolvedUrl("SettingsPage.qml") },
    { text: "About Felgo", section: "General", page: Qt.resolvedUrl("AboutFelgoPage.qml") }
  ]

  section.property: "section"

  // open configured page when clicked
  onItemSelected: {
    morePage.navigationStack.popAllExceptFirstAndPush(model[index].page)
  }
}
