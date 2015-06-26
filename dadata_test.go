package dadata_test

import (
	"testing"

	dadata "."
)

func TestNewDaData(t *testing.T) {
	apiKey := "apiKey"
	secretKey := "secretKey"
	daData := dadata.NewDaData(apiKey, secretKey)

	if daData.ApiKey != apiKey {
		t.Errorf(`Invalid api key. Expect "%s", but got "%s"`, apiKey, daData.ApiKey)
	}

	if daData.SecretKey != secretKey {
		t.Errorf(`Invalid secret key. Expect "%s", but got "%s"`, secretKey, daData.SecretKey)
	}
}
