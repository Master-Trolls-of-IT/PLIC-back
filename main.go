package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"reflect"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)

func main() {
    connStr := "host=localhost port=1234 dbname=gaia user=admin password=gaia2024 sslmode=disable"

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()

    err = db.Ping()
	fmt.Println("Connected to Google Cloud SQL instance Postgres database!")

    if err != nil {
        panic(err)
    }
	stmt , err := db.Query("SELECT tablename FROM pg_catalog.pg_tables;")
	
	if err !=nil {
		stmt.Err()
	}

	
	fmt.Println(reflect.TypeOf(stmt))
	var tables []string;
	var name string;
	for stmt.Next() {
		err := stmt.Scan(&name)
		if err != nil {
			panic(err)
		}
		tables = append(tables, name)
	}
	log.Println(tables)




	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"data": tables})
	})	

	r.Run()
}
