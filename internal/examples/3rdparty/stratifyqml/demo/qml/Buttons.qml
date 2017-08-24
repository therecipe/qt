import QtQuick 2.6
import StratifyLabs.UI 2.0

EPane {
  name: "Buttons";
  EHeading {
    title: "Buttons";
    inherits: "QtQuick.Controls 2.0 Button";
    stratifyName: "SButton";
    specialAttibutes: "icon";
  }
  ESectionTitle { text: "Description"; }
  EParagraph { text: '\
An <i>SButton</i> is used as a user input. It can \
be colored using <i>btn-primary</i>, <i>btn-secondary</i>, \
and other branding colors (see the examples below for details).';
  }
  SHLine{}

  ESectionTitle { text: "Colors"; }
  ButtonExampleColors{}
  ECodeButton { source: "ButtonExampleColors"; }

  ESectionTitle { text: "Outlines"; }
  ButtonExampleOutlineColors{}
  ECodeButton { source: "ButtonExampleOutlineColors"; }

  ESectionTitle { text: "Icons"; }
  ButtonExampleIcons{}
  ECodeButton { source: "ButtonExampleIcons"; }

  ESectionTitle { text: "Sizes"; }
  ButtonExampleSizes{}
  ECodeButton { source: "ButtonExampleSizes"; }

  ESectionTitle { text: "Naked"; }
  ButtonExampleNaked{}
  ECodeButton { source: "ButtonExampleNaked"; }

}

