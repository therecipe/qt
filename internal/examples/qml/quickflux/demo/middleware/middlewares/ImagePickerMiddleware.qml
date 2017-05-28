import QtQuick 2.0
import QtQuick.Dialogs 1.0
import QuickFlux 1.1
import "../actions"
import "../components"

Middleware {

    filterFunctionEnabled: true

    Component {
        id: imagePreview

        ImagePreview {
            onCancelled: {
                next(ActionTypes.navigateBack);
            }

            onConfirmed: {
                next(ActionTypes.navigateBack);
                next(ActionTypes.pickPhoto, {url: dialog.fileUrl.toString()});
            }
        }
    }

    FileDialog {
        id: dialog
        title: qsTr("Pick Image")
        nameFilters: [ "Image files (*.jpg *.png)"]
        onAccepted: {
            var message = {
                item: imagePreview,
                properties: {
                    source: dialog.fileUrl
                }
            }
            next(ActionTypes.navigateTo, message);
        }
    }

    function pickPhoto(message) {
        dialog.open();
    }


}
