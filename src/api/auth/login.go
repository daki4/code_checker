package auth

import (
	"net/http"
	"strings"

	"github.com/labstack/echo/v4"
	"yordanmitev.me/code-checker/auth"
	"yordanmitev.me/code-checker/db"
	"yordanmitev.me/code-checker/user"
)

type LoginData struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

const WWW_AUTHENTICATE_KINDS_BASIC = "Basic realm=Restricted"
const WWW_AUTHENTICATE_KINDS_JWT = "Bearer realm=Restricted"

func setHeadersOnFail(c echo.Context) {
	c.Response().Header().Set("WWW-Authenticate", WWW_AUTHENTICATE_KINDS_BASIC)
}

func Login(c echo.Context) error {
	// Check if the "Renew" query parameter is set to "true"
	param := c.QueryParam("refresh")
	if param == "true" {
		return renewToken(c)
	}
	return createNewToken(c)
}

func renewToken(c echo.Context) error {
	fac := auth.GetJwtConfig()
	header := c.Request().Header.Get("Authorization")

	// Check if the Authorization header is valid
	if header == "" {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid authorization header"})
	}

	content := strings.Split(header, " ")
	if content[0] != "Bearer" {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid authorization header"})
	}

	if !fac.ValidateToken(content[1]) {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid token"})
	}

	// Refresh the token
	refreshedToken, err := fac.RefreshToken(content[1])
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Internal error"})
	}

	return c.JSON(http.StatusOK, map[string]string{"token": refreshedToken})
}

func GetJwtConfig() {
	panic("unimplemented")
}

func createNewToken(c echo.Context) error {
	fac := auth.GetJwtConfig()
	db := db.GetDb()
	var loginData LoginData

	// Parse the request body
	if err := c.Bind(&loginData); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid request body"})
	}

	// Check for empty username or password
	if loginData.Username == "" || loginData.Password == "" {
		setHeadersOnFail(c)
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Empty username or password"})
	}

	var user user.User
	db.Where("username = ?", loginData.Username).First(&user)

	// Check if the user exists
	if user.Username == "" {
		setHeadersOnFail(c)
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "User not found"})
	}

	// Check if the password is correct
	if !user.CheckPassword(loginData.Password) {
		setHeadersOnFail(c)
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Wrong password"})
	}

	// Create a new token
	claim, err := fac.CreateToken(loginData.Username)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Internal error"})
	}

	return c.JSON(http.StatusOK, claim)
}

func GetBlacklistedTokens(c echo.Context) error {
	fac := auth.GetJwtConfig()
	return c.JSON(http.StatusOK, fac.GetBlacklist())
}

func BlacklistToken(c echo.Context) error {
	fac := auth.GetJwtConfig()
	c.Logger().Info(fac.GetBlacklist())

	header := c.Request().Header.Get("Authorization")

	// Check if the Authorization header is valid
	if header == "" {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid authorization header"})
	}

	content := strings.Split(header, " ")
	if content[0] != "Bearer" {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid authorization header"})
	}
	token := content[1]

	if !fac.ValidateToken(token) {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid token"})
	}

	if !fac.BlacklistToken(token) {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Internal error"})
	}

	// return c.JSON(http.StatusOK, factory.blacklist)
	c.Logger().Info(fac.GetBlacklist())

	return c.JSON(http.StatusOK, map[string]string{"message": "Token blacklisted"})
}
