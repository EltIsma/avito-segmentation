package handler

import (
	"avito-third/internal/apperror"
	"avito-third/internal/segment"
	"avito-third/internal/user"
	"encoding/csv"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

// @Summary UsersOperation
// @Tags users
// @Description add_and_delete_segments
// @ID create-userssegment
// @Accept  json
// @Produce  json
// @Param input body  user.UserSegment true "segment info"
// @Success 200 {bool} bool true
// @Failure 400,404 {object} apperror.AppError
// @Failure 500 {object} apperror.AppError
// @Failure default {object} apperror.AppError
// @Router /users/ [post]

func (h *Handler) CRUDUsersInSegment(c *gin.Context) {
	var input user.UserSegment
	if err := c.BindJSON(&input); err != nil {
		apperror.NewAppError(c, "invalid input body", http.StatusBadRequest)
		return
	}

	err := h.services.User.CRUDOperation(&input)
	if err != nil {
		apperror.NewAppError(c, err.Error(), http.StatusInternalServerError)
		return
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

// @Summary Users_slus
// @Tags users
// @Description get_active_slugs
// @ID get-slugs
// @Accept  json
// @Produce  json
// @Param input body   user.UsersActiveSlugsDTO  true "user_id info"
// @Success 200 {integer} integer 1
// @Failure 400,404 {object} apperror.AppError
// @Failure 500 {object} apperror.AppError
// @Failure default {object} apperror.AppError
// @Router /users/activeSegments [post]

func (h *Handler) GetActiveSlugs(c *gin.Context) {
	var input user.UsersActiveSlugsDTO
	if err := c.BindJSON(&input); err != nil {
		apperror.NewAppError(c, "invalid input body", http.StatusBadRequest)
		return
	}

	activeSlugs, err := h.services.User.GetActive(input.User_id)
	if err != nil {
		apperror.NewAppError(c, err.Error(), http.StatusInternalServerError)
		return
	}
	c.JSON(http.StatusOK, map[string]interface{}{
		strconv.Itoa(input.User_id): activeSlugs,
	})
	c.Writer.Header().Set("Content-type", "application/json")
	c.Writer.WriteHeader(http.StatusOK)
}

// @Summary ReportFile
// @Tags users
// @Description get_report_url
// @ID get-url
// @Accept  json
// @Produce  json
// @Param input body  user.ReportSegmentRequest  true "period info"
// @Success 200 {string} string "URL"
// @Failure 400,404 {object} apperror.AppError
// @Failure 500 {object} apperror.AppError
// @Failure default {object} apperror.AppError
// @Router /users/reports [post]
func (h *Handler) GetUrlReportFile(c *gin.Context) {
	var input user.ReportSegmentRequest
	if err := c.BindJSON(&input); err != nil {
		apperror.NewAppError(c, "invalid input body", http.StatusBadRequest)
		return
	}
	date, err := time.Parse("2006-01", input.Period)
	if err != nil {
		apperror.NewAppError(c, err.Error(), http.StatusInternalServerError)
		return
	}

	report, err := h.services.User.GetReport(date)
	if err != nil {
		apperror.NewAppError(c, err.Error(), http.StatusInternalServerError)
		return
	}

	csvFile := fmt.Sprintf("%s-report.csv", input.Period)
	pathToFile := "report/" + csvFile
	if _, err := os.Stat(pathToFile); err == nil {
		if err := os.Remove(pathToFile); err != nil {
			return
		}
	}
	currentReportFile, err := os.Create(pathToFile)
	if err != nil {
		return
	}
	defer currentReportFile.Close()
	writer := csv.NewWriter(currentReportFile)
	writer.Comma = ';'
	defer writer.Flush()
	headers := []string{"UserID", "Slug", "Operation", "TimeExecution"}
	writer.Write(headers)

	for _, reportUser := range report {
		row := []string{
			strconv.Itoa(reportUser.UserID),
			reportUser.Slug,
			reportUser.Operation,
			reportUser.TimeOperation.Format("2006-01-02 15:04:05"),
		}
		writer.Write(row)
	}

	url := fmt.Sprintf("%s/%s", "localhost:8080/users/reportForperiod", csvFile)
	c.Writer.Header().Set("Content-type", "application/json")
	c.Writer.WriteHeader(http.StatusOK)
	c.JSON(http.StatusOK, map[string]string{
		"URL": url,
	})

}

// @Summary ReportFile
// @Tags users
// @Description get_report_file
// @ID get-file
// @Accept  json
// @Produce  json
// @Router /users/reportForperiod/* [get]

func (h *Handler) GetReportFile(c *gin.Context) {
	fileName := c.Param("path")
	reportFilePath := "./report" + fileName
	c.File(reportFilePath)
	c.Writer.Header().Set("Content-type", "application/json")
	c.Writer.WriteHeader(http.StatusOK)
}
