import QtQuick 2.0
import QtQuick.Controls 1.0
import "../actions"
import "../stores"

Item {
    height: 48

    CheckBox {
        id: checkBox
        checked: false
        text: "Show Completed";
        anchors.right: parent.right
        anchors.verticalCenter: parent.verticalCenter

        onCheckedChanged: {
            AppActions.setShowCompletedTasks(checked);
        }

        Binding {
            target: checkBox
            property: "checked"
            value: MainStore.userPrefs.showCompletedTasks
        }
    }

}

