package router

import (
	"github.com/gin-gonic/gin"
	"gitlab.com/F1ukez/sample-go/handler"
)

func New() *gin.Engine {
	return gin.Default()
}

func Mount(r *gin.Engine) {
	r.GET("/", handler.CountInfectedHandler)
}
