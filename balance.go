package dadata

import (
	"context"
	"errors"
)

// DailyStat return daily statistics
// see documentation https://dadata.ru/api/stat/
func (daData *DaData) ProfileBalance(ctx context.Context) (*BalanceResponse, error) {
	result := &BalanceResponse{}
	if err := daData.sendRequestToURL(ctx, "GET", baseURL+"profile/balance", nil, result); err != nil {
		return nil, err
	}

	if result == nil {
		return nil, errors.New("dadata.ProfileBalance: cannot return balance")
	}

	return result, nil
}
