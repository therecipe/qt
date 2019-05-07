pragma Singleton

import QtQuick 2.7

import ThemeTemplate 1.0

ThemeTemplate {

  //min size to calculate with
  property real minWidth: 1024
  property real minHeight: 768

  //needed to position dialog correctly
  property real currentWidth
  property real currentHeight
}
