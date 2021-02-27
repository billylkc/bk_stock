package main

import (
	"fmt"

	"github.com/billylkc/stock/stock"
)

func main() {
	getStockPrice()
}

func getStockPrice() {
	res, _ := stock.GetStockPrice(5)
	fmt.Println(res)
}
