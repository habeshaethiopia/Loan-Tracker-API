package infrastructure

import (
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

type Config struct {
	DatabaseUrl string
	Port        int
	Jwt_secret  string
	Dbname      string
	Usercoll    string
	Loancoll    string
}

func LoadEnv() (*Config, error) {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
		return nil, err
	}

	dbURL := os.Getenv("DATABASE_URL")
	portStr := os.Getenv("PORT")
	jwtSecret := os.Getenv("JWT_SECRET")
	dbname := os.Getenv("DB_NAME")
	usercoll := os.Getenv("USER_COLLECTION")
	loancoll := os.Getenv("LOAN_COLLECTION")

	port, err := strconv.Atoi(portStr)
	if err != nil {
		log.Fatal("Invalid PORT value")
		return nil, err
	}

	config := &Config{
		DatabaseUrl: dbURL,
		Port:        port,
		Jwt_secret:  jwtSecret,
		Dbname:      dbname,
		Usercoll:    usercoll,
		Loancoll:    loancoll,
	}

	return config, nil
}
