import QtQuick 2.0
import QuickFlux 1.0
import QtQuick.Controls 1.0
import QtQuick.Layouts 1.0
import "../stores"

ScrollView {
    ListView {
        anchors.fill: parent

        model: TodoVisualModel {
            id: visualModel
        }
    }
}

