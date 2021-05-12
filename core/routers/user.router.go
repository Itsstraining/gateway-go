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
		return c.NoContent(200)
	})

	return nil, router
}
