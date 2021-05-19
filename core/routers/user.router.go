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
		ctx := core.ToContextV2(&c)
		token, err := ctx.GetAuthToken()
		if err != nil {
			return echo.ErrUnauthorized
		}
		return c.JSON(200, token.UID)
	}, server.Middlewares.Auth.Authorize)

	return nil, router
}
