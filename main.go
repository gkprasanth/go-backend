package main

import (
	"backend/config"
	"backend/controllers"
	"backend/database"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func init() {
	fmt.Println("Initialzing...")
	config.LoadEnv()
	database.InitDB()
	// database.SyncDatabase()

}

func main() {

	router := gin.Default()

	router.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "listening",
		})
	})

	router.POST("/register", controllers.Register)
	router.GET("/users", controllers.GetUsers)

	router.Run()
}
