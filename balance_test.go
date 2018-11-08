package dadata_test

import (
	"context"
	"testing"
	"time"

	. "gopkg.in/webdeskltd/dadata.v2"
)

// Тестирование нормального ответа
func TestDaDataProfileBalance(t *testing.T) {
	daData := newBalance()
	got, err := daData.ProfileBalance()
	if err != nil {
		t.Errorf("DaData.ProfileBalance() error = %s", err)
		return
	}
	if got.Balance == 0 {
		t.Error("DaData.ProfileBalance() = return 0 balance")
	}
}

// Тестирование ответа по таймауту
func TestDaDataProfileBalanceWithCancel(t *testing.T) {
	var err error
	var ctx context.Context
	var ctf context.CancelFunc
	var rsp *BalanceResponse

	ctx, ctf = context.WithTimeout(context.Background(), time.Microsecond*10)
	defer ctf()
	rsp, err = newBalance().ProfileBalanceWithCtx(ctx)
	if err == nil {
		t.Errorf("DaData.ProfileBalance() error, do not work context deadline exceeded")
		return
	}
	if rsp != nil {
		t.Error("DaData.ProfileBalance() response not nil object with error")
	}
}
