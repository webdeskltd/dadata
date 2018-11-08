package dadata // import "gopkg.in/webdeskltd/dadata.v2"

import (
	"context"
)

// ProfileBalance return daily statistics
// see documentation https://dadata.ru/api/stat/
func (daData *DaData) ProfileBalance() (*BalanceResponse, error) {
	return daData.ProfileBalanceWithCtx(context.Background())
}

// ProfileBalanceWithCtx return daily statistics
// see documentation https://dadata.ru/api/stat/
func (daData *DaData) ProfileBalanceWithCtx(ctx context.Context) (result *BalanceResponse, err error) {
	result = new(BalanceResponse)
	if err = daData.sendRequestToURL(ctx, "GET", baseURL+"profile/balance", nil, result); err != nil {
		result = nil
	}
	return
}
