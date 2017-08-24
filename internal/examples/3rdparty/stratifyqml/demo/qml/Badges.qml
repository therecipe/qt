import QtQuick 2.6
import StratifyLabs.UI 2.0

EPane {
  name: "Badges";
  EHeading {
    title: "Badges";
    inherits: "Rectangle";
    stratifyName: "SBadge";
    specialAttibutes: "icon";
  }

  ESectionTitle { text: "Description"; }

  EParagraph { text: '\
An <i>SBadge</i> is used as a decorative object to \
emphasize meaning to a small amount of information. \
An <i>SBadge</i> is always circular. It is only useful in showing
one or two digits or an icon.';
  }

  SHLine{}

  ESectionTitle { text: "Colors"; }
  BadgeExampleColors{}
  ECodeButton { source: "BadgeExampleColors"; }

  ESectionTitle { text: "Outlines"; }
  BadgeExampleOutlineColors{}
  ECodeButton { source: "BadgeExampleOutlineColors"; }

  ESectionTitle { text: "Icons"; }
  BadgeExampleIcons{}
  ECodeButton { source: "BadgeExampleIcons"; }

  ESectionTitle { text: "Sizes"; }
  BadgeExampleSizes{}
  ECodeButton { source: "BadgeExampleSizes"; }
}


