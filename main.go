package main

import (
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func main() {

	router := gin.Default()

	router.GET("/", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"msg": "Welcome to sample Go Application :)",
		})
	})

	router.GET("/data", func(ctx *gin.Context) {
		ip := ctx.RemoteIP()
		time.Sleep(time.Second / 2)
		ctx.JSON(http.StatusOK, gin.H{
			"status": 200,
			"msg":    "Your Public IP is: " + ip,
		})
	})

	router.GET("/health", func(ctx *gin.Context) {

		ctx.JSON(http.StatusOK, gin.H{
			"msg": "Everything looks good here :)",
		})
	})

	if err := router.Run(":8003"); err != nil {
		log.Println("Error starting the server on port 8003 :(", err.Error())
	}
}
