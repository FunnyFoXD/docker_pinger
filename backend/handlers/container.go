package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/funnyfoxd/docker_pinger/db"
	"github.com/funnyfoxd/docker_pinger/models"
)

func GetContainers(ctx *gin.Context) {
	var containers []models.ContainerPing

	if err := db.DB.Find(&containers).Error; err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, containers)
}

func CreatePingContainer(ctx *gin.Context) {
	var container models.ContainerPing

	if err := ctx.ShouldBindJSON(&container); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := db.DB.Create(&container).Error; err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, container)
}
