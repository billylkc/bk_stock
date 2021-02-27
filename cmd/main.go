package main

import (
	"encoding/json"
	"fmt"

	"github.com/billylkc/stock/quandl"
)

func main() {
	getByDate()
}

func dev() {
	quandl.Dev()
}

func getStock() {
	q := quandl.New()
	code := 5
	res, err := q.GetStockByCode(code)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(PrettyPrint(res))
}

func getByDate() {
	q := quandl.New()
	date := "2021-02-26"
	res, err := q.GetStockByDate(date)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(PrettyPrint(res))
}

// PrettyPrint to print struct in a readable way
func PrettyPrint(i interface{}) string {
	s, _ := json.MarshalIndent(i, "", "\t")
	return string(s)
}
