package config

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

const (
	host     = "192.168.0.100"
	port     = "5432"
	user     = "postgres"
	password = "password"
	dbname   = "shorter"
)

func ConnectDB() (*sql.DB, error) {

	// Create Connection String
	psqlInfo := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)

	// Try to Connect into DB
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}

	// Test the connection
	err = db.Ping()
	if err != nil {
		panic(err)
	}

	fmt.Println("DB: Connect to " + dbname + " as " + user)

	// Return Db
	return db, nil
}
