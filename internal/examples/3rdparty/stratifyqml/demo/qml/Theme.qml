import QtQuick 2.6
import StratifyLabs.UI 2.0

EPane {
  name: "Theme";
  ESectionTitle { text: "Introduction"; }
  EParagraph { text: '\
SL.UI themes are based on the singleton object <i>STheme</i> \
which defines colors, fonts, spacing, and shapes for the application.';
  }
  ESectionTitle { text: "Customization"; }
  EParagraph { text: '\
The following example shows how to change the primary and \
secondary brand colors. This code typically goes in the \
main <i>ApplicationWindow</i>.';
  }
  ECodeBlockInline {
    text: '\
Component.onCompleted: {
  STheme.brand_primary = "#244E99";
  STheme.brand_secondary = "#666";
}';
  }

  ESectionTitle { text: "Reference"; }
  EParagraph { text: 'See the <i>STheme.qml</i> file for a full list of variables.'; }

  SButton {
    style: "btn-outline-info block";
    text: "View STheme.qml on Github";
    icon: Fa.Icon.external_link;
    onClicked: Qt.openUrlExternally("https://github.com/StratifyLabs/StratifyQML");
  }
}
