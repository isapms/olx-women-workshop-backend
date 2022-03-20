package database

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
	_ "github.com/joho/godotenv/autoload"
)

var (
	dbConn *sql.DB
)

func CreateConn() *sql.DB {
	var err error

	user := os.Getenv("DB_USER")
	pass := os.Getenv("DB_PASS")
	host := os.Getenv("DB_HOST")
	name := os.Getenv("DB_NAME")

	credentials := fmt.Sprintf("%s:%s@(%s:33061)/%s?charset=utf8&parseTime=True", user, pass, host, name)
	dbConn, err = sql.Open("mysql", credentials)

	if err != nil {
		log.Fatal(err)
	} else {
		fmt.Println("Succeeded to connect to database")
	}

	return dbConn
}

func GetConn() *sql.DB {
	return dbConn
}
