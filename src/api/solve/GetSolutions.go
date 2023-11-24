package solve

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"yordanmitev.me/code-checker/db"
	exerpack "yordanmitev.me/code-checker/exercise"
)

func GetSolutions(c echo.Context) error {
	db := db.GetDb()
	var solutions []exerpack.Solution

	if db.Where(db.Where("user = ?", c.Get("username")).Where("exerciseid = ?", c.Param("exercise"))).Find(&solutions).RowsAffected > 0 {
		return c.JSON(http.StatusOK, solutions)
	}
	return c.JSON(http.StatusNotFound, "not found")
}
