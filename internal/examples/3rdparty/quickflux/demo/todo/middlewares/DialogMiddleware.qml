import QtQuick 2.0
import QuickFlux 1.1
import QtQuick.Dialogs 1.2
import "../actions"
import "../stores"

Middleware {

    property RootStore store: MainStore

    MessageDialog {
        id: dialog
        title: "Confirmation"
        text: "Are you sure want to show completed tasks?"
        standardButtons: StandardButton.Ok | StandardButton.Cancel

        onAccepted: {
            next(ActionTypes.setShowCompletedTasks, {value: true});
        }

        onRejected: {
            /// Trigger the changed signal even it is unchanged. It forces the checkbox to be turned off.
            store.userPrefs.showCompletedTasksChanged();
        }
    }

    function dispatch(type, message) {

        if (type === ActionTypes.setShowCompletedTasks &&
            message.value === true) {
            // If user want to show completed tasks, drop the action and show a dialog
            dialog.open();
            return;
        }

        /// Pass the action to next middleware / store
        next(type, message);
    }

}
