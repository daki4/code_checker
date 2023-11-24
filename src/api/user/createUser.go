package user

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"yordanmitev.me/code-checker/db"
	userpack "yordanmitev.me/code-checker/user"
)

func CreateUser(c echo.Context) error {
	db := db.GetDb()
	var user userpack.User
	var checkUserExistsBuf []userpack.User
	var checkUserExists userpack.User

	contentType := c.Request().Header.Get(echo.HeaderContentType)
	if contentType != "application/json" {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Requested data is of wrong format"})
	}

	if err := c.Bind(&user); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid request body"})
	}

	db.Find(&checkUserExistsBuf, "username = ?", user.Username).First(&checkUserExists)

	if checkUserExists.Username != "" {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "username already in use"})
	}

	db.Find(&checkUserExistsBuf, "email = ?", user.Email).First(&checkUserExists)

	if checkUserExists.Email != "" {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "email already in use"})
	}

	hashedPassword, err := hashPassword(user.Password)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to hash password"})
	}
	user.Password = hashedPassword

	db.Create(&user)

	return c.JSON(http.StatusCreated, user)
}
