import QtQuick 2.7
import StratifyLabs.UI 2.0

SRow {
  SText {
    span: 4;
    style: "left text-h1 top";
    text: "List";
  }

  SListGroup {
    span: 8;
    style: "block right";
    model: SJsonModel {
      id: model;
      json: "{ \"data\": [{ \"text\": \"List Item 1\" },
                { \"text\": \"List Item 2\" },
                { \"text\": \"List Item 3\" },
                { \"text\": \"List Item 4\" } ] }";
    }
  }
}
