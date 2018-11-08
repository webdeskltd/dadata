package dadata_test

import (
	"os"
	"sync"
	"testing"

	. "github.com/webdeskltd/dadata"
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

func newGeoIPDetector() GeoIPDetector {
	return instance()
}

func newAddressByIDDetector() ByIDFinder {
	return instance()
}

func newStater() Stater {
	return instance()
}

func newBalance() Balancer {
	return instance()
}

func TestNewDaData(t *testing.T) {
	apiKey := "apiKey"
	secretKey := "secretKey"
	daData := NewDaData(apiKey, secretKey)

	if daData == nil {
		t.Errorf(`NewDaData return nil`)
	}

}
