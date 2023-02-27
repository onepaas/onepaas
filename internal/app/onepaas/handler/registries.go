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

// RegistriesHandler represents registries handler
type RegistriesHandler struct {
	RegistryRepository repository.RegistryRepository
}

// CreateRegistry adds a new registry
func (h *RegistriesHandler) CreateRegistry(c *gin.Context) {
	var regSpec v1.RegistrySpec

	if err := c.ShouldBindJSON(&regSpec); err != nil {
		_, _ = problem.NewStatusProblem(http.StatusUnprocessableEntity).Write(c.Writer)
		return
	}

	if err := validator.New().Struct(regSpec); err != nil {
		_, _ = problem.NewValidationProblem(err.(validator.ValidationErrors)).Write(c.Writer)
		return
	}

	registryModel := model.NewRegistry(v1.Registry{Spec: regSpec})

	err := h.RegistryRepository.Create(c, registryModel)
	if err != nil {
		log.Error().
			Err(err).
			Send()

		_, _ = problem.NewStatusProblem(http.StatusInternalServerError).Write(c.Writer)

		return
	}

	c.JSON(http.StatusOK, registryModel.MarshalRegistryAPI())
}
