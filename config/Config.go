package config

import (
	"log"
	"os"
)

var ApiKey string
var PhoneNumber string
var SendUrl string

func init() {
	ApiKey = os.Getenv("WHATSGW_API_KEY")
	PhoneNumber = os.Getenv("WHATSGW_PHONE")
	SendUrl = os.Getenv("WHATSGW_SEND")

	if ApiKey == "" || PhoneNumber == "" {
		log.Fatal("API key or phone number not set in environment variables")
	}
}
