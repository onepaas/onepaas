package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/onepaas/onepaas/internal/app/onepaas/repository"
	v1 "github.com/onepaas/onepaas/pkg/api/v1"
	"github.com/onepaas/onepaas/pkg/problem"
	"net/http"
)

// PipelinesHandler represents pipelines handler
type PipelinesHandler struct {
	ApplicationRepository repository.ApplicationRepository
}

// RunPipelineFromGithub trigger a pipeline for an application from GitHub
func (a *PipelinesHandler) RunPipelineFromGithub(c *gin.Context) {
	var hookHeaderSpec v1.GithubHookHeaderSpec

	if err := c.ShouldBindHeader(&hookHeaderSpec); err != nil {
		_, _ = problem.NewStatusProblem(http.StatusUnprocessableEntity).Write(c.Writer)
		return
	}

	requestValidator := validator.New()

	if err := requestValidator.Struct(hookHeaderSpec); err != nil {
		_, _ = problem.NewValidationProblem(err.(validator.ValidationErrors)).Write(c.Writer)
		return
	}

	var hookSpec v1.GithubHookSpec

	if err := c.ShouldBindJSON(&hookSpec); err != nil {
		_, _ = problem.NewStatusProblem(http.StatusUnprocessableEntity).Write(c.Writer)
		return
	}

	if err := requestValidator.Struct(hookSpec); err != nil {
		_, _ = problem.NewValidationProblem(err.(validator.ValidationErrors)).Write(c.Writer)
		return
	}

	// TODO: Find and fetch application details

	// TODO: Trigger Temporal workflow

	c.String(http.StatusOK, "Pipeline created by GitHub %s (%s) event for %s repository", hookHeaderSpec.Event, hookSpec.ReferenceType, hookSpec.Repository.FullName)
}
