# Client for DaData.ru
[![Build Status](https://travis-ci.org/webdeskltd/dadata.svg)](https://travis-ci.org/webdeskltd/dadata)
[![GoDoc](https://godoc.org/gopkg.in/webdeskltd/dadata.v2?status.svg)](https://godoc.org/gopkg.in/webdeskltd/dadata.v2)

DaData API v2

Implemented [Clean](https://dadata.ru/api/clean/) and [Suggest](https://dadata.ru/api/suggest/) methods.

## Installation

`go get gopkg.in/webdeskltd/dadata.v2`

## Usage
```go
package main

import (
	"fmt"

	"gopkg.in/webdeskltd/dadata.v2"
)

func main() {
	daData := dadata.NewDaData("PUT_YOUR_API_KEY", "PUT_YOUR_SECRET_KEY")

	banks, err := daData.SuggestBanks(dadata.SuggestRequestParams{Query: "Кредитный", Count: 3})
	if nil != err {
		fmt.Println(err)
	}

	for _, bank := range banks {
		fmt.Println(bank.Data.Name.Full)
		fmt.Println(bank.Data.Bic)
	}

	// Output:
	// "МОСКОВСКИЙ КРЕДИТНЫЙ БАНК" (ПУБЛИЧНОЕ АКЦИОНЕРНОЕ ОБЩЕСТВО)
	// 044525659
	// КОММЕРЧЕСКИЙ БАНК "РЕСПУБЛИКАНСКИЙ КРЕДИТНЫЙ АЛЬЯНС" (ОБЩЕСТВО С ОГРАНИЧЕННОЙ ОТВЕТСТВЕННОСТЬЮ)
	// 044525860
	// ЖИЛИЩНО-КРЕДИТНЫЙ КОММЕРЧЕСКИЙ БАНК "ЖИЛКРЕДИТ" ОБЩЕСТВО С ОГРАНИЧЕННОЙ ОТВЕТСТВЕННОСТЬЮ
	// 044525325
}
```

more examples in [examples_test.go](./examples_test.go)

## Licence
MIT see [LICENSE](LICENSE)
