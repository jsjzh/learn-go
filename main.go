package main

import (
	"log"
	"net"

	"github.com/gin-gonic/gin"
)

// _ "learn-go/review"

// _ "learn-go/step_four"
// _ "learn-go/step_one"
// _ "learn-go/step_three"
// _ "learn-go/step_two"

// middleWares
// logger
// controllers
// routes
// services
// models
// utils
// crons
// conf

func main() {
	router := gin.Default()

	routerGroup := router.Group("/v1/api")

	routerGroup.GET("/hello", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"status": "success",
			"data":   routerGroup.BasePath(),
		})
	})

	addrs, error := net.InterfaceAddrs()

	if error != nil {
		log.Println("error", error)
	}

	log.Println("addrs", addrs)

	router.Run("127.0.0.1:7001")
}
