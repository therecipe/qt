import QtQuick 2.6
import StratifyLabs.UI 2.0

EPane {
  name: "Lists";
  EHeading {
    title: "Buttons";
    inherits: "ListView";
    stratifyName: "SList";
    specialAttibutes: "icon";
  }

  ESectionTitle { text: "Description"; }
  EParagraph { text: 'An <i>SList</i> is a \
QML <i>ListView</i> object with SL.UI styling.'; }
  SHLine{}

  ESectionTitle { text: "Example"; }
  ListExample{}
  ECodeButton { source: "ListExample"; }

  ESectionTitle { text: "ListGroup"; }
  ListExampleListGroup{}
  ECodeButton { source: "ListExampleListGroup"; }

}
