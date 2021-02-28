package main

import (
	"encoding/json"
	"fmt"

	"github.com/billylkc/stock/stock"
)

func main() {
	dev()
}

func dev() {
	date := "2021-02-28"
	res, err := stock.GetIndustryDetails(date)
	fmt.Println(res)
}

// PrettyPrint to print struct in a readable way
func PrettyPrint(i interface{}) string {
	s, _ := json.MarshalIndent(i, "", "\t")
	return string(s)
}
