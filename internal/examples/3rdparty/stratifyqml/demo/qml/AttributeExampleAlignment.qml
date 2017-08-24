import StratifyLabs.UI 2.0

SColumn {
  SContainer {
    implicitHeight: 400;
    color: STheme.brand_secondary;
    SRow {
      /*
  The parent needs to have
  a height that is larger than
  the implicit (or natural)
  height. Setting style to
  "fill" will make the SRow
  fill the parent height.
    */
      style: "fill";
      SLabel {
        span: 4;
        style: "label-info left";
        text: "Align Left";
      }
      SLabel {
        span: 4;
        style: "label-info center";
        text: "Align Center";
      }
      SLabel {
        span: 4;
        style: "label-info right";
        text: "Align Right";
      }
      SLabel {
        span: 4;
        style: "label-info top";
        text: "Align Top";
      }

      SLabel {
        span: 4;
        style: "label-info middle";
        text: "Align Middle";
      }
      SLabel {
        span: 4;
        style: "label-info bottom";
        text: "Align Bottom";
      }
    }
  }
}
