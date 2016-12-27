// ekke (Ekkehard Gentz) @ekkescorner
import QtQuick 2.6
import QtQuick.Layouts 1.3
import QtQuick.Controls 2.0
import QtQuick.Controls.Material 2.0
import QtGraphicalEffects 1.0

import "../common"
import "../tabs"

Page {
    id: navPage
    // index to get access to Loader (Destination)
    property int myIndex: index
    property string name: "colorSchemaNavPage"

    // because of https://bugreports.qt.io/browse/QTBUG-54260
    // lastCurrentIndex will remember currentIndex, so we can reset before Page becomes currentItem on StackView
    property int lastCurrentIndex: 0

    property alias currentIndex: navPane.currentIndex

    property bool tabBarIsFixed: true
    property var tabButtonModel: [{"name": "Primary"},
        {"name": "Accent"},
        {"name": "Theme"}]

    header: isLandscape? null : tabBar

    Loader {
        id: tabBar
        visible: !isLandscape
        active: !isLandscape
        source: "../tabs/TextTabBar.qml"
        onLoaded: {
            console.log("Tab Bar LOADED")
            if(item) {
                item.currentIndex = navPane.currentIndex
            }
        }
    }
    Loader {
        id: tabBarFloating
        visible: isLandscape
        anchors.top: parent.top
        anchors.left: parent.left
        anchors.right: parent.right
        active: isLandscape
        source: "../tabs/TextTabBar.qml"
        onLoaded: {
            console.log("Floating Tab Bar LOADED")
            if(item) {
                item.currentIndex = navPane.currentIndex
            }
        }
    }

    SwipeView {
        id: navPane
        focus: true
        anchors.top: isLandscape? tabBarFloating.bottom : parent.top
        anchors.topMargin: isLandscape? 6 : 0
        anchors.left: parent.left
        anchors.right: parent.right
        anchors.bottom: parent.bottom
        currentIndex: 0
        // currentIndex is the NEXT index swiped to
        onCurrentIndexChanged: {
            console.log("Swipe View current index changed: "+currentIndex)
            if(isLandscape) {
                tabBarFloating.item.currentIndex = currentIndex
            } else {
                tabBar.item.currentIndex = currentIndex
            }
        }

        Shortcut {
            sequence: "w"
            onActivated: navPane.goToPage(0)
        }
        Shortcut {
            sequence: "Alt+w"
            onActivated: navPane.goToPage(0)
        }
        Shortcut {
            sequence: "e"
            onActivated: navPane.goToPage(1)
        }
        Shortcut {
            sequence: "Alt+e"
            onActivated: navPane.goToPage(1)
        }
        Shortcut {
            sequence: "r"
            onActivated: navPane.goToPage(2)
        }
        Shortcut {
            sequence: "Alt+r"
            onActivated: navPane.goToPage(2)
        }

        function goToPage(pageIndex) {
            if(pageIndex == navPane.currentIndex) {
                // it's the current page
                return
            }
            if(pageIndex > 4 || pageIndex < 0) {
                return
            }
            navPane.currentIndex = pageIndex
        } // goToPage

        // only 3 Pages - all preloaded to be able to swipe
        Loader {
            // index 0
            id: pageOneLoader
            active: true
            source: "PrimaryColorPage.qml"
            onLoaded: item.init()
        }
        Loader {
            // index 1
            id: pageTwoLoader
            active: true
            source: "AccentColorPage.qml"
            onLoaded: item.init()
        }
        Loader {
            // index 2
            id: pageThreeLoader
            active: true
            source: "ThemePage.qml"
            onLoaded: item.init()
        }

        function switchPrimaryPalette(paletteIndex) {
            primaryPalette = myApp.primaryPalette(paletteIndex)
        }
        function switchAccentPalette(paletteIndex) {
            accentPalette = myApp.accentPalette(paletteIndex)
        }

    } // navPane

    // emitting a Signal could be another option
    Component.onDestruction: {
        cleanup()
    }

    function init() {
        console.log("INIT colorSchemaNavPage")
    }
    function cleanup() {
        console.log("CLEANUP colorSchemaNavPage")
    }

} // navPage
