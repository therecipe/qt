import StratifyLabs.UI 2.0

SContainer {
  color: STheme.brand_warning;
  SColumn {
    SContainer {
      color: STheme.brand_success;
      SRow {
        SLabel {
          span: 6;
          style: "label-info";
          text: "Hello";
        }
        SLabel {
          span: 6;
          style: "label-info";
          text: "World";
        }
      }
    }
    SContainer {
      color: STheme.brand_danger;
      SRow {
        SLabel {
          span: 6;
          style: "label-info";
          text: "Hello";
        }
        SLabel {
          span: 6;
          style: "label-info";
          text: "World";
        }
      }
    }
  }
}
