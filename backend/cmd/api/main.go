package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type formA struct {
	Foo string `json:"foo" xml:"foo" binding:"required"`
}

type formB struct {
	Bar string `json:"bar" xml:"bar" binding:"required"`
}

func SomeHandler(c *gin.Context) {
	objA := formA{}
	objB := formB{}
	// この c.ShouldBind メソッドは c.Request.Body を消費し、再利用できなくします。
	if errA := c.ShouldBind(&objA); errA == nil {
		c.String(http.StatusOK, `the body should be formA`)
		// c.Request.Body が EOF なので、常にエラーとなります。
	} else if errB := c.ShouldBind(&objB); errB == nil {
		c.String(http.StatusOK, `the body should be formB`)
	}
}

func main() {
	router := gin.Default()
	router.GET("/favicon.ico", func(c *gin.Context) {
		c.Status(204)
	})
	router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	router.GET("/someJSON", func(c *gin.Context) {
		data := map[string]interface{}{
			"lang": "GO语言",
			"tag":  "<br>",
		}

		c.AsciiJSON(http.StatusOK, data)
	})

	if err := router.Run(":8080"); err != nil {
		panic(err)
	}
}
