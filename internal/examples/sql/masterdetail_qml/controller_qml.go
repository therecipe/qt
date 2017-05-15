package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/therecipe/qt/core"
	"github.com/therecipe/qt/xml"
)

var albumData *xml.QDomDocument

type Controller struct {
	core.QObject

	_ func() `constructor:"init"`

	//qml -> go
	_ func(row int)                                              `signal:"changeArtist"`
	_ func(column int, order core.Qt__SortOrder)                 `signal:"sortTableView"`
	_ func(index *core.QModelIndex)                              `signal:"showAlbumDetails"`
	_ func()                                                     `signal:"aboutQt"`
	_ func(index *core.QModelIndex)                              `signal:"deleteAlbum"`
	_ func(artist string, title string, year int, tracks string) `signal:"addAlbum"`

	//go -> qml
	_ func()                                `signal:"showImageLabel"`
	_ func(profileLabelText string)         `signal:"showArtistProfile"`
	_ func(title string, elements []string) `signal:"showTitleAndAlbumDetails"`
	_ func()                                `signal:"resetComboBoxModel"`
}

func (c *Controller) init() {
	albumData = xml.NewQDomDocument()

	c.ConnectChangeArtist(c.changeArtist)
	c.ConnectSortTableView(sortFilterModel.Sort)
	c.ConnectShowAlbumDetails(c.showAlbumDetails)
	c.ConnectAboutQt(qApp.AboutQt)
	c.ConnectDeleteAlbum(c.deleteAlbum)
	c.ConnectAddAlbum(c.addAlbum)
}

func (c *Controller) initWith(file *core.QFile) {
	c.readAlbumData(file)
}

func (c *Controller) changeArtist(row int) {
	artist := listModel.Index(row, 0, core.NewQModelIndex()).Data(int(core.Qt__DisplayRole)).ToString()

	if row > 0 {
		sortFilterModel.SetFilterFixedString(artist)
		sortFilterModel.SetFilterKeyColumn(2)
		c.showArtistProfile(artist)
	} else if row == 0 {
		sortFilterModel.SetFilterFixedString("")
		c.ShowImageLabel()
	} else {
		return
	}
}

func (c *Controller) showArtistProfile(artist string) {
	c.ShowArtistProfile(fmt.Sprintf("Artist : %v \nNumber of Albums: %v", artist, getAlbumCountForArtist(artist)))
}

func (c *Controller) showAlbumDetails(index *core.QModelIndex) {

	artist := sortFilterModel.Index(index.Row(), 2, core.NewQModelIndex()).Data(int(core.Qt__DisplayRole)).ToString()
	title := sortFilterModel.Index(index.Row(), 1, core.NewQModelIndex()).Data(int(core.Qt__DisplayRole)).ToString()
	year := sortFilterModel.Index(index.Row(), 3, core.NewQModelIndex()).Data(int(core.Qt__DisplayRole)).ToString()
	albumId := sortFilterModel.Index(index.Row(), 0, core.NewQModelIndex()).Data(int(core.Qt__DisplayRole)).ToString()

	c.showArtistProfile(artist)

	var trackList []string

	albums := albumData.ElementsByTagName("album")
	for i := 0; i < albums.Count(); i++ {
		album := albums.Item(i)
		if album.ToElement().Attribute("id", "") == albumId {
			trackList = c.getTrackList(album.ToElement())
			break
		}
	}

	c.ShowTitleAndAlbumDetails(fmt.Sprintf("Title: %v (%v)", title, year), trackList)
}

func (c *Controller) readAlbumData(file *core.QFile) {

	if !file.Open(core.QIODevice__ReadOnly) {
		return
	}

	if !albumData.SetContent3(file, false, "", 0, 0) {
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
	c.removeAlbumFromFile(index.Data(int(core.Qt__DisplayRole)).ToInt(true))
	c.removeAlbumFromDatabase(index)
}

func (c *Controller) removeAlbumFromFile(id int) {
	albums := albumData.ElementsByTagName("album")

	for i := 0; i < albums.Count(); i++ {
		node := albums.Item(i)
		if node.ToElement().Attribute("id", "") == strconv.Itoa(id) {
			albumData.ElementsByTagName("archive").Item(0).RemoveChild(node)
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

func (c *Controller) removeAlbumFromDatabase(index *core.QModelIndex) {

	//TODO
	//inserting or removing from this model (SortFilterModel) will NOT affect the sourceModel
	//(because calls are not going through?) and it will therefore lead to glitches
	//resetModel however affectes both models and works ... but it can be slow

	//sortFilterModel.BeginRemoveRows(core.NewQModelIndex(), index.Row(), index.Row())
	sortFilterModel.BeginResetModel()
	deleteAlbum(index.Data(int(core.Qt__DisplayRole)).ToInt(true))
	sortFilterModel.EndResetModel()
	//sortFilterModel.EndRemoveRows()
}

func (c *Controller) addAlbum(artist string, title string, year int, tracks string) {

	artistId := c.findArtistId(artist)
	albumId := c.addNewAlbum(title, artistId, year)

	c.addTracks(albumId, strings.Split(tracks, ","))
}

func (c *Controller) findArtistId(artist string) int {
	if a := getArtistForName(artist); a != nil {
		return a.ID
	}
	return c.addNewArtist(artist)
}

func (c *Controller) addNewArtist(name string) int {
	artistId := getNextArtistID()

	listModel.BeginInsertRows(core.NewQModelIndex(), listModel.RowCount(core.NewQModelIndex())+1, listModel.RowCount(core.NewQModelIndex())+1)
	createNewArtist(artistId, name)
	listModel.EndInsertRows()

	c.ResetComboBoxModel()

	return artistId
}

func (c *Controller) addNewAlbum(title string, artistId int, year int) int {

	albumId := getNextAlbumID()

	//TODO
	//inserting or removing from this model (SortFilterModel) will NOT affect the sourceModel
	//(because calls are not going through?) and it will therefore lead to glitches
	//resetModel however affectes both models and works ... but it can be slow

	//sortFilterModel.BeginInsertRows(core.NewQModelIndex(), d.model.RowCount(core.NewQModelIndex())+1, d.model.RowCount(core.NewQModelIndex())+1)
	sortFilterModel.BeginResetModel()
	createNewAlbum(artistId, albumId, title, year)
	sortFilterModel.EndResetModel()
	//sortFilterModel.EndInsertRows()

	return albumId
}

func (c *Controller) addTracks(albumId int, tracks []string) {

	albumNode := albumData.CreateElement("album")
	albumNode.SetAttribute4("id", albumId)

	for i := 0; i < len(tracks); i++ {
		trackNumber := strconv.Itoa(i)
		if i < 10 {
			trackNumber = "0" + trackNumber
		}

		textNode := albumData.CreateTextNode(tracks[i])

		trackNode := albumData.CreateElement("track")
		trackNode.SetAttribute("number", trackNumber)
		trackNode.AppendChild(textNode)

		albumNode.AppendChild(trackNode)
	}

	archive := albumData.ElementsByTagName("archive")
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
