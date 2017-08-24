import QtQuick 2.6
import StratifyLabs.UI 2.0

SRow {
  property real boxHeight: 150;
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
    span: 4;
    style: "left text-h1 top";
    text: "Default";
  }

  STextBox {
    span: 8;
    implicitHeight: boxHeight;
    style: "block";
    text: loremIpsum;
  }

  SHLine {}

  SText {
    span: 4;
    style: "left text-h1 top";
    text: "Center";
  }

  STextBox {
    span: 8;
    implicitHeight: boxHeight;
    style: "block text-center";
    text: loremIpsum;
  }

  SHLine {}

  SText {
    span: 4;
    style: "left text-h1 top";
    text: "Right";
  }

  STextBox {
    span: 8;
    implicitHeight: boxHeight;
    style: "block text-right";
    text: loremIpsum;
  }

}
