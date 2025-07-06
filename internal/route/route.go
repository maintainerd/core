package route

import (
	"github.com/gin-gonic/gin"
	"github.com/maintainerd/core/internal/handler"
)

type HandlerCollection struct {
	OrganizationHandler *handler.OrganizationHandler
	ServiceHandler      *handler.ServiceHandler
}

func RegisterRoute(r *gin.Engine, h *HandlerCollection) {
	api := r.Group("/api/v1")

	RegisterOrganizationRoute(api, h.OrganizationHandler)
	RegisterServiceRoute(api, h.ServiceHandler)
}
