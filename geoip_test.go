package dadata_test

import (
	"reflect"
	"testing"
)

func TestDaData_DetectAddressByIP(t *testing.T) {
	daData := newGeoIPDetector()

	tests := []struct {
		name     string
		ip       string
		wantCity string
		wantErr  bool
	}{
		{"wrong ip", "xxx.999.ss.23", "", true},
		{"google proxy", "8.8.8.8", "", true},
		{"some msk ip", "85.235.162.70", "Москва", false},
		{"some kazan ip", "217.118.95.255", "Казань", false},
		{"some petropavl kamchats ip", "80.83.239.112", "Петропавловск-Камчатский", false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := daData.GeoIP(tt.ip)

			if (err != nil) != tt.wantErr {
				t.Errorf("DaData.GeoIP() error = %v, wantErr %v", err, tt.wantErr)
			}
			if tt.wantErr {
				// after wantErr it's make no sence check got - want
				return
			}
			if !reflect.DeepEqual(got.Location.Data.City, tt.wantCity) {
				t.Errorf("DaData.GeoIP() = %v, want %v", got.Location.Data.City, tt.wantCity)
			}
		})
	}
}
