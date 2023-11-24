package exercise

import (
	"encoding/json"
	"net/http"

	"github.com/labstack/echo/v4"
	"yordanmitev.me/code-checker/db"
	exerpack "yordanmitev.me/code-checker/exercise"
)

func UpdateExercise(c echo.Context) error {

	db := db.GetDb()
	var exercise exerpack.Exercise

	newData := make(map[string]interface{})
	err := json.NewDecoder(c.Request().Body).Decode(&newData)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "empty body"})
	}

	id := c.Param("id")
	if id == "" {
		return c.JSON(http.StatusNotFound, map[string]string{"error": "no exercise selected"})
	}
	// check if exercise exists, if exercise exists, update exercise
	if db.Find(&exercise, "id = ?", id).RowsAffected == 0 {
		return c.JSON(http.StatusNotFound, map[string]string{"error": "exercise not found"})
	}

	newData["id"] = id

	db.Model(&exercise).Where("id = ?", id).Updates(&newData)
	db.Find(&exercise, "id = ?", &exercise.Id)
	return c.JSON(http.StatusNotFound, exercise)
}
