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

// sendMessage handles the sending of a message.
// @Summary Send a message
// @Description Send a message to the service
// @Tags messages
// @Accept json
// @Produce json
// @Param message body sendMessageReq true "Message Content"
// @Success 200 {object} sendMessageRes
// @Failure 400 {object} errorResponse "Invalid request payload"
// @Failure 500 {object} errorResponse "Internal server error"
// @Router /messages [post]
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

// processMessages handles the processing of messages.
// @Summary Process messages
// @Description Trigger the processing of messages
// @Tags messages
// @Success 200 "Successfully processed messages"
// @Failure 500 {object} errorResponse "Internal server error"
// @Router /messages/process [post]
func (h *Handler) processMessages(c *gin.Context) {
	err := h.service.Message.ProcessMessages(c)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.Status(http.StatusOK)
}
