import QtQuick 2.0
import QuickFlux 1.1
import "../actions"

Store {
    bindSource: AppDispatcher
    property ListModel photoModel : ListModel {}

    function add(url) {
        photoModel.append({ url: url });
    }

    Filter {
        type: ActionTypes.pickPhoto
        onDispatched: {
            add(message.url);
        }
    }
}

