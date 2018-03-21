import QtQuick 2.7          //Rectangle
import QtQuick.Controls 2.1 //StackView

import Theme 1.0            //Theme
import ViewTemplate 1.0     //StackTemplate
import Dashboard 1.0        //Dashboard
import Files 1.0            //Files
import Hosting 1.0          //Hosting
import Wallet 1.0           //Wallet
import Terminal 1.0         //Terminal

StackTemplate {
  Rectangle {
    anchors.fill: parent
    color: Theme.darkBackground

    StackView {
      id: stackView
      anchors.fill: parent

      initialItem: wallet

      Dashboard { id: dashboard }
      Files { id: files }
      Hosting { id: hosting }
      Wallet { id: wallet }
      Terminal { id: terminal }
    }
  }

  onClicked: {
    var next = nextItem(code)
    if (next != stackView.currentItem) {
      stackView.replace(next, StackView.Immediate)
    }
  }

  function nextItem(code) {
    switch (code) {
    case "dashboard":
      return dashboard
    case "files":
      return files
    case "hosting":
      return hosting
    case "wallet":
      return wallet
    case "terminal":
      return terminal
    }
  }
}
