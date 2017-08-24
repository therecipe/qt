import StratifyLabs.UI 2.0

SRow {
  SButton {
    span: 12;
    style: "block";
    text: "Show All";
    onClicked: {
      dismissibleAlert.open();
    }
  }
  SAlert {
    span: 6;
    text: "Non-dismissible alert";
    attr.dismissible: false;
  }
  SAlert {
    id: dismissibleAlert;
    span: 6;
    text: "Dismissible alert (default)";
    attr.dismissible: true;
  }
}
