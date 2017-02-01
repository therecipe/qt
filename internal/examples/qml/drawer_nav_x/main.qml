// ekke (Ekkehard Gentz) @ekkescorner
import QtQuick 2.6
import QtQuick.Layouts 1.3
import QtQuick.Controls 2.0
import QtQuick.Controls.Material 2.0
import QtGraphicalEffects 1.0
import "common"
import "pages"
import "popups"
import "tabs"
import "navigation"

// This app demonstrates HowTo use Qt 5.7 new Qt Quick Controls 2, High DPI and more
// This app is NOT a production ready app
// This app's goal is only to help you to understand some concepts
// see blog http://j.mp/qt-x to learn about Qt 5.7 for Material - styled Android or iOS Apps
// learn about this drawer_nav_x app from this article: http://bit.ly/qt-drawer-nav-x
// ekke (Ekkehard gentz) @ekkescorner

ApplicationWindow {
    id: appWindow
    // visibile must set to true - default is false
    visible: true
    minimumWidth: 360
    minimumHeight: 520
    //
    property bool isLandscape: width > height

    // primary and accent properties:
    property variant primaryPalette: myApp.defaultPrimaryPalette()
    property color primaryLightColor: primaryPalette[0]
    property color primaryColor: primaryPalette[1]
    property color primaryDarkColor: primaryPalette[2]
    property color textOnPrimaryLight: primaryPalette[3]
    property color textOnPrimary: primaryPalette[4]
    property color textOnPrimaryDark: primaryPalette[5]
    property string iconOnPrimaryLightFolder: primaryPalette[6]
    property string iconOnPrimaryFolder: primaryPalette[7]
    property string iconOnPrimaryDarkFolder: primaryPalette[8]
    property variant accentPalette: myApp.defaultAccentPalette()
    property color accentColor: accentPalette[0]
    property color textOnAccent: accentPalette[1]
    property string iconOnAccentFolder: accentPalette[2]
    Material.primary: primaryColor
    Material.accent: accentColor
    // theme Dark vs Light properties:
    property variant themePalette: myApp.defaultThemePalette()
    property color dividerColor: themePalette[0]
    property color cardAndDialogBackground: themePalette[1]
    property real primaryTextOpacity: themePalette[2]
    property real secondaryTextOpacity: themePalette[3]
    property real dividerOpacity: themePalette[4]
    property real iconActiveOpacity: themePalette[5]
    property real iconInactiveOpacity: themePalette[6]
    property string iconFolder: themePalette[7]
    property int isDarkTheme: themePalette[8]
    property color flatButtonTextColor: themePalette[9]
    property color popupTextColor: themePalette[10]
    property real toolBarActiveOpacity: themePalette[11]
    property real toolBarInactiveOpacity: themePalette[12]
    property color toastColor: themePalette[13]
    property real toastOpacity: themePalette[14]
    // Material.dropShadowColor  OK for Light, but too dark for dark theme
    property color dropShadow: isDarkTheme? "#E4E4E4" : Material.dropShadowColor
    onIsDarkThemeChanged: {
        if(isDarkTheme == 1) {
            Material.theme = Material.Dark
        } else {
            Material.theme = Material.Light
        }
    }
    // font sizes - defaults from Google Material Design Guide
    property int fontSizeDisplay4: 112
    property int fontSizeDisplay3: 56
    property int fontSizeDisplay2: 45
    property int fontSizeDisplay1: 34
    property int fontSizeHeadline: 24
    property int fontSizeTitle: 20
    property int fontSizeSubheading: 16
    property int fontSizeBodyAndButton: 14 // is Default
    property int fontSizeCaption: 12
    property int fontSizeActiveNavigationButton: 14
    property int fontSizeInactiveNavigationButton: 12
    // fonts are grouped into primary and secondary with different Opacity
    // to make it easier to get the right property,
    // here's the opacity per size:
    property real opacityDisplay4: secondaryTextOpacity
    property real opacityDisplay3: secondaryTextOpacity
    property real opacityDisplay2: secondaryTextOpacity
    property real opacityDisplay1: secondaryTextOpacity
    property real opacityHeadline: primaryTextOpacity
    property real opacityTitle: primaryTextOpacity
    property real opacitySubheading: primaryTextOpacity
    // body can be both: primary or secondary text
    property real opacityBodyAndButton: primaryTextOpacity
    property real opacityBodySecondary: secondaryTextOpacity
    property real opacityCaption: secondaryTextOpacity
    //
    // TODO C++ ENUM
    // Destination
    // TabBar-SwipeView
    // StackView
    property var activationPolicy: {
        "NONE":0, "IMMEDIATELY":1, "LAZY":2, "WHILE_CURRENT":3
    }

    // NAVIGATION BAR PROPRTIES (a_p == activationPolicy)
    // IMMEDIATELY: Home, Flight
    // LAZY: Car, Bus, Truck, ColorSchema
    // WHILE_CURRENT: Subway, About, Settings
    // StackView: Home --> QtPage
    // SwipeView+TabBar: ColorSchema
    property var navigationModel: [
        {"type": "../navigation/DrawerNavigationButton.qml", "name": "Home", "icon": "home.png", "source": "../navigation/HomeNavigation.qml", "showCounter":false, "showMarker":false, "a_p":1},
        {"type": "../navigation/DrawerDivider.qml", "name": "", "icon": "", "source": "", "a_p":1},
        {"type": "../navigation/DrawerSubtitle.qml", "name": "Transitoptions", "icon": "", "source": "", "a_p":1},
        {"type": "../navigation/DrawerNavigationButton.qml", "name": "Car", "icon": "car.png", "source": "../pages/PageOne.qml", "showCounter":true, "showMarker":false, "a_p":2},
        {"type": "../navigation/DrawerNavigationButton.qml", "name": "Bus", "icon": "bus.png", "source": "../pages/PageTwo.qml", "showCounter":false, "showMarker":false, "a_p":2},
        {"type": "../navigation/DrawerNavigationButton.qml", "name": "Subway", "icon": "subway.png", "source": "../pages/PageThree.qml", "showCounter":false, "showMarker":true, "a_p":3},
        {"type": "../navigation/DrawerNavigationButton.qml", "name": "Truck", "icon": "truck.png", "source": "../pages/PageFour.qml", "showCounter":false, "showMarker":false, "a_p":2},
        {"type": "../navigation/DrawerNavigationButton.qml", "name": "Flight", "icon": "flight.png", "source": "../pages/PageFive.qml", "showCounter":false, "showMarker":true, "a_p":1},
        {"type": "../navigation/DrawerDivider.qml", "name": "", "icon": "", "source": "", "a_p":1},
        {"type": "../navigation/DrawerNavigationButton.qml", "name": "Settings", "icon": "settings.png", "source": "../pages/SettingsPage.qml", "showCounter":false, "showMarker":false, "a_p":3},
        {"type": "../navigation/DrawerNavigationButton.qml", "name": "Color Schema", "icon": "colors.png", "source": "../pages/ColorSchemaPage.qml", "showCounter":false, "showMarker":false, "a_p":2},
        {"type": "../navigation/DrawerNavigationTextButton.qml", "name": "About this App", "icon": "", "source": "../pages/AboutPage.qml", "showCounter":false, "showMarker":false, "a_p":3}
    ]
    // Counter: Car
    // Marker: Subway, Flight
    property var navigationData: [
        {"counter":0, "marker":""},
        {},
        {},
        {"counter":3, "marker":""},
        {"counter":0, "marker":""},
        {"counter":0, "marker":"green"},
        {"counter":0, "marker":""},
        {"counter":0, "marker":"red"},
        {},
        {"counter":0, "marker":""},
        {"counter":0, "marker":""},
        {"counter":0, "marker":""}
    ]
    // Menu Button
    // plus max 4 from drawer: home, car, flight, settings
    // favoritesModel maps to index from navigationModel
    property var favoritesModel: [
        0, 3, 7, 9
    ]
    property int navigationIndex: 0
    onNavigationIndexChanged: {
        rootPane.activateDestination(navigationIndex)
    }
    property bool hideTitleBar: false
    onHideTitleBarChanged: {
        //
    }
    property bool showFavorites: true
    property bool highlightActiveNavigationButton : true

    // header only used in PORTRAIT to provide a fixed TitleBar
    header: isLandscape || hideTitleBar ? null : titleBar

    footer: showFavorites && !isLandscape && navigationBar.position == 0 ? favoritesBar : null

    Loader {
        id: titleBar
        visible: !isLandscape && !hideTitleBar
        active: !isLandscape && !hideTitleBar
        source: "navigation/DrawerTitleBar.qml"
        onLoaded: {
            if(item) {
                item.text = qsTr("From Drawer to Destinations")
            }
        }
    }
    // in LANDSCAPE header is null and we have a floating TitleBar
    Loader {
        id: titleBarFloating
        visible: isLandscape && !hideTitleBar
        anchors.top: parent.top
        anchors.left: parent.left
        // anchors.leftMargin: sideBar.width+6
        anchors.right: parent.right
        active: isLandscape && !hideTitleBar
        source: "navigation/DrawerTitleBar.qml"
        onLoaded: {
            if(item) {
                item.text = qsTr("From Drawer to Destinations")
            }
        }
    }

    // The sliding Drawer
    DrawerNavigationBar {
        id: navigationBar
    } // navigationBar
    // Bottom Toolbar (only Portrait)
    DrawerFavoritesNavigationBar {
        id: favoritesBar
        visible: showFavorites && !isLandscape && navigationBar.position == 0
    }

    // the ROOT contains always only one Page,
    // which will be replaced if root node changed
    StackView {
        id: rootPane
        focus: true
        anchors.top: isLandscape && !hideTitleBar? titleBarFloating.bottom : parent.top
        anchors.left: parent.left
        anchors.topMargin: isLandscape && !hideTitleBar? 6 : 0
        anchors.right: parent.right
        anchors.bottom: parent.bottom

        // shows a Busy indicator - probably not noticed yet
        // but in real life app loading first Page or Pane could took some time if heavy
        Loader {
            id: initialPlaceholder
            source: "../pages/InitialItemPage.qml"
            active: true
            visible: false
            onLoaded: {
                rootPane.initialItem = item
                item.init()
            }
        }

        replaceEnter: Transition {
            PropertyAnimation {
                property: "opacity"
                from: 0
                to:1
                duration: 300
            }
        }
        replaceExit: Transition {
            PropertyAnimation {
                property: "opacity"
                from: 1
                to:0
                duration: 300
            }
        }

        // support of BACK key
        // can be used from StackView pushed on ROOT (HomeNavigation) tp pop()
        // or to exit the app
        property bool firstPageInfoRead: false
        Keys.onBackPressed: {
            if(navigationIndex == 0 && destinations.itemAt(navigationIndex).item.depth > 1) {
                destinations.itemAt(navigationIndex).item.pop()
                event.accepted = true
                return
            }

            event.accepted = !firstPageInfoRead
            // user gets Popupo Info
            // hitting again BACK will close the app
            if(!firstPageInfoRead) {
                firstPageInfoRead = true
                showInfo(qsTr("Next BACK closes APP and clears Values\n\nUse 'Android Home' Button for Fast-Restart.\n\n"))
                return
            }
            // We must cleanup loaded Pages

        }

        // TODO some Shortcus

        Repeater {
            id: destinations
            model: navigationModel
            // Destination encapsulates Loader
            // depends from activationPolicy how to load dynamically
            Destination {
                id: destinationLoader
            }
            Component.onCompleted: {
                // all destinations (Loader) created
                // all destinatation items w activationPolicy IMMEDIATELY activated
                // now show first destination (should always be IMMEDIATELY)
                rootPane.activateDestination(0)
                rootPane.firstDestinationLoaded()
            }
        }
        function firstDestinationLoaded() {
            // do some sepcial stuff here
            // first page just becomes visible
        }
        // switch to new Destination
        // Destinations are lazy loaded via Loader
        function activateDestination(navigationIndex) {
            if(destinations.itemAt(navigationIndex).status == Loader.Ready) {
                console.log("replace item on root stack: "+navigationIndex)
                replaceDestination(destinations.itemAt(navigationIndex))
            } else {
                console.log("first time item to be replaced: "+navigationIndex)
                destinations.itemAt(navigationIndex).active = true
            }
        }
        // called from activeDestination() and also from Destination.onLoaded()
        function replaceDestination(theItemLoader) {
            var previousIndex = rootPane.currentItem.myIndex
            var previousItemLoader
            if(previousIndex >= 0) {
                previousItemLoader  = destinations.itemAt(previousIndex)
            }
            // because of https://bugreports.qt.io/browse/QTBUG-54260
            // remember currentIndex before being replaced
            if(rootPane.currentItem.name == "colorSchemaNavPage") {
                rootPane.currentItem.lastCurrentIndex = rootPane.currentItem.currentIndex
            }
            // reset currentIndex to the last one
            if(theItemLoader.item.name == "colorSchemaNavPage") {
                theItemLoader.item.currentIndex = theItemLoader.item.lastCurrentIndex
            }
            // now replace the Page
            rootPane.replace(theItemLoader.item)
            // check if previous should be unloaded

            if(previousIndex >= 0) {
                if(destinations.itemAt(previousIndex).pageActivationPolicy == activationPolicy.WHILE_CURRENT) {
                    destinations.itemAt(previousIndex).active = false
                }
            } else {
                initialPlaceholder.active = false
            }
        }

        // example HowTo increase a counter (visible from Drawer and Cars Page)
        function increaseCars() {
            var counter = navigationData[3].counter + 1
            navigationData[3].counter = counter
            navigationBar.navigationButtons.itemAt(3).item.counter = counter
        }

    } // rootPane



    function switchPrimaryPalette(paletteIndex) {
        primaryPalette = myApp.primaryPalette(paletteIndex)
    }
    function switchAccentPalette(paletteIndex) {
        accentPalette = myApp.accentPalette(paletteIndex)
    }

    // we can loose the focus if Menu or Popup is opened
    function resetFocus() {
        rootPane.focus = true
    }

    function showInfo(info) {
        popupInfo.text = info
        popupInfo.buttonText = qsTr("OK")
        popupInfo.open()
    }

    PopupInfo {
        id: popupInfo
        onAboutToHide: {
            popupInfo.stopTimer()
            resetFocus()
        }
    } // popupInfo

    // PopupToast
    PopupToast {
        id: popupToast
        onAboutToHide: {
            resetFocus()
        }
    }

} // app window
