package web

import (
	"fmt"
	"github.com/jinzhu/gorm"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"go_training/config"
	"go_training/initializer"
	"go_training/web/handler"
)

type Handlers struct {
	CreateUserHandler handler.CreateUserHandler
}

const apiPrefix = "/api"

func Init(conf config.Config, db *gorm.DB) Handlers {
	repositories := initializer.InitRepositories(db)
	createUserRepository := repositories.UserRepository
	createUserHandler := handler.CreateUserHandler{
		UserRepository: createUserRepository,
	}

	e := echo.New()

	e.Use(middleware.Logger())
	//e.Use(middleware.Recover())

	e.POST(fmt.Sprintf("%s/create_user", apiPrefix), createUserHandler.CreateUser)

	e.Logger.Fatal(e.Start(":8080"))
	return Handlers{
		CreateUserHandler: createUserHandler,
	}
}