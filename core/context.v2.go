package core

import (
	"errors"

	"firebase.google.com/go/v4/auth"
	"github.com/labstack/echo/v4"
)

type ContextV2 struct {
	Context *echo.Context
}

func ToContextV2(ctx *echo.Context) *ContextV2 {
	return &ContextV2{Context: ctx}
}

func (c *ContextV2) GetAuthToken() (*auth.Token, error) {
	token := &auth.Token{}
	t := (*c.Context).Get("authToken")
	token, ok := t.(*auth.Token)
	if !ok {
		return nil, errors.New("Unauthorized user")
	}
	return token, nil
}
