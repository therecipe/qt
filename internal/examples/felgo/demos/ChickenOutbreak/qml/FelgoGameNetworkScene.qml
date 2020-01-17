import QtQuick 2.0
import Felgo 3.0

SceneBase {
  id: scene

  // needs to be accessed assigned to the FelgoGameNetwork, so it can show it
  property alias gameNetworkView: gameNetworkView

  onBackButtonPressed: {
    window.state = "main"
  }

  GameNetworkView {
    id: gameNetworkView
    anchors.fill: scene.gameWindowAnchorItem    

    // this is only used temporarily, until the font issue is fixed

    onBackClicked: {
      window.state = "main"
    }
  }
}
