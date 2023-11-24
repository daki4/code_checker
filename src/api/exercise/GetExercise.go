package exercise

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"yordanmitev.me/code-checker/db"
	exerpack "yordanmitev.me/code-checker/exercise"
)

func GetExercise(c echo.Context) error {
	db := db.GetDb()
	var exercises []exerpack.Exercise
	var exercise exerpack.Exercise
	if db.Find(&exercises, "id = ?", c.Param("id")).Preload("AllowedLanguages").Preload("Tests").First(&exercise).RowsAffected > 0 {
		return c.JSON(http.StatusOK, exercise)
	}
	return c.JSON(http.StatusNotFound, "exercise not found")
}
