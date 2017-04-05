import QtQuick 2.0
import QuickFlux 1.1

Store {

    bindSource: AppDispatcher

    property alias todo: todo

    TodoStore {
        id: todo
    }

    property alias userPrefs: userPrefs

    UserPrefsStore {
        id: userPrefs
    }

}
