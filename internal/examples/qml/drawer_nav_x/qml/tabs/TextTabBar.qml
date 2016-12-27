// ekke (Ekkehard Gentz) @ekkescorner
import QtQuick 2.6
import QtQuick.Layouts 1.3
import QtQuick.Controls 2.0
import QtQuick.Controls.Material 2.0
import QtGraphicalEffects 1.0
import "../common"

TabBar {
    id: myTabBar
    Layout.fillWidth: true
    currentIndex: 0
    onCurrentIndexChanged: {
        console.log("Tab Bar current index changed: "+ currentIndex)
        navPane.currentIndex = currentIndex
    }
    Repeater {
        model: tabButtonModel
        TabButton {
            focusPolicy: Qt.NoFocus
            text: modelData.name
            width: tabBarIsFixed? myTabBar.width / tabButtonModel.length  : Math.max(112, myTabBar.width / tabButtonModel.length)
        }
    } // repeater

}
