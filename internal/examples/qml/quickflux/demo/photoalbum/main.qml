import QtQuick 2.3
import QtQuick.Window 2.2
import QtQuick.Controls 1.3
import QuickFlux 1.0
import "./views"
import "./scripts"
import "./actions"

Window {
    visible: true
    title: "Photo Album"
    width: 640
    height: 480

    StackView {
        anchors.fill: parent
    }

    ImagePickerScript {
    }

}

