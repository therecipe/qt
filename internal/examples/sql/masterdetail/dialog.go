package main

import (
	"strconv"
	"strings"

	"github.com/therecipe/qt/core"
	"github.com/therecipe/qt/sql"
	"github.com/therecipe/qt/widgets"
	"github.com/therecipe/qt/xml"
)

var uniqueArtistId int
var uniqueAlbumId int

type Dialog struct {
	widgets.QDialog

	_ func() `constructor:"init"`

	_ func() `slot:"revert"`
	_ func() `slot:"submit"`

	model        *sql.QSqlRelationalTableModel
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

func (d *Dialog) initWith(albums *sql.QSqlRelationalTableModel, details *xml.QDomDocument, output *core.QFile, parent *widgets.QWidget) {

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

		d.increaseAlbumCount(d.indexOfArtist(artist))
		d.Accept()
	}
}

func (d *Dialog) findArtistId(artist string) int {
	artistModel := d.model.RelationModel(2)

	for row := 0; row < artistModel.RowCount(core.NewQModelIndex()); row++ {
		record := artistModel.Record(row)
		if record.Value2("artist").ToString() == artist {
			return record.Value2("id").ToInt(true)
		}
	}
	return d.addNewArtist(artist)
}

func (d *Dialog) addNewArtist(name string) int {

	artistModel := d.model.RelationModel(2)
	record := sql.NewQSqlRecord()

	id := d.generateArtistId()

	f1 := sql.NewQSqlField("id", core.QVariant__Int)
	f2 := sql.NewQSqlField("artist", core.QVariant__String)
	f3 := sql.NewQSqlField("albumcount", core.QVariant__Int)

	f1.SetValue(core.NewQVariant7(id))
	f2.SetValue(core.NewQVariant14(name))
	f3.SetValue(core.NewQVariant7(0))
	record.Append(f1)
	record.Append(f2)
	record.Append(f3)

	artistModel.InsertRecord(-1, record)
	return id
}

func (d *Dialog) addNewAlbum(title string, artistId int) int {

	id := d.generateAlbumId()
	record := sql.NewQSqlRecord()

	f1 := sql.NewQSqlField("albumid", core.QVariant__Int)
	f2 := sql.NewQSqlField("title", core.QVariant__String)
	f3 := sql.NewQSqlField("artistid", core.QVariant__Int)
	f4 := sql.NewQSqlField("year", core.QVariant__Int)

	f1.SetValue(core.NewQVariant7(id))
	f2.SetValue(core.NewQVariant14(title))
	f3.SetValue(core.NewQVariant7(artistId))
	f4.SetValue(core.NewQVariant7(d.yearEditor.Value()))
	record.Append(f1)
	record.Append(f2)
	record.Append(f3)
	record.Append(f4)

	d.model.InsertRecord(-1, record)
	return id
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

func (d *Dialog) increaseAlbumCount(artistIndex *core.QModelIndex) {
	artistModel := d.model.RelationModel(2)

	albumCountIndex := artistIndex.Sibling(artistIndex.Row(), 2)

	albumCount := albumCountIndex.Data(0).ToInt(true)
	artistModel.SetData(albumCountIndex, core.NewQVariant7(albumCount+1), 0)
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

func (d *Dialog) indexOfArtist(artist string) *core.QModelIndex {
	artistModel := d.model.RelationModel(2)

	for i := 0; i < artistModel.RowCount(core.NewQModelIndex()); i++ {
		record := artistModel.Record(i)
		if record.Value2("artist").ToString() == artist {
			return artistModel.Index(i, 1, core.NewQModelIndex())
		}
	}

	return core.NewQModelIndex()
}

func (d *Dialog) generateArtistId() int {
	uniqueArtistId++
	return uniqueArtistId
}

func (d *Dialog) generateAlbumId() int {
	uniqueAlbumId++
	return uniqueAlbumId
}
