package api

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

// envACCOUNTSID retrieves the Twilio Account SID from the environment variables.
func envACCOUNTSID() string {

	// Unmarshal .env file
	println(godotenv.Unmarshal(".env"))

	// Load and return Twilio Account SID from .env
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	return os.Getenv("TWILIO_ACCOUNT_SID")
}

// envAUTHTOKEN retrieves the Twilio Auth Token from the environment variables.
func envAUTHTOKEN() string {

	// Load and return Twilio Auth Token from .env
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	return os.Getenv("TWILIO_AUTHTOKEN")
}

// envSERVICESID retrieves the Twilio Services ID from the environment variables.
func envSERVICESID() string {

	// Load and return Twilio Services ID from .env
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	return os.Getenv("TWILIO_SERVICES_ID")
}
