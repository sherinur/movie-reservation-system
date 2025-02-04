# middleware package

```
pkg/middleware/
├── cors.go # CORS middleware
├── go.mod
├── go.sum
└── jwt.go   # JWT middleware
```

### CORS
[CORS](https://developer.mozilla.org/en-US/docs/Web/HTTP/CORS)

To use this CORS middleware package in a Gin application, first, configure the allowed origins, methods, and headers by calling SetCorsConfig with a CorsConfig struct, e.g., SetCorsConfig(&CorsConfig{AllowedOrigins: []string{"https://example.com"}, AllowedMethods: []string{"GET", "POST"}, AllowedHeaders: []string{"Content-Type"}}). Then, apply the middleware to your Gin router using router.Use(CorsMiddleware()). The middleware checks the request's Origin header against the allowed origins and sets the appropriate CORS headers if it matches. For preflight OPTIONS requests, it responds with 200 OK and aborts further processing. Otherwise, it allows the request to proceed with c.Next().

Example of usage:
```
func NewServer(cfg *Config) Server {
	r := gin.Default()

	corsConfig := &middleware.CorsConfig{
		AllowedOrigins: []string{"http://localhost:4200"},
		AllowedMethods: []string{"GET", "POST", "UPDATE", "DELETE", "OPTIONS"},
		AllowedHeaders: []string{"Content-Type", "Authorization"},
	}

	// cors middleware
	middleware.SetCorsConfig(corsConfig)
	r.Use(middleware.CorsMiddleware())
        
       ...
}
```


### JWT
[JWT](https://habr.com/en/articles/340146/)

To use this JWT middleware package in a Gin application, first, set the JWT secret by calling `SetSecret` with your secret key, e.g., `SetSecret([]byte("your-secret-key"))`. Then, apply the middleware to your Gin router using `router.Use(JwtMiddleware())`. The middleware validates the `Authorization` header, ensuring it contains a valid `Bearer` token. It parses the JWT, verifies its signature, and extracts claims like `user_id` and `role`, which are then stored in the Gin context for use in subsequent handlers. If the token is invalid or missing, it responds with appropriate error messages and aborts the request.

Example:
```
func NewServer(cfg *Config) Server {
	...

	// jwt middleware
	middleware.SetSecret([]byte(cfg.JwtSecretKey))

	...
}

func (s *server) registerRoutes() error {
	db, err := db.ConnectMongo(s.cfg.DbUri, s.cfg.DbName)
	if err != nil {
		return err
	}

	userRepository := dal.NewUserRepository(db)
	userService := service.NewUserService(userRepository, s.cfg.JwtSecretKey)
	s.userHandler = handler.NewUserHandler(userService)

	s.router.GET("/health", handler.GetHealth)

	s.router.POST("/register", s.userHandler.HandleRegister)
	s.router.POST("/login", s.userHandler.HandleLogin)
        
        # Protected routes with JWT middleware
	s.router.GET("/users/me", middleware.JwtMiddleware(), s.userHandler.HandleProfile)
	s.router.PUT("/users/me/password", middleware.JwtMiddleware(), s.userHandler.HandleUpdatePassword)
	s.router.PUT("/users/me/email", middleware.JwtMiddleware(), s.userHandler.HandleUpdatePassword)
	s.router.DELETE("/users/me", middleware.JwtMiddleware(), s.userHandler.HandleDeleteProfile)

	return nil
}
```