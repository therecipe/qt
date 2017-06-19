// +build !qml

package dialog

import (
	"github.com/therecipe/qt/core"
	"github.com/therecipe/qt/widgets"

	"github.com/therecipe/qt/internal/examples/sql/masterdetail_qml/controller"
)

type addDialogController struct {
	widgets.QDialog

	_ func() `constructor:"init"`

	//<-controller
	_ func() `signal:"addAlbumShowRequest"`

	//->controller
	_ func(artist string, title string, year int, tracks string) `signal:"addAlbum"`

	//

	artistEditor *widgets.QLineEdit
	titleEditor  *widgets.QLineEdit
	yearEditor   *widgets.QSpinBox
	tracksEditor *widgets.QLineEdit
}

func (a *addDialogController) init() {

	inputWidgetBox := a.createInputWidgets()
	buttonBox := a.createButtons()

	layout := widgets.NewQVBoxLayout()
	layout.AddWidget(inputWidgetBox, 0, 0)
	layout.AddWidget(buttonBox, 0, 0)
	a.SetLayout(layout)

	a.SetWindowTitle("Add Album")

	//<-controller
	controller.Instance.ConnectAddAlbumShowRequest(func() { a.Exec() })

	//->controller
	a.ConnectAddAlbum(controller.Instance.AddAlbum)
}

func (a *addDialogController) submit() {

	artist := a.artistEditor.Text()
	title := a.titleEditor.Text()

	if artist == "" || title == "" {
		widgets.QMessageBox_Information(nil, "Add Album",
			`Please provide both the name of the artist
and the title of the album.`, 0, 0)
	} else {
		a.AddAlbum(artist, title, a.yearEditor.Value(), a.tracksEditor.Text())
		a.revert()
		a.Close()
	}
}

func (a *addDialogController) revert() {
	a.artistEditor.Clear()
	a.titleEditor.Clear()
	a.yearEditor.SetValue(core.QDate_CurrentDate().Year())
	a.tracksEditor.Clear()
}

func (a *addDialogController) createInputWidgets() *widgets.QGroupBox {
	box := widgets.NewQGroupBox2("Add Album", nil)

	artistLabel := widgets.NewQLabel2("Artist:", nil, 0)
	titleLabel := widgets.NewQLabel2("Title:", nil, 0)
	yearLabel := widgets.NewQLabel2("Year:", nil, 0)
	tracksLabel := widgets.NewQLabel2("Tracks (seperated by comma):", nil, 0)

	a.artistEditor = widgets.NewQLineEdit(nil)
	a.titleEditor = widgets.NewQLineEdit(nil)

	a.yearEditor = widgets.NewQSpinBox(nil)
	a.yearEditor.SetMinimum(1900)
	a.yearEditor.SetMaximum(core.QDate_CurrentDate().Year())
	a.yearEditor.SetValue(a.yearEditor.Maximum())
	a.yearEditor.SetReadOnly(false)

	a.tracksEditor = widgets.NewQLineEdit(nil)

	layout := widgets.NewQGridLayout2()
	layout.AddWidget(artistLabel, 0, 0, 0)
	layout.AddWidget(a.artistEditor, 0, 1, 0)
	layout.AddWidget(titleLabel, 1, 0, 0)
	layout.AddWidget(a.titleEditor, 1, 1, 0)
	layout.AddWidget(yearLabel, 2, 0, 0)
	layout.AddWidget(a.yearEditor, 2, 1, 0)
	layout.AddWidget3(tracksLabel, 3, 0, 1, 2, 0)
	layout.AddWidget3(a.tracksEditor, 4, 0, 1, 2, 0)
	box.SetLayout(layout)

	return box
}

func (a *addDialogController) createButtons() *widgets.QDialogButtonBox {
	closeButton := widgets.NewQPushButton2("&Close", nil)
	revertButton := widgets.NewQPushButton2("&Revert", nil)
	submitButton := widgets.NewQPushButton2("&Submit", nil)

	closeButton.SetDefault(true)

	closeButton.ConnectClicked(func(bool) { a.Close() })
	revertButton.ConnectClicked(func(bool) { a.revert() })
	submitButton.ConnectClicked(func(bool) { a.submit() })

	buttonBox := widgets.NewQDialogButtonBox(nil)
	buttonBox.AddButton(submitButton, widgets.QDialogButtonBox__ResetRole)
	buttonBox.AddButton(revertButton, widgets.QDialogButtonBox__ResetRole)
	buttonBox.AddButton(closeButton, widgets.QDialogButtonBox__RejectRole)

	return buttonBox
}
