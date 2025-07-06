package route

import (
	"github.com/gin-gonic/gin"
	"github.com/maintainerd/core/internal/handler"
	"github.com/maintainerd/core/internal/middleware"
)

func RegisterServiceRoute(router *gin.RouterGroup, serviceHandler *handler.ServiceHandler) {
	protected := router.Group("/services")
	protected.Use(middleware.JWTAuthMiddleware())

	protected.POST("", serviceHandler.Create)
	protected.GET("", serviceHandler.GetAll)
	protected.GET("/:service_uuid", serviceHandler.GetByUUID)
	protected.PUT("/:service_uuid", serviceHandler.Update)
	protected.DELETE("/:service_uuid", serviceHandler.Delete)
}
