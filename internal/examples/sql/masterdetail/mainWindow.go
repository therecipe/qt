package main

import (
	"fmt"
	"strconv"

	"github.com/therecipe/qt/core"
	"github.com/therecipe/qt/gui"
	"github.com/therecipe/qt/sql"
	"github.com/therecipe/qt/widgets"
	"github.com/therecipe/qt/xml"
)

type MainWindow struct {
	widgets.QMainWindow

	_ func() `constructor:"init"`

	_ func()                            `slot:"about"`
	_ func()                            `slot:"addAlbum"`
	_ func(row int)                     `slot:"changeArtist"`
	_ func()                            `slot:"deleteAlbum"`
	_ func(index *core.QModelIndex)     `slot:"showAlbumDetails"`
	_ func(index *core.QModelIndex)     `slot:"showArtistProfile"`
	_ func(*core.QModelIndex, int, int) `slot:"updateHeader"`

	albumView  *widgets.QTableView
	artistView *widgets.QComboBox
	trackList  *widgets.QListWidget

	iconLabel    *widgets.QLabel
	imageLabel   *widgets.QLabel
	profileLabel *widgets.QLabel
	titleLabel   *widgets.QLabel

	albumData *xml.QDomDocument
	file      *core.QFile
	model     *sql.QSqlRelationalTableModel
}

func (w *MainWindow) init() {
	w.albumData = xml.NewQDomDocument()

	w.ConnectAbout(w.about)
	w.ConnectAddAlbum(w.addAlbum)
	w.ConnectChangeArtist(w.changeArtist)
	w.ConnectDeleteAlbum(w.deleteAlbum)
	w.ConnectShowAlbumDetails(w.showAlbumDetails)
	w.ConnectShowArtistProfile(w.showArtistProfile)
	w.ConnectUpdateHeader(w.updateHeader)
}

func (w *MainWindow) initWith(artistTable string, albumTable string, albumDetails *core.QFile, parent *widgets.QWidget) {
	w.file = albumDetails
	w.readAlbumData()

	w.model = sql.NewQSqlRelationalTableModel(w, db)
	w.model.SetTable(albumTable)
	w.model.SetRelation(2, sql.NewQSqlRelation2(artistTable, "id", "artist"))
	w.model.Select()

	artists := w.createArtistGroupBox()
	albums := w.createAlbumGroupBox()
	details := w.createDetailsGroupBox()

	w.artistView.SetCurrentIndex(0)
	uniqueAlbumId = w.model.RowCount(core.NewQModelIndex())
	uniqueArtistId = w.artistView.Count()

	w.model.ConnectRowsInserted(w.updateHeader)
	w.model.ConnectRowsRemoved(w.updateHeader)

	layout := widgets.NewQGridLayout2()
	layout.AddWidget(artists, 0, 0, 0)
	layout.AddWidget(albums, 1, 0, 0)
	layout.AddWidget3(details, 0, 1, 2, 1, 0)
	layout.SetColumnStretch(1, 1)
	layout.SetColumnMinimumWidth(0, 500)

	widget := widgets.NewQWidget(nil, 0)
	widget.SetLayout(layout)
	w.SetCentralWidget(widget)
	w.createMenuBar()

	w.showImageLabel()
	w.Resize2(850, 400)
	w.SetWindowTitle("Music Archive")
}

func (w *MainWindow) changeArtist(row int) {
	if row > 0 {
		index := w.model.RelationModel(2).Index(row, 1, core.NewQModelIndex())
		w.model.SetFilter("artist = '" + index.Data(0).ToString() + "'")
		w.showArtistProfile(index)
	} else if row == 0 {
		w.model.SetFilter("")
		w.showImageLabel()
	} else {
		return
	}
}

func (w *MainWindow) showArtistProfile(index *core.QModelIndex) {
	record := w.model.RelationModel(2).Record(index.Row())

	name := record.Value2("artist").ToString()
	count := record.Value2("albumcount").ToString()
	w.profileLabel.SetText(fmt.Sprintf("Artist : %v \nNumber of Albums: %v", name, count))

	w.profileLabel.Show()
	w.iconLabel.Show()

	w.titleLabel.Hide()
	w.trackList.Hide()
	w.imageLabel.Hide()
}

func (w *MainWindow) showAlbumDetails(index *core.QModelIndex) {
	record := w.model.Record(index.Row())

	artist := record.Value2("artist").ToString()
	title := record.Value2("title").ToString()
	year := record.Value2("year").ToString()
	albumId := record.Value2("albumid").ToString()

	w.showArtistProfile(w.indexOfArtist(artist))
	w.titleLabel.SetText(fmt.Sprintf("Title: %v (%v)", title, year))
	w.titleLabel.Show()

	albums := w.albumData.ElementsByTagName("album")
	for i := 0; i < albums.Count(); i++ {
		album := albums.Item(i)
		if album.ToElement().Attribute("id", "") == albumId {
			w.getTrackList(album.ToElement())
			break
		}
	}
	if w.trackList.Count() != 0 {
		w.trackList.Show()
	}
}

func (w *MainWindow) getTrackList(album *xml.QDomElement) {
	w.trackList.Clear()

	tracks := album.ChildNodes()
	var track *xml.QDomNode
	var trackNumber string

	for j := 0; j < tracks.Count(); j++ {
		track = tracks.Item(j)
		trackNumber = track.ToElement().Attribute("number", "")

		item := widgets.NewQListWidgetItem(w.trackList, 0)
		item.SetText(trackNumber + ": " + track.ToElement().Text())
	}
}

func (w *MainWindow) addAlbum() {
	dialog := NewDialog(nil, 0)
	dialog.initWith(w.model, w.albumData, w.file, w.QWidget_PTR())
	accepted := dialog.Exec()

	if accepted == 1 {
		lastRow := w.model.RowCount(core.NewQModelIndex()) - 1
		w.albumView.SelectRow(lastRow)
		w.albumView.ScrollToBottom()
		w.showAlbumDetails(w.model.Index(lastRow, 0, core.NewQModelIndex()))
	}
}

func (w *MainWindow) deleteAlbum() {
	selection := w.albumView.SelectionModel().SelectedRows(0)

	if len(selection) != 0 {
		idIndex := selection[0]
		id := idIndex.Data(0).ToInt(true)
		title := idIndex.Sibling(idIndex.Row(), 1).Data(0).ToString()
		artist := idIndex.Sibling(idIndex.Row(), 2).Data(0).ToString()

		var button widgets.QMessageBox__StandardButton
		button = widgets.QMessageBox_Question(w, "Delete Album",
			fmt.Sprintf("Are you sure you want to delete '%v' by '%v'?", title, artist),
			widgets.QMessageBox__Yes|widgets.QMessageBox__No, 0)

		if button == widgets.QMessageBox__Yes {
			w.removeAlbumFromFile(id)
			w.removeAlbumFromDatabase(idIndex)
			w.decreaseAlbumCount(w.indexOfArtist(artist))

			w.showImageLabel()
		} else {
			widgets.QMessageBox_Information(w, "Delete Album",
				"Select the album you want to delete.", 0, 0)
		}
	}
}

func (w *MainWindow) removeAlbumFromFile(id int) {
	albums := w.albumData.ElementsByTagName("album")

	for i := 0; i < albums.Count(); i++ {
		node := albums.Item(i)
		if node.ToElement().Attribute("id", "") == strconv.Itoa(id) {
			w.albumData.ElementsByTagName("archive").Item(0).RemoveChild(node)
			break
		}
	}

	/*
		The following code is commented out since the example uses an in
		memory database, i.e., altering the XML file will bring the data
		out of sync.

		if !w.file.Open(core.QIODevice__WriteOnly) {
			return
		} else {
			stream := core.NewQTextStream2(w.file)
			w.albumData.ElementsByTagName("archive").Item(0).Save(stream, 4, 0)
			w.file.Close()
		}
	*/
}

func (w *MainWindow) removeAlbumFromDatabase(index *core.QModelIndex) {
	w.model.RemoveRow(index.Row(), core.NewQModelIndex())
}

func (w *MainWindow) decreaseAlbumCount(artistIndex *core.QModelIndex) {
	row := artistIndex.Row()
	albumCountIndex := artistIndex.Sibling(row, 2)
	albumCount := albumCountIndex.Data(0).ToInt(true)

	artists := w.model.RelationModel(2)

	if albumCount == 1 {
		artists.RemoveRow(row, core.NewQModelIndex())
		w.showImageLabel()
	} else {
		artists.SetData(albumCountIndex, core.NewQVariant7(albumCount-1), 0)
	}
}

func (w *MainWindow) readAlbumData() {

	if !w.file.Open(core.QIODevice__ReadOnly) {
		return
	}

	if !w.albumData.SetContent3(w.file, false, "", 0, 0) {
		w.file.Close()
		return
	}
	w.file.Close()
}

func (w *MainWindow) createArtistGroupBox() *widgets.QGroupBox {
	w.artistView = widgets.NewQComboBox(nil)
	w.artistView.SetModel(w.model.RelationModel(2))
	w.artistView.SetModelColumn(1)

	w.artistView.ConnectCurrentIndexChanged(w.changeArtist)

	box := widgets.NewQGroupBox2("Artist", nil)

	layout := widgets.NewQGridLayout2()
	layout.AddWidget(w.artistView, 0, 0, 0)
	box.SetLayout(layout)

	return box
}

func (w *MainWindow) createAlbumGroupBox() *widgets.QGroupBox {
	box := widgets.NewQGroupBox2("Album", nil)

	w.albumView = widgets.NewQTableView(nil)
	w.albumView.SetEditTriggers(widgets.QAbstractItemView__NoEditTriggers)
	w.albumView.SetSortingEnabled(true)
	w.albumView.SetSelectionBehavior(widgets.QAbstractItemView__SelectRows)
	w.albumView.SetSelectionMode(widgets.QAbstractItemView__SingleSelection)
	w.albumView.SetShowGrid(false)
	w.albumView.VerticalHeader().Hide()
	w.albumView.SetAlternatingRowColors(true)
	w.albumView.SetModel(w.model)
	w.adjustHeader()

	locale := w.albumView.Locale()
	locale.SetNumberOptions(core.QLocale__OmitGroupSeparator)
	w.albumView.SetLocale(locale)

	w.albumView.ConnectClicked(w.showAlbumDetails)
	w.albumView.ConnectActivated(w.showAlbumDetails)

	layout := widgets.NewQVBoxLayout()
	layout.AddWidget(w.albumView, 0, 0)
	box.SetLayout(layout)

	return box
}

func (w *MainWindow) createDetailsGroupBox() *widgets.QGroupBox {
	box := widgets.NewQGroupBox2("Details", nil)

	w.profileLabel = widgets.NewQLabel(nil, 0)
	w.profileLabel.SetWordWrap(true)
	w.profileLabel.SetAlignment(core.Qt__AlignBottom)

	w.titleLabel = widgets.NewQLabel(nil, 0)
	w.titleLabel.SetWordWrap(true)
	w.titleLabel.SetAlignment(core.Qt__AlignBottom)

	w.iconLabel = widgets.NewQLabel(nil, 0)
	w.iconLabel.SetAlignment(core.Qt__AlignBottom | core.Qt__AlignRight)
	w.iconLabel.SetPixmap(gui.NewQPixmap5(":/images/icon.png", "", 0))

	w.imageLabel = widgets.NewQLabel(nil, 0)
	w.imageLabel.SetWordWrap(true)
	w.imageLabel.SetAlignment(core.Qt__AlignCenter)
	w.imageLabel.SetPixmap(gui.NewQPixmap5(":/images/image.png", "", 0))

	w.trackList = widgets.NewQListWidget(nil)

	layout := widgets.NewQGridLayout2()
	layout.AddWidget3(w.imageLabel, 0, 0, 3, 2, 0)
	layout.AddWidget(w.profileLabel, 0, 0, 0)
	layout.AddWidget(w.iconLabel, 0, 1, 0)
	layout.AddWidget3(w.titleLabel, 1, 0, 1, 2, 0)
	layout.AddWidget3(w.trackList, 2, 0, 1, 2, 0)
	layout.SetRowStretch(2, 1)
	box.SetLayout(layout)

	return box
}

func (w *MainWindow) createMenuBar() {
	addAction := widgets.NewQAction2("&Add album...", w)
	deleteAction := widgets.NewQAction2("&Delete album...", w)
	quitAction := widgets.NewQAction2("&Quit", w)
	aboutAction := widgets.NewQAction2("&About", w)
	aboutQtAction := widgets.NewQAction2("About &Qt", w)

	addAction.SetShortcut(gui.QKeySequence_FromString("Ctrl+A", 0))
	deleteAction.SetShortcut(gui.QKeySequence_FromString("Ctrl+D", 0))
	quitAction.SetShortcuts2(gui.QKeySequence__Quit)

	fileMenu := w.MenuBar().AddMenu2("&File")
	fileMenu.AddActions([]*widgets.QAction{addAction, deleteAction})
	fileMenu.AddSeparator()
	fileMenu.AddActions([]*widgets.QAction{quitAction})

	helpMenu := w.MenuBar().AddMenu2("&Help")
	helpMenu.AddActions([]*widgets.QAction{aboutAction, aboutQtAction})

	addAction.ConnectTriggered(func(bool) { w.addAlbum() })
	deleteAction.ConnectTriggered(func(bool) { w.deleteAlbum() })
	quitAction.ConnectTriggered(func(bool) { w.Close() })
	aboutAction.ConnectTriggered(func(bool) { w.about() })
	aboutQtAction.ConnectTriggered(func(bool) { qApp.AboutQt() })
}

func (w *MainWindow) showImageLabel() {
	w.profileLabel.Hide()
	w.titleLabel.Hide()
	w.iconLabel.Hide()
	w.trackList.Hide()

	w.imageLabel.Show()
}

func (w *MainWindow) indexOfArtist(artist string) *core.QModelIndex {
	artistModel := w.model.RelationModel(2)

	for i := 0; i < artistModel.RowCount(core.NewQModelIndex()); i++ {
		record := artistModel.Record(i)
		if record.Value2("artist").ToString() == artist {
			return artistModel.Index(i, 1, core.NewQModelIndex())
		}
	}
	return core.NewQModelIndex()
}

func (w *MainWindow) updateHeader(parent *core.QModelIndex, first int, last int) {
	w.adjustHeader()
}

func (w *MainWindow) adjustHeader() {
	w.albumView.HideColumn(0)
	w.albumView.HorizontalHeader().SetSectionResizeMode2(1, widgets.QHeaderView__Stretch)
	w.albumView.ResizeColumnToContents(2)
	w.albumView.ResizeColumnToContents(3)
}

func (w *MainWindow) about() {
	widgets.QMessageBox_About(w, "About Music Archive",
		`<p>The <b>Music Archive</b> example shows how to present
data from different data sources in the same application.
The album titles, and the corresponding artists and release dates,
are kept in a database, while each album's tracks are stored
in an XML file. </p><p>The example also shows how to add as
well as remove data from both the database and the
associated XML file using the API provided by the Qt SQL and
Qt XML modules, respectively.</p>`)
}
