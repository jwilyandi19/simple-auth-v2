package main

import (
	"context"
	"log"

	userHTTP "github.com/jwilyandi19/simple-auth-v2/delivery/http/user"
	"github.com/jwilyandi19/simple-auth-v2/helper"
	userRepository "github.com/jwilyandi19/simple-auth-v2/repository/user"
	userUsecase "github.com/jwilyandi19/simple-auth-v2/usecase/user"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	config, err := helper.LoadConfig(".")
	if err != nil {
		log.Fatal("cannot load config: ", err)
	}

	ctx := context.TODO()

	db, err := helper.OpenMongoDB(ctx, config)
	if err != nil {
		log.Fatal("cannot open db: ", err)
	}

	r := echo.New()
	r.Use(middleware.Recover())

	userRepo := userRepository.NewUserRepository(db)
	userUsecase := userUsecase.NewUserUsecase(userRepo)
	userHTTP.NewUserHandler(r, userUsecase)

	port := ":8080"
	r.Start(port)
}
