package main

import (
	"strconv"
	"strings"

	"github.com/therecipe/qt/charts"
	"github.com/therecipe/qt/core"
	"github.com/therecipe/qt/gui"
	"github.com/therecipe/qt/widgets"
)

func NewTableWidget() *widgets.QWidget {
	this := widgets.NewQWidget(nil, 0)

	model := NewCustomTableModel(nil)

	tableView := widgets.NewQTableView(nil)
	tableView.SetModel(model)
	tableView.HorizontalHeader().SetSectionResizeMode(widgets.QHeaderView__Stretch)
	tableView.VerticalHeader().SetSectionResizeMode(widgets.QHeaderView__Stretch)

	chart := charts.NewQChart(nil, 0)
	chart.SetAnimationOptions(charts.QChart__AllAnimations)

	series := charts.NewQLineSeries(nil)
	series.SetName("Line 1")
	mapper := charts.NewQVXYModelMapper(this)
	mapper.SetXColumn(0)
	mapper.SetYColumn(1)
	mapper.SetSeries(series)
	mapper.SetModel(model)
	chart.AddSeries(series)

	seriesColorHex := "#000000"
	sch := strconv.FormatUint(uint64(series.Pen().Color().Rgb()), 16)
	seriesColorHex = "#" + strings.ToUpper(sch[len(sch)-6:])
	model.addMapping(seriesColorHex, core.NewQRect4(0, 0, 2, model.RowCount(nil)))

	series = charts.NewQLineSeries(nil)
	series.SetName("Line 2")
	mapper = charts.NewQVXYModelMapper(this)
	mapper.SetXColumn(2)
	mapper.SetYColumn(3)
	mapper.SetSeries(series)
	mapper.SetModel(model)
	chart.AddSeries(series)

	sch = strconv.FormatUint(uint64(series.Pen().Color().Rgb()), 16)
	seriesColorHex = "#" + strings.ToUpper(sch[len(sch)-6:])
	model.addMapping(seriesColorHex, core.NewQRect4(2, 0, 2, model.RowCount(nil)))

	chart.CreateDefaultAxes()
	chartView := charts.NewQChartView(nil)
	chartView.SetChart(chart)
	chartView.SetRenderHint(gui.QPainter__Antialiasing, true)
	chartView.SetMinimumSize2(640, 480)

	mainLayout := widgets.NewQGridLayout(nil)
	mainLayout.AddWidget2(tableView, 1, 0, 0)
	mainLayout.AddWidget2(chartView, 1, 1, 0)
	mainLayout.SetColumnStretch(1, 1)
	mainLayout.SetColumnStretch(0, 0)
	this.SetLayout(mainLayout)

	return this
}
