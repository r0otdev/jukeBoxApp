package endpoints

import (
	"jukebox/api"
	app "jukebox/app"
	"sort"
	"strconv"

	"github.com/gin-gonic/gin"
)

func AddNewAlbumEndpoint(service app.JukeBoxServiceInterface) gin.HandlerFunc {
	return func(ctx *gin.Context) {

		var album app.AlbumRequest

		err := ctx.ShouldBindJSON(&album)
		if err != nil {
			api.SendInvalidRequestBodyResponse(ctx, err)
			return
		}

		err = service.AddNewAlbum(&album)
		if err != nil {
			api.SendErrorResponse(ctx, err)
			return
		}

		api.SendSuccessResponse(ctx, "success")
	}
}

func GetMusiciansByAlbumIDEndpoint(service app.JukeBoxServiceInterface) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		
		var AlbumID = ctx.Param("albumID")

		albumIDInt, err := strconv.Atoi(AlbumID)
		if err != nil || albumIDInt <= 0 {
			api.SendInvalidPropertyResponse(ctx, "albumID")
			return
		}

		musicians := service.GetMusiciansByAlbumID(albumIDInt)

		sort.Slice(musicians, func(i, j int) bool {
			return musicians[i].Name < musicians[j].Name
		})

		api.SendSuccessResponse(ctx, musicians)
	}
}

func GetAlbumsEndpoint(service app.JukeBoxServiceInterface) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		
		albums := service.GetAlbums()

		sort.Slice(albums, func(i, j int) bool {
			return albums[i].ReleaseDate.Unix() < albums[j].ReleaseDate.Unix()
		})

		api.SendSuccessResponse(ctx, albums)
	}
}

func UpdateAlbumEndpoint(service app.JukeBoxServiceInterface) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		
		var album app.AlbumRequest

		var AlbumID = ctx.Param("albumID")

		albumIDInt, err := strconv.Atoi(AlbumID)
		if err != nil || albumIDInt <= 0 {
			api.SendInvalidPropertyResponse(ctx, "albumID")
			return
		}

		err = ctx.ShouldBindJSON(&album)
		if err != nil {
			api.SendInvalidRequestBodyResponse(ctx, err)
			return
		}

		err = service.UpdateAlbum(albumIDInt, &album)
		if err != nil {
			api.SendErrorResponse(ctx, err)
			return
		}

		api.SendSuccessResponse(ctx, "success")
	}
}