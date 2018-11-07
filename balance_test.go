package dadata_test

import (
	"context"
	"testing"
)

func TestDaData_ProfileBalance(t *testing.T) {
	// for current date we can check if stat is really more than 0
	// for sure - we need some call to api
	newCleaner().CleanNames("алекс")

	daData := newBalance()
	got, err := daData.ProfileBalance(context.Background())
	if err != nil {
		t.Errorf("DaData.ProfileBalance() error = %v", err)
		return
	}
	if got.Balance == 0 {
		t.Error("DaData.ProfileBalance() = return 0 balance")
	}
}
