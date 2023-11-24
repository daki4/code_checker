package user

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"yordanmitev.me/code-checker/db"
	userpack "yordanmitev.me/code-checker/user"
)

func GetUser(c echo.Context) error {
	db := db.GetDb()
	var users []userpack.User
	var user userpack.User
	if db.Find(&users, "username = ?", c.Param("username")).First(&user).RowsAffected > 0 {
		return c.JSON(http.StatusOK, user)
	}
	return c.JSON(http.StatusNotFound, "user not found")
}
