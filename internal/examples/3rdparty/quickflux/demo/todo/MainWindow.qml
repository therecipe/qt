import QtQuick 2.3
import QtQuick.Window 2.2
import QtQuick.Layouts 1.0
import QuickFlux 1.1
import "./views"
import "./middlewares"
import "./actions"

Window {
    width: 480
    height: 640
    visible: true

    ColumnLayout {
        anchors.fill: parent
        anchors.leftMargin: 16
        anchors.rightMargin: 16

        Header {
            Layout.fillWidth: true
            Layout.fillHeight: false
        }

        TodoList {
            Layout.fillWidth: true
            Layout.fillHeight: true
        }

        Footer {
            Layout.fillWidth: true
            Layout.fillHeight: false
        }
    }
}

