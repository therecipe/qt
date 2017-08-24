import StratifyLabs.UI 2.0

SColumn {
  clip: true;

  attr.paddingVertical: 0;
  SContainer {
    style: "padding-zero";
    color: STheme.brand_success;
    SRow {
      SButton {
        span: 1;
        style: "btn-naked text-on-success left";
        icon: Fa.Icon.chevron_left;
        onClicked: {
          animationContainer.screen = "Screen1";
        }
      }

      SButton {
        span: 1;
        style: "btn-naked text-on-success right";
        icon: Fa.Icon.chevron_right;
        onClicked: {
          animationContainer.screen = "Screen2";
        }
      }
    }
  }

  SAnimator {
    id: animationContainer;
    implicitHeight: 300;

    animation: SAnimationHPush{ }
    screen: "Screen1";

    resources: [
      SContainer {
        name: "Screen1";
        style: "fill";
        color: STheme.brand_danger;
        SColumn {
          style: "block fill";
          SLabel {
            style: "text-on-danger";
            text: "Screen One";
          }
        }
      },
      SContainer {
        name: "Screen2";
        style: "fill";
        color: STheme.brand_warning;
        SColumn {
          style: "block fill";
          SLabel {
            style: "text-on-warning";
            text: "Screen Two";
          }
        }
      }
    ]
  }
}
