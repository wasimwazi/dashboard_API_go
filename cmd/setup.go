package cmd

import (
	"database/sql"
	"errors"
	"log"
	"os"

	// go package for postgres
	_ "github.com/lib/pq"
)

func prepareDatabase() (*sql.DB, error) {
	db, err := preparePostgres()
	if err != nil {
		return nil, err
	}
	log.Println("App : Database connected successfully!")
	return db, nil
}

func getServerAddr() string {
	port, ok := os.LookupEnv("SERVER_PORT")
	if !ok {
		log.Println("App : SERVER environment variable required but not set")
		os.Exit(1)
	}
	addr := ":" + port
	return addr
}

//CheckEnv : Check if the environment variables are set
func CheckEnv() error {
	_, ok := os.LookupEnv("SERVER_PORT")
	if !ok {
		return errors.New("SERVER PORT environment variable required but not set")
	}
	_, ok = os.LookupEnv("DB_URL")
	if !ok {
		return errors.New("DB_URL environment variable required but not set")
	}
	_, ok = os.LookupEnv("JWTKEY")
	if !ok {
		return errors.New("JWTKEY environment variable required but not set")
	}
	return nil
}
