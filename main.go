package main

import (
	"flag"
	"gaia-api/domain/service"
	api "gaia-api/infrastructure/controller"
	"gaia-api/infrastructure/repository"

	_ "github.com/golang-jwt/jwt/v5"
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
	userService := service.NewUserService(userRepo)
	productService := service.NewProductService(productRepo)
	mealService := service.NewMealService(mealRepo)

	//Server Instance
	ginServer := api.NewServer(userService, productService, mealService)

	ginServer.Start()
}
