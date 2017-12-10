package auth

import (
	"encoding/json"
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/rtawormy14/cakman-go/controller"
)

type (
	// LoginRequest is ...
	loginRequest struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}
)

var auth controller.AuthController

// init function
func init() {

}

// Login is used for authenticate user
func Login(ctx *gin.Context) {
	if auth == nil {
		auth = controller.NewAuthController()
	}

	var login loginRequest
	if err := ctx.ShouldBindJSON(&login); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	authData, err := auth.Login(login.Username, login.Password)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	resByte, err := json.Marshal(authData)
	if err != nil {
		return
	}
	ctx.String(http.StatusOK, string(resByte))
}

// Logout is used for remove session authenticated user
func Logout(ctx *gin.Context) {
	if auth == nil {
		auth = controller.NewAuthController()
	}
	token := ctx.GetHeader("token")
	if token == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "bad token request"})
		return
	}

	err := auth.Logout(token)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error})
		return
	}

	ctx.String(http.StatusOK, "")
}
