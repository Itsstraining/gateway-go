package routers

import (
	"main/core"

	"github.com/labstack/echo/v4"
)

type User struct {
	Router *echo.Group
}

func NewUser(server *core.Server, route string) (err error, router *User) {
	router = &User{
		Router: server.Echo.Group(route),
	}

	router.Router.GET("/profile", func(c echo.Context) error {
		return c.JSON(200, c.Get("authToken"))
	}, server.Middlewares.Auth.Authorize)

	return nil, router
}
