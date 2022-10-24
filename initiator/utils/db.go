package utils

import (
	"fmt"
	"os"
)

func DbConnectionString() (string, error) {

	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASS")
	dbname := os.Getenv("DB_NAME")
	dbsslmode := os.Getenv("DB_SSL_MODE")

	addr := fmt.Sprintf("postgres://%v:%v@%v:%v/%v?sslmode=%v", user, password, host, port, dbname, dbsslmode)
	return addr, nil
}
