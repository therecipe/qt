pragma Singleton
import QtQuick 2.0
import QuickFlux 1.0
import "./"

ActionCreator {

    /* Create action by signal */

    signal pickPhoto()

    signal previewPhoto(string url);

    /*

    signal pickPhoto(url)

    signal previewPhoto(string url);

    They are equivalent to:

    function pickPhoto(url) {
        dispatch("pickPhoto", {url: url});
    }

    function previewPhoto(url) {
        dispatch(ActionTypes.previewPhoto,{url: url});
    }

    */

    /* Create action by traditional method */


    function navigateTo(item,properties) {
        dispatch(ActionTypes.navigateTo,
                               {item: item,
                                properties: properties});
    }

    function navigateBack() {
        dispatch(ActionTypes.navigateBack);
    }
}

