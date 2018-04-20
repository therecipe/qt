import QtQuick 2.6
import StratifyLabs.UI 2.1
import QtCharts 2.2

SRow {
  SChart {
    span: 4;
    style: "text-h1 hide-legend";
    title: "Primary";
    SPieSeries {
      id: pieSeries;
      chartStyle: "primary";
      SPieSlice {
        label: "1"; value: 10;
      }
      SPieSlice {
        label: "2"; value: 10;
      }

      SPieSlice {
        label: "3"; value: 10;
      }

      SPieSlice {
        label: "4"; value: 10;
      }

      SPieSlice {
        label: "5"; value: 10;
      }

      SPieSlice {
        label: "6"; value: 10;
      }
      SPieSlice {
        label: "7"; value: 10;
      }

      SPieSlice {
        label: "8"; value: 10;
      }

      SPieSlice {
        label: "9"; value: 10;
      }

      SPieSlice {
        label: "10"; value: 10;
      }
    }
  }

  SChart {
    span: 4;
    title: "Secondary";
    style: "text-h1 hide-legend";
    ChartPie {  //this object is the same as id: pieSeries
      chartStyle: "secondary";
    }
  }

  SChart {
    span: 4;
    title: "Info";
    style: "text-h1 hide-legend";
    ChartPie {
      chartStyle: "info";
    }
  }

  SChart {
    span: 4;
    title: "Success";
    style: "text-h1 hide-legend";
    ChartPie {
      chartStyle: "success";
    }
  }

  SChart {
    span: 4;
    title: "Warning";

    style: "text-h1 hide-legend";
    ChartPie {
      chartStyle: "warning";
    }
  }

  SChart {
    span: 4;
    title: "Danger";
    style: "text-h1 hide-legend";
    ChartPie {
      chartStyle: "danger";
    }
  }

  SChart {
    span: 4;
    title: "Primary";
    style: "text-h1 hide-legend primary";
    ChartPie {}
  }

  SChart {
    span: 4;
    title: "Secondary";
    style: "text-h1 hide-legend secondary";
    ChartPie {}
  }

  SChart {
    span: 4;
    title: "Info";
    style: "text-h1 hide-legend info";
    ChartPie {}
  }

  SChart {
    span: 4;
    title: "Success";
    style: "text-h1 hide-legend success";
    ChartPie {}
  }

  SChart {
    span: 4;
    title: "Warning";
    style: "text-h1 hide-legend warning";
    ChartPie {}
  }

  SChart {
    span: 4;
    title: "Danger";
    style: "text-h1 hide-legend danger";
    ChartPie {}
  }
}
