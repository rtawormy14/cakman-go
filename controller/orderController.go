package controller

import (
	"log"

	orderModel "github.com/rtawormy14/cakman-go/model/delivery"
)

type (
	// OrderController is...
	OrderController interface {
		FindResi(string) (orderModel.Order, error)
		GetOrder(int64) (orderModel.Order, error)
		GetOrderList(int64, int64, orderModel.Order) ([]orderModel.Order, error)
	}

	// OrderCtr is ...
	OrderCtr struct{}
)

// NewOrderController is...
func NewOrderController() OrderController {
	return &OrderCtr{}
}

// FindResi is ...
func (c *OrderCtr) FindResi(resi string) (orderObj orderModel.Order, err error) {
	orderObj, err = order.FindResi(resi)
	if err != nil {
		log.Println(err)
		return
	}
	//get master data information
	countryObj, _ := country.GetCountry(orderObj.CountryID)
	provinceObj, _ := province.GetProvince(orderObj.ProvinceID)
	cityObj, _ := city.GetCity(orderObj.CityID)

	orderObj.Country = countryObj
	orderObj.Province = provinceObj
	orderObj.City = cityObj

	// Add detail delivery order -> when order status is not in WAREHOUSE
	if orderObj.Status != orderModel.StatusOrderWarehouse {
		//get delivery information
		deliveryObj, err := delivery.GetDeliveryByOrderID(orderObj.ID)
		if err != nil {
			log.Println("[OrderController][FindResi] error while getting delivery order ->", err)
			return orderObj, err
		}

		// get courier information
		courierObj, err := courier.GetCourier(deliveryObj.CourierID)
		if err != nil {
			log.Println("[OrderController][FindResi] error while getting Courier Information ->", err)
		}
		deliveryObj.Courier = &courierObj

		// get delivery history informations
		histories, err := history.GetHistory(deliveryObj.ID)
		if err != nil {
			log.Println("[OrderController][FindResi] error while getting delivery history ->", err)
		}
		deliveryObj.History = &histories

		orderObj.Delivery = deliveryObj
	}
	return
}

// GetOrder is ...
func (c *OrderCtr) GetOrder(OrderID int64) (orderObj orderModel.Order, err error) {
	orderObj, err = order.GetOrder(OrderID)
	if err != nil {
		log.Println(err)
		return
	}
	countryObj, _ := country.GetCountry(orderObj.CountryID)
	provinceObj, _ := province.GetProvince(orderObj.ProvinceID)
	cityObj, _ := city.GetCity(orderObj.CityID)

	orderObj.Country = countryObj
	orderObj.Province = provinceObj
	orderObj.City = cityObj

	return
}

// GetOrderList is ...
func (c *OrderCtr) GetOrderList(page int64, limit int64, filter orderModel.Order) (orders []orderModel.Order, err error) {
	orders, err = order.GetOrderList(page, limit, filter)
	if err != nil {
		log.Println(err)
		return
	}
	for i := 0; i < len(orders); i++ {
		countryObj, _ := country.GetCountry(orders[i].CountryID)
		provinceObj, _ := province.GetProvince(orders[i].ProvinceID)
		cityObj, _ := city.GetCity(orders[i].CityID)

		orders[i].Country = countryObj
		orders[i].Province = provinceObj
		orders[i].City = cityObj
	}
	return
}
