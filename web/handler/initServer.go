package handler

import (
	"fmt"
	"github.com/jinzhu/gorm"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"go_training/config"
	"go_training/initializer"
	"go_training/web/api_error"
)

const apiPrefix = "/api"

func InitServer(conf config.Config, db *gorm.DB) Handlers {
	repositories := initializer.InitRepositories(db)
	handlers := InitHandler(repositories)
	e := NewRouter(handlers)
	e.Logger.Fatal(e.Start(":8080"))
	return handlers
}

func NewRouter(handlers Handlers) *echo.Echo {
	e := echo.New()

	e.HTTPErrorHandler = api_error.CustomHTTPErrorHandler

	e.Use(middleware.Logger())
	//e.Use(middleware.Recover())

	e.POST(fmt.Sprintf("%s/users", apiPrefix), handlers.CreateUserHandler.CreateUser)

	return e
}
