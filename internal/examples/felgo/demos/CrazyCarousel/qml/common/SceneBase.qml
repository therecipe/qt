import Felgo 3.0
import QtQuick 2.0

Scene {
  id: sceneBase

  // by default, set the opacity to 0 - this will be changed from the main.qml with PropertyChanges
  opacity: 0
  // we set the visible property to false if opacity is 0 because the renderer skips invisible items, this is an performance improvement
  visible: opacity > 0
  // if the scene is invisible, we disable it. In Qt 5, components are also enabled if they are invisible. This means any MouseArea in the Scene would still be active even we hide the Scene, since we do not want this to happen, we disable the Scene (and therefore also its children) if it is hidden
  enabled: visible
}
