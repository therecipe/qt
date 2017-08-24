import QtQuick 2.0
import StratifyLabs.UI 2.0

ECodeBlock {
  id: code;
  property string source;


  onSourceChanged: {
    var doc = new XMLHttpRequest;
    doc.onreadystatechange = function() {
      if( doc.readyState === XMLHttpRequest.DONE ){
        code.text = doc.responseText;
      }
    }
    doc.open("GET", source + ".qml");
    doc.send();

  }
}
