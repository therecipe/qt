import StratifyLabs.UI 2.0

SRow {
  SButton {
    span: 4;
    style: "btn-primary";
    text: "Button";
    SToolTip {
      text: "Button";
    }
  }
  SCheckBox {
    span: 4;
    text: "Check Box";
    SToolTip {
      text: "Check Box";
    }
  }
  SRadioButton {
    span: 4;
    text: "Radio Button";
    SToolTip {
      text: "Radio Button";
    }
  }
  SInput {
    span: 12;
    placeholder: "placeholder";
    SToolTip {
      text: "Input Box";
    }
  }

  SDropdown {
    style: "block";
    span: 12;
    model:
        ["First",
      "Second",
      "Third",
      "Fourth"];
    SToolTip {
      text: "Dropdown";
    }
  }
}
