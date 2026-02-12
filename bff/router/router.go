package router

import (
	"jiaqizuoye2/bff/handler/service"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func Router() *gin.Engine {
	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		time.Sleep(5 * time.Second)
		c.String(http.StatusOK, "Welcome Gin Server")
	})
	r.POST("Login", service.Login)
	r.POST("Upload", service.Upload)
	r.POST("SpAdd", service.SpAdd)
	return r
}
