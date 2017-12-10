package controller

import (
	"errors"
	"log"
	"time"

	deliveryModel "github.com/rtawormy14/cakman-go/model/delivery"
	orderModel "github.com/rtawormy14/cakman-go/model/order"
	"github.com/rtawormy14/cakman-go/util/database"
)

type (
	// DeliveryController is...
	DeliveryController interface {
		GetDelivery(int64) (deliveryModel.Delivery, error)
		GetDeliveryHistory(int64, int64, int64) ([]deliveryModel.Delivery, error)
		GetDeliveryList(int64, int64, deliveryModel.Delivery) ([]deliveryModel.Delivery, error)
		PickupDelivery(int64, int64) error
		UpdateDelivery(deliveryModel.Delivery) error
	}

	// DeliveryCtr is ...
	DeliveryCtr struct{}
)

// NewDeliveryController is...
func NewDeliveryController() DeliveryController {
	return &DeliveryCtr{}
}

// GetDelivery is ...
func (d *DeliveryCtr) GetDelivery(id int64) (deliveryObj deliveryModel.Delivery, err error) {
	deliveryObj, err = delivery.GetDelivery(id)
	if err != nil {
		log.Println(err)
		return
	}

	return
}

// GetDeliveryHistory is ...
func (d *DeliveryCtr) GetDeliveryHistory(page int64, limit int64, courierID int64) (deliveries []deliveryModel.Delivery, err error) {
	deliveries, err = delivery.GetDeliveryHistory(page, limit, courierID)
	if err != nil {
		log.Println(err)
		return
	}
	for i := 0; i < len(deliveries); i++ {
		orderObj, _ := order.GetOrder(deliveries[i].OrderID)
		deliveries[i].Order = &orderObj
	}
	return
}

// GetDeliveryList is ...
func (d *DeliveryCtr) GetDeliveryList(page int64, limit int64, filter deliveryModel.Delivery) (deliveries []deliveryModel.Delivery, err error) {
	deliveries, err = delivery.GetDeliveryList(page, limit, filter)
	if err != nil {
		log.Println(err)
		return
	}
	for i := 0; i < len(deliveries); i++ {
		orderObj, _ := order.GetOrder(deliveries[i].OrderID)
		deliveries[i].Order = &orderObj
	}
	return
}

// PickupDelivery is ...
func (d *DeliveryCtr) PickupDelivery(courierID int64, orderID int64) (err error) {
	orderObj, err := order.GetOrder(orderID)
	if err != nil {
		log.Println("[Delivery][PickupDelivery] error when get order information : ", err)
		return
	}
	if orderObj.Status == orderModel.StatusInprogress {
		return errors.New("You cannot pickup order that has been picked up by another courier")
	}
	if orderObj.Status == orderModel.StatusFinish {
		return errors.New("You cannot pickup order that has been delivered")
	}

	deliveryObj := deliveryModel.Delivery{
		CourierID:  courierID,
		OrderID:    orderID,
		Status:     deliveryModel.StatusPickup,
		CreateBy:   courierID,
		CreateTime: time.Now(),
	}

	db := database.DB
	tx := db.MustBegin()
	defer tx.Rollback()

	deliveryObj, err = delivery.Insert(deliveryObj, tx)
	if err != nil {
		log.Println("[Delivery][PickupDelivery] error when inserting data delivery : ", err)
		return
	}

	orderObj.Status = orderModel.StatusInprogress
	err = order.Update(orderObj, tx)
	if err != nil {
		log.Println("[Delivery][PickupDelivery] error when update order status : ", err)
		return
	}

	historyObj := deliveryModel.History{
		DeliveryID: deliveryObj.ID,
		Status:     deliveryObj.Status,
		Note:       deliveryObj.Note,
		CreateTime: time.Now(),
	}
	err = history.Insert(historyObj, tx)
	if err != nil {
		log.Println("[Delivery][PickupDelivery] error when inserting data delivery history : ", err)
		return
	}
	err = tx.Commit()
	if err != nil {
		log.Println("[Delivery][PickupDelivery] error when commit transaction : ", err)
	}

	return
}

// UpdateDelivery is ....
func (d *DeliveryCtr) UpdateDelivery(newData deliveryModel.Delivery) (err error) {
	//get old delivery data
	deliveryObj, err := delivery.GetDelivery(newData.ID)
	if err != nil {
		log.Println("[Delivery][UpdateDelivery] error when get delivery information : ", err)
		return
	}
	// check only authorized user is permitted
	if deliveryObj.CourierID != newData.CourierID {
		return errors.New("you are not authorize to update this data")
	}
	// check status is valid or not
	if _, ok := deliveryModel.ValidStatus[newData.Status]; !ok {
		return errors.New("invalid status")
	}

	db := database.DB
	tx := db.MustBegin()
	defer tx.Rollback()

	// Update Delivery
	err = delivery.Update(newData, nil)
	if err != nil {
		log.Println("[Delivery][UpdateDelivery] error when update data delivery : ", err)
		return
	}

	// Get Order Data to update its status
	orderObj, err := order.GetOrder(deliveryObj.OrderID)
	if err != nil {
		log.Println("[Delivery][UpdateDelivery] error when get order information: ", err)
		return
	}
	if newData.Status == deliveryModel.StatusCancel {
		orderObj.Status = orderModel.StatusWarehouse
	} else if newData.Status == deliveryModel.StatusFinish {
		orderObj.Status = orderModel.StatusFinish
	} else if newData.Status == deliveryModel.StatusPickup {
		orderObj.Status = orderModel.StatusInprogress
	}
	err = order.Update(orderObj, tx)
	if err != nil {
		log.Println("[Delivery][UpdateDelivery] error when update data order : ", err)
		return
	}

	// Get Delivery Data to add delivery history
	deliveryObj, _ = delivery.GetDelivery(newData.ID)

	historyObj := deliveryModel.History{
		DeliveryID: deliveryObj.ID,
		Status:     deliveryObj.Status,
		Note:       deliveryObj.Note,
		CreateTime: time.Now(),
	}
	err = history.Insert(historyObj, tx)
	if err != nil {
		log.Println("[Delivery][UpdateDelivery] error when inserting data delivery history : ", err)
		return
	}
	err = tx.Commit()
	if err != nil {
		log.Println("[Delivery][UpdateDelivery] error when commit transaction : ", err)
	}
	return
}
