package config

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/go-chi/jwtauth"
	"github.com/subosito/gotenv"
)

func init() {
	err := gotenv.Load()
	if err != nil {
		log.Fatal("error to loading .env")
	}
}

func LoadConfigs() (*sql.DB, *jwtauth.JWTAuth, int, error) {
	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")
	jwtSecret := os.Getenv("JWT_SECRET")
	jwtExpireIn := os.Getenv("JWT_EXPIRESIN")

	expiresIn, err := strconv.Atoi(jwtExpireIn)
	if err != nil {
		log.Fatalf("err to converter jwt_expires to int %v", err)
	}

	psqlInfo := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", host, port, dbUser, dbPassword, dbName)

	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}

	err = db.Ping()
	if err != nil {
		panic(err)
	}

	configToken := jwtauth.New("HS256", []byte(jwtSecret), nil)

	return db, configToken, expiresIn, nil
}
