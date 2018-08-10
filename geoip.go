package dadata

import "fmt"

const baseSuggestURL = "https://suggestions.dadata.ru/suggestions/api/4_1/rs/"

// GeoIPResponse response for GeoIP
type GeoIPResponse struct {
	Location *ResponseAddress `json:"location"`
}

// GeoIP try to find address by IP
// see documentation on:
//    https://dadata.ru/api/detect_address_by_ip/
//    https://confluence.hflabs.ru/pages/viewpage.action?pageId=715096277
// ip string representation of ip-address (example 10.12.44.23)
// if ip=="" then dadata try to get ip-address from X-Forwarded-For header
func (daData *DaData) GeoIP(ip string) (*GeoIPResponse, error) {

	var result GeoIPResponse

	if err := daData.sendRequestToURL("GET", baseSuggestURL+"detectAddressByIp?ip="+ip, nil, &result); err != nil {
		return nil, err
	}
	if result.Location == nil {
		return nil, fmt.Errorf("dadata.GeoIP: cannot detect address by ip %v", ip)
	}
	return &result, nil
}
