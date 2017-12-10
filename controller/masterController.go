package controller

import (
	cityModel "github.com/rtawormy14/cakman-go/model/city"
	countryModel "github.com/rtawormy14/cakman-go/model/country"
	provinceModel "github.com/rtawormy14/cakman-go/model/province"
)

type (
	// MasterController is...
	MasterController interface {
		GetCountry(int64) (countryModel.Country, error)
		GetCountryList(int64, int64, countryModel.Country) ([]countryModel.Country, error)
		GetProvince(int64) (provinceModel.Province, error)
		GetProvinceList(int64, int64, provinceModel.Province) ([]provinceModel.Province, error)
		GetCity(int64) (cityModel.City, error)
		GetCityList(int64, int64, cityModel.City) ([]cityModel.City, error)
	}

	// master is ...
	MasterCtr struct{}
)

// NewMasterController is...
func NewMasterController() MasterController {

	return &MasterCtr{}
}

// GetCountry is ...
func (m *MasterCtr) GetCountry(code int64) (countryModel.Country, error) {
	return country.GetCountry(code)
}

// GetCountryList is ...
func (m *MasterCtr) GetCountryList(page int64, limit int64, filter countryModel.Country) ([]countryModel.Country, error) {
	return country.GetCountryList(page, limit, filter)
}

// GetProvince is ...
func (m *MasterCtr) GetProvince(code int64) (provinceModel.Province, error) {
	return province.GetProvince(code)
}

// GetProvinceList is ...
func (m *MasterCtr) GetProvinceList(page int64, limit int64, filter provinceModel.Province) ([]provinceModel.Province, error) {
	return province.GetProvinceList(page, limit, filter)
}

// GetCity is ...
func (m *MasterCtr) GetCity(code int64) (cityModel.City, error) {
	return city.GetCity(code)
}

// GetCityList is ...
func (m *MasterCtr) GetCityList(page int64, limit int64, filter cityModel.City) ([]cityModel.City, error) {
	return city.GetCityList(page, limit, filter)
}
