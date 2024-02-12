package helpers

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func StartServer() *gin.Engine {
	router := gin.Default()
	router.GET("/student", ActionStudent)
	fmt.Println("live server...")
	return router
}
