package main

import (
	"log"
	"world-skill-tree/config"
	_ "world-skill-tree/docs"
	"world-skill-tree/server"
	"world-skill-tree/server/routes"

	_ "github.com/joho/godotenv/autoload"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	// echoSwagger "github.com/swaggo/echo-swagger"
)

func main() {
	cfg := config.NewConfig()

	// docs.SwaggerInfo.Host = fmt.Sprintf("%s:%s", cfg.HTTP.Host, cfg.HTTP.Port)

	app := server.NewServer(cfg)

	app.Echo.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"http://localhost:3000"},
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept},
	  }))
	routes.ConfigureRoutes(app)

	err := app.Start(cfg.HTTP.Port)
	if err != nil {
		log.Fatal("Port already used")
	}
}
