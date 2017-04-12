import QtQuick 2.0
import Sailfish.Silica 1.0
import "pages"

ApplicationWindow {
    id: mainWindow
    initialPage: Qt.resolvedUrl("pages/page.qml")
    allowedOrientations: Orientation.All
    _defaultPageOrientations: Orientation.All
}
