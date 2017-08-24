import QtQuick 2.6
import StratifyLabs.UI 2.0

EPane {
  name: "Alerts";
  EHeading {
    title: "Alerts";
    inherits: "QtQuick.Controls 2.0 Rectangle";
    stratifyName: "SAlert";
    defaultSize: "block";
  }

  ESectionTitle { text: "Description"; }
  EParagraph { text: '\
An <i>SAlert</i> is used to alert the user of \
new information. The <i>warning</i>, <i>danger</i>, \
and <i>success</i> branding colors are useful \
in indicating context of the information.';
  }
  SHLine{}

  ESectionTitle { text: "Colors"; }
  AlertExampleColors{}
  ECodeButton { source: "AlertExampleColors"; }

  ESectionTitle { text: "Options"; }
  EParagraph { text: '\
By default, an <i>SAlert</i> has a dismissible button. \
The button can be disabled by setting <i>attr.dissmisible</i> \
to false.';
  }
  AlertExampleOptions{}
  ECodeButton { source: "AlertExampleOptions"; }
}


