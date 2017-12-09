package util

import (
	"github.com/gin-gonic/gin"
)

// init function
func init() {

}

// handler for ping operation
func Ping(ctx *gin.Context) {
	ctx.String(200, "pong")
}

// Default Root will return page not found
func Default(ctx *gin.Context) {
	ctx.Status(404)
}
