import QtQuick 2.6
import StratifyLabs.UI 2.0

EPane {
  name: "Animators";
  EHeading {
    title: "Animators";
    inherits: "Item";
    stratifyName: "SAnimator";
    defaultSize: "block";
  }

  ESectionTitle { text: "Description"; }
  EParagraph { text: 'An <i>SAnimator</i> is an object that \
can be used to navigate through its resources by setting the <i>screen</i> property. The <i>animation</i> \
property can be set to either <i>SAnimationFade</i> (default) or <i>SAnimationHPush</i>.';
  }

  SHLine{}

  ESectionTitle { text: "Simple Example"; }
  AnimatorsExample{}
  ECodeButton { source: "AnimatorsExample"; }
}

