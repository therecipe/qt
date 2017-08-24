import QtQuick 2.0
import StratifyLabs.UI 2.0

SGroup {
  property string tags;
  id: root;
  style: "right";
  SText {
    style: "text-sm";
    text: "Read More";
    anchors.verticalCenter: parent.verticalCenter;
  }

  function parseTags(){ return tags.split(" "); }
  function parseItem(value){ return value.split(":"); }

  onTagsChanged: {
    var items = parseTags();
    for(var i=0; i < items.length; i++){
      var value = parseItem(items[i]);

      console.log(value[0] + "+" + value[1]);

      var component = Qt.createComponent("EReadMoreItem.qml");
      component.createObject(root,
                             {
                               "text": value[1],
                               "screenName": value[0],
                             });
    }
  }
}
