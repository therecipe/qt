import QtQuick 2.7            //Image
import QtQuick.Layouts 1.3    //RowLayout
import QtQuick.Controls 2.1   //Button
import QtGraphicalEffects 1.0 //ColorOverlay

import Theme 1.0              //Theme

Button {
  id: button

  property string code
  property alias image : logo.source
  property QtObject template
  property bool isCheckable : false

  background: null

  contentItem: RowLayout {
    spacing: isCheckable ? 0 : 20

    Item {
      Layout.fillWidth: true

      Image {
        id: logo
        anchors.centerIn: parent
      }

      ColorOverlay {
        anchors.fill: logo
        source: logo
        color: button.down || button.checked || button.hovered ? Theme.accent : Theme.font
      }
    }

    Label {
      text: button.text
      color: button.down || button.checked || button.hovered ? Theme.fontHighlight : Theme.font
      font.pointSize: 16
    }

    Item {
      Layout.fillWidth: true
    }
  }

  onClicked: {
    checked = isCheckable
    template.clicked(code)
  }
}
