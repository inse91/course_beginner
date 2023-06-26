package handler

import (
	"course_beginner/internal/model"
	"course_beginner/internal/usecase"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"io"
	"log"
	"net/http"
)

type Handler struct {
	uc *usecase.UseCase
}

func New(uc *usecase.UseCase) *Handler {
	return &Handler{uc: uc}
}

func (h *Handler) Register(r *gin.Engine) {
	r.POST("/sum", h.sum)
}

func (h *Handler) sum(c *gin.Context) {

	bts, err := io.ReadAll(c.Request.Body)
	if err != nil {
		log.Fatal(err)
	}
	var input model.SumRequest
	if err = json.Unmarshal(bts, &input); err != nil {
		log.Fatal(err)
	}

	sum, err := h.uc.Sum(input.A, input.B)
	if err != nil {
		c.JSON(http.StatusBadRequest, model.SumResponse{
			Sum:     0,
			Success: false,
			Error:   err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, model.SumResponse{
		Sum:     sum,
		Success: true,
	})
	return

}
