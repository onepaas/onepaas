package handler

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/onepaas/onepaas/internal/app/onepaas/model"
	"github.com/onepaas/onepaas/internal/app/onepaas/repository"
	v1 "github.com/onepaas/onepaas/pkg/api/v1"
	"github.com/onepaas/onepaas/pkg/problem"
	"github.com/rs/zerolog/log"
	"gorm.io/gorm"
	"net/http"
)

// ApplicationsHandler represents applications handler
type ApplicationsHandler struct {
	ApplicationRepository repository.ApplicationRepository
}

// CreateApplication adds a new application.
func (a *ApplicationsHandler) CreateApplication(c *gin.Context) {
	var appSpec v1.ApplicationSpec

	if err := c.ShouldBindJSON(&appSpec); err != nil {
		_, _ = problem.NewStatusProblem(http.StatusUnprocessableEntity).Write(c.Writer)
		return
	}

	if err := validator.New().Struct(appSpec); err != nil {
		_, _ = problem.NewValidationProblem(err.(validator.ValidationErrors)).Write(c.Writer)
		return
	}

	applicationModel := model.NewApplication(v1.Application{Spec: appSpec})

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

// ListApplications returns list of all applications
func (a *ApplicationsHandler) ListApplications(c *gin.Context) {
	list, err := a.ApplicationRepository.FindAll(c)
	if err != nil {
		log.Error().Err(err).Send()
		_, _ = problem.NewStatusProblem(http.StatusInternalServerError).Write(c.Writer)
		return
	}

	apps := make([]v1.Application, 0)
	for _, app := range list {
		apps = append(apps, app.MarshalApplicationAPI())
	}

	c.JSON(http.StatusOK, v1.ApplicationList{Items: apps})
}

// GetApplication returns details of one application
func (a *ApplicationsHandler) GetApplication(c *gin.Context) {
	record, err := a.ApplicationRepository.FindByID(c, c.Param("id"))
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			_, _ = problem.NewStatusProblem(http.StatusNotFound).Write(c.Writer)
			return
		}

		log.Error().Err(err).Send()
		_, _ = problem.NewStatusProblem(http.StatusInternalServerError).Write(c.Writer)

		return
	}

	c.JSON(http.StatusOK, record.MarshalApplicationAPI())
}
