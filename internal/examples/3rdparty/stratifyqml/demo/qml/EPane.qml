import QtQuick 2.7
import QtQuick.Controls 2.1
import StratifyLabs.UI 2.0


SColumn {
  property string name;
  style: "block fill padding-zero";
  default property alias data: contents.data;

  function showCode(source){
    contentAnimationContainer.screen = "code";
    code.source = source;
  }

  function showContent() {
    contentAnimationContainer.screen = "content";
  }

  SAnimator {
    style: "block fill";

    id: contentAnimationContainer;

    animation: SAnimationHPush{}

    screen: "content";

    resources: [
      SContainer {
        name: "content";
        style: "fill padding-zero";
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
      },

      SContainer {
        name: "code";
        style: "fill padding-zero";
        EViewCode {
          id: code;
        }
      }

    ]
  }
}



