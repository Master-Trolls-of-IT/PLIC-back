package main

import (
	"flag"
	"gaia-api/domain/service"
	api "gaia-api/infrastructure/controller"
	"gaia-api/infrastructure/controller/product"
	"gaia-api/infrastructure/repository"

	_ "github.com/golang-jwt/jwt/v5"
	_ "github.com/jackc/pgx/v4/stdlib"
	_ "github.com/lib/pq"
)

func main() {
	useLocalDB := flag.Bool("local", false, "use local database")
	dbURI := flag.String("dburi", "", "database URI")

	flag.Parse()
	var db *repository.Database
	var err error
	if *useLocalDB {
		db, err = repository.NewDatabase(*dbURI)
	} else {
		db, err = repository.NewDatabase()
	}
	if err != nil {
		panic(err)
	}

	//Repository
	userRepo := repository.NewUserRepository(db)
	productRepo := repository.NewProductRepository(db)
	mealRepo := repository.NewMealRepository(db, productRepo)

	//Service
	authenticationService := service.NewAuthService(userRepo)
	OpenFoodFactsService := service.NewOpenFoodFactsService(productRepo, mealRepo)

	//API
	OpenFoodFactsAPI := product.NewOpenFoodFactsAPI()

	//Server Instance
	ginServer := api.NewServer(authenticationService, OpenFoodFactsService, OpenFoodFactsAPI)

	ginServer.Start()
}
