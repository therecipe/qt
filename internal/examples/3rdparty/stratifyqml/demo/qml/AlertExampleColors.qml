import StratifyLabs.UI 2.0

SRow {
  SButton {
    span: 12;
    style: "block";
    text: "Show All";
    onClicked: {
      alertPrimary.open();
      alertSecondary.open();
      alertInfo.open();
      alertWarning.open();
      alertDanger.open();
      alertSuccess.open();
    }
  }
  SAlert{
    id: alertPrimary;
    span: 6;
    style: "alert-primary";
    text: "Primary";
  }
  SAlert{
    id: alertSecondary;
    span: 6;
    style: "alert-secondary";
    text: "Secondary";
  }
  SAlert{
    id: alertInfo;
    span: 6;
    style: "alert-info";
    text: "Info";
  }
  SAlert{
    id: alertWarning;
    span: 6;
    style: "alert-warning";
    text: "Warning";
  }
  SAlert{
    id: alertDanger;
    span: 6;
    style: "alert-danger";
    text: "Danger";
  }
  SAlert{
    id: alertSuccess;
    span: 6;
    style: "alert-success";
    text: "Success";
  }
}
