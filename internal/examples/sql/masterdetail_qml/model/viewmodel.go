package model

import (
	"github.com/therecipe/qt/core"
)

type ViewModel struct {
	*core.QAbstractTableModel
}

func NewViewModel() *core.QAbstractItemModel {
	model := &ViewModel{core.NewQAbstractTableModel(nil)}

	model.ConnectHeaderData(model.headerData)
	model.ConnectRowCount(model.rowCount)
	model.ConnectColumnCount(model.columnCount)
	model.ConnectData(model.data)
	model.ConnectRoleNames(model.roleNames)

	return model.QAbstractItemModel_PTR()
}

func (m *ViewModel) headerData(section int, orientation core.Qt__Orientation, role int) *core.QVariant {
	if role != int(core.Qt__DisplayRole) && role < int(core.Qt__UserRole) || orientation != core.Qt__Horizontal {
		return core.NewQVariant()
	}

	switch section {
	case 0:
		return core.NewQVariant14("ID")

	case 1:
		return core.NewQVariant14("Title")

	case 2:
		return core.NewQVariant14("Artist")

	case 3:
		return core.NewQVariant14("Year")
	}

	return core.NewQVariant()
}

func (m *ViewModel) rowCount(parent *core.QModelIndex) int {
	return getAlbumCount()
}

func (m *ViewModel) columnCount(parent *core.QModelIndex) int {
	return 4 //	ID | Title | Artist | Year
}

func (m *ViewModel) data(index *core.QModelIndex, role int) *core.QVariant {
	if role != int(core.Qt__DisplayRole) && role < int(core.Qt__UserRole) {
		return core.NewQVariant()
	}

	dbArray := getModelArray()
	if index.Row() < 0 || index.Row() >= len(dbArray) {
		return core.NewQVariant()
	}

	dbItem := dbArray[index.Row()]

	switch {
	case
		index.Column() == 0 && role == int(core.Qt__DisplayRole) || //widgets
			role == int(core.Qt__UserRole)+1: //qml

		return core.NewQVariant7(dbItem.Album.ID)

	case
		index.Column() == 1 && role == int(core.Qt__DisplayRole) || //widgets
			role == int(core.Qt__UserRole)+2: //qml

		return core.NewQVariant14(dbItem.Album.Title)

	case
		index.Column() == 2 && role == int(core.Qt__DisplayRole) || //widgets
			role == int(core.Qt__UserRole)+3: //qml

		return core.NewQVariant14(dbItem.Artist.Name)

	case
		index.Column() == 3 && role == int(core.Qt__DisplayRole) || //widgets
			role == int(core.Qt__UserRole)+4: //qml

		return core.NewQVariant7(dbItem.Album.Year)
	}

	return core.NewQVariant()
}

//needed only for qml
func (m *ViewModel) roleNames() map[int]*core.QByteArray {
	return map[int]*core.QByteArray{
		int(core.Qt__UserRole) + 1: core.NewQByteArray2("ID", -1),
		int(core.Qt__UserRole) + 2: core.NewQByteArray2("Title", -1),
		int(core.Qt__UserRole) + 3: core.NewQByteArray2("Artist", -1),
		int(core.Qt__UserRole) + 4: core.NewQByteArray2("Year", -1),
	}
}
