import QtQuick 2.0
import QtQuick.Layouts 1.3
import QtQuick.Controls 2.0
import QtWebEngine 1.1

Rectangle {
  width: 400
  height: 400

  ColumnLayout {
    anchors.fill: parent

    WebEngineView {
      id: webengineview
      Layout.fillWidth: true
      Layout.fillHeight: true

      url: "http://www.qt.io"
    }

    Button {
      Layout.fillWidth: true
      Layout.preferredHeight: 50

      text: "collect cookies"
      onClicked: bridge.sendProfile(webengineview.profile)
    }
  }
}
