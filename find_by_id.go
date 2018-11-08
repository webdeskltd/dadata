package dadata // import "gopkg.in/webdeskltd/dadata.v2"

import (
	"context"
	"fmt"
)

// AddressByID find address by Fias or Kladr
// see full documentation https://confluence.hflabs.ru/pages/viewpage.action?pageId=312016944
func (daData *DaData) AddressByID(id string) (*ResponseAddress, error) {
	return daData.AddressByIDWithCtx(context.Background(), id)
}

// AddressByIDWithCtx find address by Fias or Kladr
// see full documentation https://confluence.hflabs.ru/pages/viewpage.action?pageId=312016944
func (daData *DaData) AddressByIDWithCtx(ctx context.Context, id string) (address *ResponseAddress, err error) {
	var result []ResponseAddress
	if result, err = daData.AddressesByIDWithCtx(ctx, id); err != nil {
		return
	}
	address = &result[0]
	return
}

// AddressesByID find addresses by Fias or Kladr
// see full documentation https://confluence.hflabs.ru/pages/viewpage.action?pageId=312016944
func (daData *DaData) AddressesByID(id string) ([]ResponseAddress, error) {
	return daData.AddressesByIDWithCtx(context.Background(), id)
}

// AddressesByIDWithCtx find addresses by Fias or Kladr
// see full documentation https://confluence.hflabs.ru/pages/viewpage.action?pageId=312016944
func (daData *DaData) AddressesByIDWithCtx(ctx context.Context, id string) (addresses []ResponseAddress, err error) {
	var result = &SuggestAddressResponse{}
	var req = SuggestRequestParams{Query: id}

	if err = daData.sendRequestToURL(ctx, "POST", baseSuggestURL+"findById/address", req, result); err != nil {
		return
	}
	if len(result.Suggestions) == 0 {
		err = fmt.Errorf("dadata.AddressByID: cannot detect address by id %s", id)
		return
	}
	addresses = result.Suggestions

	return
}
