import StratifyLabs.UI 2.0

SRow {
  SLabel {
    span: 2;
    text: "4 Options";
  }
  SDropup {
    style: "block";
    span: 10;
    model:
        ["First",
      "Second",
      "Third",
      "Fourth"];
  }
}
