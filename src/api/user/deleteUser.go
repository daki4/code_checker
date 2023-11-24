package user

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm/clause"
	"yordanmitev.me/code-checker/db"
	userpack "yordanmitev.me/code-checker/user"
)

func DeleteUser(c echo.Context) error {
	db := db.GetDb()
	id := c.Param("username")
	var user userpack.User

	if err := c.Bind(c.Request()); err != nil || id == "" {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid or missing User ID"})
	}

	db.Clauses(clause.Returning{}).Where("username = ?", id).Delete(&user)

	return c.JSON(http.StatusOK, user)
}
