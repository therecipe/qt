package model

import "github.com/therecipe/qt/core"

type Transaction struct {
	Status string
	Date   string
	Amount string
	Type   string
	Total  string
	ID     string
}

type WalletModel struct {
	core.QAbstractTableModel

	transactions []Transaction

	_ func() `constructor:"init"`

	_ func([]Transaction) `signal:"updateWith,auto"`
}

func (m *WalletModel) init() {
	m.ConnectRowCount(m.rowCount)
	m.ConnectColumnCount(func(*core.QModelIndex) int { return 6 })
	m.ConnectData(m.data)
	m.ConnectRoleNames(m.roleNames)
}

func (m *WalletModel) rowCount(parent *core.QModelIndex) int {
	return len(m.transactions)
}

func (m *WalletModel) data(index *core.QModelIndex, role int) *core.QVariant {
	if role != int(core.Qt__DisplayRole) && role < int(core.Qt__UserRole) {
		return core.NewQVariant()
	}

	dbArray := m.transactions
	if index.Row() < 0 || index.Row() >= len(dbArray) {
		return core.NewQVariant()
	}

	dbItem := dbArray[index.Row()]

	switch {
	case role == int(core.Qt__UserRole)+1:
		return core.NewQVariant1(dbItem.Status)

	case role == int(core.Qt__UserRole)+2:
		return core.NewQVariant1(dbItem.Date)

	case role == int(core.Qt__UserRole)+3:
		return core.NewQVariant1(dbItem.Amount)

	case role == int(core.Qt__UserRole)+4:
		return core.NewQVariant1(dbItem.Type)

	case role == int(core.Qt__UserRole)+5:
		return core.NewQVariant1(dbItem.Total)

	case role == int(core.Qt__UserRole)+6:
		return core.NewQVariant1(dbItem.ID)
	}

	return core.NewQVariant()
}

func (m *WalletModel) roleNames() map[int]*core.QByteArray {
	return map[int]*core.QByteArray{
		int(core.Qt__UserRole) + 1: core.NewQByteArray2("STATUS", -1),
		int(core.Qt__UserRole) + 2: core.NewQByteArray2("DATE AND TIME", -1),
		int(core.Qt__UserRole) + 3: core.NewQByteArray2("AMOUNT", -1),
		int(core.Qt__UserRole) + 4: core.NewQByteArray2("TYPE", -1),
		int(core.Qt__UserRole) + 5: core.NewQByteArray2("TOTAL BALANCE", -1),
		int(core.Qt__UserRole) + 6: core.NewQByteArray2("ID", -1),
	}
}

func (m *WalletModel) updateWith(transactions []Transaction) {
	if len(m.transactions) != len(transactions) {
		m.BeginResetModel()
		m.transactions = transactions
		m.EndResetModel()
	} else {
		m.transactions = transactions
		m.DataChanged(m.Index(0, 0, core.NewQModelIndex()), m.Index(len(m.transactions)-1, 0, core.NewQModelIndex()), make([]int, 0))
	}
}
