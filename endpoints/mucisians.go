package endpoints

import (
	"jukebox/api"
	app "jukebox/app"
	"sort"
	"strconv"

	"github.com/gin-gonic/gin"
)

func AddNewMusicianEndpoint(service app.JukeBoxServiceInterface) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		
		var musician app.MusicianRequest

		err := ctx.ShouldBindJSON(&musician)
		if err != nil {
			api.SendInvalidRequestBodyResponse(ctx, err)
			return
		}

		err = service.AddNewMusician(&musician)
		if err != nil {
			api.SendErrorResponse(ctx, err)
			return
		}

		api.SendSuccessResponse(ctx, "success")
	}
}

func GetAlbumsByMusicianIDEnpoint(service app.JukeBoxServiceInterface) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		
		var musicianID = ctx.Param("musicianID")

		musicianIDInt, err := strconv.Atoi(musicianID)
		if err != nil|| musicianIDInt <= 0 {
			api.SendInvalidPropertyResponse(ctx, "musicianID")
			return
		}

		albums := service.GetAlbumsByMusicianID(musicianIDInt)

		sort.Slice(albums, func(i, j int) bool {
			return albums[i].Price < albums[j].Price
		})

		api.SendSuccessResponse(ctx, albums)
	}
}

func GetMusiciansEndpoint(service app.JukeBoxServiceInterface) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		
		musicians := service.GetMusicians()

		api.SendSuccessResponse(ctx, musicians)
	}
}

func UpdateMusicianEndpoint(service app.JukeBoxServiceInterface) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		
		var musician app.MusicianRequest

		var musicianID = ctx.Param("musicianID")

		musicianIDInt, err := strconv.Atoi(musicianID)
		if err != nil || musicianIDInt <= 0 {
			api.SendInvalidPropertyResponse(ctx, "musicianID")
			return
		}

		err = ctx.ShouldBindJSON(&musician)
		if err != nil {
			api.SendInvalidRequestBodyResponse(ctx, err)
			return
		}

		err = service.UpdateMusician(musicianIDInt, &musician)
		if err != nil {
			api.SendErrorResponse(ctx, err)
			return
		}

		api.SendSuccessResponse(ctx, "success")
	}
}