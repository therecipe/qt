import QtQuick 2.7
import QtQuick.Controls 2.1
import StratifyLabs.UI 2.0

SButton {
  property string source;
  style: "btn-outline-success lg block";
  text: "View Code";
  icon: Fa.Icon.code;
  onClicked: {
    showCode(source);
  }
}
