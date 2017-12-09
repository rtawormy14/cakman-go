package controller

import (
	"errors"
	"log"
	"strings"
	"time"

	"github.com/satori/go.uuid"

	authModel "github.com/rtawormy14/cakman-go/model/authentication"
)

type (
	// AuthController is...
	AuthController interface {
		Login(string, string) (authModel.Authentication, error)
		Logout(string) error
		IsAuthenticate(string) (authModel.Authentication, bool)
	}

	// Authentication is ...
	Authentication struct{}
)

var (
	auth authModel.Authenticator
)

// NewAuthController is...
func NewAuthController() AuthController {
	auth = authModel.NewAuthenticator()

	return &Authentication{}
}

// Login is...
func (a *Authentication) Login(username string, password string) (authModel.Authentication, error) {
	// TODO change to check on courier data
	if username != "admin" || password != "password" {
		return authModel.Authentication{}, errors.New("wrong username or password")
	}
	token := generateToken()
	createTime := time.Now()
	duration := time.Second * 1800

	//TODO fix courier id get from courier table
	authData := authModel.Authentication{
		CourierID:  1,
		ExpireTime: createTime.Add(duration),
		CreateTime: createTime,
		Token:      token,
	}

	// insert to session
	err := auth.Insert(authData, nil)
	if err != nil {
		return authModel.Authentication{}, err
	}

	// get newest data
	authData, err = auth.GetAuthenticationByToken(token)
	if err != nil {
		log.Printf("[AuthController][Login] error when get newest data (token %s) : %s", token, err)
		return authModel.Authentication{}, errors.New("error when authenticating your account")
	}
	authData.ExpireIn = int64(duration / time.Second)

	return authData, nil
}

// Logout is...
func (a *Authentication) Logout(token string) (err error) {
	// Get Authentication Data
	authData, err := auth.GetAuthenticationByToken(token)
	if err != nil {
		return errors.New("session not found")
	}

	// Don't return error when error occured on remove, just log it
	err = auth.Remove(authData, nil)
	if err != nil {
		log.Println("[AuthController][Logout] error when get removing data :", err)
		err = nil
	}
	return err
}

// IsAuthenticate is...
func (a *Authentication) IsAuthenticate(token string) (authModel.Authentication, bool) {
	// Get Authentication Data
	authData, err := auth.GetAuthenticationByToken(token)
	if err != nil {
		return authModel.Authentication{}, false
	}

	// check token expiration
	if time.Now().After(authData.ExpireTime) {
		return authModel.Authentication{}, false
	}

	return authData, true
}

func generateToken() string {
	u1 := uuid.NewV4()

	return strings.Replace(u1.String(), "-", "", -1)
}
