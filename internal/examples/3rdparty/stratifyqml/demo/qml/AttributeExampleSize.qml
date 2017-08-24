import StratifyLabs.UI 2.0

SColumn {
  SContainer {
    implicitHeight: 200;
    color: STheme.brand_secondary;
    SColumn {
      /*
  The parent needs to have a
  height that is larger than
  the implicit (or natural)
  height. Setting style to "fill"
  will make the SRow fill the
  parent height.
            */
      style: "fill";
      SLabel {
        style: "label-info block";
        text: "Block Label (fill width)";
      }
      SLabel {
        /*
  For "fill" to have meaning,
  all parent must set an implicitHeight
  other than the default or set the style
  to "fill". Otherwise "fill" will simply
  fill the natural height of the item and
  look the same (and maybe cause a bind
  loop with the implicitHeight)
                */
        style: "label-info fill";
        text: "Fill Label (fill height)";
      }
      SLabel {
        style: "label-info";
        text: "Default Label Size";
      }
    }
  }
}
