import QtQuick 2.7            //Rectangle
import QtQuick.Controls 2.1   //TextField

import Theme 1.0              //Theme
import TerminalTemplate 1.0   //TerminalTemplate

TextField {
  property TerminalTemplate template
  property QtObject output

  placeholderText: "Enter a command"

  color: Theme.font
  selectionColor: Theme.accent
  font.pointSize: 16
  font.italic: displayText === "" ? true : false

  background: Rectangle {
    color: Theme.inputFieldBackground
  }

  onAccepted: {
    output.text = template.command(text)
    clear()
  }
}
