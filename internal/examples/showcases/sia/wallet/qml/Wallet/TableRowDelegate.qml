import QtQuick 2.7  //Rectangle

import Theme 1.0    //Theme

Rectangle {
  radius: 5
  height: Theme.minHeight * 0.05
  color: styleData.alternate ? Theme.walletTableAlternate : "transparent"
}
