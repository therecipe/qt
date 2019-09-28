import Felgo 3.0
import QtQuick 2.0

import "charts"
import "qmlvideofx/qml"

App {
  id: app

  property bool enable3DView: false
  property var chartData: []

  // load data from C++ at app startup
  Component.onCompleted: cppDataModel.loadData()

  // when new data is sent from C++, parse the JSON and use as chartData in QML
  Connections {
    target: cppDataModel
    // the dataLoaded signal provides a jsonDataString parameter
    onDataLoaded: chartData = JSON.parse(jsonDataString)
  }

  // UI code starts here
  NavigationStack {
    initialPage: Page {
      title: "Chart Example"
      rightBarItem: IconButtonBarItem {
        icon: !enable3DView ? IconType.cube : IconType.barchart
        onClicked: enable3DView = !enable3DView
      }

      BarChart3D {
        visible: enable3DView
      }

      BarChart2D {
        id: barChart2D
        visible: !enable3DView
      }

      // apply shader fx effect
      VideoFXEffect {
        id: fxEffect
        anchors.fill: barChart2D
        divider.anchors.bottomMargin: fxControls.parameterPanelHeight
        sourceForShaderEffect.sourceItem: barChart2D
        gripSize: fxControls.gripSize
        visible: barChart2D.visible
      }

      // shader fx controls
      VideoFXControls {
        id: fxControls
        //    anchors.topMargin: background.height - scaledMargin
        targetEffect: fxEffect
        visible: fxEffect.visible
      }

    } // Page
  } // NavigationStack
} // App
