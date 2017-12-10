package delivery

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/rtawormy14/cakman-go/controller"
	"github.com/rtawormy14/cakman-go/handler"
	"github.com/rtawormy14/cakman-go/model/delivery"
)

type (
	updateRequest struct {
		Status string `json:"status"`
		Note   string `json:"note"`
	}
)

var deliveryCtr controller.DeliveryController
var authCtr controller.AuthController

func init() {
	if deliveryCtr == nil {
		deliveryCtr = controller.NewDeliveryController()
	}
	if authCtr == nil {
		authCtr = controller.NewAuthController()
	}
}

// GetDeliveryList will return detail order.
func GetDeliveryList(ctx *gin.Context) {
	//page, limit, token is not used
	page, limit, token := handler.GetDefaultParam(ctx)

	if token == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "bad token request"})
		return
	}
	var courierID int64
	if authData, ok := authCtr.IsAuthenticate(token); ok {
		courierID = authData.CourierID
	} else {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "token is expired"})
		return
	}

	filter := delivery.Delivery{
		CourierID: courierID,
	}

	deliveries, err := deliveryCtr.GetDeliveryList(page, limit, filter)
	if err != nil {
		log.Println(err)
	}

	resByte, err := json.Marshal(deliveries)
	if err != nil {
		return
	}
	ctx.String(http.StatusOK, string(resByte))
}

// GetDeliveryHistory will return detail order.
func GetDeliveryHistory(ctx *gin.Context) {
	//page, limit, token is not used
	page, limit, token := handler.GetDefaultParam(ctx)

	if token == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "bad token request"})
		return
	}
	var courierID int64
	if authData, ok := authCtr.IsAuthenticate(token); ok {
		courierID = authData.CourierID
	} else {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "token is expired"})
		return
	}

	deliveries, err := deliveryCtr.GetDeliveryHistory(courierID, page, limit)
	if err != nil {
		log.Println(err)
	}

	resByte, err := json.Marshal(deliveries)
	if err != nil {
		return
	}
	ctx.String(http.StatusOK, string(resByte))
}

// AddDelivery will insert new delivery order
func AddDelivery(ctx *gin.Context) {
	//page, limit, token is not used
	_, _, token := handler.GetDefaultParam(ctx)

	if token == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "bad token request"})
		return
	}

	var courierID int64
	if authData, ok := authCtr.IsAuthenticate(token); ok {
		courierID = authData.CourierID
	} else {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "token is expired"})
		return
	}

	pOrderID := ctx.Param("order_id")
	orderID, _ := strconv.ParseInt(pOrderID, 10, 64)

	err := deliveryCtr.PickupDelivery(courierID, orderID)
	if err != nil {
		log.Println(err)
		ctx.String(http.StatusInternalServerError, "")
		return
	}
	ctx.String(http.StatusCreated, "")
}

// UpdateDelivery will update delivery order
func UpdateDelivery(ctx *gin.Context) {
	//page, limit is not used
	_, _, token := handler.GetDefaultParam(ctx)

	if token == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "bad token request"})
		return
	}

	var courierID int64
	if authData, ok := authCtr.IsAuthenticate(token); ok {
		courierID = authData.CourierID
	} else {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "token is expired"})
		return
	}

	pDeliveryID := ctx.Param("delivery_order_id")
	deliveryID, _ := strconv.ParseInt(pDeliveryID, 10, 64)
	var updateData updateRequest
	if err := ctx.ShouldBindJSON(&updateData); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	deliveryData := delivery.Delivery{
		ID:         deliveryID,
		CourierID:  courierID,
		Status:     updateData.Status,
		Note:       updateData.Note,
		UpdateBy:   courierID,
		UpdateTime: time.Now(),
	}

	err := deliveryCtr.UpdateDelivery(deliveryData)
	if err != nil {
		log.Println(err)
		ctx.String(http.StatusInternalServerError, "")
		return
	}
	ctx.String(http.StatusOK, "")
}
