package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/mdzakyabd/dating-app/app/models"
	"github.com/mdzakyabd/dating-app/app/usecase"
)

type SwipeHandler struct {
	swipeUseCase usecase.SwipeUseCase
}

func NewSwipeHandler(swipeUseCase usecase.SwipeUseCase) *SwipeHandler {
	return &SwipeHandler{swipeUseCase: swipeUseCase}
}

func (h *SwipeHandler) Swipe(c *gin.Context) {
	var request struct {
		UserID       string `json:"user_id" binding:"required"`
		TargetUserID string `json:"target_user_id" binding:"required"`
		Liked        bool   `json:"liked"`
	}

	if err := c.BindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	userID, err := uuid.Parse(request.UserID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
		return
	}

	targetUserID, err := uuid.Parse(request.TargetUserID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid target user ID"})
		return
	}

	swipe := &models.Swipe{
		UserID:       userID,
		TargetUserID: targetUserID,
		Liked:        request.Liked,
	}

	err = h.swipeUseCase.Swipe(swipe)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to swipe"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Swipe recorded"})
}
