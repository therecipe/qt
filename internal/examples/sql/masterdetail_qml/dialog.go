package main

import (
	"strconv"
	"strings"

	"github.com/therecipe/qt/core"
	"github.com/therecipe/qt/widgets"
	"github.com/therecipe/qt/xml"
)

type Dialog struct {
	widgets.QDialog

	_ func() `constructor:"init"`

	_ func() `slot:"revert"`
	_ func() `slot:"submit"`

	model        *core.QAbstractItemModel
	albumDetails *xml.QDomDocument
	outputFile   *core.QFile

	artistEditor *widgets.QLineEdit
	titleEditor  *widgets.QLineEdit
	yearEditor   *widgets.QSpinBox
	tracksEditor *widgets.QLineEdit
}

func (d *Dialog) init() {
	d.ConnectRevert(d.revert)
	d.ConnectSubmit(d.submit)
}

func (d *Dialog) initWith(albums *core.QAbstractItemModel, details *xml.QDomDocument, output *core.QFile, parent *widgets.QWidget) {

	d.model = albums
	d.albumDetails = details
	d.outputFile = output

	inputWidgetBox := d.createInputWidgets()
	buttonBox := d.createButtons()

	layout := widgets.NewQVBoxLayout()
	layout.AddWidget(inputWidgetBox, 0, 0)
	layout.AddWidget(buttonBox, 0, 0)
	d.SetLayout(layout)

	d.SetWindowTitle("Add Album")
}

func (d *Dialog) submit() {

	artist := d.artistEditor.Text()
	title := d.titleEditor.Text()

	if artist == "" || title == "" {
		widgets.QMessageBox_Information(d, "Add Album",
			`Please provide both the name of the artist
and the title of the album.`, 0, 0)
	} else {
		artistId := d.findArtistId(artist)
		albumId := d.addNewAlbum(title, artistId)

		tracks := strings.Split(d.tracksEditor.Text(), ",")
		d.addTracks(albumId, tracks)

		d.Accept()
	}
}

func (d *Dialog) findArtistId(artist string) int {
	if a := getArtistForName(artist); a != nil {
		return a.ID
	}
	return d.addNewArtist(artist)
}

func (d *Dialog) addNewArtist(name string) int {
	artistId := getNextArtistID()

	listModel.BeginInsertRows(core.NewQModelIndex(), listModel.RowCount(core.NewQModelIndex())+1, listModel.RowCount(core.NewQModelIndex())+1)
	createNewArtist(artistId, name)
	listModel.EndInsertRows()

	return artistId
}

func (d *Dialog) addNewAlbum(title string, artistId int) int {

	albumId := getNextAlbumID()

	//TODO
	//inserting or removing from this model (SortFilterModel) will NOT affect the sourceModel
	//(because calls are not going through?) and it will therefore lead to glitches
	//resetModel however affectes both models and works ... but it can be slow

	//d.model.BeginInsertRows(core.NewQModelIndex(), d.model.RowCount(core.NewQModelIndex())+1, d.model.RowCount(core.NewQModelIndex())+1)
	d.model.BeginResetModel()
	createNewAlbum(artistId, albumId, title, d.yearEditor.Value())
	d.model.EndResetModel()
	//d.model.EndInsertRows()

	return albumId
}

func (d *Dialog) addTracks(albumId int, tracks []string) {

	albumNode := d.albumDetails.CreateElement("album")
	albumNode.SetAttribute4("id", albumId)

	for i := 0; i < len(tracks); i++ {
		trackNumber := strconv.Itoa(i)
		if i < 10 {
			trackNumber = "0" + trackNumber
		}

		textNode := d.albumDetails.CreateTextNode(tracks[i])

		trackNode := d.albumDetails.CreateElement("track")
		trackNode.SetAttribute("number", trackNumber)
		trackNode.AppendChild(textNode)

		albumNode.AppendChild(trackNode)
	}

	archive := d.albumDetails.ElementsByTagName("archive")
	archive.Item(0).AppendChild(albumNode)

	/*
		The following code is commented out since the example uses an in
		memory database, i.e., altering the XML file will bring the data
		out of sync.

		if !d.outputFile.Open(core.QIODevice__WriteOnly) {
			return
		} else {
			stream := core.NewQTextStream2(d.outputFile)
			archive.Item(0).Save(stream, 4, 0)
			d.outputFile.Close()
		}
	*/
}

func (d *Dialog) revert() {
	d.artistEditor.Clear()
	d.titleEditor.Clear()
	d.yearEditor.SetValue(core.QDate_CurrentDate().Year())
	d.tracksEditor.Clear()
}

func (d *Dialog) createInputWidgets() *widgets.QGroupBox {
	box := widgets.NewQGroupBox2("Add Album", nil)

	artistLabel := widgets.NewQLabel2("Artist:", nil, 0)
	titleLabel := widgets.NewQLabel2("Title:", nil, 0)
	yearLabel := widgets.NewQLabel2("Year:", nil, 0)
	tracksLabel := widgets.NewQLabel2("Tracks (seperated by comma):", nil, 0)

	d.artistEditor = widgets.NewQLineEdit(nil)
	d.titleEditor = widgets.NewQLineEdit(nil)

	d.yearEditor = widgets.NewQSpinBox(nil)
	d.yearEditor.SetMinimum(1900)
	d.yearEditor.SetMaximum(core.QDate_CurrentDate().Year())
	d.yearEditor.SetValue(d.yearEditor.Maximum())
	d.yearEditor.SetReadOnly(false)

	d.tracksEditor = widgets.NewQLineEdit(nil)

	layout := widgets.NewQGridLayout2()
	layout.AddWidget(artistLabel, 0, 0, 0)
	layout.AddWidget(d.artistEditor, 0, 1, 0)
	layout.AddWidget(titleLabel, 1, 0, 0)
	layout.AddWidget(d.titleEditor, 1, 1, 0)
	layout.AddWidget(yearLabel, 2, 0, 0)
	layout.AddWidget(d.yearEditor, 2, 1, 0)
	layout.AddWidget3(tracksLabel, 3, 0, 1, 2, 0)
	layout.AddWidget3(d.tracksEditor, 4, 0, 1, 2, 0)
	box.SetLayout(layout)

	return box
}

func (d *Dialog) createButtons() *widgets.QDialogButtonBox {
	closeButton := widgets.NewQPushButton2("&Close", nil)
	revertButton := widgets.NewQPushButton2("&Revert", nil)
	submitButton := widgets.NewQPushButton2("&Submit", nil)

	closeButton.SetDefault(true)

	closeButton.ConnectClicked(func(bool) { d.Close() })
	revertButton.ConnectClicked(func(bool) { d.revert() })
	submitButton.ConnectClicked(func(bool) { d.submit() })

	buttonBox := widgets.NewQDialogButtonBox(nil)
	buttonBox.AddButton(submitButton, widgets.QDialogButtonBox__ResetRole)
	buttonBox.AddButton(revertButton, widgets.QDialogButtonBox__ResetRole)
	buttonBox.AddButton(closeButton, widgets.QDialogButtonBox__RejectRole)

	return buttonBox
}
