package main

import (
	"log"
	"net/http"
	"robots/internal/input"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	router.POST("/robots", func(c *gin.Context) {
		input := input.NewHttpGinInput(c)

		result, err := input.ExecuteInstructions()

		if err != nil {
			c.JSON(http.StatusInternalServerError, map[string]interface{}{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, map[string]interface{}{"output": result})
	})

	if err := router.Run(); err != nil {
		log.Fatal("error running web service ", err)
	}
}
