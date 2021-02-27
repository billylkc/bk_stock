package main

import (
	"encoding/json"
	"fmt"

	"github.com/billylkc/stock/quandl"
)

func main() {
	getStock()
}

func dev() {
	quandl.Dev()
}

func getStock() {
	q := quandl.New()
	code := 5
	date := "2021-02-26"
	res, err := q.GetStock(code, date)
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
