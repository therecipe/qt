import QtQuick 2.0
import Sailfish.Silica 1.0

Page {
    SilicaListView {
        id: listView
        model: PersonModel

        anchors.fill: parent
        spacing: Theme.paddingMedium

        header: PageHeader {
            title: "Names"
        }

        delegate: Text {
            color: "white"
            text: display.firstName + "  " + display.lastName
        }
    }
}
