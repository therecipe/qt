package main

import (
	"strings"

	"github.com/therecipe/qt/core"
	"github.com/therecipe/qt/gui"
	"github.com/therecipe/qt/sql"
)

type CustomSqlModel struct {
	*sql.QSqlQueryModel
}

func newCustomSqlModel(p *core.QObject) *CustomSqlModel {
	var model = &CustomSqlModel{sql.NewQSqlQueryModel(p)}

	model.ConnectData(model.data)

	return model
}

func (m *CustomSqlModel) data(index *core.QModelIndex, role int) *core.QVariant {

	var value = m.DataDefault(index, role)
	if value.IsValid() && role == int(core.Qt__DisplayRole) {
		if index.Column() == 0 {
			return core.NewQVariant14("#" + value.ToString())
		} else if index.Column() == 2 {
			return core.NewQVariant14(strings.ToUpper(value.ToString()))
		}
	}

	if role == int(core.Qt__TextColorRole) && index.Column() == 1 {
		return gui.NewQColor2(core.Qt__blue).ToVariant()
	}

	return value
}
