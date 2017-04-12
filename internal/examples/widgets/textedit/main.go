//source: http://doc.qt.io/qt-5/qtwidgets-richtext-textedit-example.html

package main

import (
	"os"

	"github.com/therecipe/qt/core"
	"github.com/therecipe/qt/widgets"
)

var qApp *widgets.QApplication

func main() {
	qApp = widgets.NewQApplication(len(os.Args), os.Args)
	core.QCoreApplication_SetOrganizationName("QtProject")
	core.QCoreApplication_SetApplicationName("Rich Text")
	core.QCoreApplication_SetApplicationVersion("1.0.0")

	var parser = core.NewQCommandLineParser()
	parser.SetApplicationDescription(core.QCoreApplication_ApplicationName())
	parser.AddHelpOption()
	parser.AddVersionOption()
	parser.AddPositionalArgument("file", "The file to open.", "")
	parser.Process2(qApp)

	var mw = initTextEdit()

	var availableGeometry = widgets.QApplication_Desktop().AvailableGeometry2(mw)
	//mw.Resize2(availableGeometry.Width()/2, (availableGeometry.Height()*2)/3)
	mw.Resize2(600, 450)
	mw.Move2((availableGeometry.Width()-mw.Width())/2,
		(availableGeometry.Height()-mw.Height())/2)

	var (
		args     = parser.PositionalArguments()
		fileName string
	)
	if len(args) >= 1 && len(args[0]) >= 1 {
		fileName = args[0]
	} else {
		fileName = ":/qml/example.html"
	}

	if !mw.load(fileName) {
		mw.fileNew()
	}

	mw.Show()
	qApp.Exec()
}
