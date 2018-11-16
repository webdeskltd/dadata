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

func TestNewDaData(t *testing.T) {
	apiKey := "apiKey"
	secretKey := "secretKey"
	daData := NewDaData(apiKey, secretKey)

	if daData == nil {
		t.Errorf(`NewDaData return nil`)
	}

}

func TestDaData_WithCustomHost(t *testing.T) {
	daData := instance()
	tests := []struct {
		host      string
		wantError bool
	}{
		{"https://dadata.ru/api/v2/", false},
		{"https://some.wrong/api/v2/", true},
		{"completely_non_uri%%$$$$$", true},
	}
	for _, tt := range tests {
		t.Run(tt.host, func(t *testing.T) {
			res, err := daData.WithCustomBaseURL(tt.host).CleanNames("Алексей Иванов")
			if (err != nil) != tt.wantError {
				t.Errorf("DaData.WithCustomBaseURL() = (%v, %v), want %v", res, err, tt.wantError)
			}
		})
	}
}

func TestDaData_WithCustomSuggestHost(t *testing.T) {
	daData := instance()
	tests := []struct {
		host      string
		wantError bool
	}{
		{"https://suggestions.dadata.ru/suggestions/api/4_1/rs/", false},
		{"https://some.wrong/api/v2/", true},
		{"completely_non_uri%%$$$$$", true},
	}
	for _, tt := range tests {
		t.Run(tt.host, func(t *testing.T) {
			res, err := daData.WithCustomSuggestURL(tt.host).AddressByID("32fcb102-2a50-44c9-a00e-806420f448ea")
			if (err != nil) != tt.wantError {
				t.Errorf("DaData.WithCustomSuggestURL() = (%v, %v), want %v", res, err, tt.wantError)
			}
		})
	}
}
