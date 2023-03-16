package main

import (
	"github.com/gin-gonic/gin"
	"io"
	"net/http"
	"os"
	"time"
)

func getTimezone(c *gin.Context) {
	zoneName := c.DefaultQuery("zone", "UTC")
	zone, err := time.LoadLocation(zoneName)

	if err != nil {
		c.AbortWithError(http.StatusBadRequest, err).
		return
	}

	now := time.Now().In(zone).Format(time.RFC3339)

	c.IndentedJSON(http.StatusOK, now)
}

func setupRouter() *gin.Engine {
	router := gin.Default()

	router.GET("/", getTimezone)

	return router
}

func main() {
	f, _ := os.Create("/home/jony/Learning/Go/go-timezone/gin.log")
	gin.DefaultWriter = io.MultiWriter(f, os.Stdout)

	router := setupRouter()
	router.Run("localhost:8080")
}
