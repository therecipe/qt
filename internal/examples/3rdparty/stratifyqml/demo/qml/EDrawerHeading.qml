import QtQuick 2.0
import StratifyLabs.UI 2.0

SContainer {
  id: root;
  property alias label: label.text;
  property alias icon: icon.icon;

  SRow {
    SText { id: label; span: 1; style: "left text-on-info text-bold"; }
    SIcon { id: icon; span: 1; style: "right text-on-info"; }
  }
}
