package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/mdzakyabd/dating-app/app/models"
	"github.com/mdzakyabd/dating-app/app/usecase"
	"github.com/pusher/pusher-http-go"
)

type MatchHandler struct {
	matchUsecase usecase.MatchUsecase
	pusherClient *pusher.Client
}

func NewMatchHandler(matchUsecase usecase.MatchUsecase, pusherClient *pusher.Client) *MatchHandler {
	return &MatchHandler{matchUsecase, pusherClient}
}

func (h *MatchHandler) GetMatchRooms(c *gin.Context) {
	userID, exists := c.Get("userID")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "user not authorized"})
		return
	}
	matchRooms, err := h.matchUsecase.GetMatchRooms(userID.(uuid.UUID))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, matchRooms)
}

func (h *MatchHandler) DeleteMatchRoom(c *gin.Context) {
	userID, exists := c.Get("userID")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "user not authorized"})
		return
	}
	id, err := uuid.Parse(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid match room ID"})
		return
	}
	err = h.matchUsecase.DeleteMatchRoom(id, userID.(uuid.UUID))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.Status(http.StatusNoContent)
}

func (h *MatchHandler) CreateMessage(c *gin.Context) {
	var request struct {
		MatchRoomID string `json:"match_room_id" binding:"required"`
		UserID      string `json:"user_id" binding:"required"`
		Content     string `json:"content" binding:"required"`
	}
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	matchRoomID, err := uuid.Parse(request.MatchRoomID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid match room ID"})
		return
	}

	userID, err := uuid.Parse(request.UserID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid match room ID"})
		return
	}

	err = h.matchUsecase.CreateMessage(&models.Message{MatchRoomID: matchRoomID, SenderID: userID, Content: request.Content})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Pusher real-time update
	data := map[string]string{"user_id": request.UserID, "content": request.Content}

	h.pusherClient.Trigger("chat_room_"+request.MatchRoomID, "new_message", data)

	c.Status(http.StatusCreated)
}

func (h *MatchHandler) GetMessages(c *gin.Context) {
	matchRoomID, err := uuid.Parse(c.Param("match_room_id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid match room ID"})
		return
	}
	messages, err := h.matchUsecase.GetMessages(matchRoomID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, messages)
}
