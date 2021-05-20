package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/onepaas/onepaas/internal/app/onepaas/repository"
	"github.com/onepaas/onepaas/internal/app/onepaas/types"
	"github.com/onepaas/onepaas/internal/pkg/database"
)

type UsersController struct{}

func (u *UsersController) Add(c *gin.Context) {
	var input types.CreateUserRequest
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	r := repository.NewUserRepository(database.GetDB())
	if _, err := r.Create(input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"status": "you are logged in"})
}
