import QtQuick 2.6
import StratifyLabs.UI 2.1
import QtCharts 2.2

SRow {
  id: control;

  SChart {
    span: 4;
    style: "text-h1 hide-legend";
    title: "Pie";
    SPieSeries {
      id: pieSeries;
      SPieSlice {
        label: "1"; value: 25;
      }
      SPieSlice {
        label: "2"; value: 25;
      }

      SPieSlice {
        label: "3"; value: 50;
      }
    }
  }

  SJsonModel {
    id: xydata;
    json: '{ "data": [ {"x": 0, "y": 0}, {"x": 1, "y": 1}, {"x": 2, "y": 4}, {"x": 3, "y": 9} ]}';
  }

  SJsonModel {
    id: bardata;
    json: '{ "categories" : ["2007", "2008", "2009", "2010"], "sets" : [ \
{ "label": "Bob", "values":[1,2,5,10] },\
{ "label": "Susan", "values":[6,2,12,8] },\
{ "label": "Jacob", "values":[5,9,15,7] },\
{ "label": "James", "values":[2,1,8,7] } ] }';
    query: "";
  }

  SChart {
    span: 4;
    style: "text-h1 hide-legend";
    title: "Line";
    SLineSeries {
      chartStyle: "success line-dot xs";
      model: xydata;
    }
  }

  SChart {
    span: 4;
    style: "text-h1 hide-legend";
    title: "Spline";
    SSplineSeries {
      chartStyle: "danger line-dash lg";
      model: xydata;
    }
  }

  SChart {
    span: 4;
    style: "text-h1";
    SBarSeries {
      chartStyle: "default";
      model: bardata;
    }
  }


}
