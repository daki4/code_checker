package exercise

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm/clause"
	"yordanmitev.me/code-checker/db"
	exerpack "yordanmitev.me/code-checker/exercise"
)

func RemoveExercise(c echo.Context) error {
	db := db.GetDb()
	id := c.Param("id")
	var exercise exerpack.Exercise

	if err := c.Bind(c.Request()); err != nil || id == "" {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid or missing exercise ID"})
	}

	db.Clauses(clause.Returning{}).Where("id = ?", id).Delete(&exercise)

	return c.JSON(http.StatusOK, exercise)
}
