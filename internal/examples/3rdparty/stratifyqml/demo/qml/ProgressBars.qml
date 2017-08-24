import QtQuick 2.6
import StratifyLabs.UI 2.0

EPane {
  name: "ProgressBars";
  EHeading {
    title: "Progress Bars";
    inherits: "QtQuick.Controls 2.0 ProgressBar";
    stratifyName: "SProgressBar";
  }

  ESectionTitle { text: "Description"; }
  EParagraph { text: 'An <i>SProgressBar</i> can show \
linear progress (not used for indeterminate progress yet).'; }

  SHLine{}

  ESectionTitle { text: "Colors"; }
  ProgressBarExampleColors{}
  ECodeButton { source: "ProgressBarExampleColors"; }

  ESectionTitle { text: "Sizes"; }
  ProgressBarExampleSizes{}
  ECodeButton { source: "ProgressBarExampleSizes"; }
}
