import QtQuick 2.6
import StratifyLabs.UI 2.0

EPane {
  name: "Wells";
  EHeading {
    title: "Wells";
    inherits: "Rectangle";
    stratifyName: "SWell";
  }
  ESectionTitle { text: "Description"; }
  EParagraph { text: 'An <i>SWell</i> is used to show a block of \
text. It responds to size changes by setting the <i>style</i> property. \
You can customize the color scheme by changing <i>attr.backgroundColor</i>, \
<i>attr.fontColor</i> and/or <i>attr.borderColor</i>.';
  }
  SHLine{}

  ESectionTitle { text: "Example"; }
  WellExample{}
  ECodeButton { source: "WellExample"; }
}
