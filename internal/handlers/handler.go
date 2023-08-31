package handler

import (
	"avito-third/internal/service"

	"github.com/gin-gonic/gin"
	 "github.com/swaggo/gin-swagger" // gin-swagger middleware
    "github.com/swaggo/files" // swagger embed files
	_"avito-third/docs"


)

type Handler struct {
	services *service.Service
}

func NewHandler(services *service.Service) *Handler {
	return &Handler{services: services}
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	segment := router.Group("/segments")
	{
		segment.POST("/", h.createSegment)
		segment.DELETE("/", h.deleteSegment)
	}
	users := router.Group("/users")
	{
		users.POST("/", h.CRUDUsersInSegment)
		users.POST("/activeSegments", h.GetActiveSlugs)
		users.POST("/reports", h.GetUrlReportFile)
		users.GET("/reportForperiod/*path", h.GetReportFile)
	}

	return router
}
