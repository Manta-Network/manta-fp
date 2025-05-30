package router

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Registry struct {
}

func NewRegistry() *Registry {
	return &Registry{}
}

func (registry *Registry) Register(r *gin.Engine) {
	r.GET("/ping", func(c *gin.Context) {
		c.Status(http.StatusOK)
	})
}
