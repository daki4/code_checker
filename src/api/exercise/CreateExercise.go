package exercise

import (
	"net/http"
	"time"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"yordanmitev.me/code-checker/db"
	exerpack "yordanmitev.me/code-checker/exercise"
)

func CreateExercise(c echo.Context) error {
	db := db.GetDb()
	var exercise exerpack.Exercise
	contentType := c.Request().Header.Get(echo.HeaderContentType)
	if contentType != "application/json" {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Requested data is of wrong format"})
	}

	if err := c.Bind(&exercise); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid request body"})
	}
	exercise.Id = uuid.New()
	exercise.Author = c.Get("username").(string)
	exercise.CreationTime = time.Now()
	exercise.AllowedExecutionDuration = int64(time.Hour.Minutes() * 2)
	exercise.AllowedMemory = 256 * 1024 * 1024
	exercise.AllowedLanguages = []exerpack.Language{
		{
			Name:    "go",
			Version: "1.21",
		},
		{
			Name:    "c",
			Version: "C11",
		}, {
			Name:    "cpp",
			Version: "CPP20",
		}, {
			Name:    "rust",
			Version: "1.75",
		}, {
			Name:    "haskell",
			Version: "8.8.4",
		}, {
			Name:    "python",
			Version: "3.11",
		}}

	db.Create(&exercise)

	return c.JSON(http.StatusCreated, exercise)
}
