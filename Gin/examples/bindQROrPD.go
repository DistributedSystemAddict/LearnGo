package main

import (
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type Person1 struct {
	Name     string    `form:"name"`
	Address  string    `form:"address"`
	Birthday time.Time `form:"birthday" time_format:"2006-01-02" time_utc:"1"`
}

func main4() {
	route := gin.Default()
	route.GET("/testing", startPage)
	route.Run(":8085")
}

func startPage(c *gin.Context) {
	var person Person1

	err := c.ShouldBind(&person)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	log.Println(person.Name)
	log.Println(person.Address)
	log.Println(person.Birthday)

	c.String(200, "Success")
}
