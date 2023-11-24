package user

import (
	"encoding/json"
	"net/http"

	"github.com/labstack/echo/v4"
	"yordanmitev.me/code-checker/db"
	userpack "yordanmitev.me/code-checker/user"
)

func UpdateUser(c echo.Context) error {

	db := db.GetDb()
	var user userpack.User

	newData := make(map[string]interface{})
	err := json.NewDecoder(c.Request().Body).Decode(&newData)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "empty body"})
	}

	username := c.Param("username")
	if username == "" {
		return c.JSON(http.StatusNotFound, map[string]string{"error": "no user selected"})
	}
	// check if user exists, if user exists, update user
	if db.Find(&user, "username = ?", username).RowsAffected == 0 {
		return c.JSON(http.StatusNotFound, map[string]string{"error": "user not found"})
	}

	newData["username"] = username

	if newData["password"] != nil {
		hashedPassword, err := hashPassword(newData["password"].(string))
		if err != nil {
			return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to hash password"})
		}
		newData["password"] = hashedPassword
	}

	db.Model(&user).Where("username = ?", username).Updates(&newData)
	db.Find(&user, "username = ?", &user.Username)
	return c.JSON(http.StatusNotFound, user)
}
