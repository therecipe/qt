package main

import (
	"github.com/therecipe/qt/widgets"
)

func itemWidgets() {

	//List Widget
	listWidget := widgets.NewQListWidget(nil)
	listWidget.SetWindowTitle("List Widget")

	list := []string{"This", "Is", "A", "List", "View"}

	listWidget.AddItems(list)
	addWidget(listWidget)

	//Tree Widget
	treeWidget := widgets.NewQTreeWidget(nil)
	treeWidget.SetWindowTitle("Tree Widget")

	rootNode := treeWidget.InvisibleRootItem()

	americaItem := widgets.NewQTreeWidgetItem2([]string{"America"}, 0)
	mexicoItem := widgets.NewQTreeWidgetItem2([]string{"Canada"}, 0)
	usaItem := widgets.NewQTreeWidgetItem2([]string{"USA"}, 0)
	bostonItem := widgets.NewQTreeWidgetItem2([]string{"Boston"}, 0)
	europeItem := widgets.NewQTreeWidgetItem2([]string{"Europe"}, 0)
	italyItem := widgets.NewQTreeWidgetItem2([]string{"Italy"}, 0)
	romeItem := widgets.NewQTreeWidgetItem2([]string{"Rome"}, 0)
	veronaItem := widgets.NewQTreeWidgetItem2([]string{"Verona"}, 0)

	rootNode.AddChild(americaItem)
	rootNode.AddChild(europeItem)
	americaItem.AddChild(mexicoItem)
	americaItem.AddChild(usaItem)
	usaItem.AddChild(bostonItem)
	europeItem.AddChild(italyItem)
	italyItem.AddChild(romeItem)
	italyItem.AddChild(veronaItem)

	treeWidget.ExpandAll()
	addWidget(treeWidget)

	//Table Widget
	tableWidget := widgets.NewQTableWidget(nil)
	tableWidget.SetWindowTitle("Table Widget")

	table := [][]string{
		0: {"This", "Is", "The", "First", "Row"},
		1: {"This", "Is", "The", "Second", "Row"},
		2: {"This", "Is", "The", "Third", "Row"},
	}

	tableWidget.SetRowCount(len(table))
	tableWidget.SetColumnCount(len(table[0]))

	for row := 0; row < tableWidget.RowCount(); row++ {
		for column := 0; column < tableWidget.ColumnCount(); column++ {
			tableWidget.SetItem(row, column, widgets.NewQTableWidgetItem2(table[row][column], 0))
		}
	}
	addWidget(tableWidget)

}
