import QtQuick 2.6
import StratifyLabs.UI 2.0

EPane {
  name: "ProgressCircles";
  EHeading {
    title: "Progress Circles";
    inherits: "Item";
    stratifyName: "SProgressCircle";
  }

  ESectionTitle { text: "Description"; }
  EParagraph { text: 'An <i>SProgressCircle</i> shows linear \
progress but with a greater prominance to the user than an <i>SProgressBar</i>.'; }

  EReadMore {
    tags: "ProgressBars:SProgressBar";
  }

  SHLine{}

  ESectionTitle { text: "Colors"; }
  ProgressCircleExampleColors{}
  ECodeButton { source: "ProgressCircleExampleColors"; }

  ESectionTitle { text: "With an Icon"; }
  ProgressCircleExample{}
  ECodeButton { source: "ProgressCircleExample"; }
}
