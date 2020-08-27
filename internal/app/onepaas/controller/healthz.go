package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type HealthzController struct{}

func (h *HealthzController) Healthz(c *gin.Context) {
	c.Status(http.StatusNoContent)

	return
}
