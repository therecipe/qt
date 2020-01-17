//source: https://doc.qt.io/qt-5/qtcharts-dynamicspline-example.html

package main

import (
	"os"

	"github.com/therecipe/qt/charts"
	"github.com/therecipe/qt/gui"
	"github.com/therecipe/qt/widgets"
)

func main() {
	widgets.NewQApplication(len(os.Args), os.Args)

	window := widgets.NewQMainWindow(nil, 0)
	chart := NewChart(nil)
	chart.SetTitle("Dynamic spline chart")
	chart.Legend().Hide()
	chart.SetAnimationOptions(charts.QChart__AllAnimations)

	chartView := charts.NewQChartView2(chart, nil)
	chartView.SetRenderHint(gui.QPainter__Antialiasing, true)
	window.SetCentralWidget(chartView)
	window.Resize2(400, 300)
	window.Show()

	widgets.QApplication_Exec()
}
