package constants

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}

func EnvMongoDbUrl() string {
	mongoDb := os.Getenv("MONGODB_URL")

	if mongoDb == "" {
		log.Fatal("MONGODB_URL not exist in .env")
	}

	return mongoDb
}

func EnvSecretKey() string {
	secretKey := os.Getenv("SECRET_KEY")

	if secretKey == "" {
		log.Fatal("SECRET_KEY not exist in .env")
	}

	return secretKey
}

func EnvPasswordPepper() string {
	passwordPepper := os.Getenv("PASSWORD_PEPPER")

	if passwordPepper == "" {
		log.Fatal("PASSWORD_PEPPER not exist in .env")
	}

	return passwordPepper
}
