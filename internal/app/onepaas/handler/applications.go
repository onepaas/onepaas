package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/onepaas/onepaas/internal/app/onepaas/model"
	"github.com/onepaas/onepaas/internal/app/onepaas/repository"
	v1 "github.com/onepaas/onepaas/pkg/api/v1"
	"github.com/onepaas/onepaas/pkg/problem"
	"github.com/rs/zerolog/log"
	"net/http"
)

// ApplicationsHandler represents applications handler
type ApplicationsHandler struct {
	ApplicationRepository repository.ApplicationRepository
}

// CreateApplication adds a new application.
func (a *ApplicationsHandler) CreateApplication(c *gin.Context) {
	var applicationAPI v1.Application

	if err := c.ShouldBindJSON(&applicationAPI); err != nil {
		_, _ = problem.NewValidationProblem(err.(validator.ValidationErrors)).Write(c.Writer)
		return
	}

	applicationModel := model.NewApplication(applicationAPI)

	err := a.ApplicationRepository.Create(c, applicationModel)
	if err != nil {
		log.Error().
			Err(err).
			Send()

		_, _ = problem.NewStatusProblem(http.StatusInternalServerError).Write(c.Writer)

		return
	}

	c.JSON(http.StatusOK, applicationModel.MarshalApplicationAPI())
}
