import QtQuick 2.0
import QtQuick.Layouts 1.1
import QtQuick.Controls 1.3
import "../stores"
import "../actions"

Rectangle {
    id: viewer
    color: "#000000"

    ColumnLayout {
        anchors.fill: parent

        Grid {
            Layout.fillWidth: true
            Layout.fillHeight: true
            columns: 3
            spacing: 0

            Repeater {
                  model: MainStore.photoModel
                  delegate: Image {
                      width: viewer.width / 3
                      height: width / 4 * 3
                      source: model.url
                      asynchronous: true
                      fillMode: Image.PreserveAspectCrop
                  }
             }
        }

        Button {
            Layout.fillWidth: true
            Layout.fillHeight: false
            text: qsTr("Pick Image")
            onClicked: {
                AppActions.pickPhoto();
            }
        }
    }

}

