// +build boltdb

package model

import (
	"encoding/json"
	"os"
	"path/filepath"
	"strconv"

	"github.com/boltdb/bolt"
)

var boltdb, _ = bolt.Open(filepath.Join(os.Getenv("HOME"), "my.db"), 0600, nil)

func init() {

	if err := boltdb.Update(func(tx *bolt.Tx) error {

		for _, artist := range db {
			artistBucket, err := tx.CreateBucketIfNotExists([]byte(strconv.Itoa(artist.ID)))
			if err != nil {
				println("failed to create artist bucket")
				return err
			}

			jsn, err := json.Marshal(artist)
			if err != nil {
				return err
			}
			artistBucket.Put([]byte("json"), jsn)
		}

		return nil
	}); err != nil {
		println("failed to create boltdb:", err.Error())
	}

}

func getArtistCount() int {
	var artistCount int

	if err := boltdb.View(func(tx *bolt.Tx) error {

		if err := tx.ForEach(func(name []byte, b *bolt.Bucket) error {
			artistCount++
			return nil
		}); err != nil {
			return err
		}

		return nil
	}); err != nil {
		println("failed to getArtistCount:", err.Error())
	}

	return artistCount
}

func getArtistForID(ID int) *Artist {
	var artist *Artist

	if err := boltdb.View(func(tx *bolt.Tx) error {

		artistBucket := tx.Bucket([]byte(strconv.Itoa(ID)))
		if err := json.Unmarshal(artistBucket.Get([]byte("json")), &artist); err != nil {
			return err
		}

		return nil
	}); err != nil {
		println("failed to getArtistForID:", err.Error())
	}

	return artist
}

func GetArtistForName(name string) *Artist {
	var artist *Artist

	if err := boltdb.View(func(tx *bolt.Tx) error {

		if err := tx.ForEach(func(_ []byte, b *bolt.Bucket) error {

			var tmpArtist *Artist
			if err := json.Unmarshal(b.Get([]byte("json")), &tmpArtist); err != nil {
				return err
			}
			if tmpArtist.Name == name {
				artist = tmpArtist
			}

			return nil
		}); err != nil {
			return err
		}

		return nil
	}); err != nil {
		println("failed to GetArtistForName:", err.Error())
	}

	return artist
}

func getAlbumCount() int {
	var albumCount int
	for i := 0; i < getArtistCount(); i++ {
		albumCount += len(getArtistForID(i).Albums)
	}
	return albumCount
}

func GetAlbumCountForArtist(name string) int {
	return len(GetArtistForName(name).Albums)
}

func DeleteAlbum(ID int) {

	if err := boltdb.Update(func(tx *bolt.Tx) error {

		if err := tx.ForEach(func(_ []byte, b *bolt.Bucket) error {

			artistBucket := b

			var artist *Artist
			if err := json.Unmarshal(artistBucket.Get([]byte("json")), &artist); err != nil {
				return err
			}

			for _, album := range artist.Albums {
				if album.ID == ID {
					delete(artist.Albums, album.ID)

					jsn, err := json.Marshal(artist)
					if err != nil {
						return err
					}
					artistBucket.Put([]byte("json"), jsn)
				}
			}

			return nil
		}); err != nil {
			return err
		}

		return nil
	}); err != nil {
		println("failed to GetArtistForName:", err.Error())
	}

}

func CreateNewArtist(artistId int, name string) {
	artist := &Artist{artistId, name, make(map[int]*Album)}

	if err := boltdb.Update(func(tx *bolt.Tx) error {

		artistBucket, err := tx.CreateBucketIfNotExists([]byte(strconv.Itoa(artist.ID)))
		if err != nil {
			println("failed to create artist bucket")
			return err
		}

		jsn, err := json.Marshal(artist)
		if err != nil {
			return err
		}
		artistBucket.Put([]byte("json"), jsn)

		return nil
	}); err != nil {
		println("failed to CreateNewArtist:", err.Error())
	}

}

func CreateNewAlbum(artistId int, albumId int, title string, year int) {
	album := &Album{albumId, title, year}

	if err := boltdb.Update(func(tx *bolt.Tx) error {

		artistBucket := tx.Bucket([]byte(strconv.Itoa(artistId)))

		var artist *Artist
		if err := json.Unmarshal(artistBucket.Get([]byte("json")), &artist); err != nil {
			return err
		}

		if artist.Albums == nil {
			artist.Albums = make(map[int]*Album)
		}

		artist.Albums[album.ID] = album

		jsn, err := json.Marshal(artist)
		if err != nil {
			return err
		}
		artistBucket.Put([]byte("json"), jsn)

		return nil
	}); err != nil {
		println("failed to CreateNewAlbum:", err.Error())
	}

}

func GetNextArtistID() int {
	var highestID int

	if err := boltdb.View(func(tx *bolt.Tx) error {

		if err := tx.ForEach(func(name []byte, b *bolt.Bucket) error {

			artistId, err := strconv.Atoi(string(name))
			if err != nil {
				return err
			}

			if artistId > highestID {
				highestID = artistId
			}

			return nil
		}); err != nil {
			return err
		}

		return nil
	}); err != nil {
		println("failed to GetNextArtistID:", err.Error())
	}

	return highestID + 1
}

func GetNextAlbumID() int {
	var highestID int

	if err := boltdb.View(func(tx *bolt.Tx) error {

		if err := tx.ForEach(func(name []byte, b *bolt.Bucket) error {

			var artist *Artist
			if err := json.Unmarshal(b.Get([]byte("json")), &artist); err != nil {
				return err
			}

			for _, album := range artist.Albums {
				if album.ID > highestID {
					highestID = album.ID
				}
			}

			return nil
		}); err != nil {
			return err
		}

		return nil
	}); err != nil {
		println("failed to GetNextAlbumID:", err.Error())
	}

	return highestID + 1
}

func getModelArray() []*dbArrayStruct {

	m := make(map[int]*dbArrayStruct)

	for i := 0; i < getArtistCount(); i++ {
		artist := getArtistForID(i)
		for _, album := range artist.Albums {
			m[album.ID] = &dbArrayStruct{artist, album}
		}
	}

	o := make([]*dbArrayStruct, 0)

	for i := 0; i <= GetNextAlbumID(); i++ {
		if s, ok := m[i]; ok {
			o = append(o, s)
		}
	}

	return o
}
