import QtQuick 2.4
import QtQuick.Layouts 1.1
import Material 0.2

Item {
    implicitHeight: column.height

    ColumnLayout {
        id: column
        anchors.centerIn: parent
        spacing: dp(32)

        TextField {
            text: "Big Field with text"
            font.pixelSize: dp(32)
            anchors.horizontalCenter: parent.horizontalCenter
        }

        TextField {
            placeholderText: "Search..."
            anchors.horizontalCenter: parent.horizontalCenter
        }

        TextField {
            text: "Text under label"
            placeholderText: "Floating label"
            floatingLabel: true
            anchors.horizontalCenter: parent.horizontalCenter
        }

        TextField {
            placeholderText: "Character limit"
            floatingLabel: true
            characterLimit: 10
            anchors.horizontalCenter: parent.horizontalCenter
        }

        TextField {
            id: passwordField
            placeholderText: "Password"
            floatingLabel: true
            echoMode: TextInput.Password
            helperText: "Hint: It's not password."
            anchors.horizontalCenter: parent.horizontalCenter

            onAccepted: {
                if (passwordField.text === "password") {
                    passwordField.helperText = "Told ya."
                    passwordField.hasError = true
                }
            }
        }
    }
}
