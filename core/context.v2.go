package core

import (
	"errors"

	"firebase.google.com/go/v4/auth"
	"github.com/labstack/echo/v4"
)

type ContextV2 struct {
	echo.Context
}

func (c *ContextV2) GetAuthToken() (*auth.Token, error) {
	token := auth.Token{}
	token, ok := c.Get("authToken").(auth.Token)
	if !ok {
		return nil, errors.New("Unauthorized user")
	}
	return &token, nil
}
