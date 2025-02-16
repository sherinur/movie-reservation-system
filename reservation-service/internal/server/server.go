package server

import (
	"os"

	"reservation-service/internal/dal"
	"reservation-service/internal/handler"
	"reservation-service/internal/service"

	"github.com/gin-gonic/gin"
	"github.com/sherinur/movie-reservation-system/pkg/db"
	"github.com/sherinur/movie-reservation-system/pkg/logging"
	"github.com/sherinur/movie-reservation-system/pkg/middleware"
)

type Server interface {
	Start() error
	Shutdown()
	registerRoutes() error
}

type server struct {
	router *gin.Engine
	cfg    *config
	log    *logging.Logger

	reservationHandler handler.ReservationHandler
	promotionHandler   handler.PromotionHandler
	paymentHandler     handler.PaymentHandler
}

func NewServer(cfg *config) Server {
	r := gin.Default()
	corsConfig := &middleware.CorsConfig{
		AllowedOrigins: []string{"http://localhost:3000"},
		AllowedMethods: []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders: []string{"Content-Type", "Authorization"},
	}

	// cors middleware
	middleware.SetCorsConfig(corsConfig)
	r.Use(middleware.CorsMiddleware())

	// jwt middleware
	middleware.SetSecret([]byte(cfg.SecretKey))
	return &server{
		router: r,
		cfg:    cfg,
		log:    logging.NewLogger("dev"),
	}
}

func (s *server) Start() error {
	err := s.registerRoutes()
	if err != nil {
		s.log.Errorf("Could not register routes: %s", err.Error())
	}

	s.log.Info("Sarting server at the port" + s.cfg.Port)

	err = s.router.Run(s.cfg.Port)
	if err != nil {
		s.log.Errorf("Error starting server: %s", err.Error())
	}

	return nil
}

func (s *server) Shutdown() {
	os.Exit(0)
}

func (s *server) registerRoutes() error {
	database, err := db.ConnectMongo(s.cfg.DBuri, s.cfg.DBname)
	if err != nil {
		return err
	}

	s.log.Info("Registering routes..")

	reservationRepository := dal.NewReservationRepository(database)
	reservationService := service.NewReservationService(reservationRepository)
	s.reservationHandler = handler.NewReservationHandler(reservationService, s.log)

	promotionRepository := dal.NewPromotionRepository(database)
	promotionService := service.NewPromotionService(promotionRepository)
	s.promotionHandler = handler.NewPromotionHandler(promotionService, s.log)

	paymentRepository := dal.NewPaymentRepository(database)
	paymentService := service.NewPaymentService(paymentRepository)
	s.paymentHandler = handler.NewPaymentHandler(paymentService, s.log)

	reservation := s.router.Group("/booking")
	reservation.Use(middleware.JwtMiddleware())
	{
		reservation.POST("/", s.reservationHandler.AddReservation)
		reservation.GET("/", s.reservationHandler.GetReservations)
		reservation.GET("/:id", s.reservationHandler.GetReservation)
		reservation.PUT("/:id", s.reservationHandler.PayReservation)
		reservation.DELETE("/delete/:id", s.reservationHandler.DeleteReservation)
	}

	promotions := s.router.Group("/promotions")
	{
		promotions.POST("/", s.promotionHandler.AddPromotion)
		promotions.GET("/", s.promotionHandler.GetPromotions)
		promotions.GET("/:id", s.promotionHandler.GetPromotion)
		promotions.PUT("/:id", s.promotionHandler.UpdatePromotion)
		promotions.DELETE("/delete/:id", s.promotionHandler.DeletePromotion)
	}

	payments := s.router.Group("/payments")
	payments.Use(middleware.JwtMiddleware())
	{
		payments.POST("/", s.paymentHandler.AddPayment)
		payments.GET("/", s.paymentHandler.GetPayments)
		payments.GET("/:id", s.paymentHandler.GetPayment)
		payments.PUT("/:id", s.paymentHandler.UpdatePayment)
		payments.DELETE("/delete/:id", s.paymentHandler.DeletePayment)
	}

	return nil
}
