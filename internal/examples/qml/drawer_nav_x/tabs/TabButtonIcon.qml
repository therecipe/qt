// ekke (Ekkehard Gentz) @ekkescorner
import QtQuick 2.6
import QtQuick.Layouts 1.3
import QtQuick.Controls 2.0
import QtQuick.Controls.Material 2.0
import QtGraphicalEffects 1.0

TabButton {
    property color theButtonColor: accentColor
    property string theIconFolder: iconFolder
    property alias hasOverlay: colorOverlay.visible
    property real theOpacity: 1.0
    focusPolicy: Qt.NoFocus
    height: 48
    contentItem:
        Item {
        Image {
            id: contentImage
            anchors.centerIn: parent
            horizontalAlignment: Image.AlignHCenter
            verticalAlignment: Image.AlignVCenter
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
    } // item
}
