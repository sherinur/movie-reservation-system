package server

import (
	"user-service/internal/dal"
	"user-service/internal/handler"
	"user-service/internal/service"

	"github.com/sherinur/movie-reservation-system/pkg/db"
	"github.com/sherinur/movie-reservation-system/pkg/middleware"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// TODO: make the middleware not global, composite in server struct

func (s *server) registerRoutes() error {
	db, err := db.ConnectMongo(s.cfg.DbUri, s.cfg.DbName)
	if err != nil {
		return err
	}

	s.router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	userRepository := dal.NewUserRepository(db)
	userService := service.NewUserService(userRepository)
	tokenService := service.NewTokenService(s.cfg.JwtAccessSecret, s.cfg.JwtRefreshSecret, s.cfg.JwtAccessExpiration, s.cfg.JwtRefreshExpiration)
	s.userHandler = handler.NewUserHandler(userService, tokenService, s.log)

	s.router.GET("/health", handler.GetHealth)

	s.router.POST("/users/register", s.userHandler.HandleRegister)
	s.router.POST("/users/login", s.userHandler.HandleLogin)
	s.router.GET("/users/me", middleware.JwtMiddleware(), s.userHandler.HandleProfile)
	s.router.PUT("/users/me/password", middleware.JwtMiddleware(), s.userHandler.HandleUpdatePassword)
	s.router.PUT("/users/me/email", middleware.JwtMiddleware(), s.userHandler.HandleUpdatePassword)
	s.router.DELETE("/users/me", middleware.JwtMiddleware(), s.userHandler.HandleDeleteProfile)

	return nil
}
