package app

import (
	"errors"
)

const musicianIDStarting = 1000

type MusicianRequest struct {
	Name string `json:"name"`
	Type string `json:"type"`
	Albums []int `json:"albums"`
}

type Musician struct {
	ID int `json:"id"`
	Name string	`json:"name"`
	Type string	`json:"type"`
	Albums []int `json:"albums"`
}

func (ms *Musician) addAlbums(albums ...int) {
	
	ms.Albums = append(ms.Albums, albums...)
	
	origAlbums := make([]int, 0)

	var albumMap = make(map[int]struct{}, len(ms.Albums))
	for _, album := range ms.Albums {
		if _, exist := albumMap[album]; !exist {
			origAlbums = append(origAlbums, album)
		}
		albumMap[album] = struct{}{}
	}


	ms.Albums = origAlbums
}

func (ms *Musician) deleteAlbum(delAlbumID int) {
	var delIndex = -1 
	for index, album := range ms.Albums {
		if album == delAlbumID {
			delIndex = index
		}
	} 

	if delIndex > -1 {
		if delIndex == len(ms.Albums)-1 {
			ms.Albums = ms.Albums[:delIndex]
		}else{
			ms.Albums = append(ms.Albums[:delIndex], ms.Albums[delIndex+1:]...)
		}
	}
}

func (ms *Musician) validateName() error {

	if ms.Name == "" {
		return errors.New("Name required")
	}
	
	if len(ms.Name) < 3 {
		return errors.New("Name must have 3 characters")
	}

	return nil	
}

func (ms *Musician) validateType() error {
	if ms.Type == "" {
		errors.New("Type required")
	}
	return nil
}

func NewMusicianStore() *MusicianStore {
	return &MusicianStore{lastIndex: musicianIDStarting, Musicians: make(map[int]*Musician, 0)}
}

type MusicianStore struct {
	lastIndex int
	Musicians map[int]*Musician
}

func (ms *MusicianStore) AddMusician(musician *Musician) error {

	if err := musician.validateName(); err != nil {
		return err
	}

	if musician.Type != "" {
		if err := musician.validateType(); err != nil {
			return err
		}
	}

	ms.lastIndex = ms.lastIndex + 1

	musician.ID = ms.lastIndex

	ms.Musicians[ms.lastIndex] = musician

	return nil
}

func (ms *MusicianStore) UpdateMusician(musicianID int, musician *Musician) error {

	musicianRecord, exists := ms.Musicians[musicianID]
	if !exists {
		return errors.New("Musician not found")
	}

	if musician.Name != "" {
		if err := musician.validateName(); err != nil {
			return err
		} 
		musicianRecord.Name = musician.Name
	}
	
	if musician.Type != "" {
		if err := musician.validateType(); err != nil {
			return err
		} 
		musicianRecord.Type = musician.Type
	}

	ms.Musicians[musicianID].Name = musicianRecord.Name
	ms.Musicians[musicianID].Type = musicianRecord.Type
	ms.Musicians[musicianID].Albums = musicianRecord.Albums

	return nil
}

func (ms *MusicianStore) GetMusiciansFiltered(musicianIDs []int) []*Musician {

	var musiciansRet []*Musician

	if len(musicianIDs) == 0 {
		for _, mObj := range ms.Musicians {
			musiciansRet = append(musiciansRet, mObj)
		}
		return musiciansRet
	}
	
	for _, mID := range musicianIDs {
		if msObj, exists := ms.Musicians[mID]; exists {
			musiciansRet = append(musiciansRet, msObj)
		}
	}
	
	return musiciansRet
}

type MusicianStoreInterface interface {
	AddMusician(musician *Musician) error
	UpdateMusician(musicianID int, musician *Musician) error
	GetMusiciansFiltered(musicianIDs []int) []*Musician
}

