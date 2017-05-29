package main

import (
	"fmt"

	"github.com/therecipe/qt/core"
	"github.com/therecipe/qt/widgets"
)

func containers() {

	//Group Box
	groupBox := widgets.NewQGroupBox2("Group Box", nil)
	groupBox.SetWindowTitle("Group Box")
	groupBoxLayout := widgets.NewQVBoxLayout2(groupBox)

	for i := 0; i < 3; i++ {
		groupBoxLayout.AddWidget(widgets.NewQPushButton2(fmt.Sprintf("PushButton: %v", i), nil), 0, 0)
	}
	addWidget(groupBox)

	//Scroll Area
	scrollArea := widgets.NewQScrollArea(nil)
	scrollArea.SetWindowTitle("Scroll Area")

	scrollAreaWidget := widgets.NewQWidget(nil, 0)
	scrollAreaWidgetLayout := widgets.NewQGridLayout(scrollAreaWidget)
	for row := 0; row < 25; row++ {
		for column := 0; column < 25; column++ {
			scrollAreaWidgetLayout.AddWidget(widgets.NewQLabel2(fmt.Sprintf("[%v:%v]", row, column), nil, 0), row, column, core.Qt__AlignCenter)
		}
	}

	scrollArea.SetWidget(scrollAreaWidget)
	addWidget(scrollArea)

	//Tool Box
	toolBox := widgets.NewQToolBox(nil, 0)
	toolBox.SetWindowTitle("Tool Box")

	toolBox.AddItem2(widgets.NewQLabel2("First Widget", nil, 0), "First")
	toolBox.AddItem2(widgets.NewQPushButton2("Second Widget", nil), "Second")
	toolBox.AddItem2(widgets.NewQGroupBox2("Third Widget", nil), "Third")
	addWidget(toolBox)

	//Tab Widget
	tabWidget := widgets.NewQTabWidget(nil)
	tabWidget.SetWindowTitle("Tab Widget")

	tabWidget.AddTab(widgets.NewQLabel2("First Widget", nil, 0), "First")
	tabWidget.AddTab(widgets.NewQPushButton2("Second Widget", nil), "Second")
	tabWidget.AddTab(widgets.NewQGroupBox2("Third Widget", nil), "Third")
	addWidget(tabWidget)

	//Stacked Widget
	centralWidget := widgets.NewQWidget(nil, 0)
	centralWidget.SetWindowTitle("Stacked Widget")
	centralWidgetLayout := widgets.NewQVBoxLayout2(centralWidget)

	stackedWidget := widgets.NewQStackedWidget(nil)
	stackedWidget.AddWidget(widgets.NewQLabel2("First Widget", nil, 0))
	stackedWidget.AddWidget(widgets.NewQPushButton2("Second Widget", nil))
	stackedWidget.AddWidget(widgets.NewQGroupBox2("Third Widget", nil))
	centralWidgetLayout.AddWidget(stackedWidget, 0, 0)

	changeStackButton := widgets.NewQPushButton2("Show Next In Stack", nil)
	changeStackButton.ConnectClicked(func(checked bool) {
		nextIndex := stackedWidget.CurrentIndex() + 1
		if nextIndex >= stackedWidget.Count() {
			stackedWidget.SetCurrentIndex(0)
		} else {
			stackedWidget.SetCurrentIndex(nextIndex)
		}
	})
	centralWidgetLayout.AddWidget(changeStackButton, 0, 0)
	addWidget(centralWidget)

	//Frame
	frame := widgets.NewQFrame(nil, 0)
	frame.SetWindowTitle("Frame")

	frameLayout := widgets.NewQVBoxLayout2(frame)

	for i := 0; i < 3; i++ {
		someFrame := widgets.NewQFrame(nil, 0)

		switch i {
		case 0:
			someFrame.SetFrameStyle(int(widgets.QFrame__Box) | int(widgets.QFrame__Raised))

		case 1:
			someFrame.SetFrameStyle(int(widgets.QFrame__Panel) | int(widgets.QFrame__Raised))

		case 2:
			someFrame.SetFrameStyle(int(widgets.QFrame__StyledPanel) | int(widgets.QFrame__Raised))
		}

		someFrameLayout := widgets.NewQVBoxLayout2(someFrame)
		someFrameLayout.AddWidget(widgets.NewQPushButton2(fmt.Sprintf("PushButton: %v", i), nil), 0, 0)
		frameLayout.AddWidget(someFrame, 0, 0)
	}
	addWidget(frame)

	//Widget
	widget := widgets.NewQWidget(nil, 0)
	widget.SetWindowTitle("Widget")

	widgetLayout := widgets.NewQVBoxLayout2(widget)

	for i := 0; i < 3; i++ {
		someWidget := widgets.NewQWidget(nil, 0)
		someWidgetLayout := widgets.NewQVBoxLayout2(someWidget)
		someWidgetLayout.AddWidget(widgets.NewQPushButton2(fmt.Sprintf("PushButton: %v", i), nil), 0, 0)
		widgetLayout.AddWidget(someWidget, 0, 0)
	}
	addWidget(widget)

	//MDI Area
	mdiArea := widgets.NewQMdiArea(nil)
	mdiArea.SetWindowTitle("MDI Area")

	subWindow := widgets.NewQMdiSubWindow(nil, 0)
	sWCentralWidget := widgets.NewQWidget(nil, 0)
	sWCentralWidgetLayout := widgets.NewQVBoxLayout2(sWCentralWidget)

	sWCentralWidgetLayout.AddWidget(widgets.NewQLabel2("Label", nil, 0), 0, 0)
	sWCentralWidgetLayout.AddWidget(widgets.NewQPushButton2("PushButton", nil), 0, 0)

	subWindow.SetWidget(sWCentralWidget)
	mdiArea.AddSubWindow(subWindow, 0)

	mdiArea.Resize2(300, 300)
	addWidget(mdiArea)

	//Dock Widget
	mainWindow := widgets.NewQMainWindow(nil, 0)
	mainWindow.SetWindowTitle("Dock Widget")

	topDockWidget := widgets.NewQDockWidget("Top Dock Widget", nil, 0)
	topDockWidget.SetAllowedAreas(core.Qt__AllDockWidgetAreas)
	topDockWidget.SetFloating(true)
	topDockWidget.SetWidget(widgets.NewQPushButton2("PushButton", nil))
	mainWindow.AddDockWidget(core.Qt__TopDockWidgetArea, topDockWidget)

	bottomDockWidget := widgets.NewQDockWidget("Bottom Dock Widget", nil, 0)
	bottomDockWidget.SetAllowedAreas(core.Qt__AllDockWidgetAreas)
	bottomDockWidget.SetFloating(true)
	bottomDockWidget.SetWidget(widgets.NewQPushButton2("PushButton", nil))
	mainWindow.AddDockWidget(core.Qt__BottomDockWidgetArea, bottomDockWidget)
	addWidget(mainWindow)
}
