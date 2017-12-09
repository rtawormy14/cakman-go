package handler

import (
	"strconv"

	"github.com/gin-gonic/gin"
)

//GetDefaultParam is used for getting default param for GET
func GetDefaultParam(ctx *gin.Context) (page int64, limit int64, token string) {
	//get page
	pPage := ctx.Query("page")
	if pPage != "" {
		page, _ = strconv.ParseInt(pPage, 10, 64)
	}

	//get limit
	pLimit := ctx.Query("limit")
	if pLimit != "" {
		limit, _ = strconv.ParseInt(pLimit, 10, 64)
	}

	//get token
	token = ctx.GetHeader("token")

	return page, limit, token
}
