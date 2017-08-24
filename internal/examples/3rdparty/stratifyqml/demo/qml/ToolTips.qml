import QtQuick 2.6
import StratifyLabs.UI 2.0

EPane {
  name: "ToolTips";
  EHeading {
    title: "Tool Tips";
    inherits: "QtQuick.Controls 2.0 ToolTip";
    stratifyName: "SToolTip";
  }

  ESectionTitle { text: "Description"; }
  EParagraph { text: 'An <i>SToolTip</i> can be added to clickable \
objects and will display a tool tip when hovered. <i>SToopTip</i> is not \
currently implemented on touch interfaces.';
  }

  SHLine{}

  ESectionTitle { text: "Example"; }
  ToolTipExample{}
  ECodeButton { source: "ToolTipExample"; }
}
