import QtQuick 2.6
import StratifyLabs.UI 2.0

EPane {
  name: "Columns";
  EHeading {
    title: "Columns";
    inherits: "SRow";
    stratifyName: "SColumn";
    defaultSize: "block";
  }

  ESectionTitle { text: "Description"; }
  EParagraph { text: 'An <i>SColumn</i> is \
similar to an <i>SRow</i> but it assumes \
all of its children span the entire width \
(which is the default value). <i>SColumn</i> does not \
respond to changes in screen size like <i>SRow</i> does.';
  }

  EReadMore {
    tags: "Rows:SRow";
  }

  SHLine{}

  ESectionTitle { text: "Example"; }
  ColumnExample{}
  ECodeButton { source: "ColumnExample"; }
}
