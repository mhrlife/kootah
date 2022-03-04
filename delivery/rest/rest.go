package rest

import (
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
	"kootah/definitions"
)

type respJson map[string]interface{}

type Rest struct {
	app       definitions.Kootah
	echo      *echo.Echo
	validator *validator.Validate
}

func NewRest(app definitions.Kootah) *Rest {
	e := echo.New()
	r := &Rest{
		app:       app,
		echo:      e,
		validator: validator.New(),
	}
	r.registerRoutes()
	return r
}
func (r *Rest) registerRoutes() {
	r.echo.POST("/url", r.CreateURL)
	r.echo.GET("/u/:uid", r.HandleRedirect)
}

func (r *Rest) Listen() {
	r.echo.Logger.Fatal(r.echo.Start(":1323"))
}

func (r *Rest) handleError(code int, c echo.Context, err error) error {
	logrus.WithError(err).Error("error in http request")
	return c.JSON(code, respJson{
		"ok":    false,
		"error": err.Error(),
	})
}

func (r *Rest) handleResponse(code int, c echo.Context, data interface{}) error {
	return c.JSON(code, respJson{
		"ok":   true,
		"data": data,
	})
}
