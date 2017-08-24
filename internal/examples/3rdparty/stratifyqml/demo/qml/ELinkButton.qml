import QtQuick 2.7
import QtQuick.Controls 2.1
import StratifyLabs.UI 2.0

SAnimationContainer {
  property string name;
  style: "fill padding-zero";
  default property alias data: contents.data;
  SPane {
    clip: false;
    ScrollBar.vertical: ScrollBar {
      visible: !STheme.isScreenSm;
    }
    SContainer {
      SColumn {
        id: contents;
      }
    }
  }
}
