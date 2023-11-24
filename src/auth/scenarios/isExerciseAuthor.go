package scenarios

import (
	"net/http"
	"strings"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"yordanmitev.me/code-checker/auth"
	"yordanmitev.me/code-checker/db"
	"yordanmitev.me/code-checker/exercise"
)

func IsExerciseOwner(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		fac := auth.GetJwtConfig()
		header := c.Request().Header.Get("Authorization")
		exerciseId, _ := uuid.FromBytes([]byte(c.Param("exercise")))

		content := strings.Split(header, " ")
		if content[0] != "Bearer" {
			return echo.NewHTTPError(http.StatusBadRequest, map[string]string{"error": "Invalid authorization header"})
		}
		if !fac.ValidateToken(content[1]) {
			return echo.NewHTTPError(http.StatusBadRequest, map[string]string{"error": "Invalid token"})
		}
		claims, _ := fac.GetClaim(content[1])
		db := db.GetDb()

		var exercise exercise.Exercise

		// find the exercise by id
		if db.Find(&exercise, "id = ?", exerciseId).RowsAffected == 0 {
			return echo.NewHTTPError(http.StatusBadRequest, map[string]string{"error": "Exercise not found"})
		}
		db.Where(db.Where("author = ?", claims.Username).Where("id = ?", exerciseId)).Find(&exercise)
		if exercise.Id != exerciseId {
			return echo.NewHTTPError(http.StatusForbidden, map[string]string{"error": "You are not authorized to access this resource"})
		}
		c.Set("author", claims.Username)
		c.Set("exercise", exerciseId)
		return next(c)
	}
}
