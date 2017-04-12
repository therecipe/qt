package main

import (
	"os"

	"github.com/therecipe/qt/widgets"
)

func main() {
	widgets.NewQApplication(len(os.Args), os.Args)

	var (
		window = widgets.NewQMainWindow(nil, 0)
		vbox   = widgets.NewQVBoxLayout2(nil)
		button = widgets.NewQPushButton2("click me!", nil)
	)
	vbox.AddWidget(button, 0, 0)

	println("are the java classes available?", Runnable__IsClassAvailable(), Runnable__IsQtClassAvailable())

	var runner = NewQtRunnable()
	runner.SetObjectName("java")
	runner.ConnectRun(func() { println("hello from", runner.ObjectName()) })

	button.ConnectClicked(func(checked bool) {
		println("hello from go")
		runner.Run()
	})

	var centralWidget = widgets.NewQWidget(window, 0)
	centralWidget.SetLayout(vbox)
	window.SetCentralWidget(centralWidget)

	window.Show()
	widgets.QApplication_Exec()
}
