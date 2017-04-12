//author: https://github.com/5k3105

package main

import (
	"os"
	"strconv"

	"github.com/emirpasic/gods/lists/arraylist"
	"github.com/therecipe/qt/core"
	"github.com/therecipe/qt/gui"
	"github.com/therecipe/qt/widgets"
)

var statusbar *widgets.QStatusBar

func main() {
	widgets.NewQApplication(len(os.Args), os.Args)

	// Main Window
	var window = widgets.NewQMainWindow(nil, 0)
	window.SetWindowTitle("Table")

	// Main Widget
	var table = NewArgsTable()

	// Set Central Widget
	window.SetCentralWidget(table)

	// Statusbar
	statusbar = widgets.NewQStatusBar(window)
	window.SetStatusBar(statusbar)

	// Run App
	widgets.QApplication_SetStyle2("fusion")
	window.Show()
	widgets.QApplication_Exec()
}

var (
	list  *arraylist.List
	model *core.QAbstractTableModel
	view  *widgets.QTableView

	textIndex  *core.QModelIndex
	currentRow int
)

type Delegate struct {
	widgets.QStyledItemDelegate //don't use *pointers or it won't work
}

func NewArgsTable() *widgets.QTableView {
	view = widgets.NewQTableView(nil)
	model = core.NewQAbstractTableModel(nil)

	delegate := InitDelegate()
	view.SetItemDelegate(delegate)
	view.SetFont(gui.NewQFont2("verdana", 13, 1, false))

	view.ConnectKeyPressEvent(keypressevent)
	view.ConnectCurrentChanged(currentchanged)

	model.ConnectRowCount(rowcount)
	model.ConnectColumnCount(columncount)
	model.ConnectData(data)
	model.ConnectSetData(setdata)
	model.ConnectInsertRows(insertrows)
	model.ConnectFlags(flags)
	model.ConnectRemoveRows(removerows)
	model.ConnectHeaderData(headerdata)

	list = arraylist.New()

	for i := 0; i < 6; i++ {
		list.Add(strconv.Itoa(i + i))
	}
	list.Add("") // blank

	view.SetModel(model)
	return view
}

func keypressevent(e *gui.QKeyEvent) {
	if e.Key() == int(core.Qt__Key_Delete) {
		if currentRow < list.Size()-1 {
			model.RemoveRows(currentRow, 1, core.NewQModelIndex())
		}
	} else {
		view.KeyPressEventDefault(e)
	}
}

func currentchanged(current *core.QModelIndex, previous *core.QModelIndex) {
	currentRow = current.Row()
}

func InitDelegate() *Delegate {
	item := NewDelegate(nil) //will be generated in moc.go
	item.ConnectCreateEditor(createEditor)
	item.ConnectSetEditorData(setEditorData)
	item.ConnectSetModelData(setModelData)
	item.ConnectUpdateEditorGeometry(updateEditorGeometry)
	return item
}

func createEditor(parent *widgets.QWidget, option *widgets.QStyleOptionViewItem, index *core.QModelIndex) *widgets.QWidget {
	editor := widgets.NewQLineEdit(parent)
	textIndex = index

	if index.Row() == list.Size()-1 {
		model.InsertRow(index.Row(), core.NewQModelIndex())
	}

	editor.ConnectTextChanged(textchanged)
	return editor.QWidget_PTR()
}

func textchanged(text string) {
	model.SetData(textIndex, core.NewQVariant14(text), 2) // edit role
}

func setEditorData(editor *widgets.QWidget, index *core.QModelIndex) {
	value, _ := list.Get(index.Row())
	lineedit := widgets.NewQLineEditFromPointer(editor.Pointer())
	lineedit.SetText(value.(string))
}

func setModelData(editor *widgets.QWidget, model *core.QAbstractItemModel, index *core.QModelIndex) {
	lineedit := widgets.NewQLineEditFromPointer(editor.Pointer())
	text := lineedit.Text()
	model.SetData(index, core.NewQVariant14(text), int(core.Qt__EditRole))

}

func updateEditorGeometry(editor *widgets.QWidget, option *widgets.QStyleOptionViewItem, index *core.QModelIndex) {
	editor.SetGeometry(option.Rect())
}

func headerdata(section int, orientation core.Qt__Orientation, role int) *core.QVariant {
	if orientation == 1 && role == 0 { // Qt__Horizontal, Qt__DisplayRole
		return core.NewQVariant14("column" + strconv.Itoa(section+1))
	}

	if orientation == 2 && role == 0 {
		if section < list.Size()-1 {
			return core.NewQVariant14(strconv.Itoa(section + 1))
		} else {
			return core.NewQVariant14("*")
		}
	}
	return core.NewQVariant()
}

func rowcount(parent *core.QModelIndex) int {
	return list.Size()
}

func columncount(parent *core.QModelIndex) int {
	return 1
}

func data(index *core.QModelIndex, role int) *core.QVariant {
	if role == 0 && index.IsValid() { // display role
		text, exists := list.Get(index.Row())
		if exists {
			switch deducedText := text.(type) {
			case int:
				{
					return core.NewQVariant7(deducedText)
				}
			case string:
				{
					return core.NewQVariant14(deducedText)
				}
			}
		}
	}
	return core.NewQVariant()
}

func setdata(index *core.QModelIndex, value *core.QVariant, role int) bool {
	if (role == 2 || role == 0) && index.IsValid() { // edit role
		list.Remove(index.Row())
		list.Insert(index.Row(), value.ToString())
		return true
	}
	return true
}

func insertrows(row int, count int, parent *core.QModelIndex) bool {
	model.BeginInsertRows(core.NewQModelIndex(), row, row)
	list.Add("")
	model.EndInsertRows()
	view.SelectRow(row)
	return true
}

func removerows(row int, count int, parent *core.QModelIndex) bool {
	model.BeginRemoveRows(core.NewQModelIndex(), row, row)
	list.Remove(row)
	model.EndRemoveRows()
	return true
}

func flags(index *core.QModelIndex) core.Qt__ItemFlag {
	return 35 // ItemIsSelectable || ItemIsEditable || ItemIsEnabled
}
