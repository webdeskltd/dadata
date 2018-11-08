package dadata // import "gopkg.in/webdeskltd/dadata.v2"

import (
	"context"
	"time"
)

// DailyStat return daily statistics
// see documentation https://dadata.ru/api/stat/
func (daData *DaData) DailyStat(date time.Time) (*StatResponse, error) {
	return daData.DailyStatWithCtx(context.Background(), date)
}

// DailyStatWithCtx return daily statistics
// see documentation https://dadata.ru/api/stat/
func (daData *DaData) DailyStatWithCtx(ctx context.Context, date time.Time) (result *StatResponse, err error) {
	var dateStr string

	result, dateStr = &StatResponse{}, date.Format("2006-01-02")
	if err = daData.sendRequestToURL(ctx, "GET", baseURL+"stat/daily?date="+dateStr, nil, result); err != nil {
		result = nil
	}

	return
}
