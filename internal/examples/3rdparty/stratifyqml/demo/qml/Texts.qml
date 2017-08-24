import QtQuick 2.6
import StratifyLabs.UI 2.0

EPane {
  name: "Texts";
  EHeading {
    title: "Buttons";
    inherits: "QtQuick.Controls 2.0 Button";
    stratifyName: "SButton";
    specialAttibutes: "icon";
  }

  ESectionTitle { text: "Description"; }

  SHLine{}

  ESectionTitle { text: "Blocks"; }
  TextExample{}
  ECodeButton { source: "TextExample"; }

  ESectionTitle { text: "Sizes"; }
  TextExampleSizes{}
  ECodeButton { source: "TextExampleSizes"; }

  ESectionTitle { text: "Styles"; }
  TextExampleStyles{}
  ECodeButton { source: "TextExampleStyles"; }
}
