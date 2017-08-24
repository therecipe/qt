import QtQuick 2.6
import StratifyLabs.UI 2.0

EPane {
  name: "Panels";
  EHeading {
    title: "Panels";
    inherits: "Item";
    stratifyName: "SPanel";
    //specialAttibutes: "Heading";
  }

  ESectionTitle { text: "Description"; }
  EParagraph { text: 'An <i>SPanel</i> has a \
heading (optional), body, and footer (optional). The body is \
populated by adding objects to the default <i>data</i> variable.'; }

  SHLine{}

  ESectionTitle { text: "Colors"; }
  PanelExampleColors{}
  ECodeButton { source: "PanelExampleColors"; }

  ESectionTitle { text: "Headings and Footers"; }
  PanelExampleHeadingFooter{}
  ECodeButton { source: "PanelExampleHeadingFooter"; }
}
