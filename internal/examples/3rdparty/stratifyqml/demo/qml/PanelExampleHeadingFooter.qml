import StratifyLabs.UI 2.0

SRow {
  SPanel {
    span: 4;
    SColumn {
      SText {
        style: "block";
        text: "No heading and no footer";
      }
    }
  }
  SPanel {
    span: 4;
    footer: "Panel Footer";
    SColumn {
      SText {
        style: "block";
        text: "Footer and no heading";
      }
    }
  }
  SPanel {
    span: 4;
    heading: "Panel Heading";
    footer: "Panel Footer";
    SColumn {
      SText {
        style: "block";
        text: "Heading and footer";
      }
    }
  }
}
