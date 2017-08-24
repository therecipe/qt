import QtQuick 2.6
import StratifyLabs.UI 2.0

EPane {
  name: "Labels";
  EHeading {
    title: "Labels";
    inherits: "QtQuick.Controls 2.0 Label";
    stratifyName: "SLabel";
  }
  ESectionTitle { text: "Description"; }
  EParagraph { text: 'An <i>SLabel</i> is a \
text label that can be used in forms or just for labeling \
information given to the user.'; }

  SHLine{}

  ESectionTitle { text: "Colors"; }
  LabelExampleColors{}
  ECodeButton { source: "LabelExampleColors"; }

  ESectionTitle { text: "Sizes"; }
  LabelExampleSizes{}
  ECodeButton { source: "LabelExampleSizes"; }

  SLabel { style: "left"; text: "Naked Labels:"; }
  LabelExampleNaked{}
  ECodeButton { source: "LabelExampleNaked"; }

}
