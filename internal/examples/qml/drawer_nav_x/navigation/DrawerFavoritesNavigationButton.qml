// ekke (Ekkehard Gentz) @ekkescorner
import QtQuick 2.6
import QtQuick.Layouts 1.3
import QtQuick.Controls 2.0
import QtQuick.Controls.Material 2.0
import QtGraphicalEffects 1.0
import "../common"

ToolButton {
    id: myButton
    property bool isActive: modelData == navigationIndex
    Layout.alignment: Qt.AlignHCenter
    focusPolicy: Qt.NoFocus
    implicitHeight: 56
    implicitWidth: (myBar.width - 56 - 6) / (favoritesModel.length)
    Column {
        spacing: 0
        topPadding: 0
        anchors.horizontalCenter: parent.horizontalCenter
        anchors.verticalCenter: parent.verticalCenter
        Item {
            anchors.horizontalCenter: parent.horizontalCenter
            width: 24
            height: 24
            Image {
                id: contentImage
                width: 24
                height: 24
                verticalAlignment: Image.AlignTop
                anchors.horizontalCenter: parent.horizontalCenter
                source: "qrc:/images/"+iconFolder+"/"+ navigationModel[modelData].icon
                opacity: isActive? myBar.activeOpacity : myBar.inactiveOpacity
            }
            ColorOverlay {
                id: colorOverlay
                visible: myButton.isActive
                anchors.fill: contentImage
                source: contentImage
                color: primaryColor
            }
        } // image and coloroverlay
        Label {
            anchors.horizontalCenter: parent.horizontalCenter
            text: navigationModel[modelData].name
            opacity: isActive? 1.0 : 0.7
            color: isActive? primaryColor : flatButtonTextColor
            font.pixelSize: fontSizeActiveNavigationButton
        } // label
    } // column
    onClicked: {
        navigationIndex = modelData
    }
} // myButton
