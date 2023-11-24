package user

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"yordanmitev.me/code-checker/db"
	userpack "yordanmitev.me/code-checker/user"
)

func GetUsers(c echo.Context) error {
	db := db.GetDb()
	var users []userpack.User
	db.Find(&users)
	return c.JSON(http.StatusOK, users)
}
