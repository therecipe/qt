// ekke (Ekkehard Gentz) @ekkescorner
import QtQuick 2.6
import QtQuick.Layouts 1.3
import QtQuick.Controls 2.0
import QtQuick.Controls.Material 2.0
import QtGraphicalEffects 1.0

// Raised Button
Button {
    id: button
    // default: textOnPrimary
    property alias textColor: buttonText.color
    // default: primaryColor
    property alias buttonColor: buttonBackground.color
    focusPolicy: Qt.NoFocus
    Layout.fillWidth: true
    Layout.preferredWidth : 1
    leftPadding: 6
    rightPadding: 6
    contentItem: Text {
        id: buttonText
        text: button.text
        opacity: enabled ? 1.0 : 0.3
        color: textOnPrimary
        horizontalAlignment: Text.AlignHCenter
        verticalAlignment: Text.AlignVCenter
        elide: Text.ElideRight
        font.capitalization: Font.AllUppercase
    }
    background:
        Rectangle {
        id: buttonBackground
        implicitHeight: 48
        color: primaryColor
        radius: 2
        opacity: button.pressed ? 0.75 : 1.0
        layer.enabled: true
        layer.effect: DropShadow {
            verticalOffset: 2
            horizontalOffset: 1
            color: dropShadow
            samples: button.pressed ? 20 : 10
            spread: 0.5
        }
    } // background
} // button
