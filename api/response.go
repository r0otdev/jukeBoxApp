package api

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func formatResponse(ctx *gin.Context, status int, data interface{}, err error){
	retMap := make(gin.H, 0)
	if data != nil {
		retMap["data"] = data
	}
	if err != nil {
		retMap["error"] = err.Error()
	}
	ctx.JSON(status, retMap)
}

func SendSuccessResponse(ctx *gin.Context, data interface{}){
	formatResponse(ctx, http.StatusOK, data, nil)
}

func SendInternalErrorResponse(ctx *gin.Context, err error){
	formatResponse(ctx, http.StatusInternalServerError, nil, err)
}

func SendErrorResponse(ctx *gin.Context, err error){
	formatResponse(ctx, http.StatusBadRequest, nil, err)
}

func SendInvalidRequestBodyResponse(ctx *gin.Context, err error){
	formatResponse(ctx, http.StatusBadRequest, nil, errors.New("Invalid request body"))
}

func SendInvalidPropertyResponse(ctx *gin.Context, property string){
	formatResponse(ctx, http.StatusBadRequest, nil, fmt.Errorf("Invalid %s supplied", property))
}