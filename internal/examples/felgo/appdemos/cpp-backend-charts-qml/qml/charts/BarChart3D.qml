import QtQuick 2.0
import QtDataVisualization 1.2
import Felgo 3.0


Item {
  anchors.fill: parent

  Bars3D {
    id: barGraph
    anchors.fill: parent

    Bar3DSeries {
      // the format when an item is selected
      itemLabelFormat: "@colLabel, @rowLabel: @valueLabel"

      ItemModelBarDataProxy {
        itemModel: ListModel { id: barGraphModel }

        // Mapping model roles to bar series rows, columns, and values.
        rowRole: "year"
        columnRole: "city"
        valueRole: "expenses"
      }
    }

    function udpateBarGraphModel() {
      barGraphModel.clear()

      for(var i = 0; i < chartData.length; i++) {
        barGraphModel.append(chartData[i])
      }
    }
  }

  // data connection
  Connections {
    target: app
    onChartDataChanged: barGraph.udpateBarGraphModel()
  }
  Component.onCompleted: barGraph.udpateBarGraphModel()

  // animations
  SequentialAnimation {
    id: cameraAnimationY
    loops: Animation.Infinite
    running: true

    NumberAnimation {
      id: firstAnimation
      target: barGraph.scene.activeCamera
      property:"yRotation"
      from: 25.0
      to: 35.0
      duration: 3000
      easing.type: Easing.InOutSine
    }

    NumberAnimation {
      target: barGraph.scene.activeCamera
      property:"yRotation"
      from: 35.0
      to: 25.0
      duration: firstAnimation.duration
      easing.type: Easing.InOutSine
    }
  }

  NumberAnimation {
    id: cameraAnimationX
    loops: Animation.Infinite
    running: true
    target: barGraph.scene.activeCamera
    property:"xRotation"
    from: 0.0
    to: 360.0
    duration: 25000
  }

  AppButton {
    text: "Toggle x animation"
    onClicked: cameraAnimationX.running = !cameraAnimationX.running
  }
}
