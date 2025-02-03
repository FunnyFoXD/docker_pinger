package routes

import (
	"github.com/gin-gonic/gin"

	"github.com/funnyfoxd/docker_pinger/handlers"
)

func SetupRouter(r *gin.Engine) {
	r.GET("/containers", handlers.GetContainers)
	r.POST("/containers", handlers.CreatePingContainer)
}
