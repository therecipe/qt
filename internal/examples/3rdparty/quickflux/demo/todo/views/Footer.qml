import QtQuick 2.0
import QtQuick.Controls 1.0
import QtQuick.Layouts 1.0
import "../actions"

Item {
    height: 56

    function add() {
        AppActions.addTask(textField.text);
        textField.text = "";
    }

    RowLayout {
        anchors.fill: parent

        TextField {
            id: textField
            Layout.fillWidth: true
            focus: true
            onAccepted: add();
        }

        Button {
            text: "ADD"
            onClicked: {
                add();
            }
        }
    }
}

