package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {

	fmt.Println("Starting main execution")
	router := gin.Default()
	router.GET("/", getResponse)
	router.Run("localhost:8080")
}
func getResponse(c *gin.Context) {

	//write your logic for the api
	c.IndentedJSON(http.StatusOK, "Hello from go-gin")
}
