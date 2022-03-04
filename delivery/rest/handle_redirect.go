package rest

import (
	"github.com/labstack/echo/v4"
	"kootah/domain"
)

func (r *Rest) HandleRedirect(c echo.Context) error {
	uid := c.Param("uid")
	url, err := r.app.GetRedirectTo(c.Request().Context(), domain.UID(uid))
	if err != nil {
		return r.handleError(404, c, err)
	}
	return c.Redirect(302, url)
}
