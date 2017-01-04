package main

import (
	"os"

	"github.com/therecipe/qt/core"
	"github.com/therecipe/qt/sql"
	"github.com/therecipe/qt/widgets"
)

func initializeModel(model *sql.QSqlQueryModel) {
	model.SetQuery2("select * from person", db)
	model.SetHeaderData(0, core.Qt__Horizontal, core.NewQVariant14("ID"), 0)
	model.SetHeaderData(1, core.Qt__Horizontal, core.NewQVariant14("First name"), 0)
	model.SetHeaderData(2, core.Qt__Horizontal, core.NewQVariant14("Last name"), 0)
}

var offset int

func createView(model sql.QSqlQueryModel_ITF, title string) *widgets.QTableView {
	var view = widgets.NewQTableView(nil)
	view.SetModel(model)

	view.SetWindowTitle(title)
	view.Move2(100+offset, 100+offset)
	offset += 20
	view.Show()

	return view
}

func main() {
	widgets.NewQApplication(len(os.Args), os.Args)

	if !createConnection() {
		return
	}

	var (
		plainModel    = sql.NewQSqlQueryModel(nil)
		editableModel = newEditableSqlModel(nil)
		customModel   = newCustomSqlModel(nil)
	)

	initializeModel(plainModel.QSqlQueryModel_PTR())
	initializeModel(editableModel.model.QSqlQueryModel_PTR())
	initializeModel(customModel.model.QSqlQueryModel_PTR())

	createView(plainModel, "Plain Query Model")
	createView(editableModel.model, "Editable Query Model")
	createView(customModel.model, "Custom Query Model")

	widgets.QApplication_Exec()
}
