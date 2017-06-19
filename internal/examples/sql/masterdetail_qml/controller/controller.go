package controller

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/therecipe/qt/core"
	"github.com/therecipe/qt/widgets"
	"github.com/therecipe/qt/xml"

	"github.com/therecipe/qt/internal/examples/sql/masterdetail_qml/model"
)

var Instance *Controller

type Controller struct {
	core.QObject

	albumData *xml.QDomDocument
	qApp      *widgets.QApplication

	_ func() `constructor:"init"`

	_ *core.QAbstractItemModel `property:"artistModel"`
	_ *core.QAbstractItemModel `property:"albumModel"`

	//<-view
	_ func() `signal:"aboutQt"`

	//<-artist
	_ func(row int) `signal:"changeArtist"`

	//<-album
	_ func(index *core.QModelIndex)              `signal:"deleteAlbum"`
	_ func(index *core.QModelIndex)              `signal:"showAlbumDetails"`
	_ func(column int, order core.Qt__SortOrder) `signal:"sortTableView"`

	//->detail
	_ func()                                `signal:"showImageLabel"`
	_ func(profileLabelText string)         `signal:"showArtistProfile"`
	_ func(title string, elements []string) `signal:"showTitleAndAlbumDetails"`

	//album<->dialog
	_ func()                     `signal:"deleteAlbumCommand"`
	_ func()                     `signal:"deleteAlbumRequest"`
	_ func(title, artist string) `signal:"deleteAlbumShowRequest"`

	//dialog<->view
	_ func()                                                     `signal:"addAlbumShowRequest"`
	_ func(artist string, title string, year int, tracks string) `signal:"addAlbum"`
}

func (c *Controller) init() {
	Instance = c
	c.albumData = xml.NewQDomDocument()

	c.SetArtistModel(model.NewListModel())
	c.SetAlbumModel(model.NewSortFilterModel())

	//<-view
	c.ConnectAboutQt(func() { c.qApp.AboutQt() })

	//<-artist
	c.ConnectChangeArtist(c.changeArtist)

	//<-album
	c.ConnectDeleteAlbum(c.deleteAlbum)
	c.ConnectShowAlbumDetails(c.showAlbumDetails)
	c.ConnectSortTableView(model.SortFilterModel.Sort)
	c.ConnectAddAlbum(c.addAlbum)

	go func() {
		count := 0
		for range time.NewTicker(5 * time.Second).C {
			c.AddAlbum("goRoutineArtist", fmt.Sprint(count), 2017, "")
			count++
		}
	}()
}

func (c *Controller) InitWith(file *core.QFile, qApp *widgets.QApplication) {
	c.readalbumData(file)
	c.qApp = qApp
}

func (c *Controller) changeArtist(row int) {
	artist := model.ListModel.Index(row, 0, core.NewQModelIndex()).Data(int(core.Qt__DisplayRole)).ToString()

	if row > 0 {
		model.SortFilterModel.SetFilterFixedString(artist)
		model.SortFilterModel.SetFilterKeyColumn(2)
		c.showArtistProfile(artist)
	} else if row == 0 {
		model.SortFilterModel.SetFilterFixedString("")
		c.ShowImageLabel()
	}
}

func (c *Controller) showArtistProfile(artist string) {
	c.ShowArtistProfile(fmt.Sprintf("Artist : %v \nNumber of Albums: %v", artist, model.GetAlbumCountForArtist(artist)))
}

func (c *Controller) showAlbumDetails(index *core.QModelIndex) {

	artist := index.Data(int(core.Qt__UserRole) + 3).ToString()
	title := index.Data(int(core.Qt__UserRole) + 2).ToString()
	year := index.Data(int(core.Qt__UserRole) + 4).ToString()
	albumId := index.Data(int(core.Qt__UserRole) + 1).ToString()

	c.showArtistProfile(artist)

	var trackList []string

	albums := c.albumData.ElementsByTagName("album")
	for i := 0; i < albums.Count(); i++ {
		album := albums.Item(i)
		if album.ToElement().Attribute("id", "") == albumId {
			trackList = c.getTrackList(album.ToElement())
			break
		}
	}

	c.ShowTitleAndAlbumDetails(fmt.Sprintf("Title: %v (%v)", title, year), trackList)
}

func (c *Controller) readalbumData(file *core.QFile) {

	if !file.Open(core.QIODevice__ReadOnly) {
		return
	}

	if !c.albumData.SetContent3(file, false, "", 0, 0) {
		file.Close()
		return
	}
	file.Close()
}

func (c *Controller) getTrackList(album *xml.QDomElement) []string {
	var out []string

	tracks := album.ChildNodes()
	var track *xml.QDomNode
	var trackNumber string

	for j := 0; j < tracks.Count(); j++ {
		track = tracks.Item(j)
		trackNumber = track.ToElement().Attribute("number", "")
		out = append(out, trackNumber+": "+track.ToElement().Text())
	}

	return out
}

func (c *Controller) deleteAlbum(index *core.QModelIndex) {
	c.removeAlbumFromFile(index.Data(int(core.Qt__UserRole) + 1).ToInt(true))
	c.removeAlbumFromDatabase(index)
}

func (c *Controller) removeAlbumFromFile(id int) {
	albums := c.albumData.ElementsByTagName("album")

	for i := 0; i < albums.Count(); i++ {
		node := albums.Item(i)
		if node.ToElement().Attribute("id", "") == strconv.Itoa(id) {
			c.albumData.ElementsByTagName("archive").Item(0).RemoveChild(node)
			break
		}
	}

	/*
		The following code is commented out since the example uses an in
		memory database, i.e., altering the XML file will bring the data
		out of sync.

		if !c.file.Open(core.QIODevice__WriteOnly) {
			return
		} else {
			stream := core.NewQTextStream2(w.file)
			c.albumData.ElementsByTagName("archive").Item(0).Save(stream, 4, 0)
			c.file.Close()
		}
	*/
}

func (c *Controller) removeAlbumFromDatabase(index *core.QModelIndex) {

	//TODO
	//inserting or removing from this model (SortFilterModel) will NOT affect the sourceModel
	//(because calls are not going through?) and it will therefore lead to glitches
	//resetModel however affectes both models and works ... but it can be slow

	//model.SortFilterModel.BeginRemoveRows(core.NewQModelIndex(), index.Row(), index.Row())
	model.SortFilterModel.BeginResetModel()
	model.DeleteAlbum(index.Data(int(core.Qt__UserRole) + 1).ToInt(true))
	model.SortFilterModel.EndResetModel()
	//model.SortFilterModel.EndRemoveRows()

	c.ShowImageLabel()
}

func (c *Controller) addAlbum(artist string, title string, year int, tracks string) {

	var artistId int
	if a := model.GetArtistForName(artist); a != nil {
		artistId = a.ID
	} else {
		artistId = c.addNewArtist(artist)
	}

	albumId := c.addNewAlbum(title, artistId, year)

	c.addTracks(albumId, strings.Split(tracks, ","))
}

func (c *Controller) addNewArtist(name string) int {
	artistId := model.GetNextArtistID()

	model.ListModel.BeginInsertRows(core.NewQModelIndex(), model.ListModel.RowCount(core.NewQModelIndex())+1, model.ListModel.RowCount(core.NewQModelIndex())+1)
	model.CreateNewArtist(artistId, name)
	model.ListModel.EndInsertRows()

	return artistId
}

func (c *Controller) addNewAlbum(title string, artistId int, year int) int {

	albumId := model.GetNextAlbumID()

	//TODO
	//inserting or removing from this model (SortFilterModel) will NOT affect the sourceModel
	//(because calls are not going through?) and it will therefore lead to glitches
	//resetModel however affectes both models and works ... but it can be slow

	//model.SortFilterModel.BeginInsertRows(core.NewQModelIndex(), d.model.RowCount(core.NewQModelIndex())+1, d.model.RowCount(core.NewQModelIndex())+1)
	model.SortFilterModel.BeginResetModel()
	model.CreateNewAlbum(artistId, albumId, title, year)
	model.SortFilterModel.EndResetModel()
	//model.SortFilterModel.EndInsertRows()

	return albumId
}

func (c *Controller) addTracks(albumId int, tracks []string) {

	albumNode := c.albumData.CreateElement("album")
	albumNode.SetAttribute4("id", albumId)

	for i := 0; i < len(tracks); i++ {
		trackNumber := strconv.Itoa(i)
		if i < 10 {
			trackNumber = "0" + trackNumber
		}

		textNode := c.albumData.CreateTextNode(tracks[i])

		trackNode := c.albumData.CreateElement("track")
		trackNode.SetAttribute("number", trackNumber)
		trackNode.AppendChild(textNode)

		albumNode.AppendChild(trackNode)
	}

	archive := c.albumData.ElementsByTagName("archive")
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
