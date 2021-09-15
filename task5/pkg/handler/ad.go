package handler

import (
	"task5"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"net/http"
	"strconv"
)

type GetAllAds struct {
	Data []task5.Ad `json:"data"`
}

func (h *Handler) createAd(c *gin.Context){
	var input task5.Ad

	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	validate := validator.New()

	if err := validate.Struct(&input); err != nil{
		return
	}

	id, err := h.services.Ad.Create(input)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"id":id,
	})
}


func (h *Handler) getAllAdsSortByPriceAsc(c *gin.Context){
	ads, err := h.services.GetAllAdsSortByPriceAsc()
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, GetAllAds{
		Data:ads,
	})
}

func (h *Handler) getAllAdsSortByPriceDesc(c *gin.Context){
	ads, err := h.services.GetAllAdsSortByPriceDesc()
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, GetAllAds{
		Data:ads,
	})
}

func (h *Handler) getAllAdsSortByDateAsc(c *gin.Context){
	ads, err := h.services.GetAllAdsSortByDateAsc()
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, GetAllAds{
		Data:ads,
	})
}

func (h *Handler) getAllAdsSortByDateDesc(c *gin.Context){
	ads, err := h.services.GetAllAdsSortByDateDesc()
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, GetAllAds{
		Data:ads,
	})
}

func (h *Handler) getAllAds(c *gin.Context){
	ads, err := h.services.GetAllAds()
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, GetAllAds{
		Data:ads,
	})
}

func (h *Handler) getAdById(c *gin.Context){
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil{
		newErrorResponse(c, http.StatusBadRequest, "invalid params")
		return
	}

	ad, err := h.services.GetAdById(id)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, ad)
}

