package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Person struct {
	Name      string    `form:"name,default=William"`
	Age       int       `form:"age,default=10"`
	Friends   []string  `form:"friends,default=Will;Bill"`
	Addresses [2]string `form:"addresses,default=foo bar" collection_format:"ssv"`
	LapTimes  []int     `form:"lap_times,default=1;2;3" collection_format:"csv"`
}

func main2() {
	r := gin.Default()
	r.POST("/person", func(c *gin.Context) {
		var req Person
		if err := c.ShouldBind(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, req)
	})
	_ = r.Run(":8080")
}
