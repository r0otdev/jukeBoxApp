package app

import (
	"errors"
)

type JukeBoxServiceInterface interface {
	GetAlbums() []*Album
	AddNewAlbum(album *AlbumRequest) error
	UpdateAlbum(albumID int, album *AlbumRequest) error
	GetMusiciansByAlbumID(albumID int) []*Musician
	
	GetMusicians() []*Musician
	AddNewMusician(musician *MusicianRequest) error
	UpdateMusician(musicianID int, musician *MusicianRequest) error
	GetAlbumsByMusicianID(musicianID int) []*Album
}

func NewJukeBoxService(musicStore MusicianStoreInterface, albumStore AlbumStoreInterface) *JukeBoxService {
	return &JukeBoxService{Musicians: musicStore, Albums: albumStore}
}

type JukeBoxService struct {
	Musicians MusicianStoreInterface
	Albums AlbumStoreInterface
}

func (jb *JukeBoxService) AddNewAlbum(album *AlbumRequest) error {

	var musicSlc []*Musician
	if len(album.Musicians) > 0 {
		musicSlc = jb.Musicians.GetMusiciansFiltered(album.Musicians)
		if len(musicSlc) != len(album.Musicians) {
			return errors.New("Invalid musician(s) input")
		}
	}

	albumNew := &Album{Name: album.Name, ReleaseDate: album.ReleaseDate, Genre: album.Genre, Price: album.Price, Description: album.Description, Musicians: album.Musicians}
	
	err := jb.Albums.AddAlbum(albumNew)

	if err != nil {
		return err
	}
	
	if len(musicSlc) == 0 {
		return nil
	}

	for _, musician := range musicSlc {
		musician.addAlbums(albumNew.ID)
	}

	return nil
}

func (jb *JukeBoxService) UpdateAlbum(albumID int, album *AlbumRequest) error {

	var musiciansSlc []*Musician
	if len(album.Musicians) > 0 {
		musiciansSlc = jb.Musicians.GetMusiciansFiltered(album.Musicians)
		if len(musiciansSlc) != len(album.Musicians) {
			return errors.New("Invalid musician(s) input")
		}
	}

	return jb.Albums.UpdateAlbum(albumID, &Album{Name: album.Name, ReleaseDate: album.ReleaseDate, Genre: album.Genre, Description: album.Description, Price: album.Price, Musicians: album.Musicians})
}

func (jb *JukeBoxService) GetAlbums() []*Album {
	
	albums := jb.Albums.GetAlbumsFiltered(nil)

	if len(albums) < 1 {
		return nil
	}

	return albums
}

func (jb *JukeBoxService) GetMusiciansByAlbumID(albumID int) []*Musician {
	
	albums := jb.Albums.GetAlbumsFiltered([]int{albumID})

	if len(albums) < 1 {
		return nil
	}

	musicSlc := jb.Musicians.GetMusiciansFiltered(albums[0].Musicians)

	return musicSlc
}

func (jb *JukeBoxService) AddNewMusician(musician *MusicianRequest) error {
	var albumsSlc []*Album
	if len(musician.Albums) > 0 {
		albumsSlc = jb.Albums.GetAlbumsFiltered(musician.Albums)
		if len(albumsSlc) != len(musician.Albums) {
			return errors.New("Invalid album(s) input")
		}
	}

	musicnNew := &Musician{Name: musician.Name, Type: musician.Type, Albums: musician.Albums}

	err := jb.Musicians.AddMusician(musicnNew)

	if err != nil {
		return err
	}
	
	if len(albumsSlc) == 0 {
		return nil
	}

	for _, album := range albumsSlc {
		album.addMusicians(musicnNew.ID)
	}

	return nil
}

func (jb *JukeBoxService) UpdateMusician(musicianID int, musician *MusicianRequest) error {
	var albumsSlc []*Album
	if len(musician.Albums) > 0 {
		albumsSlc = jb.Albums.GetAlbumsFiltered(musician.Albums)
		if len(albumsSlc) != len(musician.Albums) {
			return errors.New("Invalid album(s) input")
		}
	}

	return jb.Musicians.UpdateMusician(musicianID, &Musician{Name: musician.Name, Type: musician.Type, Albums: musician.Albums})
}

func (jb *JukeBoxService) GetMusicians() []*Musician {
	
	musicians := jb.Musicians.GetMusiciansFiltered(nil)

	if len(musicians) < 1 {
		return nil
	}

	return musicians
}

func (jb *JukeBoxService) GetAlbumsByMusicianID(musicianID int) []*Album {

	musicians := jb.Musicians.GetMusiciansFiltered([]int{musicianID})

	if len(musicians) < 1 {
		return nil
	}

	albumSlc := jb.Albums.GetAlbumsFiltered(musicians[0].Albums)

	return albumSlc
}