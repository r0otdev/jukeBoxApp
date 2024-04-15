package app

import (
	"sort"
	"testing"
	"time"
)

func loadTestAlbums() map[int]*Album {
	return map[int]*Album{
		1000 : {
			ID: 1000,
			Name: "Album 1",
			Genre: "POP",
			Price: 120,
			Description: "Album 1 is good",
			ReleaseDate: time.Now(),
			Musicians: []int{2000, 2001},
		},
		1001: {
			ID: 1000,
			Name: "Album 2",
			Genre: "ROCK",
			Price: 130,
			Description: "Album 2 is good",
			ReleaseDate: time.Now(),
			Musicians: []int{2001, 2002},
		},
		1002 : {
			ID: 1000,
			Name: "Album 3",
			Genre: "INDIE",
			Price: 140,
			Description: "Album 3 is good",
			ReleaseDate: time.Now(),
			Musicians: []int{2001},
		},
	}
}

func loadTestMusicians() map[int]*Musician {
	return map[int]*Musician{
		2000: {
			ID: 2001,
			Name: "Singer 1",
			Type: "Singer",
			Albums: []int{1000},
		},
		2001: {
			ID: 2001,
			Name: "Drummer 1",
			Type: "Drummer",
			Albums: []int{1000, 1001, 1002},
		},
		2002: {
			ID: 2001,
			Name: "Singer 2",
			Type: "Singer",
			Albums: []int{1001},
		},
	}
}

func demoDataStore() *JukeBoxService {

	testMusiciansStore := &MusicianStore{
		lastIndex: 2002,
		Musicians: loadTestMusicians(),
	}

	testAlbumStore := &AlbumStore{
		lastIndex: 1002,
		Albums: loadTestAlbums(),
	}

	return &JukeBoxService{Albums: testAlbumStore, Musicians: testMusiciansStore}
}

/* 
TestApp includes the unit test cases for albums and musicians
*/
func TestApp(t *testing.T){

	testJukeBox := demoDataStore()

	testAlbum := func (t *testing.T){

		albumAddTest := func (t *testing.T)  {
			testAlbum := &Album{
				Name: "Album 4",
				Genre: "INDIE",
				Price: 140,
				Description: "Album 4 is good",
				ReleaseDate: time.Now(),
				Musicians: []int{2001},
			}

			err := testJukeBox.Albums.AddAlbum(testAlbum)
			if err != nil {
				t.Fatalf("Adding album failed")
			}

			fetchAlbums := testJukeBox.Albums.GetAlbumsFiltered([]int{testAlbum.ID})
			if len(fetchAlbums) == 0 {
				t.Fatalf("Adding album was not successfull")
			}
		}

		albumFetchTest := func (t *testing.T)  {

			fetchAlbums := testJukeBox.Albums.GetAlbumsFiltered([]int{1000})
			if len(fetchAlbums) == 0 {
				t.Fatalf("Fetching album failed")
			}
			
			if fetchAlbums[0] == loadTestAlbums()[1000] {
				t.Fatalf("Fetching album failed")
			}
		}

		updateAlbumTest := func (t *testing.T) {
			testAlbum := &Album{
				Name: "Album 1 Updated",
				Genre: "JAZZ",
				Price: 150,
				Description: "Album 1 is good to update",
				ReleaseDate: time.Now(),
				Musicians: []int{2001,2000},
			}

			testAlbumID := 1000

			err := testJukeBox.Albums.UpdateAlbum(testAlbumID, testAlbum)
			if err != nil {
				t.Fatalf("Updating an album failed - %s", err.Error())
			}

			fetchAlbums := testJukeBox.Albums.GetAlbumsFiltered([]int{testAlbumID})
			if len(fetchAlbums) == 0 {
				t.Fatalf("Fetching album failed")
			}

			fetchAlbum := fetchAlbums[0]
			if fetchAlbum.Name != testAlbum.Name &&
			fetchAlbum.Genre != testAlbum.Genre &&
			fetchAlbum.Price != testAlbum.Price &&
			fetchAlbum.Description != testAlbum.Description &&
			fetchAlbum.ReleaseDate != testAlbum.ReleaseDate {
				t.Fatalf("Updating an album failed")
			}

			fetchAlbumMusicians := fetchAlbum.Musicians
			sort.Ints(fetchAlbumMusicians)
			testAlbumMusicians := testAlbum.Musicians
			sort.Ints(testAlbumMusicians)

			for i := range fetchAlbumMusicians {
				if fetchAlbumMusicians[i] != testAlbumMusicians[i] {
					t.Fatalf("Updating album's musicians failed")
				}
			}
		}

		t.Run("AlbumAddTest", albumAddTest)
		t.Run("AlbumFetchTest", albumFetchTest)
		t.Run("AlbumUpdateTest", updateAlbumTest)
	}

	testMusician := func (t *testing.T){
		
		musicianAddTest := func (t *testing.T)  {
			testMusician := &Musician{
				Name: "Musician 4",
				Type: "Vocalist",
				Albums: []int{1002},
			}

			err := testJukeBox.Musicians.AddMusician(testMusician)
			if err != nil {
				t.Fatalf("Adding musician failed")
			}

			fetchMusicians := testJukeBox.Musicians.GetMusiciansFiltered([]int{testMusician.ID})
			if len(fetchMusicians) == 0 {
				t.Fatalf("Adding musician was not successfull")
			}
		}

		musicianFetchTest := func (t *testing.T)  {

			fetchMusicians := testJukeBox.Musicians.GetMusiciansFiltered([]int{2000})
			if len(fetchMusicians) == 0 {
				t.Fatalf("Fetching musician failed")
			}
			
			if fetchMusicians[0] == loadTestMusicians()[1000] {
				t.Fatalf("Fetching musician failed")
			}
		}

		updateMusicianTest := func (t *testing.T) {
			testMusician := &Musician{
				Name: "Musician 1 Updated",
				Albums: []int{1000, 1001},
			}

			testMusicianID := 2000

			err := testJukeBox.Musicians.UpdateMusician(testMusicianID, testMusician)
			if err != nil {
				t.Fatalf("Updating an musician failed - %s", err.Error())
			}

			fetchMusicians := testJukeBox.Musicians.GetMusiciansFiltered([]int{testMusicianID})
			if len(fetchMusicians) == 0 {
				t.Fatalf("Fetching musician failed")
			}

			fetchMusician := fetchMusicians[0]
			if fetchMusician.Name != testMusician.Name {
				t.Fatalf("Updating an musician failed")
			}

			fetchMusicianAlbums := fetchMusician.Albums
			sort.Ints(fetchMusicianAlbums)
			testMusicianAlbums := testMusician.Albums
			sort.Ints(testMusicianAlbums)

			for i := range fetchMusicianAlbums {
				if fetchMusicianAlbums[i] != testMusicianAlbums[i] {
					t.Fatalf("Updating musician's albums failed")
				}
			}
		}

		t.Run("MusicianAddTest", musicianAddTest)
		t.Run("MusicianFetchTest", musicianFetchTest)
		t.Run("MusicianUpdateTest", updateMusicianTest)

	}

	t.Run("AlbumsTest", testAlbum)
	t.Run("MusicianTest", testMusician)
}
