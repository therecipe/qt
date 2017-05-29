package main

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"time"

	"github.com/therecipe/qt/core"
	"github.com/therecipe/qt/gui"
	"github.com/therecipe/qt/quick"
	"github.com/therecipe/qt/widgets"
)

func displayWidgets() {

	//Label
	label := widgets.NewQLabel2("This Is A Label", nil, 0)
	label.SetWindowTitle("Label")
	addWidget(label)

	//Text Browser
	textBrowser := widgets.NewQTextBrowser(nil)
	textBrowser.SetWindowTitle("Text Browser")

	textBrowser.SetText("This Is A Text Browser")
	addWidget(textBrowser)

	//Graphics View
	graphicsView := widgets.NewQGraphicsView(nil)
	graphicsView.SetWindowTitle("Graphics View")

	scene := widgets.NewQGraphicsScene(nil)
	scene.AddText("This Is A Graphics View (+ Graphics Scene)", gui.NewQFont())
	scene.AddRect2(0, 0, scene.Width(), scene.Height(), gui.NewQPen(), gui.NewQBrush())

	graphicsView.SetScene(scene)
	addWidget(graphicsView)

	//Calendar Widget
	calendarWidget := widgets.NewQCalendarWidget(nil)
	calendarWidget.SetWindowTitle("Calendar Widget")
	addWidget(calendarWidget)

	//LCD Number
	lcdNumber := widgets.NewQLCDNumber(nil)
	lcdNumber.SetWindowTitle("LCD Number")

	lcdNumber.SetDigitCount(15)
	lcdNumber.Display("0123456789.-ABC")
	addWidget(lcdNumber)

	//Progress Bar
	progressBar := widgets.NewQProgressBar(nil)
	progressBar.SetWindowTitle("Progress Bar")

	progressBar.SetMinimum(0)
	progressBar.SetMaximum(1000)
	progressBar.SetValue(progressBar.Maximum() / 2)

	progressBar.ConnectValueChanged(func(value int) {
		if value == progressBar.Maximum() {
			progressBar.SetValue(progressBar.Minimum())
		}
	})

	go func() {
		for range time.NewTicker(500 * time.Millisecond).C {
			progressBar.SetValue(progressBar.Value() + 50)
		}
	}()
	addWidget(progressBar)

	//Quick Widget
	quickWidget := quick.NewQQuickWidget(nil)
	quickWidget.SetWindowTitle("Quick Widget")

	quickWidget.SetResizeMode(quick.QQuickWidget__SizeRootObjectToView)

	path := filepath.Join(os.TempDir(), "tmpQuickWidget.qml")
	ioutil.WriteFile(path, []byte("import QtQuick 2.0\nRectangle{width: 320; height: 320; color:\"red\"}"), 0644)
	quickWidget.SetSource(core.QUrl_FromLocalFile(path))
	addWidget(quickWidget)

}
