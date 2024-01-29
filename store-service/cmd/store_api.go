package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	//router.GET("/items", getItems)
	//router.POST("/addItem", addItem)
	//router.DELETE("/removeItem", removeItem)

	router.Run("localhost:8081")
}
