package main

import (
	"encoding/json"
	"net/http"
	"os"

	authapi "yordanmitev.me/code-checker/api/auth"
	"yordanmitev.me/code-checker/api/piston"
	solve "yordanmitev.me/code-checker/api/solve"
	auth "yordanmitev.me/code-checker/auth"

	scenarios "yordanmitev.me/code-checker/auth/scenarios"

	exerciseapi "yordanmitev.me/code-checker/api/exercise"
	userapi "yordanmitev.me/code-checker/api/user"
	db "yordanmitev.me/code-checker/db"
	exercise "yordanmitev.me/code-checker/exercise"
	user "yordanmitev.me/code-checker/user"

	"github.com/joho/godotenv"
	echoSwagger "github.com/swaggo/echo-swagger"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {

	// init
	godotenv.Load(".env")
	piston.Start()
	auth.Init()
	// jwtConfig := auth.GetJwtConfig()
	err := db.Connect()
	if err != nil {
		panic(err)
	}
	// run migrations
	db := db.GetDb()
	db.AutoMigrate(&user.User{})
	db.AutoMigrate(&exercise.ControlTest{})
	db.AutoMigrate(&exercise.Exercise{})
	db.AutoMigrate(&exercise.Submission{})
	db.AutoMigrate(&exercise.Solution{})
	db.AutoMigrate(&exercise.TestOutcome{})

	e := echo.New()
	e.Logger.SetLevel(1)
	e.Use(middleware.RemoveTrailingSlash())
	api := e.Group("/api/v1", middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "method=${method}, uri=${uri}, status=${status}\n",
	}))

	// echoJwtConfig := echojwt.Config{
	// 	NewClaimsFunc: func(c echo.Context) jwt.Claims {
	// 		return new(auth.Claim)
	// 	},
	// 	SigningKey: []byte(jwtConfig.SecretKey),
	// }

	// api.Use(echojwt.WithConfig(echoJwtConfig))

	e.GET("/", hello)

	api.GET("/swagger/*", echoSwagger.WrapHandler)

	// users
	users := api.Group("/users")
	users.POST("", userapi.CreateUser)
	users.GET("", userapi.GetUsers)

	loggedInUsers := users.Group("/:username", scenarios.IsSelf)
	loggedInUsers.GET("", userapi.GetUser)
	loggedInUsers.PUT("", userapi.UpdateUser)
	loggedInUsers.DELETE("", userapi.DeleteUser)

	// submissions
	submissions := loggedInUsers.Group("/exercises/:exercise", scenarios.IsLoggedIn, scenarios.IsSelf)
	submissions.GET("/submissions", solve.GetSubmissions)
	submissions.GET("/:solution", solve.GetSolution)
	submissions.GET("/solutions", solve.GetSolutions)

	// exercises
	exercises := api.Group("/exercises", scenarios.IsLoggedIn)
	exercises.GET("", exerciseapi.GetExercises)
	exercises.POST("", exerciseapi.CreateExercise)
	exercises.GET("/:id", exerciseapi.GetExercise)
	exercises.POST("/:id/solve", solve.SubmitSolution)

	authorExercises := exercises.Group("", scenarios.IsExerciseOwner)
	authorExercises.DELETE("/:id", exerciseapi.RemoveExercise)
	authorExercises.PUT("/:i/", exerciseapi.UpdateExercise)

	// auth
	api.POST("/login", authapi.Login)
	api.GET("/login/blacklist", authapi.GetBlacklistedTokens)
	api.POST("/login/blacklist", authapi.BlacklistToken)

	e.Debug = true
	data, _ := json.MarshalIndent(e.Routes(), "", "  ")
	os.WriteFile("routes.json", data, 0644)

	e.Logger.Fatal(e.Start(":31415"))
}

func hello(c echo.Context) error {
	return c.HTML(http.StatusOK, "<h1>Hello, World from code-checker!</h1>")
}
