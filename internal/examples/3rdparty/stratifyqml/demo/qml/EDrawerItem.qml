import QtQuick 2.6
import StratifyLabs.UI 2.0

SContainer {
  id: root;
  property alias label: label.text;

  property bool active: false;
  property alias clickable: mouseArea.enabled;

  color: active ? Qt.darker(STheme.brand_info, 1.3) : "transparent";

  SRow {
    SText { id: label; span: 12; style: "left text-on-info"; }
  }
  MouseArea {
    id: mouseArea;
    enabled: true;
    anchors.fill: parent;
    onClicked: {
      drawer.shut(root);
      screen = root.label;
    }
  }
}
