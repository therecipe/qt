package main

import (
	"os"

	"github.com/therecipe/qt/core"
	"github.com/therecipe/qt/gui"
	"github.com/therecipe/qt/widgets"
)

var (
	centralLayout       *widgets.QGridLayout
	centralLayoutRow    int
	centralLayoutColumn int
)

func main() {
	widgets.NewQApplication(len(os.Args), os.Args)

	mainWindow := widgets.NewQMainWindow(nil, 0)
	mainWindow.SetWindowTitle("Common Widgets")

	scrollWidget := widgets.NewQScrollArea(nil)

	centralWidget := widgets.NewQWidget(nil, 0)
	centralLayout = widgets.NewQGridLayout(centralWidget)

	layouts()
	buttons()
	itemViews()
	itemWidgets()
	containers()
	inputWidgets()
	displayWidgets()

	scrollWidget.SetWidget(centralWidget)

	mainWindow.SetCentralWidget(scrollWidget)
	mainWindow.ShowMaximized()

	//->needed to work around some iOS issue: https://github.com/therecipe/qt/issues/451
	gui.QGuiApplication_Screens()[0].SetOrientationUpdateMask(core.Qt__PrimaryOrientation | core.Qt__LandscapeOrientation | core.Qt__PortraitOrientation | core.Qt__InvertedLandscapeOrientation | core.Qt__InvertedPortraitOrientation)
	gui.QGuiApplication_Screens()[0].ConnectOrientationChanged(func(core.Qt__ScreenOrientation) {
		mainWindow.Hide()
		go mainWindow.Show()
	})
	//<-

	widgets.QApplication_Exec()
}

func addWidget(widget widgets.QWidget_ITF) {

	if centralLayoutColumn > 6 {
		centralLayoutColumn = 0
		centralLayoutRow++
	}

	wrappedWidget := widgets.NewQGroupBox2(widget.QWidget_PTR().WindowTitle(), nil)
	wrappedWidgetLayout := widgets.NewQVBoxLayout2(wrappedWidget)
	wrappedWidgetLayout.AddWidget(widget, 0, core.Qt__AlignCenter)
	wrappedWidget.SetFixedSize2(250, 250)

	centralLayout.AddWidget2(wrappedWidget, centralLayoutRow, centralLayoutColumn, core.Qt__AlignCenter)

	centralLayoutColumn++
}
