package routers

import (
	"main/access"
	"main/business"
	"main/core"
	"main/models"
	"net/http"

	"github.com/labstack/echo/v4"
)

type Challenge struct {
	Router *echo.Group
}

func NewChallenge(server *core.Server, route string) (router *Challenge, err error) {
	router = &Challenge{
		Router: server.Echo.Group(route),
	}

	challengeBusiness := business.Challenge{
		ChallengeAccess: &access.Challenge{
			DB:             server.Firestore,
			CollectionName: "challenges",
		},
	}

	router.Router.POST("/", func(c echo.Context) error {
		ctx := core.ToContextV2(&c)
		challenge := new(models.Challenge)
		if err = ctx.Bind(challenge); err != nil {
			return echo.ErrBadRequest
		}
		err = challengeBusiness.Create(ctx.GetAuthToken(), challenge)
		if err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, err.Error())
		}
		return c.NoContent(201)
	}, server.Middlewares.Auth.Authorize)

	return router, nil
}
