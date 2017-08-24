import QtQuick 2.6
import StratifyLabs.UI 2.0

EPane {
  name: "FontAwesome";
  ESectionTitle { text: "Introduction"; }
  EParagraph { style: "block"; text: 'SL.UI icons are based on \
the latest release of Font Awesome. \
The icons can be used using <i>SIcon</i>, \
<i>SButton</i>, and <i>SBadge</i>. You can \
also set the <i>font.family</i> of any text to <i>STheme.fontFontAwesome.name</i> \
to use the Font Awesome in other objects.'; }

  EReadMore {
    tags: "Badges:SBadge Buttons:SButton Icons:SIcon";
  }

  FontAwesomeExampleObjects{}
  ECodeButton { source: "FontAwesomeExampleObjects"; }

  ESectionTitle { text: "The Icons"; }

  SText {
    id: fontList;
    font.family: STheme.fontFontAwesome.name;
    style: "block text-lg";
    Component.onCompleted: {
      var value = "";
      for(var ico in Fa.Icon){
        value += Fa.Icon[ico] + " ";
      }
      fontList.text = value;
    }
  }




  EParagraph { text: 'The current icon set is Font Awesome 4.7. The icons are available as \
<i>Fa.Icon.[name]</i> where name is the name in the Font Awesome cheatsheet with "-" changed to "_".  Also, only \
the original names are supported rather than the alias (e.g. use <i>Fa.Icon.times</i> rather than <i>Fa.Icon.remove</i>; see the link \
below for more information).';
  }

  SButton {
    style: "btn-outline-info block";
    text: "Font Awesome Cheatsheet";
    icon: Fa.Icon.external_link;
    onClicked: Qt.openUrlExternally("http://fontawesome.io/cheatsheet/");
  }
}
