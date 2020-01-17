package main

import (
	"fmt"

	"github.com/therecipe/qt/core"
)

type CppDataModel struct {
	core.QObject

	_ func() `constructor:"init"`

	_ func() `slot:"loadData,auto"`

	_ func(jsonDataString string) `signal:"dataLoaded"`
}

func (m *CppDataModel) init() {
	// set timer to send new data every 3 seconds
	timer := core.NewQTimer(nil)
	timer.ConnectTimeout(m.update)
	timer.Start(500)
}

func (m *CppDataModel) loadData() {
	// load data asynchronously or do C++ calculations here
	// for simplicity this example uses dummy data
	dummyData := m.createRandomData()

	// signal that new data is available
	m.DataLoaded(dummyData)
}

func (m *CppDataModel) createRandomData() string {
	dummyData := "["
	dummyData += m.createRandomEntry(2017, "Burghausen") + ","
	dummyData += m.createRandomEntry(2017, "Braunau") + ","
	dummyData += m.createRandomEntry(2017, "Munich") + ","
	dummyData += m.createRandomEntry(2017, "Vienna") + ","
	dummyData += m.createRandomEntry(2018, "Burghausen") + ","
	dummyData += m.createRandomEntry(2018, "Braunau") + ","
	dummyData += m.createRandomEntry(2018, "Munich") + ","
	dummyData += m.createRandomEntry(2018, "Vienna")
	dummyData += "]"
	return dummyData
}

func (m *CppDataModel) createRandomEntry(year int, city string) string {
	return fmt.Sprintf("{ \"year\": \"%v\", \"city\": \"%v\", \"expenses\": %v }", year, city, core.QRandomGenerator_Global().Bounded5(1e4, 1e5))
}

func (m *CppDataModel) update() {
	m.loadData()
}
