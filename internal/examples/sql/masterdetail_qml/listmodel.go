package main

import (
	"github.com/therecipe/qt/core"
)

var listModel *ListModel

type ListModel struct {
	*core.QAbstractListModel
}

func NewListModel() *core.QAbstractItemModel {
	listModel = &ListModel{core.NewQAbstractListModel(nil)}

	listModel.ConnectRowCount(listModel.rowCount)
	listModel.ConnectData(listModel.data)

	return listModel.QAbstractItemModel_PTR()
}

func (m *ListModel) rowCount(parent *core.QModelIndex) int {
	return getArtistCount()
}

func (m *ListModel) data(index *core.QModelIndex, role int) *core.QVariant {
	if role != int(core.Qt__DisplayRole) {
		return core.NewQVariant()
	}
	return core.NewQVariant14(getArtistForID(index.Row()).Name)
}
