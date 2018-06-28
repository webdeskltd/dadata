package dadata_test

import (
	"os"
	"sync"
	"testing"

	. "github.com/mekegi/dadata"
)

var (
	client *DaData
	guard  sync.Once
)

func instance() *DaData {
	guard.Do(func() {
		apiKey := os.Getenv("API_KEY")
		secretKey := os.Getenv("SECRET_KEY")

		client = NewDaData(apiKey, secretKey)
	})

	return client
}

func newCleaner() Cleaner {
	return instance()
}

func newSuggester() Suggester {
	return instance()
}

func TestNewDaData(t *testing.T) {
	apiKey := "apiKey"
	secretKey := "secretKey"
	daData := NewDaData(apiKey, secretKey)

	if daData.ApiKey != apiKey {
		t.Errorf(`Invalid api key. Expect "%s", but got "%s"`, apiKey, daData.ApiKey)
	}

	if daData.SecretKey != secretKey {
		t.Errorf(`Invalid secret key. Expect "%s", but got "%s"`, secretKey, daData.SecretKey)
	}
}
