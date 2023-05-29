package main

import (
	"context"
	"log"

	userHTTP "github.com/jwilyandi19/simple-auth-v2/delivery/http/user"
	domain "github.com/jwilyandi19/simple-auth-v2/domain/user"
	"github.com/jwilyandi19/simple-auth-v2/helper"
	userRepository "github.com/jwilyandi19/simple-auth-v2/repository/user"
	jwtUsecase "github.com/jwilyandi19/simple-auth-v2/usecase/jwt"
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

	userSeed := &domain.User{
		Name:     "Admin",
		Username: config.AdminUsername,
		Password: config.AdminPassword,
	}
	_, err = userUsecase.Create(ctx, userSeed)
	if err != nil {
		log.Fatal("cannot create seed user ", err)
	}

	jwtUsecase := jwtUsecase.NewJWTUsecase(userRepo, config)
	userHTTP.NewUserHandler(r, userUsecase, config, jwtUsecase)

	port := ":8080"
	r.Start(port)
}
