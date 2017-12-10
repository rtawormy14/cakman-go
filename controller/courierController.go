package controller

import (
	"log"
	"time"

	courierModel "github.com/rtawormy14/cakman-go/model/courier"
)

type (
	// CourierController is...
	CourierController interface {
		GetCourier(int64) (courierModel.Courier, error)
		UpdateCourier(int64, courierModel.Courier) error
		UpdateLocation(int64, float64, float64) error
	}

	// CourierCtr is ...
	CourierCtr struct{}
)

// NewCourierController is...
func NewCourierController() CourierController {
	return &CourierCtr{}
}

// GetCourier is ...
func (c *CourierCtr) GetCourier(courierID int64) (courier courierModel.Courier, err error) {
	return courier.GetCourier(courierID)
}

// UpdateCourier is ...
func (c *CourierCtr) UpdateCourier(courierID int64, newData courierModel.Courier) (err error) {

	courierObject, err := courier.GetCourier(courierID)
	if err != nil {
		log.Println("[CourierController][UpdateCourier] Error when getting current data courier :", courierID)
		return err
	}

	if newData.Username != "" {
		courierObject.Username = newData.Username
	}
	if newData.Password != "" {
		courierObject.Password = newData.Password
	}
	if newData.Name != "" {
		courierObject.Name = newData.Name
	}
	if newData.Email != "" {
		courierObject.Email = newData.Email
	}
	if newData.Phone != "" {
		courierObject.Phone = newData.Phone
	}

	courierObject.UpdateBy = courierID
	courierObject.UpdateTime = time.Now()

	err = courier.Update(courierObject, nil)
	if err != nil {
		log.Printf("[CourierController][UpdateCourier] Error when updating data courier : %v", courierObject)
	}
	return err
}

// UpdateLocation is ...
func (c *CourierCtr) UpdateLocation(courierID int64, longitude, lattitude float64) (err error) {
	courierObject, err := courier.GetCourier(courierID)
	if err != nil {
		log.Println("[CourierController][UpdateLocation] Error when getting current data courier :", courierID)
		return err
	}

	if longitude > 0 {
		courierObject.Longitude = longitude
	}
	if lattitude > 0 {
		courierObject.Lattitude = lattitude
	}

	courierObject.LocUpdateTime = time.Now()

	err = courier.UpdateLocation(courierObject, nil)
	if err != nil {
		log.Printf("[CourierController][UpdateLocation] Error when updating data courier : %v", courierObject)
	}
	return err
}
