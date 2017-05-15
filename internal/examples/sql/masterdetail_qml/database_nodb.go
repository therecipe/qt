// +build !boltdb

package main

func getArtistCount() int {
	return len(db)
}

func getArtistForID(ID int) *Artist {
	return db[ID]
}

func getArtistForName(name string) *Artist {
	for _, artist := range db {
		if artist.Name == name {
			return artist
		}
	}
	return nil
}

func getAlbumCount() int {
	var albumCount int
	for _, artists := range db {
		albumCount += len(artists.Albums)
	}
	return albumCount
}

func getAlbumCountForArtist(name string) int {
	return len(getArtistForName(name).Albums)
}

func deleteAlbum(ID int) {
	for _, artist := range db {
		for _, album := range artist.Albums {
			if album.ID == ID {
				delete(artist.Albums, album.ID)
			}
		}
	}
}

func createNewArtist(artistId int, name string) {
	db[artistId] = &Artist{artistId, name, make(map[int]*Album)}
}

func createNewAlbum(artistId int, albumId int, title string, year int) {
	db[artistId].Albums[albumId] = &Album{albumId, title, year}
}

func getNextArtistID() int {
	var highestID int
	for _, artist := range db {
		if artist.ID > highestID {
			highestID = artist.ID
		}
	}
	return highestID + 1
}

func getNextAlbumID() int {
	var highestID int
	for _, artist := range db {
		for _, album := range artist.Albums {
			if album.ID > highestID {
				highestID = album.ID
			}
		}
	}
	return highestID + 1
}

func getModelArray() []*dbArrayStruct {

	m := make(map[int]*dbArrayStruct)

	for _, artist := range db {
		for _, album := range artist.Albums {
			m[album.ID] = &dbArrayStruct{artist, album}
		}
	}

	o := make([]*dbArrayStruct, 0)

	for i := 0; i <= getNextAlbumID(); i++ {
		if s, ok := m[i]; ok {
			o = append(o, s)
		}
	}

	return o
}
