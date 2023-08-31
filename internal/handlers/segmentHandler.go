package handler

import (
	"avito-third/internal/apperror"
	"avito-third/internal/segment"
	"encoding/json"
	_ "fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

// @Summary SegmentOperation
// @Tags segment
// @Description create segment
// @ID create-segment
// @Accept  json
// @Produce  json
// @Param input body segment.SegmentDTO true "segment info"
// @Success 200 {bool} bool true
// @Failure 400,404 {object} apperror.AppError
// @Failure 500 {object} apperror.AppError
// @Failure default {object} apperror.AppError
// @Router /segments/ [post]
func (h *Handler) createSegment(c *gin.Context) {
	var input segment.SegmentDTO
	if err := c.BindJSON(&input); err != nil {
		apperror.NewAppError(c, "invalid input body", http.StatusBadRequest)
		return
	}

	err := h.services.Segment.Create(&input)
	if err != nil {
		apperror.NewAppError(c, err.Error(), http.StatusInternalServerError)
	}
	response, err := json.Marshal(segment.SegmentResponseDTO{Result: true})
	if err != nil {
		c.Writer.WriteHeader(http.StatusInternalServerError)
		return
	}
	c.Writer.Header().Set("Content-type", "application/json")
	c.Writer.WriteHeader(http.StatusOK)
	c.Writer.Write(response)
}

// @Summary SegmentOperation
// @Tags segment
// @Description delete segment
// @ID delete-segment
// @Accept  json
// @Produce  json
// @Param input body segment.SegmentDTO true "segment info"
// @Success 200 {bool} bool true
// @Failure 400,404 {object} apperror.AppError
// @Failure 500 {object} apperror.AppError
// @Failure default {object} apperror.AppError
// @Router /segments/ [delete]
func (h *Handler) deleteSegment(c *gin.Context) {
	var input segment.SegmentDTO
	if err := c.BindJSON(&input); err != nil {
		return
	}

	err := h.services.Segment.Delete(&input)
	if err != nil {
		return
	}
	response, err := json.Marshal(segment.SegmentResponseDTO{Result: true})
	if err != nil {
		return
	}
	c.Writer.Header().Set("Content-type", "application/json")
	c.Writer.WriteHeader(http.StatusOK)
	c.Writer.Write(response)
}
