import QtQuick.Controls 1.4         //TableView
import QtQuick.Controls.Styles 1.4  //TableViewStyle

import Theme 1.0                     //Theme
import WalletTemplate 1.0            //WalletTemplate

TableView {
  id: tableView

  property WalletTemplate template

  backgroundVisible: false
  frameVisible: false

  horizontalScrollBarPolicy: Qt.ScrollBarAlwaysOff
  verticalScrollBarPolicy: Qt.ScrollBarAlwaysOff

  model: WalletModel

  style: TableViewStyle {

    headerDelegate: TableHeaderDelegate {}

    rowDelegate: TableRowDelegate {}
  }

  TableColumnDelegate { role: "STATUS" }

  TableColumnDelegate { role: "DATE AND TIME" }

  TableColumnDelegate { role: "AMOUNT" }

  TableColumnDelegate { role: "TYPE" }

  TableColumnDelegate { role: "TOTAL BALANCE" }

  onDoubleClicked: template.doubleClicked(tableView.model.data(tableView.model.index(tableView.currentRow, 0), Qt.UserRole + 6))
}
