package main

import (
	"fmt"

	"github.com/billylkc/stock/quandl"
)

func main() {
	q := quandl.New()
	date := "2021-02-26"
	res, err := q.GetStockByDate(date)
	if err != nil {
		fmt.Println(err)
	}
	err = q.Insert(res)

}
