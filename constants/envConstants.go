package constants

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func init() {
	mode := os.Getenv("GIN_MODE")
	if mode != "release" {
		err := godotenv.Load()
		if err != nil {
			log.Fatal("Error loading .env file")
		}
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

func EnvGoogleClientID() string {
	googleClientID := os.Getenv("GOOGLE_CLIENT_ID")

	if googleClientID == "" {
		log.Fatal("GOOGLE_CLIENT_ID not exist in .env")
	}

	return googleClientID
}

func EnvGoogleClientSecret() string {
	googleClientSecret := os.Getenv("GOOGLE_CLIENT_SECRET")

	if googleClientSecret == "" {
		log.Fatal("GOOGLE_CLIENT_SECRET not exist in .env")
	}

	return googleClientSecret
}

func EnvGoogleRedirectURI() string {
	googleRedirectURI := os.Getenv("GOOGLE_REDIRECT_URL")
	if googleRedirectURI == "" {
		log.Fatal("GOOGLE_REDIRECT_URL not exist in .env")
	}

	return googleRedirectURI
}
