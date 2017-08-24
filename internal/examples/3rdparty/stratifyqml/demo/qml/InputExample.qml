import StratifyLabs.UI 2.0

SRow {
  SLabel {
    span: 2;
    text: "username";
  }
  SInput {
    span: 10;
  }
  SLabel {
    span: 2;
    style: "left";
    text: "password"
  }
  SPassword {
    span: 10;
  }
  SIcon {
    span: 2;
    icon: Fa.Icon.user;
  }
  SInput {
    span: 10;
  }
  SIcon {
    span: 2;
    icon: Fa.Icon.lock;
  }
  SPassword {
    span: 10;
  }
  SInput {
    span: 12;
    placeholder: "username";
  }
  SPassword {
    span: 12;
    placeholder: "password";
  }
}
