package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type sendMessageReq struct {
	Content string `json:"content"`
}

type sendMessageRes struct {
	ID int `json:"id"`
}

func (h *Handler) sendMessage(c *gin.Context) {
	var req sendMessageReq
	if err := c.BindJSON(&req); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	id, err := h.service.Message.SendMessage(c, req.Content)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	res := sendMessageRes{ID: id}
	c.JSON(http.StatusOK, res)
}

func (h *Handler) processMessages(c *gin.Context) {
	err := h.service.Message.ProcessMessages(c)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.Status(http.StatusOK)
}
