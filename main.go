package main

import (
	"gaia-api/domain/services"
	api "gaia-api/infrastructure/controllers"
	"gaia-api/infrastructure/repositories"

	_ "github.com/jackc/pgx/v4/stdlib"
	_ "github.com/lib/pq"
)

func main() {
	db, err := repositories.NewDatabase()
	if err != nil {
		panic(err)
	}
	user_repo := repositories.NewUserRepository(db)
	authentication_service := services.NewAuthService(user_repo)
	gin_server := api.NewServer(authentication_service)

	gin_server.Start()
}
