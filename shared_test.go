package dadata_test

import (
	"os"

	dadata "."
)

func newCleaner() dadata.Cleaner {
	apiKey := os.Getenv("API_KEY")
	secretKey := os.Getenv("SECRET_KEY")

	return dadata.NewDaData(apiKey, secretKey)
}
