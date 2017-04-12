// ekke (Ekkehard Gentz) @ekkescorner
import QtQuick 2.6
import QtQuick.Layouts 1.3
import QtQuick.Controls 2.0
import QtQuick.Controls.Material 2.0
import QtGraphicalEffects 1.0

import "../pages"

StackView {
    id: navPane
    property string name: "HomeNavPane"
    // index to get access to Loader (Destination)
    property int myIndex: index

    initialItem: HomePage{}

    Loader {
        id: qtPageLoader
        active: true
        visible: false
        source: "../pages/QtPage.qml"
    }

    function pushQtPage() {
        navPane.push(qtPageLoader.item)
    }

    Component.onDestruction: {
        cleanup()
    }

    function init() {
        console.log("INIT HomeNavPane")
    }
    function cleanup() {
        console.log("CLEANUP HomeNavPane")
    }

} // navPane
