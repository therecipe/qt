import QtQuick 2.0
import QuickFlux 1.0
import "../stores"
import "../actions"

VisualDataModel {

    model: MainStore.todo.model

    filterOnGroup: MainStore.userPrefs.showCompletedTasks ? "" : "nonCompleted"

    groups: [
        VisualDataGroup {
            name: "nonCompleted"
            includeByDefault: true
        }
    ]

    delegate: TodoItem {
        id:item
        uid: model.uid
        title: model.title
        checked: model.done

        Component.onCompleted: {
            item.VisualDataModel.inNonCompleted = Qt.binding(function() { return !model.done})
        }
    }
}

