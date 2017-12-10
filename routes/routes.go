package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/rtawormy14/cakman-go/handler/auth"
	"github.com/rtawormy14/cakman-go/handler/courier"
	"github.com/rtawormy14/cakman-go/handler/delivery"
	"github.com/rtawormy14/cakman-go/handler/master"
	"github.com/rtawormy14/cakman-go/handler/order"
	"github.com/rtawormy14/cakman-go/handler/util"
)

// InitRoute used for define route list
// All Route registration should be done here
func InitRoute(app *gin.Engine) {

	// Util Routes
	// default service
	app.GET("/", util.Default)
	// ping service
	app.GET("/ping", util.Ping)

	// Authentication Routes
	authRoute := app.Group("auth")
	// login service
	authRoute.POST("/login", auth.Login)
	// logout service
	authRoute.POST("/logout", auth.Logout)

	// Courier Routes
	courierRoute := app.Group("/users")
	// show profile service
	courierRoute.GET("/:courier_id", courier.ShowProfile)
	// update profile service
	courierRoute.PATCH("/:courier_id", courier.UpdateProfile)
	// update location service
	courierRoute.PATCH("/:courier_id/location", courier.UpdateLocation)

	// Master Data Routes
	// get country list
	app.GET("/countries", master.GetCountries)
	// get country data by code
	app.GET("/countries/:country_code", master.GetCountryByCode)
	// get province list
	app.GET("/provinces", master.GetProvinces)
	// get province data by code
	app.GET("/provinces/:province_code", master.GetProvinceByCode)
	// get city list
	app.GET("/cities", master.GetCities)
	// get city data by code
	app.GET("/cities/:city_code", master.GetCityByCode)

	// Order Routes
	orderRoute := app.Group("/orders")
	// get order list
	orderRoute.GET("", order.GetOrders)
	// get order detail
	orderRoute.GET("/:order_id", order.GetOrderDetail)
	// add new order
	orderRoute.POST("/:order_id", order.AddOrder)
	// update order
	orderRoute.PATCH("/:order_id", order.UpdateOrder)

	// Delivery Order Routes
	deliveryRoute := app.Group("/deliveries")
	// get current delivery list
	deliveryRoute.GET("/", delivery.GetDeliveryList)
	// get delivery history
	deliveryRoute.GET("/history", delivery.GetDeliveryHistory)
	// add new delivery order
	deliveryRoute.POST("/:order_id", delivery.AddDelivery)
	// update delivery order
	deliveryRoute.PATCH("/:delivery_order_id", delivery.UpdateDelivery)

}
