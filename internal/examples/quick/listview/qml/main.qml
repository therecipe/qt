import QtQuick 2.0

ListView {
    id: listView
    model: PersonModel
    delegate: Text {
        text: firstName + "  " + lastName
    }
}
