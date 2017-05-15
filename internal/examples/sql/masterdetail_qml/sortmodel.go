package main

import (
	"github.com/therecipe/qt/core"
)

var sortFilterModel *SortFilterModel

type SortFilterModel struct {
	*core.QSortFilterProxyModel
}

func NewSortFilterModel(source *core.QAbstractItemModel) *core.QAbstractItemModel {
	sortFilterModel = &SortFilterModel{core.NewQSortFilterProxyModel(nil)}

	sortFilterModel.SetSourceModel(source)

	return sortFilterModel.QAbstractItemModel_PTR()
}
