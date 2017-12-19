package order

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rtawormy14/cakman-go/controller"
	"github.com/rtawormy14/cakman-go/handler"
	"github.com/rtawormy14/cakman-go/model/delivery"
)

var orderCtr controller.OrderController
var authCtr controller.AuthController

func init() {
	if orderCtr == nil {
		orderCtr = controller.NewOrderController()
	}
	if authCtr == nil {
		authCtr = controller.NewAuthController()
	}
}

// GetOrders will return a list of order.
func GetOrders(ctx *gin.Context) {
	//token is not used
	page, limit, token := handler.GetDefaultParam(ctx)

	resi := ctx.Query("resi")

	var resByte []byte
	if resi != "" {
		GetOrderDetail(ctx)
	} else {
		if _, ok := authCtr.IsAuthenticate(token); !ok {
			ctx.JSON(http.StatusUnauthorized, gin.H{"error": "token is expired"})
			return
		}

		orders, err := orderCtr.GetOrderList(page, limit, delivery.Order{})
		if err != nil {
			log.Println(err)
		}
		resByte, err = json.Marshal(orders)
		if err != nil {
			return
		}
		ctx.String(http.StatusOK, string(resByte))
	}
}

// GetOrderDetail will return detail order.
func GetOrderDetail(ctx *gin.Context) {
	resi := ctx.Query("resi")

	var resByte []byte
	if resi != "" {
		orderObj, err := orderCtr.FindResi(resi)
		if err != nil {
			log.Println(err)
		}
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": "resi not found"})
			return
		}
		resByte, err = json.Marshal(orderObj)
		if err != nil {
			return
		}
	}

	ctx.String(http.StatusOK, string(resByte))
}

// AddOrder will update order data
func AddOrder(ctx *gin.Context) {

}

// UpdateOrder will update order data
func UpdateOrder(ctx *gin.Context) {

}
