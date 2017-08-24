import QtQuick 2.0
import StratifyLabs.UI 2.0

SButton {
  id: root;
  property string screenName;
  style: "sm btn-outline-info";
  onClicked: {
    animationContainer.screen = root.screenName;
  }
}
