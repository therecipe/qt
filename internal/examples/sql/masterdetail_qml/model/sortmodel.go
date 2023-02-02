package model

import (
	"github.com/bluszcz/cutego/core"
)

var SortFilterModel *sortFilterModel

type sortFilterModel struct {
	*core.QSortFilterProxyModel
}

func NewSortFilterModel() *core.QAbstractItemModel {
	SortFilterModel = &sortFilterModel{core.NewQSortFilterProxyModel(nil)}

	SortFilterModel.SetSourceModel(NewViewModel())

	return SortFilterModel.QAbstractItemModel_PTR()
}
