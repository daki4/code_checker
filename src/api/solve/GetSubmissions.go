package solve

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"yordanmitev.me/code-checker/db"
	exerpack "yordanmitev.me/code-checker/exercise"
)

func GetSubmissions(c echo.Context) error {
	db := db.GetDb()
	var submissions []exerpack.Submission

	if db.Where(db.Where("username = ?", c.Get("username")).Where("id = ?", c.Param("exercise"))).Find(&submissions).RowsAffected > 0 {
		return c.JSON(http.StatusOK, submissions)
	}
	return c.JSON(http.StatusNotFound, "not found")
}
