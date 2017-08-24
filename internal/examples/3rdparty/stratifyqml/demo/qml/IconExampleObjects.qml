import StratifyLabs.UI 2.0

SRow {
  SIcon {
    span: 4;
    icon: Fa.Icon.adjust;
  }
  SIcon {
    span: 4;
    icon: Fa.Icon.bomb;
    label: "Label";
  }
  SIcon {
    span: 4;
    icon: Fa.Icon.clock_o;
  }

  SIcon {
    style: "icon-spin";
    span: 4;
    icon: Fa.Icon.safari;
    label: "Slow Spin";
    attr.animationPeriod: 5000;
  }
  SIcon {
    style: "icon-spin";
    span: 4;
    icon: Fa.Icon.refresh;
    label: "Medium Spin";
    attr.animationPeriod: 2000;
  }
  SIcon {
    style: "icon-spin";
    span: 4;
    icon: Fa.Icon.cog;
    label: "Fast Spin";
    attr.animationPeriod: 1000;
  }

  SIcon {
    style: "icon-pulse";
    span: 4;
    icon: Fa.Icon.circle_o_notch;
    label: "Slow Pulse";
    attr.animationPeriod: 5000;
  }
  SIcon {
    style: "icon-pulse";
    span: 4;
    icon: Fa.Icon.spinner;
    label: "Medium Pulse";
    attr.animationPeriod: 2000;
  }
  SIcon {
    style: "icon-pulse";
    span: 4;
    icon: Fa.Icon.location_arrow;
    label: "Fast Pulse";
    attr.animationPeriod: 1000;
  }
}
