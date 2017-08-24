import QtQuick 2.6
import StratifyLabs.UI 2.0

EPane {
  name: "TextBoxes";
  EHeading {
    title: "Text Boxes";
    inherits: "Rectangle (Flickable/TextArea)";
    stratifyName: "STextBox";
    defaultSize: "block";
  }

  SLabel { style: "left"; text: "Text Box Examples"; }
  TextBoxExample{}
  ECodeButton { source: "TextBoxExample"; }
}
