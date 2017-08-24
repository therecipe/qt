import QtQuick 2.6
import StratifyLabs.UI 2.0

EPane {
  name: "JsonModels";
  EHeading {
    title: "Json Models";
    inherits: "ListModel";
    stratifyName: "SJsonModel";
  }

  ESectionTitle { text: "Description"; }
  SText {
    style: "block";
    text: '\
The <i>SJsonModel</i> object allows you to use JSON objects \
as list model (and table model) information. The example \
below shows how the JSON is formatted. By default, <i>SJsonModel</i> \
looks for a root item named <i>data</i> contains an array \
of objects. Each named string value in the object is made \
available to the delegate. In the example, below,
the delegate item uses <i>model.text</i>. \
';
  }

  SHLine{}

  ESectionTitle { text: "Example"; }
  JsonModelExample{}
  ECodeButton { source: "JsonModelExample"; }

}
