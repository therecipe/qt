package main

import (
	"github.com/therecipe/qt/core"
	"github.com/therecipe/qt/gui"
)

func itemViews() {

	//Abstract List Model Used By The List And Table View
	listDB := []map[string]string{
		{"name": "Apple", "cost": "2.45"},
		{"name": "Orange", "cost": "3.25"},
		{"name": "Banana", "cost": "1.95"},
	}
	listRoles := map[int]*core.QByteArray{
		int(core.Qt__UserRole) + 1: core.NewQByteArray2("name", -1),
		int(core.Qt__UserRole) + 2: core.NewQByteArray2("cost", -1),
	}

	listModel := core.NewQAbstractListModel(nil)
	listModel.ConnectRowCount(func(parent *core.QModelIndex) int {
		return len(listDB)
	})
	listModel.ConnectData(func(index *core.QModelIndex, role int) *core.QVariant {
		if _, ok := listRoles[role]; !ok {
			return core.NewQVariant()
		}
		return core.NewQVariant14(listDB[index.Row()][listRoles[role].ConstData()])
	})
	listModel.ConnectRoleNames(func() map[int]*core.QByteArray {
		return listRoles
	})

	//Standard Item Model Used By The Tree View
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

	//List View + List Model
	createWidget("List View + List Model", `
	ListView {
		model: ListModel {

			ListElement {
				name: "Apple"
				cost: 2.45
			}

			ListElement {
				name: "Orange"
				cost: 3.25
			}

			ListElement {
				name: "Banana"
				cost: 1.95
			}
		}

		delegate: RowLayout {
			Text { text: "Fruit: " + name }
			Text { text: "Cost: $" + cost }
		}
	}`)

	//List View + JSON Model
	createWidget("List View + JSON Model", `
	ListView {
		model: ListModel {

			Component.onCompleted: {
				var elements =
				[
					{ "name": "Apple", "cost": "2.45" },
					{ "name": "Orange", "cost": "3.25" },
					{ "name": "Banana", "cost": "1.95" }
				]

				for (var i = 0; i < elements.length; i++) {
					append(elements[i])
				}
			}
		}

		delegate: RowLayout {
			Text { text: "Fruit: " + name }
			Text { text: "Cost: $" + cost }
		}
	}`)

	//List View + Go Model
	createWidgetWithContext("List View + Go Model", `
	ListView {
		model: goListModel

		delegate: RowLayout {
			Text { text: "Fruit: " + name }
			Text { text: "Cost: $" + cost }
		}
	}`, "goListModel", listModel)

	//List View + Go String List Model
	createWidgetWithContext("List View + Go String List Model", `
	ListView {
		model: goListModel
		delegate: Text { text: display }
	}`, "goListModel", core.NewQStringListModel2([]string{"Some", "String", "List", "Items"}, nil))

	//Tree View + List Model
	//TODO:

	//Tree View + JSON Model
	//TODO:

	//Tree View + Go Model
	createWidgetWithContext("Tree View + Go Model", `
	TreeView {
		model: goTreeModel

		TableViewColumn {
			role: "display"
		}
	}`, "goTreeModel", treeModel)

	//Table View + List Model
	createWidget("Table View + List Model", `
	TableView {
		model: ListModel {

			ListElement {
				name: "Apple"
				cost: 2.45
			}

			ListElement {
				name: "Orange"
				cost: 3.25
			}

			ListElement {
				name: "Banana"
				cost: 1.95
			}
		}

		TableViewColumn {
			role: "name"
			title: role
		}

		TableViewColumn {
			role: "cost"
			title: role
		}
	}`)

	//Table View + JSON Model
	createWidget("Table View + JSON Model", `
	TableView {
		model: ListModel {

			Component.onCompleted: {
				var elements =
				[
					{ "name": "Apple", "cost": "2.45" },
					{ "name": "Orange", "cost": "3.25" },
					{ "name": "Banana", "cost": "1.95" }
				]

				for (var i = 0; i < elements.length; i++) {
					append(elements[i])
				}
			}
		}

		TableViewColumn {
			role: "name"
			title: role
		}

		TableViewColumn {
			role: "cost"
			title: role
		}
	}`)

	//Table View + Go Model
	createWidgetWithContext("Table View + Go Model", `
	TableView {
		model: goListModel

		TableViewColumn {
			role: "name"
			title: role
		}

		TableViewColumn {
			role: "cost"
			title: role
		}
	}`, "goListModel", listModel)

	//Grid View + List Model
	createWidget("Grid View + List Model", `
	GridView {
		cellWidth: width / (model.count / 2)
		cellHeight: height / (model.count / 2)

		model: ListModel {

			ListElement {
				color: "blue"
			}

			ListElement {
				color: "red"
			}

			ListElement {
				color: "green"
			}

			ListElement {
				color: "yellow"
			}
		}

		delegate: Rectangle {
			width: cellWidth
			height: cellHeight

			color: model.color
		}
	}`)

	//Grid View + JSON Model
	createWidget("Grid View + JSON Model", `
	GridView {
		cellWidth: width / (model.count / 2)
		cellHeight: height / (model.count / 2)

		model: ListModel {

			Component.onCompleted: {
				var elements =
				[
					{ "color": "blue" },
					{ "color": "red" },
					{ "color": "green" },
					{ "color": "yellow" }
				]

				for (var i = 0; i < elements.length; i++) {
					append(elements[i])
				}
			}
		}

		delegate: Rectangle {
			width: cellWidth
			height: cellHeight

			color: model.color
		}
	}`)

	//Grid View + Go Model
	createWidgetWithContext("Grid View + Go Model", `
	GridView {
		cellWidth: width / (model.count / 2)
		cellHeight: height / (model.count / 2)

		model: goListModel

		delegate: Rectangle {
			width: cellWidth
			height: cellHeight

			color: display
		}
	}`, "goListModel", core.NewQStringListModel2([]string{"blue", "red", "green", "yellow"}, nil))

}
