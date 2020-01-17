import QtQuick 2.0
import Felgo 3.0

// base component for all 3 scenes in the game: MainScene, ChickenOutbreakScene and GameOverScene
Scene {
  id: sceneBase
  width: 320
  height: 480

  // this is an important performance improvement, as renderer can skip invisible items (and all its children)
  // also, the focus-property for key handling relies on the visible-property
  visible: opacity>0


  // this fades in and out automatically, when the opacity gets changed from 0 to 1 in ChickenOutbreakMain
  Behavior on opacity {
    // the cross-fade animation should last 350ms
    NumberAnimation { duration: 350 }
  }

}
