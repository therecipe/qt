import QtQuick 2.6
import StratifyLabs.UI 2.0

SColumn {
  SLabel {
    id: menuLabel;
    style: "label-info";
    text: "Nothing yet";
  }

  SHLine {}
  SContainer {
    SRow {
      SLabel {
        span: 1;
        style: "label-danger left";
        text: "Menu Item 1";
      }
      SIcon {
        span: 1;
        style: "right";
        icon: Fa.Icon.chevron_right;
      }
    }
    MouseArea {
      anchors.fill: parent;
      onClicked: menuLabel.text = "Menu Item 1";
    }
  }
  SHLine {}
  SContainer {
    SRow {
      SLabel {
        span: 1;
        style: "label-success left";
        text: "Menu Item 2";
      }
      SIcon {
        span: 1;
        style: "right";
        icon: Fa.Icon.chevron_right;
      }
    }
    MouseArea {
      anchors.fill: parent;
      onClicked: menuLabel.text = "Menu Item 2";
    }
  }
  SHLine {}

}

