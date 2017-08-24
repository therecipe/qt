import QtQuick 2.6
import StratifyLabs.UI 2.0

EPane {
  name: "Icons";
  EHeading {
    title: "Icons";
    inherits: "Item";
    stratifyName: "SIcon";
    defaultSize: "block";
  }

  ESectionTitle { text: "Description"; }
  EParagraph { text: 'An <i>SIcon</i> \
draws an optional icon with an optional label. It does not \
support text alignment using <i>text-left</i> or <i>text-right</i>.'; }
  SHLine{}

  ESectionTitle { text: "Example"; }
  IconExampleObjects{}
  ECodeButton { source: "IconExampleObjects"; }

  ESectionTitle { text: "Other Objects"; }
  IconExampleOther{}
  ECodeButton { source: "IconExampleOther"; }
}
