package master

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"

	"github.com/rtawormy14/cakman-go/controller"

	"github.com/rtawormy14/cakman-go/handler"
	"github.com/rtawormy14/cakman-go/model/city"
	"github.com/rtawormy14/cakman-go/model/country"
	"github.com/rtawormy14/cakman-go/model/province"
)

var master controller.MasterController

// init function
func init() {

}

func GetCountries(ctx *gin.Context) {
	if master == nil {
		master = controller.NewMasterController()
	}
	//token is not used
	page, limit, _ := handler.GetDefaultParam(ctx)

	pName := ctx.Query("country_name")
	filter := country.Country{
		Name: pName,
	}

	countries, err := master.GetCountryList(page, limit, filter)
	if err != nil {
		log.Println(err)
	}

	resByte, err := json.Marshal(countries)
	if err != nil {
		return
	}
	ctx.String(http.StatusOK, string(resByte))
}

func GetCountryByCode(ctx *gin.Context) {
	if master == nil {
		master = controller.NewMasterController()
	}

	//page, limit, token is not used
	_, _, _ = handler.GetDefaultParam(ctx)

	pCode := ctx.Param("country_code")
	code, _ := strconv.ParseInt(pCode, 10, 64)

	country, err := master.GetCountry(code)
	if err != nil {
		log.Println(err)
	}

	resByte, err := json.Marshal(country)
	if err != nil {
		return
	}
	ctx.String(http.StatusOK, string(resByte))
}

func GetProvinces(ctx *gin.Context) {
	if master == nil {
		master = controller.NewMasterController()
	}
	//token is not used
	page, limit, _ := handler.GetDefaultParam(ctx)

	pCode, _ := strconv.ParseInt(ctx.Query("country_code"), 10, 64)
	pName := ctx.Query("province_name")

	filter := province.Province{
		CountryCode: pCode,
		Name:        pName,
	}

	provinces, err := master.GetProvinceList(page, limit, filter)
	if err != nil {
		log.Println(err)
	}

	resByte, err := json.Marshal(provinces)
	if err != nil {
		return
	}
	ctx.String(http.StatusOK, string(resByte))
}

func GetProvinceByCode(ctx *gin.Context) {
	if master == nil {
		master = controller.NewMasterController()
	}

	//page, limit, token is not used
	_, _, _ = handler.GetDefaultParam(ctx)

	time.Sleep(time.Second)

	pCode := ctx.Param("province_code")
	code, _ := strconv.ParseInt(pCode, 10, 64)

	province, err := master.GetProvince(code)
	if err != nil {
		log.Println(err)
	}

	resByte, err := json.Marshal(province)
	if err != nil {
		return
	}
	ctx.String(http.StatusOK, string(resByte))
}

func GetCities(ctx *gin.Context) {
	if master == nil {
		master = controller.NewMasterController()
	}

	//token is not used
	page, limit, _ := handler.GetDefaultParam(ctx)

	pCode, _ := strconv.ParseInt(ctx.Query("province_code"), 10, 64)
	pName := ctx.Query("city_name")

	filter := city.City{
		ProvinceCode: pCode,
		Name:         pName,
	}

	cities, err := master.GetCityList(page, limit, filter)
	if err != nil {
		log.Println(err)
	}

	resByte, err := json.Marshal(cities)
	if err != nil {
		return
	}

	ctx.String(http.StatusOK, string(resByte))

}

func GetCityByCode(ctx *gin.Context) {
	if master == nil {
		master = controller.NewMasterController()
	}
	//page, limit, token is not used
	_, _, _ = handler.GetDefaultParam(ctx)

	time.Sleep(time.Second)

	pCode := ctx.Param("city_code")
	code, _ := strconv.ParseInt(pCode, 10, 64)

	city, err := master.GetCity(code)
	if err != nil {
		log.Println(err)
	}

	resByte, err := json.Marshal(city)
	if err != nil {
		return
	}
	ctx.String(http.StatusOK, string(resByte))
}
