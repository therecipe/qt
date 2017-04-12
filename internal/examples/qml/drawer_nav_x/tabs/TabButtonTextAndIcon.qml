// ekke (Ekkehard Gentz) @ekkescorner
import QtQuick 2.6
import QtQuick.Layouts 1.3
import QtQuick.Controls 2.0
import QtQuick.Controls.Material 2.0
import QtGraphicalEffects 1.0
import "../common"

TabButton {
    property color theButtonColor: accentColor
    property string theIconFolder: iconFolder
    property alias hasOverlay: colorOverlay.visible
    property real theOpacity: 1.0
    focusPolicy: Qt.NoFocus
    height: 72
    contentItem:
        Item {
        Image {
            id: contentImage
            // anchors.centerIn: parent
            anchors.top: parent.top
            anchors.horizontalCenter: parent.horizontalCenter
            // y: 12
            horizontalAlignment: Image.AlignHCenter
            source: "qrc:/images/"+theIconFolder+"/"+modelData.icon
            opacity: colorOverlay.visible? 1.0 : theOpacity
        }
        ColorOverlay {
            id: colorOverlay
            visible: true
            anchors.fill: contentImage
            source: contentImage
            color: index == navPane.currentIndex ? theButtonColor : Qt.lighter(theButtonColor, 1.35)

        }
        LabelBody {
            id: myLabel
            topPadding: 6
            anchors.top: contentImage.bottom
            anchors.left: parent.left
            anchors.right: parent.right
            horizontalAlignment: Text.AlignHCenter
            font.capitalization: Font.AllUppercase
            font.weight: Font.Medium
            color: theButtonColor
            opacity: theOpacity
            text: modelData.name
        }
    } // content item



}
