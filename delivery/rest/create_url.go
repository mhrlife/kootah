package rest

import (
	"fmt"
	"github.com/labstack/echo/v4"
)

type requestCreateURL struct {
	URL string `json:"url" validate:"required,url"`
}

func (r *Rest) CreateURL(c echo.Context) error {
	var req requestCreateURL
	if err := c.Bind(&req); err != nil {
		return r.handleError(400, c, err)
	}
	if err := r.validator.Struct(req); err != nil {
		return r.handleError(400, c, err)
	}

	uid, err := r.app.CreateURL(c.Request().Context(), req.URL)
	if err != nil {
		return r.handleError(500, c, err)
	}
	return r.handleResponse(201, c, fmt.Sprintf("http://localhost:1323/u/%s", uid))
}
