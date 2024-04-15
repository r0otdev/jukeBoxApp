package server

import (
	"jukebox/app"
	"jukebox/endpoints"

	"github.com/gin-gonic/gin"
)

/*
RunServer will start the api server
*/
func RunServer(){

	router := gin.Default()

	defaultJB := app.NewJukeBoxService(app.NewMusicianStore(), app.NewAlbumStore())

	router.POST("/v1/api/albums/", endpoints.AddNewAlbumEndpoint(defaultJB))
	router.POST("/v1/api/musicians/", endpoints.AddNewMusicianEndpoint(defaultJB))
	
	router.PATCH("/v1/api/albums/:albumID", endpoints.UpdateAlbumEndpoint(defaultJB))
	router.PATCH("/v1/api/musicians/:musicianID", endpoints.UpdateMusicianEndpoint(defaultJB))
	
	router.GET("/v1/api/albums/", endpoints.GetAlbumsEndpoint(defaultJB))
	router.GET("/v1/api/musicians/", endpoints.GetMusiciansEndpoint(defaultJB))
	
	router.GET("/v1/api/albums/:albumID/musicians", endpoints.GetMusiciansByAlbumIDEndpoint(defaultJB))
	router.GET("/v1/api/musicians/:musicianID/albums", endpoints.GetAlbumsByMusicianIDEnpoint(defaultJB))

	router.Run(":8077")
}