package main

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"time"

	"assignment_task/postgres"
	"assignment_task/task"
	"assignment_task/user"

	"github.com/labstack/echo/v4/middleware"

	_ "github.com/joho/godotenv/autoload"
	"github.com/labstack/echo/v4"
)

func main() {

	p := postgres.New()
	userHandler := user.New(p)
	taskHander := task.New(p)

	e := echo.New()
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"http://localhost:5173"},
		AllowMethods: []string{http.MethodGet, http.MethodPost, http.MethodPut, http.MethodDelete},
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAuthorization},
	}))

	{
		r := e.Group("/api/v1")
		r.POST("/register", userHandler.Register)
		r.POST("/login", userHandler.Login)

		r.GET("/task", taskHander.GetTask)
		r.POST("/task", taskHander.CraeteTask)
		r.PUT("/task", taskHander.UpdateTask)
		r.DELETE("/task", taskHander.DeleteTaskByID)
	}

	go func() {
		if err := e.Start(":" + os.Getenv("PORT")); err != nil && err != http.ErrServerClosed {
			e.Logger.Fatal(e.Start(":" + os.Getenv("PORT")))
		}
	}()

	shutdown := make(chan os.Signal, 1)
	signal.Notify(shutdown, os.Interrupt)
	<-shutdown
	fmt.Println("shutting down the server")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	if err := e.Shutdown(ctx); err != nil {
		e.Logger.Fatal(err)
	}
}
