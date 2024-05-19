package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/mdzakyabd/dating-app/app/handler"
	"github.com/mdzakyabd/dating-app/app/middleware"
)

type AppRouteHandlers struct {
	UserHandler    handler.UserHandler
	ProfileHandler handler.ProfileHandler
	SwipeHandler   handler.SwipeHandler
	MatchHandler   handler.MatchHandler
}

func Routes(router *gin.Engine, handlers AppRouteHandlers, jwtSecret string) {
	router.POST("/signup", handlers.UserHandler.Register)
	router.POST("/login", handlers.UserHandler.Login)

	users := router.Group("/user")
	users.Use(middleware.JWTAuth(jwtSecret))
	{
		users.PUT("", handlers.UserHandler.UpdateUser)
		users.POST("/subscribe", handlers.UserHandler.SubscribePremium)
	}

	profile := router.Group("/profile")
	profile.Use(middleware.JWTAuth(jwtSecret))
	{
		profile.POST("", handlers.ProfileHandler.CreateProfile)
		profile.GET("", handlers.ProfileHandler.ViewProfiles)
		profile.GET("/:id", handlers.ProfileHandler.GetProfileByID)
		profile.PUT("/:id", handlers.ProfileHandler.UpdateProfile)
	}

	swipe := router.Group("/swipes")
	swipe.Use(middleware.JWTAuth(jwtSecret))
	{
		swipe.POST("", handlers.SwipeHandler.Swipe)
	}

	chatRoom := router.Group("/chat-rooms")
	chatRoom.Use(middleware.JWTAuth(jwtSecret))
	{
		chatRoom.GET("", handlers.MatchHandler.GetMatchRooms)
		chatRoom.DELETE("/:id", handlers.MatchHandler.DeleteMatchRoom)
		chatRoom.POST("/messages", handlers.MatchHandler.CreateMessage)
		chatRoom.GET("/:match_room_id/messages", handlers.MatchHandler.GetMessages)
	}
}
