package main

import (
	"github.com/gin-gonic/gin"

	"github.com/rtawormy14/cakman-go/controller"
	"github.com/rtawormy14/cakman-go/routes"
	"github.com/rtawormy14/cakman-go/util/database"
)

func main() {
	r := gin.Default()

	// init Route
	routes.InitRoute(r)

	// init Database
	database.InitDB()

	// init controller
	controller.InitController()

	// run on port 4040
	r.Run(":4040")
}
