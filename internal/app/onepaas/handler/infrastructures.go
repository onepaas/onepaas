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

// InfrastructuresHandler represents infra-structures handler
type InfrastructuresHandler struct {
	InfraRepository repository.InfrastructureRepository
}

// Create creates a new infra-structure
func (h *InfrastructuresHandler) Create(c *gin.Context) {
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

// List returns list of all infras
func (h *InfrastructuresHandler) List(c *gin.Context) {
	list, err := h.InfraRepository.FindAll(c)
	if err != nil {
		log.Error().Err(err).Send()
		_, _ = problem.NewStatusProblem(http.StatusInternalServerError).Write(c.Writer)
		return
	}

	results := make([]v1.Infrastructure, 0)
	for _, reg := range list {
		record, err := reg.MarshalInfrastructureAPI()
		if err != nil {
			log.Error().Err(err).Str("id", reg.Id).Msg("invalid infrastructure found")
			continue
		}

		results = append(results, record)
	}

	c.JSON(http.StatusOK, v1.InfrastructureList{Items: results})
}

// Get returns details of one infra
func (h *InfrastructuresHandler) Get(c *gin.Context) {
	record, err := h.InfraRepository.FindByID(c, c.Param("id"))
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			_, _ = problem.NewStatusProblem(http.StatusNotFound).Write(c.Writer)
			return
		}

		log.Error().Err(err).Send()
		_, _ = problem.NewStatusProblem(http.StatusInternalServerError).Write(c.Writer)

		return
	}

	result, err := record.MarshalInfrastructureAPI()
	if err != nil {
		_, _ = problem.NewStatusProblem(http.StatusUnprocessableEntity).Write(c.Writer)
		return
	}

	c.JSON(http.StatusOK, result)
}
