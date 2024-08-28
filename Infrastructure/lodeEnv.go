package infrastructure

import (
	"log"
	"os"
	"strconv"
	"time"

	"github.com/joho/godotenv"
)

type Config struct {
	DatabaseUrl   string
	Port          int
	Jwt_secret    string
	Dbname        string
	Usercoll      string
	Bookcoll      string
	Borrowcoll    string
	Logcoll       string
	SmtpHost      string
	SmtpPort      int
	EmailFrom     string
	EmailPassword string
	ServerHost    string
	TokenTTL      time.Duration
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
	bookcoll := os.Getenv("BOOK_COLLECTION")
	borrowcoll := os.Getenv("BORROW_COLLECTION")
	logcoll := os.Getenv("LOG_COLLECTION")
	smtpHost := os.Getenv("SMTP_HOST")
	smtpPortStr := os.Getenv("SMTP_PORT")
	emailFrom := os.Getenv("EMAIL_FROM")
	emailPassword := os.Getenv("EMAIL_PASSWORD")
	serverHost := os.Getenv("SERVER_HOST")
	tokenTTLStr := os.Getenv("TOKEN_TTL")

	port, err := strconv.Atoi(portStr)
	if err != nil {
		log.Fatal("Invalid PORT value")
		return nil, err
	}
	smtpPort, err := strconv.Atoi(smtpPortStr)
	if err != nil {
		log.Fatalf("Error converting SMTP_PORT to int: %v", err)
		return nil, err
	}

	tokenTTL, err := time.ParseDuration(tokenTTLStr)
	if err != nil {
		log.Fatalf("Error parsing TOKEN_TTL: %v", err)
		return nil, err
	}

	config := &Config{
		DatabaseUrl:   dbURL,
		Port:          port,
		Jwt_secret:    jwtSecret,
		Dbname:        dbname,
		Usercoll:      usercoll,
		Bookcoll:      bookcoll,
		Borrowcoll:    borrowcoll,
		Logcoll:       logcoll,
		SmtpHost:      smtpHost,
		SmtpPort:      smtpPort,
		EmailFrom:     emailFrom,
		EmailPassword: emailPassword,
		ServerHost:    serverHost,
		TokenTTL:      tokenTTL,
	}

	return config, nil
}
