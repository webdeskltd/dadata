package dadata_test

import (
	"reflect"
	"testing"

	. "gopkg.in/webdeskltd/dadata.v2"
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
		{"some belgorod ip", "217.118.95.255", "Белгород", false},
		{"some vladivostok ip", "80.83.239.112", "Владивосток", false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := daData.GeoIP(tt.ip)

			if (err != nil) != tt.wantErr {
				t.Errorf("DaData.GeoIP() error = %v, wantErr %v", err, tt.wantErr)
			}
			if tt.wantErr {
				// after wantErr it's make no sense check got - want
				return
			}
			if !reflect.DeepEqual(got.Location.Data.City, tt.wantCity) {
				t.Errorf("DaData.GeoIP() = %v, want %v", got.Location.Data.City, tt.wantCity)
			}
		})
	}
}

func TestDaData_GeolocateAddresses(t *testing.T) {
	tt := []struct {
		name       string
		query      GeolocateRequest
		wantStreet string
		wantErr    bool
	}{
		{
			"ГБОУ Школа №2120",
			GeolocateRequest{Lat: 55.6010979, Lon: 37.3593894},
			"77000006000003600",
			false,
		},
		{
			"МФЮА",
			GeolocateRequest{Lat: 55.722853, Lon: 37.692247},
			"77000000000001300",
			false,
		},
	}
	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			daData := instance()
			got, err := daData.GeolocateAddress(tc.query)
			if (err != nil) != tc.wantErr {
				t.Errorf("DaData.GeolocateAddress() error = %v, wantErr %v", err, tc.wantErr)
				return
			}
			kladr := got[0].Data.StreetKladrID
			if !reflect.DeepEqual(kladr, tc.wantStreet) {
				t.Errorf("DaData.GeolocateAddress() = %#v, want %#v", kladr, tc.wantStreet)
			}
		})
	}
}
