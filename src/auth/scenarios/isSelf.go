package scenarios

import (
	"net/http"
	"strings"

	"github.com/labstack/echo/v4"
	"yordanmitev.me/code-checker/auth"
)

func IsSelf(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		fac := auth.GetJwtConfig()
		header := c.Request().Header.Get("Authorization")
		username := c.Param("username")

		content := strings.Split(header, " ")
		if content[0] != "Bearer" {
			return echo.NewHTTPError(http.StatusBadRequest, map[string]string{"error": "Invalid authorization header"})
		}
		if !fac.ValidateToken(content[1]) {
			return echo.NewHTTPError(http.StatusBadRequest, map[string]string{"error": "Invalid token"})
		}
		claims, _ := fac.GetClaim(content[1])
		if claims.Username != username {
			return echo.NewHTTPError(http.StatusForbidden, map[string]string{"error": "You are not authorized to access this resource"})
		}
		c.Set("username", username)
		return next(c)
	}
}
