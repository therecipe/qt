import QtQuick 2.6
import StratifyLabs.UI 2.0

SColumn {

  property string loremIpsum: "Lorem ipsum dolor sit amet, \
consectetur adipiscing elit, sed do eiusmod \
tempor incididunt ut labore et dolore magna \
aliqua. Ut enim ad minim veniam, quis nostrud \
exercitation ullamco laboris nisi ut aliquip \
ex ea commodo consequat. Duis aute irure dolor \
in reprehenderit in voluptate velit esse cillum \
dolore eu fugiat nulla pariatur. Excepteur \
sint occaecat cupidatat non proident, sunt in \
culpa qui officia deserunt mollit anim id est laborum."

  SText {
    style: "left text-h1";
    text: "Default";
  }

  SText {
    style: "block";
    text: loremIpsum;
  }

  SHLine {}

  SText {
    style: "left text-h1";
    text: "Center";
  }

  SText {
    style: "block text-center";
    text: loremIpsum;
  }

  SHLine {}

  SText {
    style: "left text-h1";
    text: "Right";
  }

  SText {
    style: "block text-right";
    text: loremIpsum;
  }

}


