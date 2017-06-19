// +build !qml

package view

import (
	"github.com/therecipe/qt/gui"
	"github.com/therecipe/qt/widgets"
)

type menuBarController struct {
	widgets.QMenuBar

	_ func() `constructor:"init"`
}

func (m *menuBarController) init() {

	addAction := widgets.NewQAction2("&Add album...", m)
	deleteAction := widgets.NewQAction2("&Delete album...", m)
	quitAction := widgets.NewQAction2("&Quit", m)
	aboutAction := widgets.NewQAction2("&About", m)
	aboutQtAction := widgets.NewQAction2("About &Qt", m)

	addAction.SetShortcut(gui.QKeySequence_FromString("Ctrl+A", 0))
	deleteAction.SetShortcut(gui.QKeySequence_FromString("Ctrl+D", 0))
	quitAction.SetShortcuts2(gui.QKeySequence__Quit)

	fileMenu := m.AddMenu2("&File")
	fileMenu.AddActions([]*widgets.QAction{addAction, deleteAction})
	fileMenu.AddSeparator()
	fileMenu.AddActions([]*widgets.QAction{quitAction})

	helpMenu := m.AddMenu2("&Help")
	helpMenu.AddActions([]*widgets.QAction{aboutAction, aboutQtAction})

	addAction.ConnectTriggered(func(bool) { ViewControllerInstance.AddAlbumShowRequest() })
	deleteAction.ConnectTriggered(func(bool) { ViewControllerInstance.DeleteAlbumRequest() })
	quitAction.ConnectTriggered(func(bool) { ViewControllerInstance.Close() })
	aboutAction.ConnectTriggered(func(bool) { m.about() })
	aboutQtAction.ConnectTriggered(func(bool) { ViewControllerInstance.AboutQt() })
}

func (m *menuBarController) about() {
	widgets.QMessageBox_About(ViewControllerInstance, "About Music Archive",
		`<p>The <b>Music Archive</b> example shows how to present
data from different data sources in the same application.
The album titles, and the corresponding artists and release dates,
are kept in a database, while each album's tracks are stored
in an XML file. </p><p>The example also shows how to add as
well as remove data from both the database and the
associated XML file using the API provided by the Qt SQL and
Qt XML modules, respectively.</p>`)
}
