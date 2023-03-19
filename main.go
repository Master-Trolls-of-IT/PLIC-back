package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	_ "github.com/jackc/pgx/v4/stdlib"
	_ "github.com/lib/pq"
)

func connectUnixSocket() (*sql.DB, error) {
	mustGetenv := func(k string) string {
		v := os.Getenv(k)
		if v == "" {
			log.Fatalf("Fatal Error in connect_unix.go: %s environment variable not set.\n", k)
		}
		return v
	}
	var (
		dbUser         = mustGetenv("my-db-user")     // e.g. 'my-db-user'
		dbPwd          = mustGetenv("my-db-password") // e.g. 'my-db-password'
		unixSocketPath = mustGetenv("instance")       // e.g. '/cloudsql/project:region:instance'
		dbName         = mustGetenv("my-database")    // e.g. 'my-database'
	)

	dbURI := fmt.Sprintf("user=%s password=%s database=%s host=%s",
		dbUser, dbPwd, dbName, unixSocketPath)

	// dbPool is the pool of database connections.
	db, err := sql.Open("pgx", dbURI)
	if err != nil {
		return nil, fmt.Errorf("sql.Open: %v", err)
	}

	return db, nil
}

func main() {
	db, err := connectUnixSocket()
	if err != nil {
		panic(err)
	}

	stmt, err := db.Query("SELECT tablename FROM pg_catalog.pg_tables;")

	if err != nil {
		stmt.Err()
	}

	var tables []string
	var name string
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
