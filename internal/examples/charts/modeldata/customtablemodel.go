package main

import (
	"github.com/therecipe/qt/core"
	"github.com/therecipe/qt/gui"
)

type CustomTableModel struct {
	core.QAbstractTableModel

	_ func() `constructor:"init"`

	m_data        [][]float64
	m_mapping     map[string]*core.QRect
	m_columnCount int
	m_rowCount    int
}

func (c *CustomTableModel) init() {
	c.m_data = make([][]float64, 0)
	c.m_mapping = make(map[string]*core.QRect)
	c.m_columnCount = 4
	c.m_rowCount = 15

	for i := 0; i < c.m_rowCount; i++ {
		dataVec := make([]float64, c.m_columnCount)
		for k := 0; k < len(dataVec); k++ {
			if k%2 == 0 {
				dataVec[k] = float64(i*50) + core.QRandomGenerator_Global().Bounded(20)
			} else {
				dataVec[k] = core.QRandomGenerator_Global().Bounded(100)
			}
		}
		c.m_data = append(c.m_data, dataVec)
	}

	c.ConnectRowCount(func(*core.QModelIndex) int { return len(c.m_data) })
	c.ConnectColumnCount(func(*core.QModelIndex) int { return c.m_columnCount })

	c.ConnectHeaderData(func(section int, orientation core.Qt__Orientation, role int) *core.QVariant {
		if role != int(core.Qt__DisplayRole) {
			return core.NewQVariant()
		}

		if orientation == core.Qt__Horizontal {
			if section%2 == 0 {
				return core.NewQVariant1("x")
			}
			return core.NewQVariant1("y")
		}
		return core.NewQVariant1(section + 1)
	})

	c.ConnectData(func(index *core.QModelIndex, role int) *core.QVariant {
		if role == int(core.Qt__DisplayRole) || role == int(core.Qt__EditRole) {
			return core.NewQVariant1(int(c.m_data[index.Row()][index.Column()]))
		} else if role == int(core.Qt__BackgroundRole) {
			for i, rect := range c.m_mapping {
				if rect.Contains3(index.Column(), index.Row()) {
					return core.NewQVariant1(gui.NewQColor6(i))
				}
			}
			return core.NewQVariant1(gui.NewQColor2(core.Qt__white))
		}
		return core.NewQVariant()
	})

	c.ConnectSetData(func(index *core.QModelIndex, value *core.QVariant, role int) bool {
		if index.IsValid() && role == int(core.Qt__EditRole) {
			c.m_data[index.Row()][index.Column()] = value.ToDouble(nil)
			c.DataChanged(index, index, nil)
			return true
		}
		return false
	})

	c.ConnectFlags(func(index *core.QModelIndex) core.Qt__ItemFlag {
		return c.FlagsDefault(index) | core.Qt__ItemIsEditable
	})
}

func (c *CustomTableModel) clearMapping()                             { c.m_mapping = make(map[string]*core.QRect) }
func (c *CustomTableModel) addMapping(color string, area *core.QRect) { c.m_mapping[color] = area }
