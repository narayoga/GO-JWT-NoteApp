package main

import (
	"go-jwt/controllers"
	"go-jwt/initializer"
	"go-jwt/middleware"
	"os"

	"github.com/gin-gonic/gin"
)

func init() {
	initializer.LoadVariableEnv()
	initializer.ConnectToDb()
	initializer.AutoMigrationDb()
}

func main() {

	router := gin.Default()

	router.POST("/user", controllers.SignUp)
	router.POST("/login", controllers.Login)
	router.GET("/", middleware.Authorization, controllers.UserProfile)
	router.POST("/profile", controllers.CreateProfile)
	router.GET("/profile/:id", controllers.GetUserProfileById)
	router.POST("/note", controllers.CreateNote)
	router.GET("/note", controllers.GetNotes)

	router.Run(os.Getenv("PORT"))

}
