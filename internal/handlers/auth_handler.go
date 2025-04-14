package handlers

import (
	"net/http"
	"socialNetworks/internal/models"

	"github.com/gin-gonic/gin"
)

func Register(c *gin.Context) {
	var input models.RegisterInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}

}
