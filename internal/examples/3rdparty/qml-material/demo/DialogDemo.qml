import QtQuick 2.4
import QtQuick.Controls 1.3 as QuickControls
import Material 0.2
import Material.Extras 0.1

Item {

    Dialog {
        id: alertNoTitleBar
        width: dp(300)
        text: "Discard draft?"
        hasActions: true
        positiveButtonText: "discard"
        negativeButtonText: "cancel"
    }

    Dialog {
        id: alertWithTitleBar
        width: dp(300)
        title: "Use Google's location service?"
        text: "Let Google help apps determine location. This means sending anonymous location data to Google, even when no apps are running."
        hasActions: true
        positiveButtonText: "agree"
        negativeButtonText: "disagree"
    }

    Dialog {
        id: textFieldDialog
        title: "Change Text"
        hasActions: true

        TextField {
            id: optionText
            width: parent.width
            placeholderText: "New Option to Confirm"
        }

        onAccepted: {
            dialogSnackBar.open("You entered: %1".arg(optionText.text))
        }
    }

    Dialog {
        id: scrollingDialog
        title: "Phone ringtone"

        QuickControls.ExclusiveGroup {
            id: optionGroup
        }

        RadioButton {
            text: "None"
            checked: true
            exclusiveGroup: optionGroup
        }
        RadioButton {
            text: "Callisto"
            exclusiveGroup: optionGroup
        }
        RadioButton {
            text: "Dione"
            exclusiveGroup: optionGroup
        }
        RadioButton {
            text: "Ganymede"
            exclusiveGroup: optionGroup
        }
        RadioButton {
            text: "Hangouts Call"
            exclusiveGroup: optionGroup
        }
        RadioButton {
            text: "Luna"
            exclusiveGroup: optionGroup
        }
        RadioButton {
            text: "Oberon"
            exclusiveGroup: optionGroup
        }
        RadioButton {
            text: "Phobos"
            exclusiveGroup: optionGroup
        }
        RadioButton {
            text: "Pyxis"
            exclusiveGroup: optionGroup
        }
    }

    Column {
        anchors.centerIn: parent
        spacing: dp(20)

        Button {
            text: "Show alert without title bar"
            anchors.horizontalCenter: parent.horizontalCenter
            elevation: 1
            onClicked: {
                alertNoTitleBar.show()
            }
        }

        Button {
            text: "Show alert with title bar"
            anchors.horizontalCenter: parent.horizontalCenter
            elevation: 1
            onClicked: {
                alertWithTitleBar.show()
            }
        }

        Button {
            text: "Show text field dialog"
            anchors.horizontalCenter: parent.horizontalCenter
            elevation: 1
            onClicked: {
                textFieldDialog.show()
            }
        }

        Button {
            text: "Show scrolling dialog"
            anchors.horizontalCenter: parent.horizontalCenter
            elevation: 1
            onClicked: {
                scrollingDialog.show()
            }
        }
    }

    Snackbar {
        id: dialogSnackBar
    }
}
