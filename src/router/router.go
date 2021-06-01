package router

import (
	"github.com/gin-gonic/gin"
	"main/src/cli/searching"
	"net/http"
)

func InitRouter() *gin.Engine {
	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	r.GET("/ping", func(c *gin.Context) {
		c.Header("Access-Control-Allow-Origin", "*")
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.GET("/search", func(c *gin.Context) {

		result := searching.SearchByKeyWord(c)
		c.Header("Access-Control-Allow-Origin", "*")
		c.JSON(http.StatusOK, gin.H{"data": result})
	})
	r.POST("/_index", func(c *gin.Context) {

		searching.IndexDocument(c)
		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Methods", "GET,POST")
		c.JSON(http.StatusOK, gin.H{"data": "ok"})
	})

	r.GET("/_delete", func(c *gin.Context) {

		searching.DeleteIndex(c)
		c.Header("Access-Control-Allow-Origin", "*")
		c.JSON(http.StatusOK, gin.H{"data": "ok"})
	})

	return r
}
