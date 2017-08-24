import QtQuick 2.6
import StratifyLabs.UI 2.0

EPane {
  name: "Tables";
  EHeading {
    title: "Tables";
    inherits: "QtQuick.Controls 1.4 TableView";
    stratifyName: "STable";
    defaultSize: "block";
  }

  ESectionTitle { text: "Description"; }

  SHLine{}

  ESectionTitle { text: "Example"; }
  TableExample{}
  ECodeButton { source: "TableExample"; }

  ESectionTitle { text: "Options"; }
  TableExampleOptions{}
  ECodeButton { source: "TableExampleOptions"; }
}
