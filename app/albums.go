package app

import (
	"errors"
	"time"
)

const albumIDStarting = 1000

type AlbumRequest struct {
	Name string `json:"name"`
	ReleaseDate time.Time `json:"release_date"`
	Genre string `json:"genre"`
	Price float64 `json:"price"`
	Description string `json:"description"`
	Musicians []int `json:"musicians"`
}

type Album struct {
	ID int `json:"id"`
	Name string `json:"name"`
	ReleaseDate time.Time `json:"release_date"`
	Genre string `json:"genre"`
	Price float64 `json:"price"`
	Description string `json:"description"`
	Musicians []int `json:"musicians"`
}

func (al *Album) addMusicians(musicians ...int) {
	al.Musicians = append(al.Musicians, al.Musicians...)
	
	origAlbums := make([]int, 0)

	var albumMap = make(map[int]struct{}, len(al.Musicians))
	for _, album := range al.Musicians {
		if _, exist := albumMap[album]; !exist {
			origAlbums = append(origAlbums, album)
		}
		albumMap[album] = struct{}{}
	}

	al.Musicians = origAlbums
}

func (al *Album) deleteMusicians(delMusician int) {
	var delIndex = -1 
	for index, Musicians := range al.Musicians {
		if Musicians == delMusician {
			delIndex = index
		}
	} 

	if delIndex > -1 {
		if delIndex == len(al.Musicians)-1 {
			al.Musicians = al.Musicians[:delIndex]
		}else{
			al.Musicians = append(al.Musicians[:delIndex], al.Musicians[delIndex+1:]...)
		}
	}
}

func (al *Album) validateName() error {
	if al.Name == "" {
		return errors.New("Name required")
	}
	
	if len(al.Name) < 5 {
		return errors.New("Name must have 5 characters")
	}

	return nil	
}

func (al *Album) validateReleaseDate() error {

	if al.ReleaseDate.IsZero() {
		return errors.New("Invalid date")
	}
	
	if al.ReleaseDate.After(time.Now()) {
		return errors.New("Date in future is not allowed")
	}
	return nil
}

func (al *Album) validateGenre() error {
	// TODO
	return nil
}

func (al *Album) validatePrice() error {
	if al.Price < 100 || al.Price > 1000 {
		errors.New("Invalid price")
	}
	return nil
}

func (al *Album) validateDescription() error {
	// TODO
	return nil
}

func NewAlbumStore() *AlbumStore {
	return &AlbumStore{lastIndex: albumIDStarting, Albums: make(map[int]*Album, 0)}
}

type AlbumStore struct {
	lastIndex int
	Albums map[int]*Album
}

func (as *AlbumStore) AddAlbum(album *Album) error {

	if err := album.validateName(); err != nil {
		return err
	} 

	if err := album.validateReleaseDate(); err != nil {
		return err
	}
	
	if err := album.validateGenre(); err != nil {
		return err
	} 

	if album.Genre != "" {
		if err := album.validateGenre(); err != nil {
			return err
		} 
	}

	if err := album.validatePrice(); err != nil {
		return err
	} 
	
	if album.Description != "" {
		if err := album.validateDescription(); err != nil {
			return err
		} 
	}

	as.lastIndex = as.lastIndex + 1

	album.ID = as.lastIndex

	as.Albums[as.lastIndex] = album

	return nil
}

func (as *AlbumStore) UpdateAlbum(albumID int, album *Album) error {
	
	albumRecord, exists := as.Albums[albumID]
	if !exists {
		return errors.New("Album not found")
	}

	if album.Name != "" {
		if err := album.validateName(); err != nil {
			return err
		} 
		albumRecord.Name = album.Name
	}
	
	if !album.ReleaseDate.IsZero() {
		if err := album.validateReleaseDate(); err != nil {
			return err
		} 
		albumRecord.ReleaseDate = album.ReleaseDate
	}

	if album.Genre != "" {
		if err := album.validateGenre(); err != nil {
			return err
		} 
		albumRecord.Genre = album.Genre
	}

	if album.Genre != "" {
		if err := album.validateGenre(); err != nil {
			return err
		} 
		albumRecord.Genre = album.Genre
	}

	if album.Price != 0 {
		if err := album.validatePrice(); err != nil {
			return err
		} 
		albumRecord.Price = album.Price
	}
	
	if album.Description != "" {
		if err := album.validateDescription(); err != nil {
			return err
		} 
		albumRecord.Description = album.Description
	}

	if album.Description != "" {
		if err := album.validateDescription(); err != nil {
			return err
		} 
		albumRecord.Description = album.Description
	}

	if len(album.Musicians) == 0 {
		albumRecord.Musicians = album.Musicians
	}

	return nil
}

func (al *AlbumStore) GetAlbumsFiltered(AlbumIDs []int) []*Album {
	
	var AlbumsRet []*Album

	if len(AlbumIDs) == 0 {
		for _, mObj := range al.Albums {
			AlbumsRet = append(AlbumsRet, mObj)
		}
		return AlbumsRet
	}
	
	for _, mID := range AlbumIDs {
		if msObj, exists := al.Albums[mID]; exists {
			AlbumsRet = append(AlbumsRet, msObj)
		}
	}
	
	return AlbumsRet
}

type AlbumStoreInterface interface {
	AddAlbum(album *Album) error
	UpdateAlbum(albumID int, album *Album) error
	GetAlbumsFiltered(albumIDs []int) []*Album
}

