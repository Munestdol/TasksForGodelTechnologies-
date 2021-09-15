package handler

import (
	"task5/pkg/service"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	services *service.Service
}

func NewHandler(services *service.Service) *Handler{
	return &Handler{services: services}
}

func (h *Handler) InitRoutes() *gin.Engine  {
	router := gin.New()

	api := router.Group("/api")
	{
		ad := api.Group("/ad")
		{
			ad.POST("/",h.createAd)
			ad.GET("/",h.getAllAds)
			ad.GET("/sortbyprice.asc",h.getAllAdsSortByPriceAsc)
			ad.GET("/sortbyprice.desc",h.getAllAdsSortByPriceDesc)
			ad.GET("/sortbydate.asc",h.getAllAdsSortByDateAsc)
			ad.GET("/sortbydate.desc",h.getAllAdsSortByDateDesc)
			ad.GET("/:id",h.getAdById)
		}
	}

	return router
}
