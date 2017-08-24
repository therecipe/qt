import QtQuick 2.6
import StratifyLabs.UI 2.0

SColumn {
  SProgressCircle {
    id: progressCircle;
    style: "primary";
    value: 0.2;

    SIcon {
      anchors.centerIn: parent;
      style: "text-h1";
      icon: Fa.Icon.cloud;
    }

    Timer {
      id: timer;
      interval: 100;
      repeat: true;
      running: true;
      onTriggered: {
        progressCircle.value += 0.01;
        if( progressCircle.value >= 1.0 ){
          progressCircle.value = 0.0;
        }
      }
    }
  }
}
