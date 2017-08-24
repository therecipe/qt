import StratifyLabs.UI 2.0


SColumn {
  id: object;
  property bool isSpinning: true;

  SIcon {
    attr.spin: object.isSpinning;
    icon: Fa.Icon.location_arrow;
  }
  SButton {
    text:
        isSpinning ?
        "Stop Spinning" :
        "Start Spinning";
    onClicked: {
      isSpinning = !isSpinning;
    }
  }
}
