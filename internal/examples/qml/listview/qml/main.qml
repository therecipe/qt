import QtQuick 2.2
import QtQuick.Controls 2.0
import QtQuick.Layouts 1.3

ApplicationWindow {
    visible: true

    ListView {
        id: listView
        model: PersonModel
        delegate: Text {
            text: firstName + "  " + lastName
        }
    }
}
