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

// InfrastructuresHandler represents infra-structures handler
type InfrastructuresHandler struct {
	InfraRepository repository.InfrastructureRepository
}

// CreateInfra creates a new infra-structure
func (h *InfrastructuresHandler) CreateInfra(c *gin.Context) {
	var spec v1.InfrastructureSpec

	if err := c.ShouldBindJSON(&spec); err != nil {
		_, _ = problem.NewStatusProblem(http.StatusUnprocessableEntity).Write(c.Writer)
		return
	}

	if err := validator.New().Struct(spec); err != nil {
		_, _ = problem.NewValidationProblem(err.(validator.ValidationErrors)).Write(c.Writer)
		return
	}

	infraModel, err := model.NewInfrastructure(v1.Infrastructure{Spec: spec})
	if err != nil {
		log.Error().Err(err).Msg("could not create db model")
		_, _ = problem.NewStatusProblem(http.StatusInternalServerError).Write(c.Writer)
		return
	}

	err = h.InfraRepository.Create(c, infraModel)
	if err != nil {
		log.Error().Err(err).Msg("could not save db model")
		_, _ = problem.NewStatusProblem(http.StatusInternalServerError).Write(c.Writer)
		return
	}

	respObject, err := infraModel.MarshalInfrastructureAPI()
	if err != nil {
		log.Error().Err(err).Msg("could not marshal response object")
		_, _ = problem.NewStatusProblem(http.StatusInternalServerError).Write(c.Writer)
		return
	}

	c.JSON(http.StatusOK, respObject)
}
