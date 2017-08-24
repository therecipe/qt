import QtQuick 2.6
import StratifyLabs.UI 2.0

EPane {
  name: "CheckBoxes";
  EHeading {
    title: "Check Boxes";
    inherits: "QtQuick.Controls 2.0 CheckBox";
    stratifyName: "SCheckBox";
  }
  ESectionTitle { text: "Description"; }
  EParagraph { text: 'An <i>SCheckbox</i> can be used for non-exclusive \
boolean input items. Use <i>SRadioButton</i> for exclusive boolean \
input items.';
  }

  EReadMore {
    tags: "RadioButtons:SRadioButton";
  }

  SHLine{}

  ESectionTitle { text: "Example"; }
  CheckBoxExample{}
  ECodeButton { source: "CheckBoxExample"; }

  ESectionTitle { text: "Custom Icons"; }
  CheckBoxExampleCustom{}
  ECodeButton { source: "CheckBoxExampleCustom"; }
}
