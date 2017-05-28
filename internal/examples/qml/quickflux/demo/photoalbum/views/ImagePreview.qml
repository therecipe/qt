import QtQuick 2.0
import QtQuick.Layouts 1.1
import QtQuick.Controls 1.2
import QuickFlux 1.0
import "../actions"

Rectangle {
    color: "#000000"

    property alias source: image.source

    ColumnLayout {
        anchors.fill: parent

        Image {
            id: image
            asynchronous: true
            Layout.fillWidth: true
            Layout.fillHeight: true

            fillMode: Image.PreserveAspectFit
        }

        RowLayout {
            Layout.fillWidth: true
            Layout.fillHeight: false
            Layout.maximumHeight: cancelButton.implicitHeight

            Button {
                id: cancelButton
                text: qsTr("Cancel")
                onClicked: {
                    AppActions.navigateBack();
                }
            }

            Button {
                id: confirmButton
                text: qsTr("Confirm");
                onClicked: {
                    AppActions.pickPhoto(source);
                }
            }

        }
    }

}

