import QtQuick 2.0
import StratifyLabs.UI 2.0

SPanel {
  id: panel;
  property string title;
  property string inherits;
  property string stratifyName;
  property string specialAttibutes;
  property string defaultSize;

  style: "panel-info block";
  heading: "<b>" + title + "</b>";

  SRow {
    SLabel { span: 6; style: "left"; text: "Name:"; }
    SText { span: 6; style: "right"; text: panel.stratifyName; }
    SLabel { span: 6; style: "left"; text: "Inherits:"; }
    SText { span: 6; style: "right"; text: panel.inherits; }
    //SLabel { span: 6; style: "left"; text: "Special Attributes:"; visible: panel.specialAttibutes !== ""; }
    //SText { span: 6; style: "right"; text: panel.specialAttibutes; visible: panel.specialAttibutes !== ""; }
    //SLabel { span: 6; style: "left"; text: "Default Size:"; visible: panel.defaultSize !== ""; }
    //SText { span: 6; style: "right"; text: panel.defaultSize; visible: panel.defaultSize !== ""; }
  }

}
