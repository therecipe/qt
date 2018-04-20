import QtQuick 2.6
import StratifyLabs.UI 2.0

EPane {
  name: "Charts";
  EHeading {
    title: "Charts";
    inherits: "QtCharts 2.2 ChartView";
    stratifyName: "SChart";
  }
  ESectionTitle { text: "Description"; }
  EParagraph { text: '\
An <i>SChart</i> is for drawing charts and graphs.';
  }
  SHLine{}

  ESectionTitle { text: "Colors"; }
  ChartExampleColors{}
  ECodeButton { source: "ChartExampleColors"; }

  ESectionTitle { text: "Chart Types"; }
  ChartExampleTypes{}
  ECodeButton { source: "ChartExampleTypes"; }


}

