import QtQuick 2.6
import StratifyLabs.UI 2.0

EPane {
  name: "Dropdowns";
  EHeading {
    title: "Dropdown";
    inherits: "QtQuick.Controls 2.0 ComboBox";
    stratifyName: "SDropdown or SDropup";
    defaultSize: "block";
  }

  ESectionTitle { text: "Description"; }
  EParagraph { text: 'An <i>SDropdown</i> allows \
the user to select from a list of items. It can be populated \
using an <i>SJsonModel</i> object or simply a QML <i>ListModel</i>.'; }

  EReadMore {
    tags: "JsonModels:SJsonModel";
  }

  SHLine{}

  ESectionTitle { text: "Example"; }
  DropdownExample{}
  ECodeButton { source: "DropdownExample"; }

  ESectionTitle { text: "Dropup"; }
  SText {
    style: "block";
    text: 'You can use "SDropup" if the combo box is at the bottom of the window.'
  }
  ECodeButton { source: "DropupExample"; }
  DropupExample{}

}
