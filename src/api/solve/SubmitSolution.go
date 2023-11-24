package solve

import (
	"net/http"
	"time"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"yordanmitev.me/code-checker/db"
	exerpack "yordanmitev.me/code-checker/exercise"
)

func SubmitSolution(c echo.Context) error {
	db := db.GetDb()
	var submission exerpack.Submission
	contentType := c.Request().Header.Get(echo.HeaderContentType)
	if contentType != "application/json" {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Requested data is of wrong format"})
	}

	if err := c.Bind(&submission); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid request body"})
	}

	submission.SubmissionTime = time.Now()
	submission.Id = uuid.New()
	var exercise exerpack.Exercise
	var exercises []exerpack.Exercise
	db.Find(&exercises, "id = ?", c.Param("id")).Preload("AllowedLanguages").Preload("Tests").First(&exercise)
	if exercise.Id == uuid.Nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Exercise not found"})
	}
	submission.Exercise = exercise
	db.Omit("Exercise").Create(&submission)

	db.Omit("Exercise").Create(submission.RunTests())

	return c.JSON(http.StatusCreated, submission)
}
