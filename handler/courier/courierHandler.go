package courier

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/rtawormy14/cakman-go/controller"
	"github.com/rtawormy14/cakman-go/handler"
	"github.com/rtawormy14/cakman-go/model/courier"
)

type (
	locationRequest struct {
		Longitude float64 `json:"longitude"`
		Lattitude float64 `json:"lattitude"`
	}

	updateRequest struct {
		Username       string `json:"username"`
		Password       string `json:"password"`
		NewPassword    string `json:"new_password"`
		RepeatPassword string `json:"repeat_password"`
		Name           string `json:"name"`
		Phone          string `json:"phone"`
		Email          string `json:"email"`
	}
)

var courierCtr controller.CourierController
var authCtr controller.AuthController

// init function
func init() {

}

// ShowProfile will show profile data
func ShowProfile(ctx *gin.Context) {
	if courierCtr == nil {
		courierCtr = controller.NewCourierController()
	}

	//page, limit, token is not used
	_, _, _ = handler.GetDefaultParam(ctx)

	pCode := ctx.Param("courier_id")
	code, _ := strconv.ParseInt(pCode, 10, 64)

	courier, err := courierCtr.GetCourier(code)
	if err != nil {
		log.Println(err)
	}

	resByte, err := json.Marshal(courier)
	if err != nil {
		return
	}
	ctx.String(http.StatusOK, string(resByte))
}

// UpdateProfile will update profile
func UpdateProfile(ctx *gin.Context) {
	if courierCtr == nil {
		courierCtr = controller.NewCourierController()
	}
	if authCtr == nil {
		authCtr = controller.NewAuthController()
	}

	token := ctx.GetHeader("token")
	if token == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "bad token request"})
		return
	}

	pCourierID := ctx.Param("courier_id")
	courierID, _ := strconv.ParseInt(pCourierID, 10, 64)

	if authData, ok := authCtr.IsAuthenticate(token); ok {
		if courierID != authData.CourierID {
			ctx.JSON(http.StatusForbidden, gin.H{"error": "you are not allowed to do this action"})
			return
		}
	} else {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "token is expired"})
		return
	}

	var updateData updateRequest
	if err := ctx.ShouldBindJSON(&updateData); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	newData := courier.Courier{
		Name:     updateData.Name,
		Email:    updateData.Email,
		Phone:    updateData.Phone,
		Username: updateData.Username,
	}

	// update password event
	if updateData.NewPassword != "" {
		// check new password and repeat password is same
		if updateData.NewPassword != updateData.RepeatPassword {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": "confirm password is not same with new password"})
			return
		}
		// check if existing password is matched
		courier, _ := courierCtr.GetCourier(courierID)
		if courier.Password != updateData.Password {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": "wrong old password"})
			return
		}
		newData.Password = updateData.NewPassword
	}

	err := courierCtr.UpdateCourier(courierID, newData)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	courier, err := courierCtr.GetCourier(courierID)
	if err != nil {
		log.Printf("[Handler][UpdateCourier] error when getting new data : %s \n data : %v", err, courier)
		ctx.String(http.StatusInternalServerError, "")
		return
	}
	resByte, err := json.Marshal(courier)
	if err != nil {
		log.Printf("[Handler][UpdateCourier] error when marshalling json : %s \n data : %v", err, courier)
		ctx.String(http.StatusInternalServerError, "")
		return
	}
	ctx.String(http.StatusOK, string(resByte))
}

// UpdateLocation will update location data
func UpdateLocation(ctx *gin.Context) {
	if courierCtr == nil {
		courierCtr = controller.NewCourierController()
	}
	if authCtr == nil {
		authCtr = controller.NewAuthController()
	}

	token := ctx.GetHeader("token")
	if token == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "bad token request"})
		return
	}

	pCourierID := ctx.Param("courier_id")
	courierID, _ := strconv.ParseInt(pCourierID, 10, 64)

	if authData, ok := authCtr.IsAuthenticate(token); ok {
		if courierID != authData.CourierID {
			ctx.JSON(http.StatusForbidden, gin.H{"error": "you are not allowed to do this action"})
			return
		}
	} else {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "token is expired"})
		return
	}

	var location locationRequest
	if err := ctx.ShouldBindJSON(&location); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := courierCtr.UpdateLocation(courierID, location.Longitude, location.Lattitude)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	courier, err := courierCtr.GetCourier(courierID)
	if err != nil {
		log.Printf("[Handler][UpdateCourier] error when getting new data : %s \n data : %v", err, courier)
		ctx.JSON(http.StatusInternalServerError, "")
		return
	}
	resByte, err := json.Marshal(courier)
	if err != nil {
		log.Printf("[Handler][UpdateCourier] error when marshalling json : %s \n data : %v", err, courier)
		ctx.JSON(http.StatusInternalServerError, "")
		return
	}
	ctx.String(http.StatusOK, string(resByte))
}
