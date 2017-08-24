import QtQuick 2.7
import StratifyLabs.UI 2.0

SRow {

  SText {
    span: 4;
    style: "left text-h1 top";
    text: "List";
  }

  SList {
    id: list;
    span: 8;
    style: "block right padding-zero";
    model: SJsonModel {
      id: model;
      json: "{ \"data\": [{ \"text\": \"List Item 1\" },
                { \"text\": \"List Item 2\" },
                { \"text\": \"List Item 3\" },
                { \"text\": \"List Item 4\" } ] }";
    }

    delegate: Component {
      SContainer {
        style: "padding-zero";
        color:
            model.index === list.currentIndex ?
              STheme.brand_secondary :
              Qt.lighter(STheme.brand_secondary, 2.0);
        SRow {
          SLabel {
            style: "left text-on-secondary";
            span: 1;
            text: model.text;
          }
          SIcon {
            span: 1;
            style: "right text-on-secondary";
            icon: Fa.Icon.chevron_right;
          }
        }
        MouseArea {
          anchors.fill: parent;
          onClicked: {
            list.currentIndex = model.index;
          }
        }
      }
    }
  }
}
