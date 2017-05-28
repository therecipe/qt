import QtQuick 2.0
import QtQuick.Controls 1.3
import QuickFlux 1.0
import "../actions"

Item {

    StackView {
        id: stack
        anchors.fill: parent

        initialItem: ImageViewer {

        }
    }

    AppListener {
        filter: ActionTypes.navigateTo
        onDispatched: {
            stack.push(message.item,message.properties);
        }
    }

    AppListener {
        filter: ActionTypes.navigateBack
        onDispatched: {
            stack.pop();
        }
    }

}

