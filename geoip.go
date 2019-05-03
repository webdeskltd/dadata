package dadata // import "gopkg.in/webdeskltd/dadata.v2"

import (
	"context"
	"fmt"
	"net/http"
)

// GeoIP try to find address by IP
// see documentation on:
//    https://dadata.ru/api/detect_address_by_ip/
//    https://confluence.hflabs.ru/pages/viewpage.action?pageId=715096277
// ip string representation of ip-address (example 10.12.44.23)
// if ip=="" then dadata try to get ip-address from X-Forwarded-For header
func (daData *DaData) GeoIP(ip string) (*GeoIPResponse, error) {
	return daData.GeoIPWithCtx(context.Background(), ip)
}

// GeoIPWithCtx try to find address by IP
// see documentation on:
//    https://dadata.ru/api/detect_address_by_ip/
//    https://confluence.hflabs.ru/pages/viewpage.action?pageId=715096277
// ip string representation of ip-address (example 10.12.44.23)
// if ip=="" then dadata try to get ip-address from X-Forwarded-For header
func (daData *DaData) GeoIPWithCtx(ctx context.Context, ip string) (result *GeoIPResponse, err error) {
	result = &GeoIPResponse{}
	if err = daData.sendRequestToURL(ctx, "GET", baseSuggestURL+"detectAddressByIp?ip="+ip, nil, &result); err != nil {
		result = nil
	} else if result.Location == nil {
		result, err = nil, fmt.Errorf("dadata.GeoIP: cannot detect address by ip %s", ip)
	}
	return
}

// GeolocateAddress try to find address by coordinates
func (daData *DaData) GeolocateAddress(req GeolocateRequest) ([]ResponseAddress, error) {
	return daData.GeolocateAddressWithCtx(context.Background(), req)
}

// GeolocateAddressWithCtx try to find address by coordinates
func (daData *DaData) GeolocateAddressWithCtx(ctx context.Context, req GeolocateRequest) (ret []ResponseAddress, err error) {
	var result = &SuggestAddressResponse{}

	err = daData.sendRequestToURL(ctx, http.MethodPost, baseSuggestURL+"geolocate/address", &req, result)
	if err != nil {
		ret = nil
		return
	}
	if len(result.Suggestions) == 0 {
		ret, err = nil, fmt.Errorf("dadata.GeolocateAddress: cannot detect addresses by coordinates %+v", req)
		return
	}
	ret = result.Suggestions

	return
}
