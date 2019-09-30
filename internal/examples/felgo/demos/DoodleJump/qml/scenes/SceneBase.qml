import Felgo 3.0
import QtQuick 2.0

Scene {
  id: sceneBase

  width: 320
  height: 480

  // by default, set the opacity to 0. We handle this property from Main.qml via PropertyChanges.
  opacity: 0

  // the scene is only visible if the opacity is > 0. This improves performance.
  visible: opacity > 0

  // only enable scene if it is visible.
  enabled: visible
}
