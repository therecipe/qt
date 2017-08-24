import StratifyLabs.UI 2.0

SRow {
  SLabel {
    span: 2;
    text: "1 Option";
  }
  SDropdown {
    style: "block";
    span: 10;
    model: ["First"];
  }
  SLabel {
    span: 2;
    text: "2 Options";
  }
  SDropdown {
    style: "block";
    span: 10;
    model:
        ["First",
      "Second"];
  }
  SLabel {
    span: 2;
    text: "3 Options";
  }
  SDropdown {
    style: "block";
    span: 10;
    model:
        ["First",
      "Second",
      "Third"
    ];
  }
  SLabel {
    span: 2;
    text: "4 Options";
  }
  SDropdown {
    style: "block";
    span: 10;
    model:
        ["First",
      "Second",
      "Third",
      "Fourth"];
  }
}
