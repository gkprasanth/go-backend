package main

import (
	"backend/config"
	"backend/database"
	"fmt"
	"net/http"
	"github.com/gin-gonic/gin"
)

func init() {
	fmt.Println("Initialzing...")
	config.LoadEnv()
	database.InitDB()

}

func main() {

	router := gin.Default()

	router.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "listening",
		})
	})

	router.Run()
}
