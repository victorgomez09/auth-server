package internal

import (
	"github.com/labstack/echo/v4"
)

func ValidateToken(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		_, err := Oauth2Server.ValidationBearerToken(c.Request())
		if err != nil {
			c.Error(err)

			return err
		}

		return next(c)
	}
}
