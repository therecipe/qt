import QtQuick 2.6
import StratifyLabs.UI 2.0

EPane {
  name: "RadioButtons";
  EHeading {
    title: "Radio Button";
    inherits: "QtQuick.Controls 2.0 RadioButton";
    stratifyName: "SRadioButton";
    specialAttibutes: "icon,iconChecked";
  }

  ESectionTitle { text: "Description"; }

  SHLine{}

  ESectionTitle { text: "Sizes"; }
  RadioButtonExample{}
  ECodeButton { source: "RadioButtonExample"; }

  ESectionTitle { text: "Custom Icons"; }
  RadioButtonExampleCustom{}
  ECodeButton { source: "RadioButtonExampleCustom"; }
}
