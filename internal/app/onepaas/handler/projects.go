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

// ProjectsHandler represents projects handler
type ProjectsHandler struct {
	ProjectRepository repository.ProjectRepository
}

// CreateProject adds a new project.
// swagger:operation POST /v1/projects project createProject
//
// # Create a new Project
//
// ---
// produces:
// - application/json
// - application/problem+json
//
// consumes:
// - application/json
//
// parameters:
//   - description: Add project
//     in: body
//     name: project
//     required: true
//     schema:
//     $ref: "#/definitions/v1-Project"
//
// responses:
//
//	 '201':
//	   description: "Created"
//	   schema:
//		    "$ref": "#/definitions/v1-Project"
//	 '422':
//	   description: "Unprocessable Entity"
//	   schema:
//		    "$ref": "#/definitions/v1-Problem"
//	 '500':
//	   description: "Internal Server Error"
//	   schema:
//		    "$ref": "#/definitions/v1-Problem"
func (p *ProjectsHandler) CreateProject(c *gin.Context) {
	var spec v1.ProjectSpec

	if err := c.ShouldBindJSON(&spec); err != nil {
		_, _ = problem.NewStatusProblem(http.StatusUnprocessableEntity).Write(c.Writer)
		return
	}

	if err := validator.New().Struct(spec); err != nil {
		_, _ = problem.NewValidationProblem(err.(validator.ValidationErrors)).Write(c.Writer)
		return
	}

	newModel := model.NewProject(v1.Project{Spec: spec})

	err := p.ProjectRepository.Create(c, newModel)
	if err != nil {
		log.Error().
			Err(err).
			Send()

		_, _ = problem.NewStatusProblem(http.StatusInternalServerError).Write(c.Writer)

		return
	}

	c.JSON(http.StatusOK, newModel.MarshalProjectAPI())
}

// GetProject read the specified project.
// swagger:operation GET /v1/projects/{id} project readProject
//
// # Read the specified project
//
// ---
// produces:
// - application/json
// - application/problem+json
//
// parameters:
//   - name: id
//     in: path
//     description: project id
//     type: string
//     required: true
//
// responses:
//
//	 '200':
//	   description: "OK"
//	   schema:
//		    "$ref": "#/definitions/v1-Project"
//	 '404':
//	   description: "Not Found"
//	   schema:
//		    "$ref": "#/definitions/v1-Problem"
//	 '500':
//	   description: "Internal Server Error"
//	   schema:
//		    "$ref": "#/definitions/v1-Problem"
func (p *ProjectsHandler) GetProject(c *gin.Context) {
	project, err := p.ProjectRepository.FindByID(c, c.Param("id"))
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			problem.NewStatusProblem(http.StatusNotFound).
				Write(c.Writer)

			return
		}

		log.Error().
			Err(err).
			Send()

		problem.NewStatusProblem(http.StatusInternalServerError).
			Write(c.Writer)

		return
	}

	c.JSON(http.StatusOK, project.MarshalProjectAPI())
}

// ReplaceProject replace a project.
// swagger:operation PUT /v1/projects/{id} project replaceProject
//
// # Replace the specified project
//
// ---
// produces:
// - application/json
// - application/problem+json
//
// consumes:
// - application/json
//
// parameters:
//   - name: id
//     in: path
//     description: project id
//     type: string
//     required: true
//   - name: project
//     in: body
//     required: true
//     schema:
//     "$ref": "#/definitions/v1-Project"
//
// responses:
//
//	 '200':
//	   description: "OK"
//	   schema:
//		    "$ref": "#/definitions/v1-Project"
//	 '404':
//	   description: "Not Found"
//	   schema:
//		    "$ref": "#/definitions/v1-Problem"
//	 '422':
//	   description: "Unprocessable Entity"
//	   schema:
//		    "$ref": "#/definitions/v1-Problem"
//	 '500':
//	   description: "Internal Server Error"
//	   schema:
//		    "$ref": "#/definitions/v1-Problem"
func (p *ProjectsHandler) ReplaceProject(c *gin.Context) {
	var projectAPI v1.Project

	if err := c.ShouldBindJSON(&projectAPI); err != nil {
		problem.NewValidationProblem(err.(validator.ValidationErrors)).Write(c.Writer)
		return
	}

	project, err := p.ProjectRepository.FindByID(c, c.Param("id"))
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			_, _ = problem.NewStatusProblem(http.StatusNotFound).Write(c.Writer)

			return
		}

		log.Error().
			Err(err).
			Send()

		_, _ = problem.NewStatusProblem(http.StatusInternalServerError).Write(c.Writer)

		return
	}

	modelData := model.Project{}
	modelData.UnmarshalProjectAPI(projectAPI)

	err = p.ProjectRepository.Update(c, &project, modelData)
	if err != nil {
		if err, ok := err.(model.ErrorChangedColumn); ok {
			problem.NewStatusProblem(http.StatusUnprocessableEntity, problem.WithDetail(err.Error())).
				Write(c.Writer)

			return
		}

		log.Error().
			Err(err).
			Send()

		problem.NewStatusProblem(http.StatusInternalServerError).
			Write(c.Writer)

		return
	}

	c.JSON(http.StatusOK, project.MarshalProjectAPI())
}

// ListProjects list of projects.
// swagger:operation GET /v1/projects project listProject
//
// # List of projects
//
// ---
// produces:
// - application/json
// - application/problem+json
//
// responses:
//
//	 '200':
//	   description: "OK"
//	   schema:
//		    "$ref": "#/definitions/v1-ProjectList"
//	 '500':
//	   description: "Internal Server Error"
//	   schema:
//		    "$ref": "#/definitions/v1-Problem"
func (p *ProjectsHandler) ListProjects(c *gin.Context) {
	projectEntities, err := p.ProjectRepository.FindAll(c)
	if err != nil {
		log.Error().Err(err).Send()
		_, _ = problem.NewStatusProblem(http.StatusInternalServerError).Write(c.Writer)
		return
	}

	projects := make([]v1.Project, 0)
	for _, projectEntity := range projectEntities {
		projects = append(projects, projectEntity.MarshalProjectAPI())
	}

	c.JSON(http.StatusOK, v1.ProjectList{Items: projects})
}
