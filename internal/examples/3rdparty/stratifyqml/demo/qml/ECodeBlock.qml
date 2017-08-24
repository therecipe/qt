import QtQuick 2.0
import QtQuick.Controls 2.1
import StratifyLabs.UI 2.0

STextBox {
  id: code;
  style: "sm block";

  implicitHeight: 50;
  attr.fontText: STheme.font_family_monospace.name;

  textArea.selectByMouse: {
    (Qt.platform.os === "ios") || (Qt.platform.os === "ios") ?
       false : true;
  }
  textArea.readOnly: true;
  horizontalScroll.visible: true;
  verticalScroll.visible: true;
  textArea.wrapMode: TextArea.NoWrap;


}
