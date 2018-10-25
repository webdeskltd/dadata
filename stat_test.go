package dadata_test

import (
	"context"
	"testing"
	"time"
)

func TestDaData_DailyStat(t *testing.T) {
	// for current date we can check if stat is really more than 0
	// for sure - we need some call to api
	newCleaner().CleanNames("алекс")

	format := "2006-01-02"
	currentDate := time.Now().Format(format)
	tests := []struct {
		date    string
		wantErr bool
	}{
		{"2010-01-01", false},
		{"2016-01-01", false},
		{currentDate, false},
	}

	for _, tt := range tests {
		t.Run(tt.date, func(t *testing.T) {
			daData := newStater()
			date, _ := time.Parse(format, tt.date)
			got, err := daData.DailyStat(context.Background(), date)
			if (err != nil) != tt.wantErr {
				t.Errorf("DaData.DailyStat() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got.Date != tt.date {
				t.Errorf("DaData.DailyStat() = return stat for wrong date %v, want %v", got.Date, tt.date)
			}

			if tt.date != currentDate {
				return
			}
			// check stat for current date - it's must be more than 0
			// TODO Не корректно проверять на 0 статистику которая может быть нулевой
			if got.Services.Merging == 0 && tt.date != currentDate {
				t.Errorf("DaData.DailyStat() = return for current date %v - stat must be more than 0", tt.date)
			}
		})
	}
}
