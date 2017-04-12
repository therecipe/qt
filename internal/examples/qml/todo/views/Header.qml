import QtQuick 2.0
import QtQuick.Controls 1.0
import "../actions"

Item {
    height: 48

    CheckBox {
        checked: false
        text: "Show Completed";
        anchors.right: parent.right
        anchors.verticalCenter: parent.verticalCenter

        onCheckedChanged: {
            AppActions.setShowCompletedTasks(checked);
        }
    }

}

