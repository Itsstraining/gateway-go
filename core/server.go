package core

import (
	"context"
	"main/middlewares"
	"main/models"

	"encoding/json"
	"io/ioutil"

	firebase "firebase.google.com/go/v4"
	"firebase.google.com/go/v4/auth"
	"github.com/labstack/echo/v4"
)

type Server struct {
	Echo        *echo.Echo
	Firebase    *firebase.App
	Auth        *auth.Client
	Middlewares struct {
		Auth *middlewares.Auth
	}
}

func NewServer(ip *string) (svr *Server, err error) {
	e := echo.New()

	// Read configuration file

	config := &models.Config{}

	file, err := ioutil.ReadFile("./config.json")

	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(file, config)
	if err != nil {
		return nil, err
	}

	// Load Firebase config file

	file, err = ioutil.ReadFile(config.FirebaseConfig)

	if err != nil {
		return nil, err
	}

	firebaseConfig := &firebase.Config{}

	err = json.Unmarshal(file, firebaseConfig)

	if err != nil {
		return nil, err
	}

	//Init Firebase admin

	firebaseApp, err := firebase.NewApp(context.Background(), firebaseConfig)
	if err != nil {
		return nil, err
	}
	auth, err := firebaseApp.Auth(context.Background())
	if err != nil {
		return nil, err
	}

	svr = &Server{
		Echo:     e,
		Firebase: firebaseApp,
		Auth:     auth,
		Middlewares: struct{ Auth *middlewares.Auth }{
			Auth: middlewares.NewAuth(auth),
		},
	}
	return svr, nil
}