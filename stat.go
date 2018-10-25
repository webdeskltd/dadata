package dadata

import (
	"context"
	"fmt"
	"time"
)

// DailyStat return daily statistics
// see documentation https://dadata.ru/api/stat/
func (daData *DaData) DailyStat(ctx context.Context, date time.Time) (*StatResponse, error) {
	result := &StatResponse{}
	dateStr := date.Format("2006-01-02")
	if err := daData.sendRequestToURL(ctx, "GET", baseURL+"stat/daily?date="+dateStr, nil, result); err != nil {
		return nil, err
	}

	if result == nil {
		return nil, fmt.Errorf("dadata.Stat: cannot return daily stat for date %v", dateStr)
	}

	return result, nil
}
