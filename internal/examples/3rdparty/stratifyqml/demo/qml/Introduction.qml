import QtQuick 2.7
import QtQuick.Controls 2.1
import StratifyLabs.UI 2.0

EPane {
  name: "Introduction";
  ESectionTitle { text: "Introduction"; }
  EParagraph { text: '\
Stratify Labs UI (SL.UI) 2.0 is a QML framework that works, looks, \
and feels much like the twitter bootstrap HTML framework. The \
colors and shapes of components can be customized using a theme \
which uses familiar bootstrap variables. The UI also provides a responsive grid \
system that mimics bootstrap\'s ability to fit any screen.';
  }
  ESectionTitle { text: "Theme"; }
  EParagraph { text: '\
The look and feel are defined by the theme which is contained \
in the QML singleton object <i>STheme</i>. The following example shows \
how to change the primary and secondary brand colors. This code \
typically goes in the main <i>ApplicationWindow</i>.';
  }
  ECodeBlockInline {
    text: 'Component.onCompleted: {
  STheme.brand_primary = "#244E99";
  STheme.brand_secondary = "#666";
}';
  }

  EParagraph { text: '\
Individual components can also be styled and customized using \
the <i>style</i> or <i>attr</i> properties. Here is an example of making a \
button that uses the <i>warning</i> color scheme.';
  }

  IntroductionWarningButtonExample{}
  ECodeButton { source: "IntroductionWarningButtonExample"; }

  EParagraph { text: '\
SL.UI objects can be further customized by directly accessing \
the <i>SAttributes</i> property via <i>attr</i>.'; }

  EReadMore {
    tags: "Attributes:SAttributes";
  }

  IntroductionRoundButtonExample{}
  ECodeButton { source: "IntroductionRoundButtonExample"; }


  ESectionTitle { text: "Grid System"; }
  EParagraph { text: '\
The responsive grid is based on two SL.UI objects: <i>SRow</i> and \
<i>SColumn</i>. <i>SRow</i> directly inherits <i>GridLayout</i> while <i>SColumn</i> \
inherits <i>SRow</i>.  <i>SRow</i> requires its children to set the <i>span</i> \
value while <i>SColumn</i> assumes all items have a maximum <i>span</i> \
value.  Also, <i>SRow</i> will convert rows to columns based on the \
screen width. For the conversion to work, the following code \
must be in the main <i>ApplicationWindow</i> object.';
  }

  EReadMore {
    tags: "Rows:SRow Columns:SColumn";
  }

  ECodeBlockInline {
    text: '\
onWidthChanged: {
  STheme.updateScreenSize(width);
}';
  }

  EParagraph { text: '\
The following example shows how two columns can change \
to one when the screen is narrow. You can change the width of \
the window (desktop) or switch from portrait to landscape (mobile) \
to see the change.';
  }
  EParagraph { text: 'Small screen mode is <b>' + STheme.isScreenSm + '</b>.'; }


  IntroductionButtonExample{}
  ECodeButton { source: "IntroductionButtonExample"; }

  ESectionTitle { text: "Icons"; }
  EParagraph { text: '\
SL.UI comes with FontAwesome 4.7 integrated. Icons can be used \
in <i>SBadge</i>, <i>SButton</i> and <i>SIcon</i>. Here is a sample.';
  }

  EReadMore {
    tags: "Badges:SBadge Buttons:SButton Icons:SIcon";
  }

  IntroductionIconExample{}
  ECodeButton { source: "IntroductionIconExample"; }

}

