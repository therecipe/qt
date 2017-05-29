package main

import (
	"github.com/therecipe/qt/core"
	"github.com/therecipe/qt/gui"
	"github.com/therecipe/qt/widgets"
)

func itemViews() {

	//List View
	listView := widgets.NewQListView(nil)
	listView.SetWindowTitle("List View")

	list := []string{"This", "Is", "A", "List", "View"}

	listModel := core.NewQAbstractListModel(nil)
	listModel.ConnectRowCount(func(parent *core.QModelIndex) int {
		return len(list)
	})
	listModel.ConnectData(func(index *core.QModelIndex, role int) *core.QVariant {
		if role != int(core.Qt__DisplayRole) {
			return core.NewQVariant()
		}
		return core.NewQVariant14(list[index.Row()])
	})
	listView.SetModel(listModel)
	addWidget(listView)

	//List View + String List Model
	listViewS := widgets.NewQListView(nil)
	listViewS.SetWindowTitle("List View + String List Model")
	listViewS.SetModel(core.NewQStringListModel2(list, nil))
	addWidget(listViewS)

	//Tree View
	treeView := widgets.NewQTreeView(nil)
	treeView.SetWindowTitle("Tree View")

	treeModel := gui.NewQStandardItemModel(nil)
	rootNode := treeModel.InvisibleRootItem()

	americaItem := gui.NewQStandardItem2("America")
	mexicoItem := gui.NewQStandardItem2("Canada")
	usaItem := gui.NewQStandardItem2("USA")
	bostonItem := gui.NewQStandardItem2("Boston")
	europeItem := gui.NewQStandardItem2("Europe")
	italyItem := gui.NewQStandardItem2("Italy")
	romeItem := gui.NewQStandardItem2("Rome")
	veronaItem := gui.NewQStandardItem2("Verona")

	rootNode.AppendRow2(americaItem)
	rootNode.AppendRow2(europeItem)
	americaItem.AppendRow2(mexicoItem)
	americaItem.AppendRow2(usaItem)
	usaItem.AppendRow2(bostonItem)
	europeItem.AppendRow2(italyItem)
	italyItem.AppendRow2(romeItem)
	italyItem.AppendRow2(veronaItem)

	treeView.SetModel(treeModel)
	treeView.ExpandAll()
	addWidget(treeView)

	//Table View
	tableView := widgets.NewQTableView(nil)
	tableView.SetWindowTitle("Table View")

	table := [][]string{
		0: {"This", "Is", "The", "First", "Row"},
		1: {"This", "Is", "The", "Second", "Row"},
		2: {"This", "Is", "The", "Third", "Row"},
	}

	tableModel := core.NewQAbstractTableModel(nil)
	tableModel.ConnectRowCount(func(parent *core.QModelIndex) int {
		return len(table)
	})
	tableModel.ConnectColumnCount(func(parent *core.QModelIndex) int {
		return len(table[0])
	})
	tableModel.ConnectData(func(index *core.QModelIndex, role int) *core.QVariant {
		if role != int(core.Qt__DisplayRole) {
			return core.NewQVariant()
		}
		return core.NewQVariant14(table[index.Row()][index.Column()])
	})
	tableView.SetModel(tableModel)
	addWidget(tableView)

	//Column View
	columnView := widgets.NewQColumnView(nil)
	columnView.SetWindowTitle("Column View")

	columnView.SetModel(treeModel)
	addWidget(columnView)

}
