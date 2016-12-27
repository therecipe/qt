// ekke (Ekkehard Gentz) @ekkescorner
import QtQuick 2.6
import QtQuick.Layouts 1.3
import QtQuick.Controls 2.0
import QtQuick.Controls.Material 2.0
import QtGraphicalEffects 1.0
import "../common"

Drawer {
    id: myBar
    z: 1
    leftPadding: 0
    property alias navigationButtons: navigationButtonRepeater
    property real activeOpacity: iconFolder == "black" ?  0.87 : 1.0
    property real inactiveOpacity: iconFolder == "black" ?  0.56 : 0.87 //  0.26 : 0.56
    width: Math.min(240,  Math.min(appWindow.width, appWindow.height) / 3 * 2 )
    height: appWindow.height

    Flickable {
        contentHeight: myButtons.height
        anchors.fill: parent
        clip: true

        ColumnLayout {
            id: myButtons
            focus: false
            anchors.left: parent.left
            anchors.right: parent.right
            spacing: 0
            Item {
                anchors.left: parent.left
                anchors.right: parent.right
                height: 120
                Rectangle {
                    anchors.fill: parent
                    color: primaryColor
                }
                Item {
                    anchors.left: parent.left
                    anchors.right: parent.right
                    anchors.bottom: parent.bottom
                    height: 56
                    Column {
                        anchors.verticalCenter: parent.verticalCenter
                        leftPadding: 16
                        rightPadding: 16
                        LabelBody {
                            text: "ekke's Example APPs"
                            font.weight: Font.Medium
                            color: textOnPrimary
                        }
                        LabelBody {
                            text: "@ekkescorner"
                            color: textOnPrimary
                        }
                    }
                }
                Item {
                    // space between content - see google material guide
                    height: 8
                }
                Item {
                    width: 48
                    height: 48
                    x: 16
                    y: 12
                    IconColored {
                        imageName: "directions.png"
                        imageSize: 48
                        color: textOnPrimary
                    }
                }
            }
            Item {
                // space between content - see google material guide
                height: 8
            }
            Repeater {
                id: navigationButtonRepeater
                model: navigationModel
                Loader {
                    Layout.fillWidth: true
                    source: modelData.type
                    active: true
                }
            } // repeater
            //
        } // ColumnLayout myButtons
        ScrollIndicator.vertical: ScrollIndicator { }

    } // Flickable

} // drawer
