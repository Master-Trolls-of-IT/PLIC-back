package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)

func connectUnixSocket() (*sql.DB, error) {
	mustGetenv := func(k string) string {
		v := os.Getenv(k)
		if v == "" {
			log.Fatalf("Fatal Error in connect_unix.go: %s environment variable not set.", k)
		}
		return v
	}
	// Note: Saving credentials in environment variables is convenient, but not
	// secure - consider a more secure solution such as
	// Cloud Secret Manager (https://cloud.google.com/secret-manager) to help
	// keep secrets safe.
	var (
		dbUser         = mustGetenv("admin")                                          // e.g. 'my-db-user'
		dbPwd          = mustGetenv("gaia2024")                                       // e.g. 'my-db-password'
		dbName         = mustGetenv("gaia")                                           // e.g. 'my-database'
		unixSocketPath = mustGetenv("/cloudsql/gaia-api-380213:europe-west9:gaia-db") // e.g. '/cloudsql/project:region:instance'
	)

	dbURI := fmt.Sprintf("%s:%s@unix(%s)/%s?parseTime=true",
		dbUser, dbPwd, unixSocketPath, dbName)

	// dbPool is the pool of database connections.
	dbPool, err := sql.Open("postgres", dbURI)
	if err != nil {
		return nil, fmt.Errorf("sql.Open: %v", err)
	}

	// ...

	return dbPool, nil
}

func main() {
	/*connStr := "host= port= dbname=gaia user=admin password=gaia2024 sslmode=disable"

		db, err := sql.Open("postgres", connStr)
		if err != nil {
			panic(err.Error())
		}
		defer db.Close()

	    err = db.Ping()

	    if err != nil {
	        panic(err)
	    }

		fmt.Println("Connected to Google Cloud SQL instance Postgres database!")
	*/
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
