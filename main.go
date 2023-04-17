package main

import (
	"flag"
	"gaia-api/domain/services"
	api "gaia-api/infrastructure/controllers"
	"gaia-api/infrastructure/repositories"

	_ "github.com/golang-jwt/jwt/v5"
	_ "github.com/jackc/pgx/v4/stdlib"
	_ "github.com/lib/pq"
)

func main() {
	useLocalDB := flag.Bool("local", false, "use local database")
	dbURI := flag.String("dburi", "", "database URI")

	flag.Parse()
	var db *repositories.Database
	var err error
	if *useLocalDB {
		db, err = repositories.NewDatabase(*dbURI)
	} else {
		db, err = repositories.NewDatabase()
	}
	if err != nil {
		panic(err)
	}

	user_repo := repositories.NewUserRepository(db)
	authentication_service := services.NewAuthService(user_repo)
	gin_server := api.NewServer(authentication_service)

	gin_server.Start()
}
