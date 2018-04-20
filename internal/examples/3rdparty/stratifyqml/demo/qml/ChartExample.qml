import QtQuick 2.6
import StratifyLabs.UI 2.1
import QtCharts 2.2

SRow {
  SChart {
    style: "text-h1 hide-legend";
    title: "Primary";
    SPieSeries {
      id: pieSeries;
      style: "primary";
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
      holeSize: 0.3;
    }
  }

  SChart {
    title: "Secondary";
    style: "text-h1 hide-legend";
    ChartPie {  //this object is the same as id: pieSeries
      chartStyle: "secondary";
    }
  }

  SChart {
    title: "Info";
    style: "text-h1 hide-legend";
    ChartPie {  //this object is the same as id: pieSeries
      chartStyle: "info";
    }
  }

  SChart {
    title: "Success";
    style: "text-h1 hide-legend";
    ChartPie {  //this object is the same as id: pieSeries
      chartStyle: "success";
    }
  }

  SChart {
    title: "Warning";

    style: "text-h1 hide-legend";
    ChartPie {  //this object is the same as id: pieSeries
      chartStyle: "warning";
    }
  }

  SChart {
    title: "Danger";

    style: "text-h1 hide-legend";
    ChartPie {  //this object is the same as id: pieSeries
      chartStyle: "danger";
    }
  }

  SChart {
    title: "Primary";
    style: "text-h1 hide-legend primary";
    ChartPie {  //this object is the same as id: pieSeries
      chartStyle: "primary";
    }
  }

  SChart {
    title: "Secondary";
    style: "text-h1 hide-legend secondary";
    ChartPie {  //this object is the same as id: pieSeries
      chartStyle: "primary";
    }
  }

  SChart {
    title: "Info";
    style: "text-h1 hide-legend info";
    ChartPie {  //this object is the same as id: pieSeries
      chartStyle: "primary";
    }
  }

  SChart {
    title: "Success";
    style: "text-h1 hide-legend success";
    ChartPie {  //this object is the same as id: pieSeries
      chartStyle: "primary";
    }
  }

  SChart {
    title: "Warning";
    style: "text-h1 hide-legend warning";
    ChartPie {  //this object is the same as id: pieSeries
      chartStyle: "primary";
    }
  }

  SChart {
    title: "Danger";
    style: "text-h1 hide-legend danger";
    ChartPie {  //this object is the same as id: pieSeries
      chartStyle: "primary";
    }
  }


  SChart {

    style: "text-h1 legend-h3 legend-bottom";
    title: "Line Series";

    SJsonModel {
      id: seriesData;
      json: '{
  "data": [
    { "x": 0, "y": 0 },
    { "x": 1, "y": 1 },
    { "x": 2, "y": 4 },
    { "x": 3, "y": 9 },
    { "x": 4, "y": 16 },
    { "x": 5, "y": 25 },
    { "x": 6, "y": 36 }
  ]
}';
    }

    SLineSeries {
      chartStyle: "success line-dash-dot cap-flat";
      name: "LineSeries";
      model: seriesData.model;
    }
  }
}
