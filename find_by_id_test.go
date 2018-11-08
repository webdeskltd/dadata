package dadata_test

import (
	"reflect"
	"testing"
)

func TestDaData_AddressByID(t *testing.T) {
	daData := newAddressByIDDetector()

	tests := []struct {
		name     string
		id       string
		wantCity string
		wantErr  bool
	}{
		{"wrong id", "non-fias-non-kladr-111", "", true},
		{"samara by fias", "bb035cc3-1dc2-4627-9d25-a1bf2d4b936b", "Самара", false},
		{"samara by kladr", "6300000100000", "Самара", false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := daData.AddressByID(tt.id)

			if (err != nil) != tt.wantErr {
				t.Errorf("DaData.AddressByID() error = %v, wantErr %v", err, tt.wantErr)
			}
			if tt.wantErr {
				// after wantErr it's make no sense check got - want
				return
			}
			if got == nil {
				t.Error("DaData.AddressByID() return nil - should be non nil")
				return
			}
			if !reflect.DeepEqual(got.Data.City, tt.wantCity) {
				t.Errorf("DaData.AddressByID() = %v, want %v", got.Data.City, tt.wantCity)
			}
		})
	}
}
