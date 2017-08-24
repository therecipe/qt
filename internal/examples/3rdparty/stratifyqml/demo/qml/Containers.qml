import QtQuick 2.6
import StratifyLabs.UI 2.0

EPane {
  name: "Containers";
  EHeading {
    title: "Containers";
    inherits: "Rectangle";
    stratifyName: "SContainer";
    defaultSize: "block";
  }

  ESectionTitle { text: "Description"; }
  EParagraph { text: 'An <i>SContainer</i> provides \
a padded, colored background for a group of SL.UI objects. An \
<i>SContainer</i> may be used on its own or within an <i>SRow</i> or \
<i>SColumn</i> layout.'; }

  EReadMore {
    tags: "Rows:SRow Columns:SColumn";
  }

  SHLine{}

  ESectionTitle { text: "Background Colors"; }
  ContainerExample{}
  ECodeButton { source: "ContainerExample"; }

  ESectionTitle { text: "Mouse Example"; }
  ContainerExampleMouse{}
  ECodeButton { source: "ContainerExampleMouse"; }
}
