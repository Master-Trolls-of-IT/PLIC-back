package main

import (
	"flag"
	"gaia-api/application/interfaces"
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

	userRepo := repositories.NewUserRepository(db)
	productRepo := repositories.NewProductRepository(db)
	_ = productRepo
	authenticationService := services.NewAuthService(userRepo)
	returnAPIData := interfaces.NewReturnAPIData()
	ginServer := api.NewServer(authenticationService, returnAPIData)

	ginServer.Start()
}
