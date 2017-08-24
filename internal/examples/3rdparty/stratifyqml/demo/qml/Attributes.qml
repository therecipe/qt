import QtQuick 2.6
import StratifyLabs.UI 2.0

EPane {
  name: "Attributes";
  ESectionTitle { text: "Introduction"; }
  EParagraph { text: '\
SL.UI allows you to quickly layout and style \
QML applications. It does this by including an <i>SAttributes</i> \
object which then defines the size and color scheme. \
There are two ways to change an object\'s attributes: \
1) using <i>style</i> 2) using <i>attr</i> directly.';
  }

  EParagraph { text: '\
Additionally, each item has quick access to <i>attr.span</i> through \
the <i>span</i> property which \
defines the width of a SL.UI object in an <i>SRow</i>. Each SL.UI object \
defines the properties listed below.';
  }

  EReadMore {
    tags: "Rows:SRow";
  }

  ECodeBlockInline {
    text: '\
property alias style:
  attr.style;
property alias span:
  attr.span;
property alias attr:
  attr;
SAttributes {
  id: attr;
  type: "button";
}
';
  }

  EParagraph { text: '\
Most SL.UI objects use <i>SAttributes</i> as the <i>attr</i> object. \
SL.UI objects that don\'t have any color (e.g. <i>SPane</i>) use \
<i>SSizeAttributes</i> (which <i>SAttributes</i> inherits). If an object \
has additional attibutes, the object will be <i>S{Object}Attributes</i> \
(such as SAlertAttributes). <i>SAlertAttributes</i> adds the properties \
"dismissible" and "animationPeriod" to allow easy customization \
of the object.'; }

  ESectionTitle { text: "Changing the Style"; }
  EParagraph { text: '\
When a new value is assigned to the <i>style</i> property, \
the string is parsed (left to right) and new colors and sizes are bound \
to the object\'s attibutes.';
  }
  EParagraph { text: '\
For example, <i>style: "padding-zero"</i> will cause \
<i>paddingHorizontal</i> and <i>paddingVertical</i> to be \
set to zero.';
  }

  ESubSectionTitle { text: 'SSizeAttributes'; }
  EParagraph { text: '\
<i>SSizeAttributes</i> (all objects) support the following style \
options.';
  }

  ESubSectionTitle { text: "Padding"; }
  EBullet { text: "padding-zero"; }
  EBullet { text: "padding-xs"; }
  EBullet { text: "padding-sm"; }
  EBullet { text: "padding-normal"; }
  EBullet { text: "padding-lg"; }

  SHLine {}

  ESubSectionTitle { text: "Horizontal Alignment"; }
  EBullet { text: "left"; }
  EBullet { text: "right"; }
  EBullet { text: "center"; }

  ESubSectionTitle { text: "Vertical Alignment"; }
  EBullet { text: "top"; }
  EBullet { text: "bottom"; }
  EBullet { text: "middle"; }


  AttributeExampleAlignment{}
  ECodeButton { source: "AttributeExampleAlignment"; }

  ESubSectionTitle { text: "Sizing"; }
  EBullet { text: "fill: fill height"; }
  EBullet { text: "block: fill width"; }
  EBullet { text: "implicit-size: no fill"; }


  AttributeExampleSize{}
  ECodeButton { source: "AttributeExampleSize"; }


  SText { style: "block text-bold"; text: 'SAttributes'; }
  EParagraph { text: 'SAttributes (most objects) support the following style \
options.'; }

  ESubSectionTitle { text: "Size"; }
  SText { style: "left text-sm"; text: "Changes text size and padding"; }
  EBullet { text: "xs"; }
  EBullet { text: "sm"; }
  EBullet { text: "lg"; }

  SHLine {}

  ESubSectionTitle { text: "Text Sizes"; }
  EBullet { text: "text-h1"; }
  EBullet { text: "text-h2"; }
  EBullet { text: "text-h3"; }
  EBullet { text: "text-h4"; }
  EBullet { text: "text-h5"; }
  EBullet { text: "text-h6"; }
  EBullet { text: "text-sm"; }
  EBullet { text: "text-lg"; }


  AttributeExampleTextSize{}
  ECodeButton { source: "AttributeExampleTextSize"; }


  ESubSectionTitle { text: "Text Alignment"; }

  EBullet { text: "text-right"; }
  EBullet { text: "text-left"; }
  EBullet { text: "text-center"; }
  EBullet { text: "text-top"; }
  EBullet { text: "text-bottom"; }
  EBullet { text: "text-middle"; }

  AttributeExampleTextAlignment{}
  ECodeButton { source: "AttributeExampleTextAlignment"; }


  SHLine {}

  ESubSectionTitle { text: "Text Weight"; }
  EBullet { text: "text-bold"; }
  EBullet { text: "text-light"; }
  EBullet { text: "text-demi-bold"; }
  EBullet { text: "text-extra-bold"; }
  EBullet { text: "text-extra-light"; }
  SRow {
    SText { span: 4; style: "text-bold"; text: "Text Bold"; }
    SText { span: 4; style: "text-light"; text: "Text Light"; }
    SText { span: 4; style: "text-demi-bold"; text: "Text Demi Bold"; }
    SText { span: 4; style: "text-extra-bold"; text: "Text Extra Bold"; }
    SText { span: 4; style: "text-extra-ligh"; text: "Text Extra Light"; }
    SText { span: 4; text: "Text"; }
  }

  SHLine {}

  ESubSectionTitle { text: "Text Color"; }
  EBullet { text: "text-on-primary"; }
  EBullet { text: "text-on-secondary"; }
  EBullet { text: "text-on-info"; }
  EBullet { text: "text-on-warning"; }
  EBullet { text: "text-on-danger"; }
  EBullet { text: "text-on-success"; }
  EBullet { text: "text-primary"; }
  EBullet { text: "text-secondary"; }
  EBullet { text: "text-info"; }
  EBullet { text: "text-warning"; }
  EBullet { text: "text-danger"; }
  EBullet { text: "text-success"; }


  SHLine {}

  ESubSectionTitle { text: "Color Schemes"; }
  SText { style: "left text-sm"; text: "Change font, object, and border colors"; }
  EBullet { text: "primary"; }
  EBullet { text: "secondary"; }
  EBullet { text: "info"; }
  EBullet { text: "warning"; }
  EBullet { text: "danger"; }
  EBullet { text: "success"; }

  EParagraph { text: '\
Each SL.UI object may add additional styling options. For example \
SButton uses "btn-outline-primary" for setting the button\'s color scheme.';
  }



  ESectionTitle { text: "<i>attr</i>"; }

  EParagraph { text: '\
If the styling options above don\'t meet your needs, \
you can directly access an item\'s attributes using the \
<i>attr</i> property. This is useful if you want to bind to \
specific attributes. The following example binds the icon <i>spin</i> \
attribute to a variable.';
  }

  AttributeExampleCustomAttributes{}
  ECodeButton { source: "AttributeExampleCustomAttributes"; }
}



