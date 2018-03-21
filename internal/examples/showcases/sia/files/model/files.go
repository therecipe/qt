package mfiles

import "github.com/therecipe/qt/core"

type File struct {
	Name          string
	Size          string
	Redundancy    string
	Available     bool
	ProgressText  string
	ProgressValue float64
	Error         string
}

type FilesModel struct {
	core.QAbstractTableModel

	files []File

	Filter *core.QSortFilterProxyModel

	_ func() `constructor:"init"`

	_ func([]File) `signal:"updateWith"`
}

func (m *FilesModel) init() {
	m.Filter = core.NewQSortFilterProxyModel(nil)
	m.Filter.SetSourceModel(m)
	m.Filter.SetFilterKeyColumn(0)
	m.Filter.SetFilterCaseSensitivity(core.Qt__CaseInsensitive)

	m.ConnectRowCount(m.rowCount)
	m.ConnectColumnCount(func(*core.QModelIndex) int { return 4 }) //needed for sort filter model
	m.ConnectData(m.data)
	m.ConnectRoleNames(m.roleNames)
	m.ConnectUpdateWith(m.updateWith)
}

func (m *FilesModel) rowCount(parent *core.QModelIndex) int {
	return len(m.files)
}

func (m *FilesModel) data(index *core.QModelIndex, role int) *core.QVariant {
	if role != int(core.Qt__DisplayRole) && role < int(core.Qt__UserRole) {
		return core.NewQVariant()
	}

	dbArray := m.files
	if index.Row() < 0 || index.Row() >= len(dbArray) {
		return core.NewQVariant()
	}

	dbItem := dbArray[index.Row()]

	switch {
	case index.Column() == 0 && role == int(core.Qt__DisplayRole) || //needed for sort filter model
		role == int(core.Qt__UserRole)+1:
		return core.NewQVariant14(dbItem.Name)

	case role == int(core.Qt__UserRole)+2:
		return core.NewQVariant14(dbItem.Size)

	case role == int(core.Qt__UserRole)+3:
		return core.NewQVariant14(dbItem.Redundancy)

	case role == int(core.Qt__UserRole)+4:
		return core.NewQVariant25(map[string]*core.QVariant{
			"available": core.NewQVariant11(dbItem.Available),
			"text":      core.NewQVariant14(dbItem.ProgressText),
			"value":     core.NewQVariant12(dbItem.ProgressValue),
			"error":     core.NewQVariant11(dbItem.Error != ""),
		})
	}

	return core.NewQVariant()
}

func (m *FilesModel) roleNames() map[int]*core.QByteArray {
	return map[int]*core.QByteArray{
		int(core.Qt__UserRole) + 1: core.NewQByteArray2("NAME", -1),
		int(core.Qt__UserRole) + 2: core.NewQByteArray2("SIZE", -1),
		int(core.Qt__UserRole) + 3: core.NewQByteArray2("REDUNDANCY", -1),
		int(core.Qt__UserRole) + 4: core.NewQByteArray2("ACTIONS", -1),
	}
}

func (m *FilesModel) updateWith(files []File) {
	if len(m.files) != len(files) {
		m.BeginResetModel()
		m.files = files
		m.EndResetModel()
	} else {
		m.files = files
		m.DataChanged(m.Index(0, 0, core.NewQModelIndex()), m.Index(len(m.files)-1, 0, core.NewQModelIndex()), make([]int, 0))
	}
}
