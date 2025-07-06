package route

import (
	"github.com/gin-gonic/gin"
	"github.com/maintainerd/core/internal/handler"
	"github.com/maintainerd/core/internal/middleware"
)

func RegisterOrganizationRoute(router *gin.RouterGroup, organizationHandler *handler.OrganizationHandler) {
	protected := router.Group("/organizations")
	protected.Use(middleware.JWTAuthMiddleware())

	protected.POST("", organizationHandler.Create)
	protected.GET("", organizationHandler.GetAll)
	protected.GET("/:organization_uuid", organizationHandler.GetByUUID)
	protected.PUT("/:organization_uuid", organizationHandler.Update)
	protected.DELETE("/:organization_uuid", organizationHandler.Delete)
}
