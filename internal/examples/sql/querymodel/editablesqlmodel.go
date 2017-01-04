package main

import (
	"github.com/therecipe/qt/core"
	"github.com/therecipe/qt/sql"
)

//TODO: really subclass with qtmoc
type EditableSqlModel struct {
	model *sql.QSqlQueryModel
}

func newEditableSqlModel(p *core.QObject) *EditableSqlModel {
	var model = &EditableSqlModel{model: sql.NewQSqlQueryModel(p)}

	model.model.ConnectFlags(model.flags)
	model.model.ConnectSetData(model.setData)

	return model
}

func (m *EditableSqlModel) flags(index *core.QModelIndex) core.Qt__ItemFlag {
	var flags = m.model.FlagsDefault(index)
	if index.Column() == 1 || index.Column() == 2 {
		flags |= core.Qt__ItemIsEditable
	}

	return flags
}

func (m *EditableSqlModel) setData(index *core.QModelIndex, value *core.QVariant, role int) bool {

	if index.Column() < 1 || index.Column() > 2 {
		return false
	}

	var primaryKeyIndex = m.model.Index(index.Row(), 0, core.NewQModelIndex())
	var id = m.model.Data(primaryKeyIndex, 0).ToInt(true)

	m.model.Clear()

	var ok bool

	if index.Column() == 1 {
		ok = m.setFirstName(id, value.ToString())
	} else {
		ok = m.setLastName(id, value.ToString())
	}

	m.refresh()

	return ok
}

func (m *EditableSqlModel) refresh() {
	m.model.SetQuery2("select * from person", db)
	m.model.SetHeaderData(0, core.Qt__Horizontal, core.NewQVariant14("ID"), 0)
	m.model.SetHeaderData(1, core.Qt__Horizontal, core.NewQVariant14("First Name"), 0)
	m.model.SetHeaderData(2, core.Qt__Horizontal, core.NewQVariant14("Last Name"), 0)
}

func (m *EditableSqlModel) setFirstName(personId int, firstName string) bool {
	var query = sql.NewQSqlQuery3(db)
	query.Prepare("update person set firstname = ? where id = ?")
	query.AddBindValue(core.NewQVariant14(firstName), 0)
	query.AddBindValue(core.NewQVariant7(personId), 0)
	return query.Exec2()
}

func (m *EditableSqlModel) setLastName(personId int, lastName string) bool {
	var query = sql.NewQSqlQuery3(db)
	query.Prepare("update person set lastname = ? where id = ?")
	query.AddBindValue(core.NewQVariant14(lastName), 0)
	query.AddBindValue(core.NewQVariant7(personId), 0)
	return query.Exec2()
}
