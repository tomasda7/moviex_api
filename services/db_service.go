package services

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

type DBdata struct {
	Host        string
	Port        string
	DBname      string
	RolName     string
	RolPassword string
}

func LoadEnv() (DBdata, error) {
	var err error
	if err = godotenv.Load(".env"); err != nil {
		log.Fatalf("Error trying to load the file .env: %v", err)
		return DBdata{}, err
	}

	return DBdata{
		Host: os.Getenv("HOST"),
		Port: os.Getenv("PORT"),
		DBname: os.Getenv("DBNAME"),
		RolName: os.Getenv("ROLNAME"),
		RolPassword: os.Getenv("ROLPASSWORD"),
	}, nil
}

var DB *sql.DB

func EstablishDbConn() error {
	dbData, err := LoadEnv()

	if err != nil {
		log.Fatal(err)
	}

	psqlInfo := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", dbData.Host, dbData.Port, dbData.RolName, dbData.RolPassword, dbData.DBname)

	dbConn, err := sql.Open("postgres", psqlInfo)

	if err != nil {
		log.Fatal(err)
	}

	DB = dbConn
	fmt.Println("Successful connection to the database:", DB)

	if err := DB.Ping(); err != nil {
		DB.Close()
		fmt.Println("Error trying to Ping db.")
		return err
	}

	return nil
}
