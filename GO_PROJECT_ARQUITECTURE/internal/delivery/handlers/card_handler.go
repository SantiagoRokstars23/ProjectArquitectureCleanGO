package handlers

import (
	"github.com/santiago/cards-service/internal/domain"
	"github.com/santiago/cards-service/internal/usecase"
	"net/http"
	"strconv"
	"github.com/gin-gonic/gin"
)

type CardHandler struct {
	uc *usecase.CardUseCase
}

func NewCardHandler(uc *usecase.CardUseCase) *CardHandler {
	return &CardHandler{uc: uc}
}

func (h *CardHandler) List(c *gin.Context) {
	data, err := h.uc.List()
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, data)
}

func (h *CardHandler) Get(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))

	card, err := h.uc.Get(id)
	if err != nil {
		c.JSON(http.StatusNotFound, err.Error())
		return
	}
	c.JSON(http.StatusOK, card)
}

func (h *CardHandler) Create(c *gin.Context) {
	var card domain.Card

	if err := c.ShouldBindJSON(&card); err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	err := h.uc.Create(card)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	c.Status(http.StatusCreated)
}