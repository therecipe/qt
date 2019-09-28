import QtQuick 2.0
import Felgo 3.0
import "../common"

FlickablePage {
  id: examplePage
  property bool activated: true

  // configure custom rightBarItem for navigation bar of example page
  rightBarItem:  NavigationBarRow {
    ActivityIndicatorBarItem {
      id: busyBarItem
      animating: examplePage.activated
      showItem: showItemAlways
    }
    IconButtonBarItem {
      icon: examplePage.activated ? IconType.toggleon : IconType.toggleoff
      onClicked: examplePage.activated = !examplePage.activated
      title: "Toggle"
    }
  }

  flickable.contentHeight: scrollCol.height + dp(8)

  Column  {
    id: scrollCol
    width: parent.width

    SectionHeader { icon: IconType.thumbsup; text: "Rich User Interface, Less Code" }
    SectionDescription { text: "This example only requires about 120 lines code." }
    Repeater {
      model: [
        { text: "Default UI Elements",
          detailText: "No custom components are used here.",
          icon: IconType.tablet },
        { text: "Swipe a list row for more options",
          detailText: "This SwipeOptionsContainer has a left option and a right option.",
          icon: IconType.list },
        { text: "What else is shown?",
          detailText: "SimpleRow, SwipeButton, ActivityIndicator, AppSwitch, AppButton, AppText",
          icon: IconType.question }
      ]
      delegate: SwipeOptionsContainer {
        id: container
        height: row.height
        enabled: !(isSnapped && isRight)

        //the actual list item
        SimpleRow {
          id: row
          style.showDisclosure: false
        }

        //left swipe option (when swiping list item to right)
        leftOption: SwipeButton {
          text: "Option"
          icon: IconType.gear
          height: row.height
          onClicked: {
            row.item.text = "Option clicked"
            row.itemChanged()

            //hide left option when clicked
            container.hideOptions()
          }
        }

        //right swipe option (when swiping list item to left)
        rightOption: AppActivityIndicator {
          width: row.height
          anchors.centerIn: parent
        }

        //hide right option after timer finishes
        onRightOptionShown: hideTimer.start()

        property Timer hideTimer: Timer {
          running: false
          interval: 1000
          onTriggered: container.hideOptions()
        }
      }
    }

    SectionSpacer { }
    SectionHeader { text: "Signals and Property Binding" }
    SectionDescription { text: "Press the switch below or in the navigation bar. Both items are linked together and also affect both of the activity indicators.\n
This is achieved by binding the checked and activated states of the controls in a single code line." }
    SectionContent { contentItem: Row {
        anchors.horizontalCenter: parent.horizontalCenter

        AppSwitch {
          id: sw
          anchors.verticalCenter: parent.verticalCenter
          checked: examplePage.activated
          updateChecked: false //always keep the property binding
          onToggled: examplePage.activated = !examplePage.activated
        }

        AppActivityIndicator {
          anchors.verticalCenter: parent.verticalCenter
          animating: examplePage.activated
        }
      }
    }
  } // flickable content
}
