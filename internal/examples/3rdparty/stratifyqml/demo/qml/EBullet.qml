import StratifyLabs.UI 2.0

SGroup {
  property alias text: text.text;
  style: "left";
  SIcon {
    id: bullet;
    attr.fontSize: 8;
    icon: Fa.Icon.circle;
  }
  SText {
    anchors.verticalCenter: bullet.verticalCenter;
    id: text;
    text: "padding-zero";
  }
}
