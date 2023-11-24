package exercise

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"yordanmitev.me/code-checker/db"
	exerpack "yordanmitev.me/code-checker/exercise"
)

func GetExercises(c echo.Context) error {
	db := db.GetDb()
	var exercises []exerpack.Exercise
	db.Model(exerpack.Exercise{}).Preload("AllowedLanguages").Preload("Tests").Find(&exercises)
	return c.JSON(http.StatusOK, exercises)
}
