// ekke (Ekkehard Gentz) @ekkescorner
import QtQuick 2.6
import QtQuick.Layouts 1.3
import QtQuick.Controls 2.0

// special Loader to load Destinations:
// Pages, Panes, StackView, SwipeView, TabBar
// loaded from Navigation at root: Bottom/Side Navigation or Drawer
Loader {
    id: pageLoader
    property int pageActivationPolicy: modelData.a_p
    active: pageActivationPolicy == activationPolicy.IMMEDIATELY
    // Loader itself is invisible - Item will be pushed on stack
    // or used by SwipeView
    visible: false

    source: modelData.source
    onLoaded: {
        item.init()
        if(pageActivationPolicy != activationPolicy.IMMEDIATELY) {
            rootPane.replaceDestination(pageLoader)
        }
    }
}
