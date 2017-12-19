package controller

import (
	authModel "github.com/rtawormy14/cakman-go/model/authentication"
	cityModel "github.com/rtawormy14/cakman-go/model/city"
	countryModel "github.com/rtawormy14/cakman-go/model/country"
	courierModel "github.com/rtawormy14/cakman-go/model/courier"
	deliveryModel "github.com/rtawormy14/cakman-go/model/delivery"
	provinceModel "github.com/rtawormy14/cakman-go/model/province"
)

var (
	country  countryModel.Countrier
	province provinceModel.Provincer
	city     cityModel.Citier
	auth     authModel.Authenticator
	courier  courierModel.Courierer
	order    deliveryModel.Orderer
	delivery deliveryModel.Deliverier
	history  deliveryModel.Historier
)

// InitController is ....
func InitController() {
	country = countryModel.NewCountry()
	province = provinceModel.NewProvince()
	city = cityModel.NewCity()
	auth = authModel.NewAuthentication()
	courier = courierModel.NewCourier()
	order = deliveryModel.NewOrder()
	delivery = deliveryModel.NewDelivery()
	history = deliveryModel.NewHistory()
}
