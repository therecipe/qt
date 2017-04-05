/** Todo Item Store

  A centralized data store of Todo list.

 */
import QtQuick 2.0
import QuickFlux 1.1
import "../actions"

Store {
    property alias model: model

    property int nextUid: 4;

    ListModel {
        // Initial data
        id: model

        ListElement {
            uid: 0
            title: "Task Zero"
            done: true
        }

        ListElement {
            uid: 1
            title: "Task A"
            done: false
        }

        ListElement {
            uid: 2
            title: "Task B"
            done: false
        }

        ListElement {
            uid: 3
            title: "Task C"
            done: false
        }
    }

    Filter {
        // Filter - Add a filter rule to parent

        // Filter item listens on parent's dispatched signal,
        // if a dispatched signal match with its type, it will
        // emit its own "dispatched" signal. Otherwise, it will
        // simply ignore the signal.

        type: ActionTypes.addTask
        onDispatched: {
            var item = {
                uid: nextUid++,
                title: message.title,
                done: false
            }
            model.append(item);
        }
    }

    Filter {
        type: ActionTypes.setTaskDone
        onDispatched: {
            for (var i = 0 ; i < model.count ; i++) {
                var item  = model.get(i);
                if (item.uid === message.uid) {
                    model.setProperty(i,"done",message.done);
                    break;
                }
            }
        }
    }

}



