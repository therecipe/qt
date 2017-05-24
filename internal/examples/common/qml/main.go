package main

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"

	"github.com/therecipe/qt/core"
	"github.com/therecipe/qt/quick"
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
	mainWindow.SetWindowTitle("Common QML (+ Quick Controls 1)")

	scrollWidget := widgets.NewQScrollArea(nil)

	centralWidget := widgets.NewQWidget(nil, 0)
	centralLayout = widgets.NewQGridLayout(centralWidget)

	layouts()
	buttons()
	itemViews()
	containers()
	inputWidgets()
	displayWidgets()

	scrollWidget.SetWidget(centralWidget)

	mainWindow.SetCentralWidget(scrollWidget)
	mainWindow.ShowMaximized()

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

	widget.QWidget_PTR().SetFixedSize2(200, 200)

	wrappedWidget.SetFixedSize2(250, 250)

	centralLayout.AddWidget(wrappedWidget, centralLayoutRow, centralLayoutColumn, core.Qt__AlignCenter)

	centralLayoutColumn++
}

func createWidgetWithContext(name, code, ctxName string, ctx core.QObject_ITF) {
	quickWidget := quick.NewQQuickWidget(nil)
	quickWidget.SetWindowTitle(name)

	quickWidget.RootContext().SetContextProperty(ctxName, ctx)

	quickWidget.SetResizeMode(quick.QQuickWidget__SizeRootObjectToView)

	path := filepath.Join(os.TempDir(), "tmp"+strings.Replace(name, " ", "", -1)+".qml")
	ioutil.WriteFile(path, []byte("import QtQuick 2.0\nimport QtQuick.Layouts 1.3\nimport QtQuick.Controls 1.4\n"+code), 0644)
	quickWidget.SetSource(core.QUrl_FromLocalFile(path))
	addWidget(quickWidget)
}

func createWidget(name, code string) {
	createWidgetWithContext(name, code, "", nil)
}
