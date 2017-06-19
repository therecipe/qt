package model

import (
	"github.com/therecipe/qt/core"
)

var ListModel *listModel

type listModel struct {
	*core.QAbstractListModel
}

func NewListModel() *core.QAbstractItemModel {
	ListModel = &listModel{core.NewQAbstractListModel(nil)}

	ListModel.ConnectRowCount(ListModel.rowCount)
	ListModel.ConnectData(ListModel.data)

	return ListModel.QAbstractItemModel_PTR()
}

func (m *listModel) rowCount(parent *core.QModelIndex) int {
	return getArtistCount()
}

func (m *listModel) data(index *core.QModelIndex, role int) *core.QVariant {
	if role != int(core.Qt__DisplayRole) {
		return core.NewQVariant()
	}
	return core.NewQVariant14(getArtistForID(index.Row()).Name)
}
