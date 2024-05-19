package main

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/mdzakyabd/dating-app/app/handler"
	"github.com/mdzakyabd/dating-app/app/models"
	"github.com/mdzakyabd/dating-app/app/repository"
	"github.com/mdzakyabd/dating-app/app/routes"
	"github.com/mdzakyabd/dating-app/app/scheduler"
	"github.com/mdzakyabd/dating-app/app/usecase"
	"github.com/mdzakyabd/dating-app/config"
)

func main() {
	db, err := config.ConfigDB()

	if err != nil {
		panic(err)
	}

	// Migrate the schema
	db.AutoMigrate(&models.User{}, &models.Profile{}, &models.MatchRoom{}, &models.Message{}, &models.Swipe{})

	pusherClient, err := config.ConfigPusher()
	if err != nil {
		panic(err)
	}

	jwtSecret, err := config.ConfigJWT()
	if err != nil {
		panic(err)
	}

	userRepo := repository.NewUserRepository(db)
	userUC := usecase.NewUserUseCase(userRepo)
	userHandler := handler.NewUserHandler(userUC)

	matchRepo := repository.NewMatchRepository(db)
	matchUC := usecase.NewMatchUsecase(matchRepo)
	matchHandler := handler.NewMatchHandler(matchUC, pusherClient)

	swipeRepo := repository.NewSwipeRepository(db)
	swipeUC := usecase.NewSwipeUseCase(swipeRepo, matchRepo)
	swipeHandler := handler.NewSwipeHandler(swipeUC)

	profileRepo := repository.NewProfileRepository(db)
	profileUC := usecase.NewProfileUseCase(profileRepo, userRepo, swipeRepo, matchRepo)
	profileHandler := handler.NewProfileHandler(profileUC)

	routeHandler := routes.AppRouteHandlers{
		UserHandler:    *userHandler,
		ProfileHandler: *profileHandler,
		SwipeHandler:   *swipeHandler,
		MatchHandler:   *matchHandler,
	}

	r := gin.Default()
	r.Use(cors.Default())
	routes.Routes(r, routeHandler, jwtSecret)

	// Start the scheduler
	checkExpiredScheduler := scheduler.NewScheduler(userRepo)
	checkExpiredScheduler.Start()

	r.Run()
}
